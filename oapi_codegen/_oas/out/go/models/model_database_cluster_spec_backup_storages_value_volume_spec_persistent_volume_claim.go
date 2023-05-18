package models

// DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim - PersistentVolumeClaim to specify PVC spec for the volume for mysql data. It has the highest level of precedence, followed by HostPath and EmptyDir. And represents the PVC specification.
type DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim struct {

	// accessModes contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1
	AccessModes []string `json:"accessModes,omitempty"`

	DataSource DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSource `json:"dataSource,omitempty"`

	DataSourceRef DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRef `json:"dataSourceRef,omitempty"`

	Resources DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimResources `json:"resources,omitempty"`

	Selector DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimSelector `json:"selector,omitempty"`

	// storageClassName is the name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1
	StorageClassName string `json:"storageClassName,omitempty"`

	// volumeMode defines what type of volume is required by the claim. Value of Filesystem is implied when not included in claim spec.
	VolumeMode string `json:"volumeMode,omitempty"`

	// volumeName is the binding reference to the PersistentVolume backing this claim.
	VolumeName string `json:"volumeName,omitempty"`
}
