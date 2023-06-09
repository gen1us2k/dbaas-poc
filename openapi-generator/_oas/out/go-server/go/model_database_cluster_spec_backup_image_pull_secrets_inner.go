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



// DatabaseClusterSpecBackupImagePullSecretsInner - LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace.
type DatabaseClusterSpecBackupImagePullSecretsInner struct {

	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?
	Name string `json:"name,omitempty"`
}

// UnmarshalJSON sets *m to a copy of data while respecting defaults if specified.
func (m *DatabaseClusterSpecBackupImagePullSecretsInner) UnmarshalJSON(data []byte) error {

	type Alias DatabaseClusterSpecBackupImagePullSecretsInner // To avoid infinite recursion
    return json.Unmarshal(data, (*Alias)(m))
}

// AssertDatabaseClusterSpecBackupImagePullSecretsInnerRequired checks if the required fields are not zero-ed
func AssertDatabaseClusterSpecBackupImagePullSecretsInnerRequired(obj DatabaseClusterSpecBackupImagePullSecretsInner) error {
	return nil
}

// AssertDatabaseClusterSpecBackupImagePullSecretsInnerConstraints checks if the values respects the defined constraints
func AssertDatabaseClusterSpecBackupImagePullSecretsInnerConstraints(obj DatabaseClusterSpecBackupImagePullSecretsInner) error {
	return nil
}
