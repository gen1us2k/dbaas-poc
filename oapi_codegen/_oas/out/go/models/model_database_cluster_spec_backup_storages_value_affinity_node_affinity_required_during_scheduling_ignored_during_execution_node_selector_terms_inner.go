package models

// DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsInner - A null or empty node selector term matches no objects. The requirements of them are ANDed. The TopologySelectorTerm type implements a subset of the NodeSelectorTerm.
type DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsInner struct {

	// A list of node selector requirements by node's labels.
	MatchExpressions []DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionInnerPreferenceMatchExpressionsInner `json:"matchExpressions,omitempty"`

	// A list of node selector requirements by node's fields.
	MatchFields []DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionInnerPreferenceMatchExpressionsInner `json:"matchFields,omitempty"`
}
