package models

// DatabaseCluster - DatabaseCluster is the Schema for the databases API.
type DatabaseCluster struct {

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion string `json:"apiVersion,omitempty"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind string `json:"kind,omitempty"`

	Metadata map[string]interface{} `json:"metadata,omitempty"`

	Spec DatabaseClusterSpec `json:"spec,omitempty"`

	Status DatabaseClusterStatus `json:"status,omitempty"`
}
