package models

// DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionInner - An empty preferred scheduling term matches all objects with implicit weight 0 (i.e. it's a no-op). A null preferred scheduling term matches no objects (i.e. is also a no-op).
type DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionInner struct {

	Preference DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionInnerPreference `json:"preference"`

	// Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.
	Weight int32 `json:"weight"`
}
