package models

// DatabaseClusterSpecBackupStoragesValueVolumeSpecEmptyDirSizeLimit - sizeLimit is the total amount of local storage required for this EmptyDir volume. The size limit is also applicable for memory medium. The maximum usage on memory medium EmptyDir would be the minimum value between the SizeLimit specified here and the sum of memory limits of all containers in a pod. The default is nil which means that the limit is undefined. More info: http://kubernetes.io/docs/user-guide/volumes#emptydir
type DatabaseClusterSpecBackupStoragesValueVolumeSpecEmptyDirSizeLimit struct {
}
