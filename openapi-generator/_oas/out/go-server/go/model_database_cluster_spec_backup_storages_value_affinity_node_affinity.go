/*
 * Percona Everest schema
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi


import (
	"encoding/json"
)



// DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinity - Describes node affinity scheduling rules for the pod.
type DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinity struct {

	// The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding \"weight\" to the sum if the node matches the corresponding matchExpressions; the node(s) with the highest sum are the most preferred.
	PreferredDuringSchedulingIgnoredDuringExecution []DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionInner `json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`

	RequiredDuringSchedulingIgnoredDuringExecution DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecution `json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
}

// UnmarshalJSON sets *m to a copy of data while respecting defaults if specified.
func (m *DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinity) UnmarshalJSON(data []byte) error {

	type Alias DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinity // To avoid infinite recursion
    return json.Unmarshal(data, (*Alias)(m))
}

// AssertDatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityRequired checks if the required fields are not zero-ed
func AssertDatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityRequired(obj DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinity) error {
	for _, el := range obj.PreferredDuringSchedulingIgnoredDuringExecution {
		if err := AssertDatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionInnerRequired(el); err != nil {
			return err
		}
	}
	if err := AssertDatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionRequired(obj.RequiredDuringSchedulingIgnoredDuringExecution); err != nil {
		return err
	}
	return nil
}

// AssertDatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityConstraints checks if the values respects the defined constraints
func AssertDatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityConstraints(obj DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinity) error {
	return nil
}