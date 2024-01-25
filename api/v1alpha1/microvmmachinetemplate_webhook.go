// Copyright 2024 Liquid Metal Authors. All Rights Reserved.
// SPDX-License-Identifier: MPL-2.0

package v1alpha1

import (
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

var _ = logf.Log.WithName("microvmmachinetemplate-resource")

func (r *MicrovmMachineTemplate) SetupWebhookWithManager(mgr ctrl.Manager) error {
	if err := ctrl.NewWebhookManagedBy(mgr).For(r).Complete(); err != nil {
		return fmt.Errorf("unable to setup webhook with manager: %w", err)
	}

	return nil
}

// +kubebuilder:webhook:verbs=create;update,path=/validate-infrastructure-cluster-x-k8s-io-v1alpha1-microvmmachinetemplate,mutating=false,failurePolicy=fail,matchPolicy=Equivalent,groups=infrastructure.cluster.x-k8s.io,resources=microvmmachinetemplates,versions=v1alpha1,name=validation.microvmmachinetemplate.infrastructure.cluster.x-k8s.io,sideEffects=None,admissionReviewVersions=v1

var _ webhook.Validator = &MicrovmMachineTemplate{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type.
func (r *MicrovmMachineTemplate) ValidateCreate() (admission.Warnings, error) {
	return nil, nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type.
func (r *MicrovmMachineTemplate) ValidateDelete() (admission.Warnings, error) {
	return nil, nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type.
func (r *MicrovmMachineTemplate) ValidateUpdate(old runtime.Object) (admission.Warnings, error) {
	return nil, nil
}
