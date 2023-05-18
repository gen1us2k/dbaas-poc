package models

// DatabaseClusterSpecBackupStoragesValueAffinity - Affinity is a group of affinity scheduling rules.
type DatabaseClusterSpecBackupStoragesValueAffinity struct {

	NodeAffinity DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinity `json:"nodeAffinity,omitempty"`

	PodAffinity DatabaseClusterSpecBackupStoragesValueAffinityPodAffinity `json:"podAffinity,omitempty"`

	PodAntiAffinity DatabaseClusterSpecBackupStoragesValueAffinityPodAntiAffinity `json:"podAntiAffinity,omitempty"`
}
