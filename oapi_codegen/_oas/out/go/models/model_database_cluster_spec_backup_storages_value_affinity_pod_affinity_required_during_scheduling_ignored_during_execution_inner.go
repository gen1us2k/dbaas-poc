package models

// DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionInner - Defines a set of pods (namely those matching the labelSelector relative to the given namespace(s)) that this pod should be co-located (affinity) or not co-located (anti-affinity) with, where co-located is defined as running on a node whose value of the label with key <topologyKey> matches that of any node on which a pod of the set of pods is running
type DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionInner struct {

	LabelSelector DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionInnerPodAffinityTermLabelSelector `json:"labelSelector,omitempty"`

	NamespaceSelector DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionInnerPodAffinityTermNamespaceSelector `json:"namespaceSelector,omitempty"`

	// namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means \"this pod's namespace\".
	Namespaces []string `json:"namespaces,omitempty"`

	// This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.
	TopologyKey string `json:"topologyKey"`
}
