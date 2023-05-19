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



// DatabaseClusterSpec - DatabaseSpec defines the desired state of Database.
type DatabaseClusterSpec struct {

	Backup DatabaseClusterSpecBackup `json:"backup,omitempty"`

	// ClusterSize is amount of nodes that required for the cluster. A database starts in cluster mode if clusterSize >= 3.
	ClusterSize int32 `json:"clusterSize"`

	// DatabaseConfig contains a config settings for the specified database.
	DatabaseConfig string `json:"databaseConfig"`

	// DatabaseVersion sets from version service and uses the recommended version by default.
	DatabaseImage string `json:"databaseImage"`

	// Database type stands for supported databases by the PMM API Now it's pxc or psmdb types but we can extend it.
	DatabaseType string `json:"databaseType"`

	DbInstance DatabaseClusterSpecDbInstance `json:"dbInstance"`

	LoadBalancer DatabaseClusterSpecLoadBalancer `json:"loadBalancer,omitempty"`

	Monitoring DatabaseClusterSpecMonitoring `json:"monitoring,omitempty"`

	// Pause represents is a cluster paused or not.
	Pause bool `json:"pause,omitempty"`

	// SecretsName contains name of a secrets file for a database cluster.
	SecretsName string `json:"secretsName,omitempty"`
}

// UnmarshalJSON sets *m to a copy of data while respecting defaults if specified.
func (m *DatabaseClusterSpec) UnmarshalJSON(data []byte) error {

	type Alias DatabaseClusterSpec // To avoid infinite recursion
    return json.Unmarshal(data, (*Alias)(m))
}

// AssertDatabaseClusterSpecRequired checks if the required fields are not zero-ed
func AssertDatabaseClusterSpecRequired(obj DatabaseClusterSpec) error {
	elements := map[string]interface{}{
		"clusterSize": obj.ClusterSize,
		"databaseConfig": obj.DatabaseConfig,
		"databaseImage": obj.DatabaseImage,
		"databaseType": obj.DatabaseType,
		"dbInstance": obj.DbInstance,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertDatabaseClusterSpecBackupRequired(obj.Backup); err != nil {
		return err
	}
	if err := AssertDatabaseClusterSpecDbInstanceRequired(obj.DbInstance); err != nil {
		return err
	}
	if err := AssertDatabaseClusterSpecLoadBalancerRequired(obj.LoadBalancer); err != nil {
		return err
	}
	if err := AssertDatabaseClusterSpecMonitoringRequired(obj.Monitoring); err != nil {
		return err
	}
	return nil
}

// AssertDatabaseClusterSpecConstraints checks if the values respects the defined constraints
func AssertDatabaseClusterSpecConstraints(obj DatabaseClusterSpec) error {
	return nil
}