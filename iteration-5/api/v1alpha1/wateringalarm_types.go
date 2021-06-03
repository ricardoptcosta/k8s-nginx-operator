/*


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

// WateringAlarmSpec defines the desired state of WateringAlarm
type WateringAlarmSpec struct {
	Plant string `json:"plant,omitempty"`

        //+kubebuilder:validation:Minimum=0
	TimeInterval int `json:"timeinterval,omitempty"`
}

// WateringAlarmStatus defines the observed state of WateringAlarm
type WateringAlarmStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	LastWateringDate string `json:"lastwateringdate,omitempty"`
	NextWateringDate string `json:"nextwateringdate,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:JSONPath=".spec.plant",name=PLANT,type=string
// +kubebuilder:printcolumn:JSONPath=".spec.timeinterval",name=TIME INTERVAL,type=integer
// +kubebuilder:printcolumn:JSONPath=".status.lastwateringdate",name=LAST WATERING DATE,type=string
// +kubebuilder:printcolumn:JSONPath=".spec.nextwateringdate",name=NEXT WATERING DATE,type=string

// WateringAlarm is the Schema for the wateringalarms API
type WateringAlarm struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WateringAlarmSpec   `json:"spec,omitempty"`
	Status WateringAlarmStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WateringAlarmList contains a list of WateringAlarm
type WateringAlarmList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WateringAlarm `json:"items"`
}

func init() {
	SchemeBuilder.Register(&WateringAlarm{}, &WateringAlarmList{})
}
