/*
Copyright 2020 The Kubernetes Authors.

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

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PacketMachineTemplateSpec defines the desired state of PacketMachineTemplate.
type PacketMachineTemplateSpec struct {
	Template PacketMachineTemplateResource `json:"template"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=packetmachinetemplates,shortName=pmt,scope=Namespaced,categories=cluster-api
// +kubebuilder:storageversion

// PacketMachineTemplate is the Schema for the packetmachinetemplates API.
type PacketMachineTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec PacketMachineTemplateSpec `json:"spec,omitempty"`
}

// +kubebuilder:object:root=true

// PacketMachineTemplateList contains a list of PacketMachineTemplate.
type PacketMachineTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PacketMachineTemplate `json:"items"`
}

func init() {
	objectTypes = append(objectTypes, &PacketMachineTemplate{}, &PacketMachineTemplateList{})
}
