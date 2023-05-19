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

// checks if the DatabaseClusterSpecBackupResources type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatabaseClusterSpecBackupResources{}

// DatabaseClusterSpecBackupResources ResourceRequirements describes the compute resource requirements.
type DatabaseClusterSpecBackupResources struct {
	// Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.   This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.   This field is immutable. It can only be set for containers.
	Claims []DatabaseClusterSpecBackupResourcesClaimsInner `json:"claims,omitempty"`
	// Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
	Limits *map[string]DatabaseClusterSpecBackupResourcesLimitsValue `json:"limits,omitempty"`
	// Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
	Requests *map[string]DatabaseClusterSpecBackupResourcesLimitsValue `json:"requests,omitempty"`
}

// NewDatabaseClusterSpecBackupResources instantiates a new DatabaseClusterSpecBackupResources object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatabaseClusterSpecBackupResources() *DatabaseClusterSpecBackupResources {
	this := DatabaseClusterSpecBackupResources{}
	return &this
}

// NewDatabaseClusterSpecBackupResourcesWithDefaults instantiates a new DatabaseClusterSpecBackupResources object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatabaseClusterSpecBackupResourcesWithDefaults() *DatabaseClusterSpecBackupResources {
	this := DatabaseClusterSpecBackupResources{}
	return &this
}

// GetClaims returns the Claims field value if set, zero value otherwise.
func (o *DatabaseClusterSpecBackupResources) GetClaims() []DatabaseClusterSpecBackupResourcesClaimsInner {
	if o == nil || IsNil(o.Claims) {
		var ret []DatabaseClusterSpecBackupResourcesClaimsInner
		return ret
	}
	return o.Claims
}

// GetClaimsOk returns a tuple with the Claims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseClusterSpecBackupResources) GetClaimsOk() ([]DatabaseClusterSpecBackupResourcesClaimsInner, bool) {
	if o == nil || IsNil(o.Claims) {
		return nil, false
	}
	return o.Claims, true
}

// HasClaims returns a boolean if a field has been set.
func (o *DatabaseClusterSpecBackupResources) HasClaims() bool {
	if o != nil && !IsNil(o.Claims) {
		return true
	}

	return false
}

// SetClaims gets a reference to the given []DatabaseClusterSpecBackupResourcesClaimsInner and assigns it to the Claims field.
func (o *DatabaseClusterSpecBackupResources) SetClaims(v []DatabaseClusterSpecBackupResourcesClaimsInner) {
	o.Claims = v
}

// GetLimits returns the Limits field value if set, zero value otherwise.
func (o *DatabaseClusterSpecBackupResources) GetLimits() map[string]DatabaseClusterSpecBackupResourcesLimitsValue {
	if o == nil || IsNil(o.Limits) {
		var ret map[string]DatabaseClusterSpecBackupResourcesLimitsValue
		return ret
	}
	return *o.Limits
}

// GetLimitsOk returns a tuple with the Limits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseClusterSpecBackupResources) GetLimitsOk() (*map[string]DatabaseClusterSpecBackupResourcesLimitsValue, bool) {
	if o == nil || IsNil(o.Limits) {
		return nil, false
	}
	return o.Limits, true
}

// HasLimits returns a boolean if a field has been set.
func (o *DatabaseClusterSpecBackupResources) HasLimits() bool {
	if o != nil && !IsNil(o.Limits) {
		return true
	}

	return false
}

// SetLimits gets a reference to the given map[string]DatabaseClusterSpecBackupResourcesLimitsValue and assigns it to the Limits field.
func (o *DatabaseClusterSpecBackupResources) SetLimits(v map[string]DatabaseClusterSpecBackupResourcesLimitsValue) {
	o.Limits = &v
}

// GetRequests returns the Requests field value if set, zero value otherwise.
func (o *DatabaseClusterSpecBackupResources) GetRequests() map[string]DatabaseClusterSpecBackupResourcesLimitsValue {
	if o == nil || IsNil(o.Requests) {
		var ret map[string]DatabaseClusterSpecBackupResourcesLimitsValue
		return ret
	}
	return *o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseClusterSpecBackupResources) GetRequestsOk() (*map[string]DatabaseClusterSpecBackupResourcesLimitsValue, bool) {
	if o == nil || IsNil(o.Requests) {
		return nil, false
	}
	return o.Requests, true
}

// HasRequests returns a boolean if a field has been set.
func (o *DatabaseClusterSpecBackupResources) HasRequests() bool {
	if o != nil && !IsNil(o.Requests) {
		return true
	}

	return false
}

// SetRequests gets a reference to the given map[string]DatabaseClusterSpecBackupResourcesLimitsValue and assigns it to the Requests field.
func (o *DatabaseClusterSpecBackupResources) SetRequests(v map[string]DatabaseClusterSpecBackupResourcesLimitsValue) {
	o.Requests = &v
}

func (o DatabaseClusterSpecBackupResources) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatabaseClusterSpecBackupResources) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Claims) {
		toSerialize["claims"] = o.Claims
	}
	if !IsNil(o.Limits) {
		toSerialize["limits"] = o.Limits
	}
	if !IsNil(o.Requests) {
		toSerialize["requests"] = o.Requests
	}
	return toSerialize, nil
}

type NullableDatabaseClusterSpecBackupResources struct {
	value *DatabaseClusterSpecBackupResources
	isSet bool
}

func (v NullableDatabaseClusterSpecBackupResources) Get() *DatabaseClusterSpecBackupResources {
	return v.value
}

func (v *NullableDatabaseClusterSpecBackupResources) Set(val *DatabaseClusterSpecBackupResources) {
	v.value = val
	v.isSet = true
}

func (v NullableDatabaseClusterSpecBackupResources) IsSet() bool {
	return v.isSet
}

func (v *NullableDatabaseClusterSpecBackupResources) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatabaseClusterSpecBackupResources(val *DatabaseClusterSpecBackupResources) *NullableDatabaseClusterSpecBackupResources {
	return &NullableDatabaseClusterSpecBackupResources{value: val, isSet: true}
}

func (v NullableDatabaseClusterSpecBackupResources) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatabaseClusterSpecBackupResources) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

