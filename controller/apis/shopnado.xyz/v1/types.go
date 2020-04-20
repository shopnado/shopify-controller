package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Event struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EventSpec   `json:"spec"`
	Status EventStatus `json:"status"`
}

type EventSpec struct {
	Test string
}

type EventStatus struct {
	Topic   string
	Payload string
}
