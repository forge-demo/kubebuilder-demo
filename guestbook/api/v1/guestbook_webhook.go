/*
Copyright 2022.

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

package v1

import (
	"strings"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var guestbooklog = logf.Log.WithName("guestbook-resource")

func (r *Guestbook) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

//+kubebuilder:webhook:path=/mutate-webapp-forge-dev-v1-guestbook,mutating=true,failurePolicy=fail,sideEffects=None,groups=webapp.forge.dev,resources=guestbooks,verbs=create;update,versions=v1,name=mguestbook.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &Guestbook{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *Guestbook) Default() {
	guestbooklog.Info("default", "name", r.Name)

	// TODO(user): fill in your defaulting logic.
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-webapp-forge-dev-v1-guestbook,mutating=false,failurePolicy=fail,sideEffects=None,groups=webapp.forge.dev,resources=guestbooks,verbs=create;update,versions=v1,name=vguestbook.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &Guestbook{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *Guestbook) ValidateCreate() error {
	guestbooklog.Info("validate create", "name", r.Name)

	return r.validateGuestbook()
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *Guestbook) ValidateUpdate(old runtime.Object) error {
	guestbooklog.Info("validate update", "name", r.Name)

	return r.validateGuestbook()
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *Guestbook) ValidateDelete() error {
	guestbooklog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}

func (r *Guestbook) validateGuestbook() error {
	var allErrs field.ErrorList

	if r.Spec.Name == "" {
		allErrs = append(
			allErrs,
			field.Invalid(
				field.NewPath("spec").Child("name"),
				r.Spec.Name, "name is required",
			),
		)
	}

	if r.Spec.ReservationNumber == "" {
		allErrs = append(
			allErrs,
			field.Invalid(
				field.NewPath("spec").Child("reservationNumber"),
				r.Spec.ReservationNumber, "reservationNumber is required",
			),
		)
	} else {
		if !strings.HasPrefix(r.Spec.ReservationNumber, "room-") {
			allErrs = append(
				allErrs,
				field.Invalid(
					field.NewPath("spec").Child("reservationNumber"),
					r.Spec.ReservationNumber, "reservationNumber must start with room-",
				),
			)
		}
	}

	return apierrors.NewInvalid(
		schema.GroupKind{Group: "webapp.forge.dev", Kind: "Guestbook"},
		r.Name, allErrs,
	)
}
