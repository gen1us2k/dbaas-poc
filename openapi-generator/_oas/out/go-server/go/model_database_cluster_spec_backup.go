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



// DatabaseClusterSpecBackup - Backup contains backup settings.
type DatabaseClusterSpecBackup struct {

	Annotations map[string]string `json:"annotations,omitempty"`

	ContainerSecurityContext DatabaseClusterSpecBackupContainerSecurityContext `json:"containerSecurityContext,omitempty"`

	Enabled bool `json:"enabled,omitempty"`

	Image string `json:"image,omitempty"`

	// PullPolicy describes a policy for if/when to pull a container image
	ImagePullPolicy string `json:"imagePullPolicy,omitempty"`

	ImagePullSecrets []DatabaseClusterSpecBackupImagePullSecretsInner `json:"imagePullSecrets,omitempty"`

	InitImage string `json:"initImage,omitempty"`

	Labels map[string]string `json:"labels,omitempty"`

	Resources DatabaseClusterSpecBackupResources `json:"resources,omitempty"`

	Schedule []DatabaseClusterSpecBackupScheduleInner `json:"schedule,omitempty"`

	ServiceAccountName string `json:"serviceAccountName,omitempty"`

	Storages map[string]DatabaseClusterSpecBackupStoragesValue `json:"storages,omitempty"`
}

// UnmarshalJSON sets *m to a copy of data while respecting defaults if specified.
func (m *DatabaseClusterSpecBackup) UnmarshalJSON(data []byte) error {

	type Alias DatabaseClusterSpecBackup // To avoid infinite recursion
    return json.Unmarshal(data, (*Alias)(m))
}

// AssertDatabaseClusterSpecBackupRequired checks if the required fields are not zero-ed
func AssertDatabaseClusterSpecBackupRequired(obj DatabaseClusterSpecBackup) error {
	if err := AssertDatabaseClusterSpecBackupContainerSecurityContextRequired(obj.ContainerSecurityContext); err != nil {
		return err
	}
	for _, el := range obj.ImagePullSecrets {
		if err := AssertDatabaseClusterSpecBackupImagePullSecretsInnerRequired(el); err != nil {
			return err
		}
	}
	if err := AssertDatabaseClusterSpecBackupResourcesRequired(obj.Resources); err != nil {
		return err
	}
	for _, el := range obj.Schedule {
		if err := AssertDatabaseClusterSpecBackupScheduleInnerRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertDatabaseClusterSpecBackupConstraints checks if the values respects the defined constraints
func AssertDatabaseClusterSpecBackupConstraints(obj DatabaseClusterSpecBackup) error {
	return nil
}
