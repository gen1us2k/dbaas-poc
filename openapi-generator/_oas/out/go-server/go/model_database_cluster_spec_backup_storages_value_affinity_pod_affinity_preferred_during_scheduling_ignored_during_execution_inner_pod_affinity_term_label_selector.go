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



// DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionInnerPodAffinityTermLabelSelector - A label query over a set of resources, in this case pods.
type DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionInnerPodAffinityTermLabelSelector struct {

	// matchExpressions is a list of label selector requirements. The requirements are ANDed.
	MatchExpressions []DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionInnerPodAffinityTermLabelSelectorMatchExpressionsInner `json:"matchExpressions,omitempty"`

	// matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is \"key\", the operator is \"In\", and the values array contains only \"value\". The requirements are ANDed.
	MatchLabels map[string]string `json:"matchLabels,omitempty"`
}

// UnmarshalJSON sets *m to a copy of data while respecting defaults if specified.
func (m *DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionInnerPodAffinityTermLabelSelector) UnmarshalJSON(data []byte) error {

	type Alias DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionInnerPodAffinityTermLabelSelector // To avoid infinite recursion
    return json.Unmarshal(data, (*Alias)(m))
}

// AssertDatabaseClusterSpecBackupStoragesValueAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionInnerPodAffinityTermLabelSelectorRequired checks if the required fields are not zero-ed
func AssertDatabaseClusterSpecBackupStoragesValueAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionInnerPodAffinityTermLabelSelectorRequired(obj DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionInnerPodAffinityTermLabelSelector) error {
	for _, el := range obj.MatchExpressions {
		if err := AssertDatabaseClusterSpecBackupStoragesValueAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionInnerPodAffinityTermLabelSelectorMatchExpressionsInnerRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertDatabaseClusterSpecBackupStoragesValueAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionInnerPodAffinityTermLabelSelectorConstraints checks if the values respects the defined constraints
func AssertDatabaseClusterSpecBackupStoragesValueAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionInnerPodAffinityTermLabelSelectorConstraints(obj DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionInnerPodAffinityTermLabelSelector) error {
	return nil
}
