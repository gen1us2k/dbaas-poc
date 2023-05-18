package models

// DatabaseClusterStatus - DatabaseClusterStatus defines the observed state of Database.
type DatabaseClusterStatus struct {

	Host string `json:"host,omitempty"`

	Message string `json:"message,omitempty"`

	Ready int32 `json:"ready,omitempty"`

	Size int32 `json:"size,omitempty"`

	// AppState is used to represent cluster's state.
	Status string `json:"status,omitempty"`
}
