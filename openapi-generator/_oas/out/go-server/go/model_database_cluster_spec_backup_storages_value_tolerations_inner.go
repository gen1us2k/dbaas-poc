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



// DatabaseClusterSpecBackupStoragesValueTolerationsInner - The pod this Toleration is attached to tolerates any taint that matches the triple <key,value,effect> using the matching operator <operator>.
type DatabaseClusterSpecBackupStoragesValueTolerationsInner struct {

	// Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.
	Effect string `json:"effect,omitempty"`

	// Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.
	Key string `json:"key,omitempty"`

	// Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.
	Operator string `json:"operator,omitempty"`

	// TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.
	TolerationSeconds int64 `json:"tolerationSeconds,omitempty"`

	// Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.
	Value string `json:"value,omitempty"`
}

// UnmarshalJSON sets *m to a copy of data while respecting defaults if specified.
func (m *DatabaseClusterSpecBackupStoragesValueTolerationsInner) UnmarshalJSON(data []byte) error {

	type Alias DatabaseClusterSpecBackupStoragesValueTolerationsInner // To avoid infinite recursion
    return json.Unmarshal(data, (*Alias)(m))
}

// AssertDatabaseClusterSpecBackupStoragesValueTolerationsInnerRequired checks if the required fields are not zero-ed
func AssertDatabaseClusterSpecBackupStoragesValueTolerationsInnerRequired(obj DatabaseClusterSpecBackupStoragesValueTolerationsInner) error {
	return nil
}

// AssertDatabaseClusterSpecBackupStoragesValueTolerationsInnerConstraints checks if the values respects the defined constraints
func AssertDatabaseClusterSpecBackupStoragesValueTolerationsInnerConstraints(obj DatabaseClusterSpecBackupStoragesValueTolerationsInner) error {
	return nil
}