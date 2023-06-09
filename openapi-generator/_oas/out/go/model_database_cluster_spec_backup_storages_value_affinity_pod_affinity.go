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

// checks if the DatabaseClusterSpecBackupStoragesValueAffinityPodAffinity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatabaseClusterSpecBackupStoragesValueAffinityPodAffinity{}

// DatabaseClusterSpecBackupStoragesValueAffinityPodAffinity Describes pod affinity scheduling rules (e.g. co-locate this pod in the same node, zone, etc. as some other pod(s)).
type DatabaseClusterSpecBackupStoragesValueAffinityPodAffinity struct {
	// The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding \"weight\" to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.
	PreferredDuringSchedulingIgnoredDuringExecution []DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionInner `json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`
	// If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.
	RequiredDuringSchedulingIgnoredDuringExecution []DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionInner `json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
}

// NewDatabaseClusterSpecBackupStoragesValueAffinityPodAffinity instantiates a new DatabaseClusterSpecBackupStoragesValueAffinityPodAffinity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatabaseClusterSpecBackupStoragesValueAffinityPodAffinity() *DatabaseClusterSpecBackupStoragesValueAffinityPodAffinity {
	this := DatabaseClusterSpecBackupStoragesValueAffinityPodAffinity{}
	return &this
}

// NewDatabaseClusterSpecBackupStoragesValueAffinityPodAffinityWithDefaults instantiates a new DatabaseClusterSpecBackupStoragesValueAffinityPodAffinity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatabaseClusterSpecBackupStoragesValueAffinityPodAffinityWithDefaults() *DatabaseClusterSpecBackupStoragesValueAffinityPodAffinity {
	this := DatabaseClusterSpecBackupStoragesValueAffinityPodAffinity{}
	return &this
}

// GetPreferredDuringSchedulingIgnoredDuringExecution returns the PreferredDuringSchedulingIgnoredDuringExecution field value if set, zero value otherwise.
func (o *DatabaseClusterSpecBackupStoragesValueAffinityPodAffinity) GetPreferredDuringSchedulingIgnoredDuringExecution() []DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionInner {
	if o == nil || IsNil(o.PreferredDuringSchedulingIgnoredDuringExecution) {
		var ret []DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionInner
		return ret
	}
	return o.PreferredDuringSchedulingIgnoredDuringExecution
}

// GetPreferredDuringSchedulingIgnoredDuringExecutionOk returns a tuple with the PreferredDuringSchedulingIgnoredDuringExecution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseClusterSpecBackupStoragesValueAffinityPodAffinity) GetPreferredDuringSchedulingIgnoredDuringExecutionOk() ([]DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionInner, bool) {
	if o == nil || IsNil(o.PreferredDuringSchedulingIgnoredDuringExecution) {
		return nil, false
	}
	return o.PreferredDuringSchedulingIgnoredDuringExecution, true
}

// HasPreferredDuringSchedulingIgnoredDuringExecution returns a boolean if a field has been set.
func (o *DatabaseClusterSpecBackupStoragesValueAffinityPodAffinity) HasPreferredDuringSchedulingIgnoredDuringExecution() bool {
	if o != nil && !IsNil(o.PreferredDuringSchedulingIgnoredDuringExecution) {
		return true
	}

	return false
}

// SetPreferredDuringSchedulingIgnoredDuringExecution gets a reference to the given []DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionInner and assigns it to the PreferredDuringSchedulingIgnoredDuringExecution field.
func (o *DatabaseClusterSpecBackupStoragesValueAffinityPodAffinity) SetPreferredDuringSchedulingIgnoredDuringExecution(v []DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionInner) {
	o.PreferredDuringSchedulingIgnoredDuringExecution = v
}

// GetRequiredDuringSchedulingIgnoredDuringExecution returns the RequiredDuringSchedulingIgnoredDuringExecution field value if set, zero value otherwise.
func (o *DatabaseClusterSpecBackupStoragesValueAffinityPodAffinity) GetRequiredDuringSchedulingIgnoredDuringExecution() []DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionInner {
	if o == nil || IsNil(o.RequiredDuringSchedulingIgnoredDuringExecution) {
		var ret []DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionInner
		return ret
	}
	return o.RequiredDuringSchedulingIgnoredDuringExecution
}

// GetRequiredDuringSchedulingIgnoredDuringExecutionOk returns a tuple with the RequiredDuringSchedulingIgnoredDuringExecution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseClusterSpecBackupStoragesValueAffinityPodAffinity) GetRequiredDuringSchedulingIgnoredDuringExecutionOk() ([]DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionInner, bool) {
	if o == nil || IsNil(o.RequiredDuringSchedulingIgnoredDuringExecution) {
		return nil, false
	}
	return o.RequiredDuringSchedulingIgnoredDuringExecution, true
}

// HasRequiredDuringSchedulingIgnoredDuringExecution returns a boolean if a field has been set.
func (o *DatabaseClusterSpecBackupStoragesValueAffinityPodAffinity) HasRequiredDuringSchedulingIgnoredDuringExecution() bool {
	if o != nil && !IsNil(o.RequiredDuringSchedulingIgnoredDuringExecution) {
		return true
	}

	return false
}

// SetRequiredDuringSchedulingIgnoredDuringExecution gets a reference to the given []DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionInner and assigns it to the RequiredDuringSchedulingIgnoredDuringExecution field.
func (o *DatabaseClusterSpecBackupStoragesValueAffinityPodAffinity) SetRequiredDuringSchedulingIgnoredDuringExecution(v []DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionInner) {
	o.RequiredDuringSchedulingIgnoredDuringExecution = v
}

func (o DatabaseClusterSpecBackupStoragesValueAffinityPodAffinity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatabaseClusterSpecBackupStoragesValueAffinityPodAffinity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PreferredDuringSchedulingIgnoredDuringExecution) {
		toSerialize["preferredDuringSchedulingIgnoredDuringExecution"] = o.PreferredDuringSchedulingIgnoredDuringExecution
	}
	if !IsNil(o.RequiredDuringSchedulingIgnoredDuringExecution) {
		toSerialize["requiredDuringSchedulingIgnoredDuringExecution"] = o.RequiredDuringSchedulingIgnoredDuringExecution
	}
	return toSerialize, nil
}

type NullableDatabaseClusterSpecBackupStoragesValueAffinityPodAffinity struct {
	value *DatabaseClusterSpecBackupStoragesValueAffinityPodAffinity
	isSet bool
}

func (v NullableDatabaseClusterSpecBackupStoragesValueAffinityPodAffinity) Get() *DatabaseClusterSpecBackupStoragesValueAffinityPodAffinity {
	return v.value
}

func (v *NullableDatabaseClusterSpecBackupStoragesValueAffinityPodAffinity) Set(val *DatabaseClusterSpecBackupStoragesValueAffinityPodAffinity) {
	v.value = val
	v.isSet = true
}

func (v NullableDatabaseClusterSpecBackupStoragesValueAffinityPodAffinity) IsSet() bool {
	return v.isSet
}

func (v *NullableDatabaseClusterSpecBackupStoragesValueAffinityPodAffinity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatabaseClusterSpecBackupStoragesValueAffinityPodAffinity(val *DatabaseClusterSpecBackupStoragesValueAffinityPodAffinity) *NullableDatabaseClusterSpecBackupStoragesValueAffinityPodAffinity {
	return &NullableDatabaseClusterSpecBackupStoragesValueAffinityPodAffinity{value: val, isSet: true}
}

func (v NullableDatabaseClusterSpecBackupStoragesValueAffinityPodAffinity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatabaseClusterSpecBackupStoragesValueAffinityPodAffinity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


