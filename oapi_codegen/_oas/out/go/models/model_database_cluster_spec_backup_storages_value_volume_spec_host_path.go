package models

// DatabaseClusterSpecBackupStoragesValueVolumeSpecHostPath - HostPath to use as data volume for mysql. HostPath represents a pre-existing file or directory on the host machine that is directly exposed to the container.
type DatabaseClusterSpecBackupStoragesValueVolumeSpecHostPath struct {

	// path of the directory on the host. If the path is a symlink, it will follow the link to the real path. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath
	Path string `json:"path"`

	// type for HostPath Volume Defaults to \"\" More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath
	Type string `json:"type,omitempty"`
}
