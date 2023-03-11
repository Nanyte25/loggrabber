/*
Copyright 2023.

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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// LoggrabberSpec defines the desired state of Loggrabber
type LoggrabberSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of Loggrabber. Edit loggrabber_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// LoggrabberStatus defines the observed state of Loggrabber
type LoggrabberStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Loggrabber is the Schema for the loggrabbers API
type Loggrabber struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LoggrabberSpec   `json:"spec,omitempty"`
	Status LoggrabberStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// LoggrabberList contains a list of Loggrabber
type LoggrabberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Loggrabber `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Loggrabber{}, &LoggrabberList{})
}
