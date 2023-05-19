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



// DatabaseClusterSpecBackupContainerSecurityContextSeLinuxOptions - The SELinux context to be applied to the container. If unspecified, the container runtime will allocate a random SELinux context for each container.  May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is windows.
type DatabaseClusterSpecBackupContainerSecurityContextSeLinuxOptions struct {

	// Level is SELinux level label that applies to the container.
	Level string `json:"level,omitempty"`

	// Role is a SELinux role label that applies to the container.
	Role string `json:"role,omitempty"`

	// Type is a SELinux type label that applies to the container.
	Type string `json:"type,omitempty"`

	// User is a SELinux user label that applies to the container.
	User string `json:"user,omitempty"`
}

// UnmarshalJSON sets *m to a copy of data while respecting defaults if specified.
func (m *DatabaseClusterSpecBackupContainerSecurityContextSeLinuxOptions) UnmarshalJSON(data []byte) error {

	type Alias DatabaseClusterSpecBackupContainerSecurityContextSeLinuxOptions // To avoid infinite recursion
    return json.Unmarshal(data, (*Alias)(m))
}

// AssertDatabaseClusterSpecBackupContainerSecurityContextSeLinuxOptionsRequired checks if the required fields are not zero-ed
func AssertDatabaseClusterSpecBackupContainerSecurityContextSeLinuxOptionsRequired(obj DatabaseClusterSpecBackupContainerSecurityContextSeLinuxOptions) error {
	return nil
}

// AssertDatabaseClusterSpecBackupContainerSecurityContextSeLinuxOptionsConstraints checks if the values respects the defined constraints
func AssertDatabaseClusterSpecBackupContainerSecurityContextSeLinuxOptionsConstraints(obj DatabaseClusterSpecBackupContainerSecurityContextSeLinuxOptions) error {
	return nil
}
