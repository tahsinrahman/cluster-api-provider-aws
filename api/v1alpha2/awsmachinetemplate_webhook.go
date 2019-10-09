/*
Copyright 2019 The Kubernetes Authors.

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

package v1alpha2

import (
	"errors"
	"reflect"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

func (r *AWSMachineTemplate) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// +kubebuilder:webhook:verbs=create;update,path=/validate-infrastructure-cluster-x-k8s-io-v1alpha2-awsmachinetemplate,mutating=false,failurePolicy=fail,groups=infrastructure.cluster.x-k8s.io,resources=awsmachinetemplates,versions=v1alpha2,name=vawsmachinetemplate.kb.io

var _ webhook.Validator = &AWSMachineTemplate{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *AWSMachineTemplate) ValidateCreate() error {
	return errors.New("hello")
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *AWSMachineTemplate) ValidateUpdate(old runtime.Object) error {
	oldAWSMachineTemplate := old.(*AWSMachineTemplate)
	return validateUpdate(r, oldAWSMachineTemplate)
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *AWSMachineTemplate) ValidateDelete() error {
	return nil
}

// validateUpdate checks that AWSMachineTemplate.Spec is not updated
func validateUpdate(newAWSMachineTemplate *AWSMachineTemplate, oldAWSMachineTemplate *AWSMachineTemplate) error {
	if !reflect.DeepEqual(newAWSMachineTemplate.Spec, oldAWSMachineTemplate.Spec) {
		return errors.New("awsMachineTemplateSpec is immutable")
	}

	return nil
}
