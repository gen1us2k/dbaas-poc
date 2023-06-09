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

// checks if the DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinity{}

// DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinity Describes node affinity scheduling rules for the pod.
type DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinity struct {
	// The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding \"weight\" to the sum if the node matches the corresponding matchExpressions; the node(s) with the highest sum are the most preferred.
	PreferredDuringSchedulingIgnoredDuringExecution []DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionInner `json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`
	RequiredDuringSchedulingIgnoredDuringExecution *DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecution `json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
}

// NewDatabaseClusterSpecBackupStoragesValueAffinityNodeAffinity instantiates a new DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatabaseClusterSpecBackupStoragesValueAffinityNodeAffinity() *DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinity {
	this := DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinity{}
	return &this
}

// NewDatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityWithDefaults instantiates a new DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityWithDefaults() *DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinity {
	this := DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinity{}
	return &this
}

// GetPreferredDuringSchedulingIgnoredDuringExecution returns the PreferredDuringSchedulingIgnoredDuringExecution field value if set, zero value otherwise.
func (o *DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinity) GetPreferredDuringSchedulingIgnoredDuringExecution() []DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionInner {
	if o == nil || IsNil(o.PreferredDuringSchedulingIgnoredDuringExecution) {
		var ret []DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionInner
		return ret
	}
	return o.PreferredDuringSchedulingIgnoredDuringExecution
}

// GetPreferredDuringSchedulingIgnoredDuringExecutionOk returns a tuple with the PreferredDuringSchedulingIgnoredDuringExecution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinity) GetPreferredDuringSchedulingIgnoredDuringExecutionOk() ([]DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionInner, bool) {
	if o == nil || IsNil(o.PreferredDuringSchedulingIgnoredDuringExecution) {
		return nil, false
	}
	return o.PreferredDuringSchedulingIgnoredDuringExecution, true
}

// HasPreferredDuringSchedulingIgnoredDuringExecution returns a boolean if a field has been set.
func (o *DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinity) HasPreferredDuringSchedulingIgnoredDuringExecution() bool {
	if o != nil && !IsNil(o.PreferredDuringSchedulingIgnoredDuringExecution) {
		return true
	}

	return false
}

// SetPreferredDuringSchedulingIgnoredDuringExecution gets a reference to the given []DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionInner and assigns it to the PreferredDuringSchedulingIgnoredDuringExecution field.
func (o *DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinity) SetPreferredDuringSchedulingIgnoredDuringExecution(v []DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionInner) {
	o.PreferredDuringSchedulingIgnoredDuringExecution = v
}

// GetRequiredDuringSchedulingIgnoredDuringExecution returns the RequiredDuringSchedulingIgnoredDuringExecution field value if set, zero value otherwise.
func (o *DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinity) GetRequiredDuringSchedulingIgnoredDuringExecution() DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecution {
	if o == nil || IsNil(o.RequiredDuringSchedulingIgnoredDuringExecution) {
		var ret DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecution
		return ret
	}
	return *o.RequiredDuringSchedulingIgnoredDuringExecution
}

// GetRequiredDuringSchedulingIgnoredDuringExecutionOk returns a tuple with the RequiredDuringSchedulingIgnoredDuringExecution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinity) GetRequiredDuringSchedulingIgnoredDuringExecutionOk() (*DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecution, bool) {
	if o == nil || IsNil(o.RequiredDuringSchedulingIgnoredDuringExecution) {
		return nil, false
	}
	return o.RequiredDuringSchedulingIgnoredDuringExecution, true
}

// HasRequiredDuringSchedulingIgnoredDuringExecution returns a boolean if a field has been set.
func (o *DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinity) HasRequiredDuringSchedulingIgnoredDuringExecution() bool {
	if o != nil && !IsNil(o.RequiredDuringSchedulingIgnoredDuringExecution) {
		return true
	}

	return false
}

// SetRequiredDuringSchedulingIgnoredDuringExecution gets a reference to the given DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecution and assigns it to the RequiredDuringSchedulingIgnoredDuringExecution field.
func (o *DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinity) SetRequiredDuringSchedulingIgnoredDuringExecution(v DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecution) {
	o.RequiredDuringSchedulingIgnoredDuringExecution = &v
}

func (o DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PreferredDuringSchedulingIgnoredDuringExecution) {
		toSerialize["preferredDuringSchedulingIgnoredDuringExecution"] = o.PreferredDuringSchedulingIgnoredDuringExecution
	}
	if !IsNil(o.RequiredDuringSchedulingIgnoredDuringExecution) {
		toSerialize["requiredDuringSchedulingIgnoredDuringExecution"] = o.RequiredDuringSchedulingIgnoredDuringExecution
	}
	return toSerialize, nil
}

type NullableDatabaseClusterSpecBackupStoragesValueAffinityNodeAffinity struct {
	value *DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinity
	isSet bool
}

func (v NullableDatabaseClusterSpecBackupStoragesValueAffinityNodeAffinity) Get() *DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinity {
	return v.value
}

func (v *NullableDatabaseClusterSpecBackupStoragesValueAffinityNodeAffinity) Set(val *DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinity) {
	v.value = val
	v.isSet = true
}

func (v NullableDatabaseClusterSpecBackupStoragesValueAffinityNodeAffinity) IsSet() bool {
	return v.isSet
}

func (v *NullableDatabaseClusterSpecBackupStoragesValueAffinityNodeAffinity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatabaseClusterSpecBackupStoragesValueAffinityNodeAffinity(val *DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinity) *NullableDatabaseClusterSpecBackupStoragesValueAffinityNodeAffinity {
	return &NullableDatabaseClusterSpecBackupStoragesValueAffinityNodeAffinity{value: val, isSet: true}
}

func (v NullableDatabaseClusterSpecBackupStoragesValueAffinityNodeAffinity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatabaseClusterSpecBackupStoragesValueAffinityNodeAffinity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


