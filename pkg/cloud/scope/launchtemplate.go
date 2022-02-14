/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package scope

import (
	"context"
	"encoding/json"

	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/klog/v2/klogr"
	"sigs.k8s.io/controller-runtime/pkg/client"

	infrav1 "sigs.k8s.io/cluster-api-provider-aws/api/v1beta1"
	expinfrav1 "sigs.k8s.io/cluster-api-provider-aws/exp/api/v1beta1"
	expclusterv1 "sigs.k8s.io/cluster-api/exp/api/v1beta1"
	"sigs.k8s.io/cluster-api/util/conditions"
)

const (
	// AWSManagedControlPlaneKind is the Kind of AWSManagedControlPlane.
	AWSManagedControlPlaneKind = "AWSManagedControlPlane"
)

// LaunchTemplateScope defines a scope defined around a launch template.
type LaunchTemplateScope struct {
	logr.Logger
	Client client.Client

	AWSLaunchTemplate             *expinfrav1.AWSLaunchTemplate
	MachinePool                   *expclusterv1.MachinePool
	MachinePoolWithLaunchTemplate MachinePoolWithLaunchTemplate
	InfraCluster                  EC2Scope
	name                          string
	additionalTags                infrav1.Tags
}

type MachinePoolWithLaunchTemplate interface {
	GetObjectMeta() *metav1.ObjectMeta

	GetSetter() conditions.Setter

	PatchObject() error

	GetLaunchTemplateIDStatus() string
	SetLaunchTemplateIDStatus(id string)
	GetLaunchTemplateLatestVersionStatus() string
	SetLaunchTemplateLatestVersionStatus(version string)
}

type ResourceServiceToUpdate struct {
	ResourceID      *string
	ResourceService ResourceService
}

type ResourceService interface {
	UpdateResourceTags(resourceID *string, create, remove map[string]string) error
}

// LaunchTemplateScopeParams defines a scope defined around a launch template.
type LaunchTemplateScopeParams struct {
	Client client.Client
	Logger *logr.Logger

	AWSLaunchTemplate             *expinfrav1.AWSLaunchTemplate
	MachinePool                   *expclusterv1.MachinePool
	MachinePoolWithLaunchTemplate MachinePoolWithLaunchTemplate
	InfraCluster                  EC2Scope
	Name                          string
	AdditionalTags                infrav1.Tags
}

// NewLaunchTemplateScope creates a new LaunchTemplateScope from the supplied parameters.
func NewLaunchTemplateScope(params LaunchTemplateScopeParams) (*LaunchTemplateScope, error) {
	if params.Client == nil {
		return nil, errors.New("client is required when creating a MachinePoolScope")
	}
	if params.MachinePool == nil {
		return nil, errors.New("machinepool is required when creating a LaunchTemplateScope")
	}
	if params.InfraCluster == nil {
		return nil, errors.New("cluster is required when creating a LaunchTemplateScope")
	}

	if params.Logger == nil {
		log := klogr.New()
		params.Logger = &log
	}

	return &LaunchTemplateScope{
		Logger: *params.Logger,
		Client: params.Client,

		AWSLaunchTemplate:             params.AWSLaunchTemplate,
		MachinePool:                   params.MachinePool,
		MachinePoolWithLaunchTemplate: params.MachinePoolWithLaunchTemplate,
		InfraCluster:                  params.InfraCluster,
		name:                          params.Name,
		additionalTags:                params.AdditionalTags,
	}, nil
}

func (m *LaunchTemplateScope) IsEKSManaged() bool {
	return m.InfraCluster.InfraCluster().GetObjectKind().GroupVersionKind().Kind == AWSManagedControlPlaneKind
}

func (m *LaunchTemplateScope) Name() string {
	return m.name
}

func (m *LaunchTemplateScope) AdditionalTags() infrav1.Tags {
	tags := make(infrav1.Tags)

	// Start with the cluster-wide tags...
	tags.Merge(m.InfraCluster.AdditionalTags())
	// ... and merge in the Machine's
	tags.Merge(m.additionalTags)

	return tags
}

func (m *LaunchTemplateScope) GetRawBootstrapData() ([]byte, error) {
	if m.MachinePool.Spec.Template.Spec.Bootstrap.DataSecretName == nil {
		return nil, errors.New("error retrieving bootstrap data: linked Machine's bootstrap.dataSecretName is nil")
	}

	secret := &corev1.Secret{}
	key := types.NamespacedName{Namespace: m.MachinePool.Namespace, Name: *m.MachinePool.Spec.Template.Spec.Bootstrap.DataSecretName}

	if err := m.Client.Get(context.TODO(), key, secret); err != nil {
		return nil, errors.Wrapf(err, "failed to retrieve bootstrap data secret for MachinePool %s/%s", m.MachinePool.Namespace, m.MachinePool.Name)
	}

	value, ok := secret.Data["value"]
	if !ok {
		return nil, errors.New("error retrieving bootstrap data: secret value key is missing")
	}

	return value, nil
}

func (m *LaunchTemplateScope) MachinePoolAnnotationJSON(annotation string) (map[string]interface{}, error) {
	out := map[string]interface{}{}

	jsonAnnotation := m.machinePoolAnnotation(annotation)
	if len(jsonAnnotation) == 0 {
		return out, nil
	}

	err := json.Unmarshal([]byte(jsonAnnotation), &out)
	if err != nil {
		return out, err
	}

	return out, nil
}

func (m *LaunchTemplateScope) machinePoolAnnotation(annotation string) string {
	return m.MachinePoolWithLaunchTemplate.GetObjectMeta().GetAnnotations()[annotation]
}

func (m *LaunchTemplateScope) UpdateMachinePoolAnnotationJSON(annotation string, content map[string]interface{}) error {
	b, err := json.Marshal(content)
	if err != nil {
		return err
	}

	m.updateMachinePoolAnnotation(annotation, string(b))
	return nil
}

func (m *LaunchTemplateScope) updateMachinePoolAnnotation(annotation, content string) {
	// Get the annotations
	annotations := m.MachinePoolWithLaunchTemplate.GetObjectMeta().GetAnnotations()

	if annotations == nil {
		annotations = make(map[string]string)
	}

	// Set our annotation to the given content.
	annotations[annotation] = content

	// Update the machine object with these annotations
	m.MachinePoolWithLaunchTemplate.GetObjectMeta().SetAnnotations(annotations)
}
