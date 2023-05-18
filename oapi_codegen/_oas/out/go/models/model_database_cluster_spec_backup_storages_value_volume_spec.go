package models

// DatabaseClusterSpecBackupStoragesValueVolumeSpec - VolumeSpec represents a specification to configure volume for underlying database.
type DatabaseClusterSpecBackupStoragesValueVolumeSpec struct {

	EmptyDir DatabaseClusterSpecBackupStoragesValueVolumeSpecEmptyDir `json:"emptyDir,omitempty"`

	HostPath DatabaseClusterSpecBackupStoragesValueVolumeSpecHostPath `json:"hostPath,omitempty"`

	PersistentVolumeClaim DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim `json:"persistentVolumeClaim,omitempty"`
}
