/*
 * Percona Everest schema
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi


import (
	"encoding/json"
)



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

// UnmarshalJSON sets *m to a copy of data while respecting defaults if specified.
func (m *DatabaseCluster) UnmarshalJSON(data []byte) error {

	type Alias DatabaseCluster // To avoid infinite recursion
    return json.Unmarshal(data, (*Alias)(m))
}

// AssertDatabaseClusterRequired checks if the required fields are not zero-ed
func AssertDatabaseClusterRequired(obj DatabaseCluster) error {
	if err := AssertDatabaseClusterSpecRequired(obj.Spec); err != nil {
		return err
	}
	if err := AssertDatabaseClusterStatusRequired(obj.Status); err != nil {
		return err
	}
	return nil
}

// AssertDatabaseClusterConstraints checks if the values respects the defined constraints
func AssertDatabaseClusterConstraints(obj DatabaseCluster) error {
	return nil
}
