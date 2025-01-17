/*
Copyright 2019 The Knative Authors.

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

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"knative.dev/pkg/apis"
	duckv1beta1 "knative.dev/pkg/apis/duck/v1beta1"
)

var serverlessServiceCondSet = apis.NewLivingConditionSet(
	ServerlessServiceConditionEndspointsPopulated,
)

// GetGroupVersionKind returns the GVK for the ServerlessService.
func (ss *ServerlessService) GetGroupVersionKind() schema.GroupVersionKind {
	return SchemeGroupVersion.WithKind("ServerlessService")
}

// GetCondition returns the value of the condition `t`.
func (sss *ServerlessServiceStatus) GetCondition(t apis.ConditionType) *apis.Condition {
	return serverlessServiceCondSet.Manage(sss).GetCondition(t)
}

// InitializeConditions initializes the conditions.
func (sss *ServerlessServiceStatus) InitializeConditions() {
	serverlessServiceCondSet.Manage(sss).InitializeConditions()
}

// MarkEndpointsReady marks the ServerlessServiceStatus endpoints populated condition to true.
func (sss *ServerlessServiceStatus) MarkEndpointsReady() {
	serverlessServiceCondSet.Manage(sss).MarkTrue(ServerlessServiceConditionEndspointsPopulated)
}

// MarkEndpointsNotOwned marks that we don't own K8s service.
func (sss *ServerlessServiceStatus) MarkEndpointsNotOwned(kind, name string) {
	serverlessServiceCondSet.Manage(sss).MarkFalse(
		ServerlessServiceConditionEndspointsPopulated, "NotOwned",
		"Resource %s of type %s is not owned by SKS", name, kind)
}

// MarkActivatorEndpointsPopulated is setting the ActivatorEndpointsPopulated to True.
func (sss *ServerlessServiceStatus) MarkActivatorEndpointsPopulated() {
	serverlessServiceCondSet.Manage(sss).SetCondition(apis.Condition{
		Type:     ActivatorEndpointsPopulated,
		Status:   corev1.ConditionTrue,
		Severity: apis.ConditionSeverityInfo,
		Reason:   "ActivatorEndpointsPopulated",
		Message:  "Revision is backed by Activator",
	})
}

// MarkActivatorEndpointsRemoved is setting the ActivatorEndpointsPopulated to False.
func (sss *ServerlessServiceStatus) MarkActivatorEndpointsRemoved() {
	serverlessServiceCondSet.Manage(sss).SetCondition(apis.Condition{
		Type:     ActivatorEndpointsPopulated,
		Status:   corev1.ConditionFalse,
		Severity: apis.ConditionSeverityInfo,
		Reason:   "ActivatorEndpointsPopulated",
		Message:  "Revision is backed by Activator",
	})
}

// MarkEndpointsNotReady marks the ServerlessServiceStatus endpoints populated condition to unknown.
func (sss *ServerlessServiceStatus) MarkEndpointsNotReady(reason string) {
	serverlessServiceCondSet.Manage(sss).MarkUnknown(
		ServerlessServiceConditionEndspointsPopulated, reason,
		"K8s Service is not ready")
}

// IsReady returns true if ServerlessService is ready.
func (sss *ServerlessServiceStatus) IsReady() bool {
	return serverlessServiceCondSet.Manage(sss).IsHappy()
}

func (sss *ServerlessServiceStatus) duck() *duckv1beta1.Status {
	return &sss.Status
}
