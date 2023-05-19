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

// checks if the DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim{}

// DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim PersistentVolumeClaim to specify PVC spec for the volume for mysql data. It has the highest level of precedence, followed by HostPath and EmptyDir. And represents the PVC specification.
type DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim struct {
	// accessModes contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1
	AccessModes []string `json:"accessModes,omitempty"`
	DataSource *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSource `json:"dataSource,omitempty"`
	DataSourceRef *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRef `json:"dataSourceRef,omitempty"`
	Resources *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimResources `json:"resources,omitempty"`
	Selector *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimSelector `json:"selector,omitempty"`
	// storageClassName is the name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1
	StorageClassName *string `json:"storageClassName,omitempty"`
	// volumeMode defines what type of volume is required by the claim. Value of Filesystem is implied when not included in claim spec.
	VolumeMode *string `json:"volumeMode,omitempty"`
	// volumeName is the binding reference to the PersistentVolume backing this claim.
	VolumeName *string `json:"volumeName,omitempty"`
}

// NewDatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim instantiates a new DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim() *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim {
	this := DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim{}
	return &this
}

// NewDatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimWithDefaults instantiates a new DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimWithDefaults() *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim {
	this := DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim{}
	return &this
}

// GetAccessModes returns the AccessModes field value if set, zero value otherwise.
func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim) GetAccessModes() []string {
	if o == nil || IsNil(o.AccessModes) {
		var ret []string
		return ret
	}
	return o.AccessModes
}

// GetAccessModesOk returns a tuple with the AccessModes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim) GetAccessModesOk() ([]string, bool) {
	if o == nil || IsNil(o.AccessModes) {
		return nil, false
	}
	return o.AccessModes, true
}

// HasAccessModes returns a boolean if a field has been set.
func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim) HasAccessModes() bool {
	if o != nil && !IsNil(o.AccessModes) {
		return true
	}

	return false
}

// SetAccessModes gets a reference to the given []string and assigns it to the AccessModes field.
func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim) SetAccessModes(v []string) {
	o.AccessModes = v
}

// GetDataSource returns the DataSource field value if set, zero value otherwise.
func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim) GetDataSource() DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSource {
	if o == nil || IsNil(o.DataSource) {
		var ret DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSource
		return ret
	}
	return *o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim) GetDataSourceOk() (*DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSource, bool) {
	if o == nil || IsNil(o.DataSource) {
		return nil, false
	}
	return o.DataSource, true
}

// HasDataSource returns a boolean if a field has been set.
func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim) HasDataSource() bool {
	if o != nil && !IsNil(o.DataSource) {
		return true
	}

	return false
}

// SetDataSource gets a reference to the given DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSource and assigns it to the DataSource field.
func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim) SetDataSource(v DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSource) {
	o.DataSource = &v
}

// GetDataSourceRef returns the DataSourceRef field value if set, zero value otherwise.
func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim) GetDataSourceRef() DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRef {
	if o == nil || IsNil(o.DataSourceRef) {
		var ret DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRef
		return ret
	}
	return *o.DataSourceRef
}

// GetDataSourceRefOk returns a tuple with the DataSourceRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim) GetDataSourceRefOk() (*DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRef, bool) {
	if o == nil || IsNil(o.DataSourceRef) {
		return nil, false
	}
	return o.DataSourceRef, true
}

// HasDataSourceRef returns a boolean if a field has been set.
func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim) HasDataSourceRef() bool {
	if o != nil && !IsNil(o.DataSourceRef) {
		return true
	}

	return false
}

// SetDataSourceRef gets a reference to the given DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRef and assigns it to the DataSourceRef field.
func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim) SetDataSourceRef(v DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRef) {
	o.DataSourceRef = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim) GetResources() DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimResources {
	if o == nil || IsNil(o.Resources) {
		var ret DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimResources
		return ret
	}
	return *o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim) GetResourcesOk() (*DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimResources, bool) {
	if o == nil || IsNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim) HasResources() bool {
	if o != nil && !IsNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimResources and assigns it to the Resources field.
func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim) SetResources(v DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimResources) {
	o.Resources = &v
}

// GetSelector returns the Selector field value if set, zero value otherwise.
func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim) GetSelector() DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimSelector {
	if o == nil || IsNil(o.Selector) {
		var ret DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimSelector
		return ret
	}
	return *o.Selector
}

