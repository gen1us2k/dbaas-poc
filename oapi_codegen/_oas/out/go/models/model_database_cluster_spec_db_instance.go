package models

// DatabaseClusterSpecDbInstance - DBInstance represents resource requests for a database cluster.
type DatabaseClusterSpecDbInstance struct {

	Cpu DatabaseClusterSpecBackupResourcesLimitsValue `json:"cpu,omitempty"`

	DiskSize DatabaseClusterSpecBackupResourcesLimitsValue `json:"diskSize,omitempty"`

	Memory DatabaseClusterSpecBackupResourcesLimitsValue `json:"memory,omitempty"`

	StorageClassName string `json:"storageClassName,omitempty"`
}
