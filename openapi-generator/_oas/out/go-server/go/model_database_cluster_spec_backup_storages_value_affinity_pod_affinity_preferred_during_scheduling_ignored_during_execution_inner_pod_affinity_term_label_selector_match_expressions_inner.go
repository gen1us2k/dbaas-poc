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



// DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionInnerPodAffinityTermLabelSelectorMatchExpressionsInner - A label selector requirement is a selector that contains values, a key, and an operator that relates the key and values.
type DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionInnerPodAffinityTermLabelSelectorMatchExpressionsInner struct {

	// key is the label key that the selector applies to.
	Key string `json:"key"`

	// operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.
	Operator string `json:"operator"`

	// values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.
	Values []string `json:"values,omitempty"`
}

// UnmarshalJSON sets *m to a copy of data while respecting defaults if specified.
func (m *DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionInnerPodAffinityTermLabelSelectorMatchExpressionsInner) UnmarshalJSON(data []byte) error {

	type Alias DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionInnerPodAffinityTermLabelSelectorMatchExpressionsInner // To avoid infinite recursion
    return json.Unmarshal(data, (*Alias)(m))
}

// AssertDatabaseClusterSpecBackupStoragesValueAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionInnerPodAffinityTermLabelSelectorMatchExpressionsInnerRequired checks if the required fields are not zero-ed
func AssertDatabaseClusterSpecBackupStoragesValueAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionInnerPodAffinityTermLabelSelectorMatchExpressionsInnerRequired(obj DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionInnerPodAffinityTermLabelSelectorMatchExpressionsInner) error {
	elements := map[string]interface{}{
		"key": obj.Key,
		"operator": obj.Operator,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertDatabaseClusterSpecBackupStoragesValueAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionInnerPodAffinityTermLabelSelectorMatchExpressionsInnerConstraints checks if the values respects the defined constraints
func AssertDatabaseClusterSpecBackupStoragesValueAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionInnerPodAffinityTermLabelSelectorMatchExpressionsInnerConstraints(obj DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionInnerPodAffinityTermLabelSelectorMatchExpressionsInner) error {
	return nil
}
