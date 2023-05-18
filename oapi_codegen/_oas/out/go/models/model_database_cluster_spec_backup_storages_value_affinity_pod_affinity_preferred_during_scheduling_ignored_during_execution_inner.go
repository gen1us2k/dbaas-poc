package models

// DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionInner - The weights of all of the matched WeightedPodAffinityTerm fields are added per-node to find the most preferred node(s)
type DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionInner struct {

	PodAffinityTerm DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionInnerPodAffinityTerm `json:"podAffinityTerm"`

	// weight associated with matching the corresponding podAffinityTerm, in the range 1-100.
	Weight int32 `json:"weight"`
}
