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

// checks if the DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRef type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRef{}

// DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRef dataSourceRef specifies the object from which to populate the volume with data, if a non-empty volume is desired. This may be any object from a non-empty API group (non core object) or a PersistentVolumeClaim object. When this field is specified, volume binding will only succeed if the type of the specified object matches some installed volume populator or dynamic provisioner. This field will replace the functionality of the dataSource field and as such if both fields are non-empty, they must have the same value. For backwards compatibility, when namespace isn't specified in dataSourceRef, both fields (dataSource and dataSourceRef) will be set to the same value automatically if one of them is empty and the other is non-empty. When namespace is specified in dataSourceRef, dataSource isn't set to the same value and must be empty. There are three important differences between dataSource and dataSourceRef: * While dataSource only allows two specific types of objects, dataSourceRef allows any non-core object, as well as PersistentVolumeClaim objects. * While dataSource ignores disallowed values (dropping them), dataSourceRef preserves all values, and generates an error if a disallowed value is specified. * While dataSource only allows local objects, dataSourceRef allows objects in any namespaces. (Beta) Using this field requires the AnyVolumeDataSource feature gate to be enabled. (Alpha) Using the namespace field of dataSourceRef requires the CrossNamespaceVolumeDataSource feature gate to be enabled.
type DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRef struct {
	// APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.
	ApiGroup *string `json:"apiGroup,omitempty"`
	// Kind is the type of resource being referenced
	Kind string `json:"kind"`
	// Name is the name of resource being referenced
	Name string `json:"name"`
	// Namespace is the namespace of resource being referenced Note that when a namespace is specified, a gateway.networking.k8s.io/ReferenceGrant object is required in the referent namespace to allow that namespace's owner to accept the reference. See the ReferenceGrant documentation for details. (Alpha) This field requires the CrossNamespaceVolumeDataSource feature gate to be enabled.
	Namespace *string `json:"namespace,omitempty"`
}

// NewDatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRef instantiates a new DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRef object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRef(kind string, name string) *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRef {
	this := DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRef{}
	this.Kind = kind
	this.Name = name
	return &this
}

// NewDatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRefWithDefaults instantiates a new DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRef object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRefWithDefaults() *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRef {
	this := DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRef{}
	return &this
}

// GetApiGroup returns the ApiGroup field value if set, zero value otherwise.
func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRef) GetApiGroup() string {
	if o == nil || IsNil(o.ApiGroup) {
		var ret string
		return ret
	}
	return *o.ApiGroup
}

// GetApiGroupOk returns a tuple with the ApiGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRef) GetApiGroupOk() (*string, bool) {
	if o == nil || IsNil(o.ApiGroup) {
		return nil, false
	}
	return o.ApiGroup, true
}

// HasApiGroup returns a boolean if a field has been set.
func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRef) HasApiGroup() bool {
	if o != nil && !IsNil(o.ApiGroup) {
		return true
	}

	return false
}

// SetApiGroup gets a reference to the given string and assigns it to the ApiGroup field.
func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRef) SetApiGroup(v string) {
	o.ApiGroup = &v
}

// GetKind returns the Kind field value
func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRef) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRef) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRef) SetKind(v string) {
	o.Kind = v
}

// GetName returns the Name field value
func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRef) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRef) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRef) SetName(v string) {
	o.Name = v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRef) GetNamespace() string {
	if o == nil || IsNil(o.Namespace) {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRef) GetNamespaceOk() (*string, bool) {
	if o == nil || IsNil(o.Namespace) {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRef) HasNamespace() bool {
	if o != nil && !IsNil(o.Namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRef) SetNamespace(v string) {
	o.Namespace = &v
}

func (o DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRef) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRef) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApiGroup) {
		toSerialize["apiGroup"] = o.ApiGroup
	}
	toSerialize["kind"] = o.Kind
	toSerialize["name"] = o.Name
	if !IsNil(o.Namespace) {
		toSerialize["namespace"] = o.Namespace
	}
	return toSerialize, nil
}

type NullableDatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRef struct {
	value *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRef
	isSet bool
}

func (v NullableDatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRef) Get() *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRef {
	return v.value
}

func (v *NullableDatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRef) Set(val *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRef) {
	v.value = val
	v.isSet = true
}

func (v NullableDatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRef) IsSet() bool {
	return v.isSet
}

func (v *NullableDatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRef) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRef(val *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRef) *NullableDatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRef {
	return &NullableDatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRef{value: val, isSet: true}
}

func (v NullableDatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRef) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRef) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