// GetSelectorOk returns a tuple with the Selector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim) GetSelectorOk() (*DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimSelector, bool) {
	if o == nil || IsNil(o.Selector) {
		return nil, false
	}
	return o.Selector, true
}

// HasSelector returns a boolean if a field has been set.
func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim) HasSelector() bool {
	if o != nil && !IsNil(o.Selector) {
		return true
	}

	return false
}

// SetSelector gets a reference to the given DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimSelector and assigns it to the Selector field.
func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim) SetSelector(v DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimSelector) {
	o.Selector = &v
}

// GetStorageClassName returns the StorageClassName field value if set, zero value otherwise.
func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim) GetStorageClassName() string {
	if o == nil || IsNil(o.StorageClassName) {
		var ret string
		return ret
	}
	return *o.StorageClassName
}

// GetStorageClassNameOk returns a tuple with the StorageClassName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim) GetStorageClassNameOk() (*string, bool) {
	if o == nil || IsNil(o.StorageClassName) {
		return nil, false
	}
	return o.StorageClassName, true
}

// HasStorageClassName returns a boolean if a field has been set.
func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim) HasStorageClassName() bool {
	if o != nil && !IsNil(o.StorageClassName) {
		return true
	}

	return false
}

// SetStorageClassName gets a reference to the given string and assigns it to the StorageClassName field.
func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim) SetStorageClassName(v string) {
	o.StorageClassName = &v
}

// GetVolumeMode returns the VolumeMode field value if set, zero value otherwise.
func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim) GetVolumeMode() string {
	if o == nil || IsNil(o.VolumeMode) {
		var ret string
		return ret
	}
	return *o.VolumeMode
}

// GetVolumeModeOk returns a tuple with the VolumeMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim) GetVolumeModeOk() (*string, bool) {
	if o == nil || IsNil(o.VolumeMode) {
		return nil, false
	}
	return o.VolumeMode, true
}

// HasVolumeMode returns a boolean if a field has been set.
func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim) HasVolumeMode() bool {
	if o != nil && !IsNil(o.VolumeMode) {
		return true
	}

	return false
}

// SetVolumeMode gets a reference to the given string and assigns it to the VolumeMode field.
func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim) SetVolumeMode(v string) {
	o.VolumeMode = &v
}

// GetVolumeName returns the VolumeName field value if set, zero value otherwise.
func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim) GetVolumeName() string {
	if o == nil || IsNil(o.VolumeName) {
		var ret string
		return ret
	}
	return *o.VolumeName
}

// GetVolumeNameOk returns a tuple with the VolumeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim) GetVolumeNameOk() (*string, bool) {
	if o == nil || IsNil(o.VolumeName) {
		return nil, false
	}
	return o.VolumeName, true
}

// HasVolumeName returns a boolean if a field has been set.
func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim) HasVolumeName() bool {
	if o != nil && !IsNil(o.VolumeName) {
		return true
	}

	return false
}

// SetVolumeName gets a reference to the given string and assigns it to the VolumeName field.
func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim) SetVolumeName(v string) {
	o.VolumeName = &v
}

func (o DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessModes) {
		toSerialize["accessModes"] = o.AccessModes
	}
	if !IsNil(o.DataSource) {
		toSerialize["dataSource"] = o.DataSource
	}
	if !IsNil(o.DataSourceRef) {
		toSerialize["dataSourceRef"] = o.DataSourceRef
	}
	if !IsNil(o.Resources) {
		toSerialize["resources"] = o.Resources
	}
	if !IsNil(o.Selector) {
		toSerialize["selector"] = o.Selector
	}
	if !IsNil(o.StorageClassName) {
		toSerialize["storageClassName"] = o.StorageClassName
	}
	if !IsNil(o.VolumeMode) {
		toSerialize["volumeMode"] = o.VolumeMode
	}
	if !IsNil(o.VolumeName) {
		toSerialize["volumeName"] = o.VolumeName
	}
	return toSerialize, nil
}

type NullableDatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim struct {
	value *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim
	isSet bool
}

func (v NullableDatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim) Get() *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim {
	return v.value
}

func (v *NullableDatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim) Set(val *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim) {
	v.value = val
	v.isSet = true
}

func (v NullableDatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim) IsSet() bool {
	return v.isSet
}

func (v *NullableDatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim(val *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim) *NullableDatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim {
	return &NullableDatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim{value: val, isSet: true}
}

func (v NullableDatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


