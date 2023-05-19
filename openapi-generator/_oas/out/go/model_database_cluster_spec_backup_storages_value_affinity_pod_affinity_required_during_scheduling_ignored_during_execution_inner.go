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

// checks if the DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionInner{}

// DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionInner Defines a set of pods (namely those matching the labelSelector relative to the given namespace(s)) that this pod should be co-located (affinity) or not co-located (anti-affinity) with, where co-located is defined as running on a node whose value of the label with key <topologyKey> matches that of any node on which a pod of the set of pods is running
type DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionInner struct {
	LabelSelector *DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionInnerPodAffinityTermLabelSelector `json:"labelSelector,omitempty"`
	NamespaceSelector *DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionInnerPodAffinityTermNamespaceSelector `json:"namespaceSelector,omitempty"`
	// namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means \"this pod's namespace\".
	Namespaces []string `json:"namespaces,omitempty"`
	// This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.
	TopologyKey string `json:"topologyKey"`
}

// NewDatabaseClusterSpecBackupStoragesValueAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionInner instantiates a new DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatabaseClusterSpecBackupStoragesValueAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionInner(topologyKey string) *DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionInner {
	this := DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionInner{}
	this.TopologyKey = topologyKey
	return &this
}

// NewDatabaseClusterSpecBackupStoragesValueAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionInnerWithDefaults instantiates a new DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatabaseClusterSpecBackupStoragesValueAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionInnerWithDefaults() *DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionInner {
	this := DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionInner{}
	return &this
}

// GetLabelSelector returns the LabelSelector field value if set, zero value otherwise.
func (o *DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionInner) GetLabelSelector() DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionInnerPodAffinityTermLabelSelector {
	if o == nil || IsNil(o.LabelSelector) {
		var ret DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionInnerPodAffinityTermLabelSelector
		return ret
	}
	return *o.LabelSelector
}

// GetLabelSelectorOk returns a tuple with the LabelSelector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionInner) GetLabelSelectorOk() (*DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionInnerPodAffinityTermLabelSelector, bool) {
	if o == nil || IsNil(o.LabelSelector) {
		return nil, false
	}
	return o.LabelSelector, true
}

// HasLabelSelector returns a boolean if a field has been set.
func (o *DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionInner) HasLabelSelector() bool {
	if o != nil && !IsNil(o.LabelSelector) {
		return true
	}

	return false
}

// SetLabelSelector gets a reference to the given DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionInnerPodAffinityTermLabelSelector and assigns it to the LabelSelector field.
func (o *DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionInner) SetLabelSelector(v DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionInnerPodAffinityTermLabelSelector) {
	o.LabelSelector = &v
}

// GetNamespaceSelector returns the NamespaceSelector field value if set, zero value otherwise.
func (o *DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionInner) GetNamespaceSelector() DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionInnerPodAffinityTermNamespaceSelector {
	if o == nil || IsNil(o.NamespaceSelector) {
		var ret DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionInnerPodAffinityTermNamespaceSelector
		return ret
	}
	return *o.NamespaceSelector
}

// GetNamespaceSelectorOk returns a tuple with the NamespaceSelector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionInner) GetNamespaceSelectorOk() (*DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionInnerPodAffinityTermNamespaceSelector, bool) {
	if o == nil || IsNil(o.NamespaceSelector) {
		return nil, false
	}
	return o.NamespaceSelector, true
}

// HasNamespaceSelector returns a boolean if a field has been set.
func (o *DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionInner) HasNamespaceSelector() bool {
	if o != nil && !IsNil(o.NamespaceSelector) {
		return true
	}

	return false
}

// SetNamespaceSelector gets a reference to the given DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionInnerPodAffinityTermNamespaceSelector and assigns it to the NamespaceSelector field.
func (o *DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionInner) SetNamespaceSelector(v DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionInnerPodAffinityTermNamespaceSelector) {
	o.NamespaceSelector = &v
}

// GetNamespaces returns the Namespaces field value if set, zero value otherwise.
func (o *DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionInner) GetNamespaces() []string {
	if o == nil || IsNil(o.Namespaces) {
		var ret []string
		return ret
	}
	return o.Namespaces
}

// GetNamespacesOk returns a tuple with the Namespaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionInner) GetNamespacesOk() ([]string, bool) {
	if o == nil || IsNil(o.Namespaces) {
		return nil, false
	}
	return o.Namespaces, true
}

// HasNamespaces returns a boolean if a field has been set.
func (o *DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionInner) HasNamespaces() bool {
	if o != nil && !IsNil(o.Namespaces) {
		return true
	}

	return false
}

// SetNamespaces gets a reference to the given []string and assigns it to the Namespaces field.
func (o *DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionInner) SetNamespaces(v []string) {
	o.Namespaces = v
}

// GetTopologyKey returns the TopologyKey field value
func (o *DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionInner) GetTopologyKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TopologyKey
}

// GetTopologyKeyOk returns a tuple with the TopologyKey field value
// and a boolean to check if the value has been set.
func (o *DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionInner) GetTopologyKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TopologyKey, true
}

// SetTopologyKey sets field value
func (o *DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionInner) SetTopologyKey(v string) {
	o.TopologyKey = v
}

func (o DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LabelSelector) {
		toSerialize["labelSelector"] = o.LabelSelector
	}
	if !IsNil(o.NamespaceSelector) {
		toSerialize["namespaceSelector"] = o.NamespaceSelector
	}
	if !IsNil(o.Namespaces) {
		toSerialize["namespaces"] = o.Namespaces
	}
	toSerialize["topologyKey"] = o.TopologyKey
	return toSerialize, nil
}

type NullableDatabaseClusterSpecBackupStoragesValueAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionInner struct {
	value *DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionInner
	isSet bool
}

func (v NullableDatabaseClusterSpecBackupStoragesValueAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionInner) Get() *DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionInner {
	return v.value
}

func (v *NullableDatabaseClusterSpecBackupStoragesValueAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionInner) Set(val *DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionInner) {
	v.value = val
	v.isSet = true
}

func (v NullableDatabaseClusterSpecBackupStoragesValueAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionInner) IsSet() bool {
	return v.isSet
}

func (v *NullableDatabaseClusterSpecBackupStoragesValueAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatabaseClusterSpecBackupStoragesValueAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionInner(val *DatabaseClusterSpecBackupStoragesValueAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionInner) *NullableDatabaseClusterSpecBackupStoragesValueAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionInner {
	return &NullableDatabaseClusterSpecBackupStoragesValueAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionInner{value: val, isSet: true}
}

func (v NullableDatabaseClusterSpecBackupStoragesValueAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatabaseClusterSpecBackupStoragesValueAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


