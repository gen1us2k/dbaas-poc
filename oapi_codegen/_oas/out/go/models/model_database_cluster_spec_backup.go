package models

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
