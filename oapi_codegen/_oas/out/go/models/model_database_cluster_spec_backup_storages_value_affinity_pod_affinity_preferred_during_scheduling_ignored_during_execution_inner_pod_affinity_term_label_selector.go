package models

// DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionInnerPodAffinityTermLabelSelector - A label query over a set of resources, in this case pods.
type DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionInnerPodAffinityTermLabelSelector struct {

	// matchExpressions is a list of label selector requirements. The requirements are ANDed.
	MatchExpressions []DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionInnerPodAffinityTermLabelSelectorMatchExpressionsInner `json:"matchExpressions,omitempty"`

	// matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is \"key\", the operator is \"In\", and the values array contains only \"value\". The requirements are ANDed.
	MatchLabels map[string]string `json:"matchLabels,omitempty"`
}
