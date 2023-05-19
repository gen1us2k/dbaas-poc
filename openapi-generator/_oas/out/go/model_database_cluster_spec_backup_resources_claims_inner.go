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

// checks if the DatabaseClusterSpecBackupResourcesClaimsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatabaseClusterSpecBackupResourcesClaimsInner{}

// DatabaseClusterSpecBackupResourcesClaimsInner ResourceClaim references one entry in PodSpec.ResourceClaims.
type DatabaseClusterSpecBackupResourcesClaimsInner struct {
	// Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.
	Name string `json:"name"`
}

// NewDatabaseClusterSpecBackupResourcesClaimsInner instantiates a new DatabaseClusterSpecBackupResourcesClaimsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatabaseClusterSpecBackupResourcesClaimsInner(name string) *DatabaseClusterSpecBackupResourcesClaimsInner {
	this := DatabaseClusterSpecBackupResourcesClaimsInner{}
	this.Name = name
	return &this
}

// NewDatabaseClusterSpecBackupResourcesClaimsInnerWithDefaults instantiates a new DatabaseClusterSpecBackupResourcesClaimsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatabaseClusterSpecBackupResourcesClaimsInnerWithDefaults() *DatabaseClusterSpecBackupResourcesClaimsInner {
	this := DatabaseClusterSpecBackupResourcesClaimsInner{}
	return &this
}

// GetName returns the Name field value
func (o *DatabaseClusterSpecBackupResourcesClaimsInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DatabaseClusterSpecBackupResourcesClaimsInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DatabaseClusterSpecBackupResourcesClaimsInner) SetName(v string) {
	o.Name = v
}

func (o DatabaseClusterSpecBackupResourcesClaimsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatabaseClusterSpecBackupResourcesClaimsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableDatabaseClusterSpecBackupResourcesClaimsInner struct {
	value *DatabaseClusterSpecBackupResourcesClaimsInner
	isSet bool
}

func (v NullableDatabaseClusterSpecBackupResourcesClaimsInner) Get() *DatabaseClusterSpecBackupResourcesClaimsInner {
	return v.value
}

func (v *NullableDatabaseClusterSpecBackupResourcesClaimsInner) Set(val *DatabaseClusterSpecBackupResourcesClaimsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableDatabaseClusterSpecBackupResourcesClaimsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableDatabaseClusterSpecBackupResourcesClaimsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatabaseClusterSpecBackupResourcesClaimsInner(val *DatabaseClusterSpecBackupResourcesClaimsInner) *NullableDatabaseClusterSpecBackupResourcesClaimsInner {
	return &NullableDatabaseClusterSpecBackupResourcesClaimsInner{value: val, isSet: true}
}

func (v NullableDatabaseClusterSpecBackupResourcesClaimsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatabaseClusterSpecBackupResourcesClaimsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


