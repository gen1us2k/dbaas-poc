package models

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
