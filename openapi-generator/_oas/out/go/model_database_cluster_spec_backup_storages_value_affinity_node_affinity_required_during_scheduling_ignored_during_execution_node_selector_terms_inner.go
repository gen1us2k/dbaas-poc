/*
Percona Everest schema

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsInner{}

// DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsInner A null or empty node selector term matches no objects. The requirements of them are ANDed. The TopologySelectorTerm type implements a subset of the NodeSelectorTerm.
type DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsInner struct {
	// A list of node selector requirements by node's labels.
	MatchExpressions []DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionInnerPreferenceMatchExpressionsInner `json:"matchExpressions,omitempty"`
	// A list of node selector requirements by node's fields.
	MatchFields []DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionInnerPreferenceMatchExpressionsInner `json:"matchFields,omitempty"`
}

// NewDatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsInner instantiates a new DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsInner() *DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsInner {
	this := DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsInner{}
	return &this
}

// NewDatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsInnerWithDefaults instantiates a new DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsInnerWithDefaults() *DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsInner {
	this := DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsInner{}
	return &this
}

// GetMatchExpressions returns the MatchExpressions field value if set, zero value otherwise.
func (o *DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsInner) GetMatchExpressions() []DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionInnerPreferenceMatchExpressionsInner {
	if o == nil || IsNil(o.MatchExpressions) {
		var ret []DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionInnerPreferenceMatchExpressionsInner
		return ret
	}
	return o.MatchExpressions
}

// GetMatchExpressionsOk returns a tuple with the MatchExpressions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsInner) GetMatchExpressionsOk() ([]DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionInnerPreferenceMatchExpressionsInner, bool) {
	if o == nil || IsNil(o.MatchExpressions) {
		return nil, false
	}
	return o.MatchExpressions, true
}

// HasMatchExpressions returns a boolean if a field has been set.
func (o *DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsInner) HasMatchExpressions() bool {
	if o != nil && !IsNil(o.MatchExpressions) {
		return true
	}

	return false
}

// SetMatchExpressions gets a reference to the given []DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionInnerPreferenceMatchExpressionsInner and assigns it to the MatchExpressions field.
func (o *DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsInner) SetMatchExpressions(v []DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionInnerPreferenceMatchExpressionsInner) {
	o.MatchExpressions = v
}

// GetMatchFields returns the MatchFields field value if set, zero value otherwise.
func (o *DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsInner) GetMatchFields() []DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionInnerPreferenceMatchExpressionsInner {
	if o == nil || IsNil(o.MatchFields) {
		var ret []DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionInnerPreferenceMatchExpressionsInner
		return ret
	}
	return o.MatchFields
}

// GetMatchFieldsOk returns a tuple with the MatchFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsInner) GetMatchFieldsOk() ([]DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionInnerPreferenceMatchExpressionsInner, bool) {
	if o == nil || IsNil(o.MatchFields) {
		return nil, false
	}
	return o.MatchFields, true
}

// HasMatchFields returns a boolean if a field has been set.
func (o *DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsInner) HasMatchFields() bool {
	if o != nil && !IsNil(o.MatchFields) {
		return true
	}

	return false
}

// SetMatchFields gets a reference to the given []DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionInnerPreferenceMatchExpressionsInner and assigns it to the MatchFields field.
func (o *DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsInner) SetMatchFields(v []DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionInnerPreferenceMatchExpressionsInner) {
	o.MatchFields = v
}

func (o DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MatchExpressions) {
		toSerialize["matchExpressions"] = o.MatchExpressions
	}
	if !IsNil(o.MatchFields) {
		toSerialize["matchFields"] = o.MatchFields
	}
	return toSerialize, nil
}

type NullableDatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsInner struct {
	value *DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsInner
	isSet bool
}

func (v NullableDatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsInner) Get() *DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsInner {
	return v.value
}

func (v *NullableDatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsInner) Set(val *DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableDatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableDatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsInner(val *DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsInner) *NullableDatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsInner {
	return &NullableDatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsInner{value: val, isSet: true}
}

func (v NullableDatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


