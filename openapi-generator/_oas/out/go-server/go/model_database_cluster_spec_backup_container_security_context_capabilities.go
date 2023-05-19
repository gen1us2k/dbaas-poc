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



// DatabaseClusterSpecBackupContainerSecurityContextCapabilities - The capabilities to add/drop when running containers. Defaults to the default set of capabilities granted by the container runtime. Note that this field cannot be set when spec.os.name is windows.
type DatabaseClusterSpecBackupContainerSecurityContextCapabilities struct {

	// Added capabilities
	Add []string `json:"add,omitempty"`

	// Removed capabilities
	Drop []string `json:"drop,omitempty"`
}

// UnmarshalJSON sets *m to a copy of data while respecting defaults if specified.
func (m *DatabaseClusterSpecBackupContainerSecurityContextCapabilities) UnmarshalJSON(data []byte) error {

	type Alias DatabaseClusterSpecBackupContainerSecurityContextCapabilities // To avoid infinite recursion
    return json.Unmarshal(data, (*Alias)(m))
}

// AssertDatabaseClusterSpecBackupContainerSecurityContextCapabilitiesRequired checks if the required fields are not zero-ed
func AssertDatabaseClusterSpecBackupContainerSecurityContextCapabilitiesRequired(obj DatabaseClusterSpecBackupContainerSecurityContextCapabilities) error {
	return nil
}

// AssertDatabaseClusterSpecBackupContainerSecurityContextCapabilitiesConstraints checks if the values respects the defined constraints
func AssertDatabaseClusterSpecBackupContainerSecurityContextCapabilitiesConstraints(obj DatabaseClusterSpecBackupContainerSecurityContextCapabilities) error {
	return nil
}
