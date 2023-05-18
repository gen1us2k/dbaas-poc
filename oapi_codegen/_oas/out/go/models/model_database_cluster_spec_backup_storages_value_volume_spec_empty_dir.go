package models

// DatabaseClusterSpecBackupStoragesValueVolumeSpecEmptyDir - EmptyDir to use as data volume for mysql. EmptyDir represents a temporary directory that shares a pod's lifetime.
type DatabaseClusterSpecBackupStoragesValueVolumeSpecEmptyDir struct {

	// medium represents what type of storage medium should back this directory. The default is \"\" which means to use the node's default medium. Must be an empty string (default) or Memory. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir
	Medium string `json:"medium,omitempty"`

	SizeLimit DatabaseClusterSpecBackupStoragesValueVolumeSpecEmptyDirSizeLimit `json:"sizeLimit,omitempty"`
}
