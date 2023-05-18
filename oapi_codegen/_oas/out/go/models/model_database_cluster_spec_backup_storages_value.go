package models

// DatabaseClusterSpecBackupStoragesValue - BackupStorageSpec represents set of settings to configure backup storage.
type DatabaseClusterSpecBackupStoragesValue struct {

	Affinity DatabaseClusterSpecBackupStoragesValueAffinity `json:"affinity,omitempty"`

	Annotations map[string]string `json:"annotations,omitempty"`

	ContainerSecurityContext DatabaseClusterSpecBackupContainerSecurityContext `json:"containerSecurityContext,omitempty"`

	Labels map[string]string `json:"labels,omitempty"`

	NodeSelector map[string]string `json:"nodeSelector,omitempty"`

	PodSecurityContext DatabaseClusterSpecBackupStoragesValuePodSecurityContext `json:"podSecurityContext,omitempty"`

	PriorityClassName string `json:"priorityClassName,omitempty"`

	Resources DatabaseClusterSpecBackupResources `json:"resources,omitempty"`

	RuntimeClassName string `json:"runtimeClassName,omitempty"`

	SchedulerName string `json:"schedulerName,omitempty"`

	StorageProvider DatabaseClusterSpecBackupStoragesValueStorageProvider `json:"storageProvider,omitempty"`

	Tolerations []DatabaseClusterSpecBackupStoragesValueTolerationsInner `json:"tolerations,omitempty"`

	// BackupStorageType represents backup storage type.
	Type string `json:"type"`

	VerifyTLS bool `json:"verifyTLS,omitempty"`

	VolumeSpec DatabaseClusterSpecBackupStoragesValueVolumeSpec `json:"volumeSpec,omitempty"`
}
