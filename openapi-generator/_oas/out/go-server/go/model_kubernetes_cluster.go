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



// KubernetesCluster - kubernetes object
type KubernetesCluster struct {

	Name string `json:"name,omitempty"`
}

// UnmarshalJSON sets *m to a copy of data while respecting defaults if specified.
func (m *KubernetesCluster) UnmarshalJSON(data []byte) error {

	type Alias KubernetesCluster // To avoid infinite recursion
    return json.Unmarshal(data, (*Alias)(m))
}

// AssertKubernetesClusterRequired checks if the required fields are not zero-ed
func AssertKubernetesClusterRequired(obj KubernetesCluster) error {
	return nil
}

// AssertKubernetesClusterConstraints checks if the values respects the defined constraints
func AssertKubernetesClusterConstraints(obj KubernetesCluster) error {
	return nil
}
