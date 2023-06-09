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

// checks if the DatabaseClusterSpecBackupStoragesValueStorageProvider type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatabaseClusterSpecBackupStoragesValueStorageProvider{}

// DatabaseClusterSpecBackupStoragesValueStorageProvider BackupStorageProviderSpec represents set of settings to configure cloud provider.
type DatabaseClusterSpecBackupStoragesValueStorageProvider struct {
	Bucket *string `json:"bucket,omitempty"`
	// A container name is a valid DNS name that conforms to the Azure naming rules.
	ContainerName *string `json:"containerName,omitempty"`
	CredentialsSecret string `json:"credentialsSecret"`
	EndpointUrl *string `json:"endpointUrl,omitempty"`
	Prefix *string `json:"prefix,omitempty"`
	Region *string `json:"region,omitempty"`
	// STANDARD, NEARLINE, COLDLINE, ARCHIVE for GCP Hot (Frequently accessed or modified data), Cool (Infrequently accessed or modified data), Archive (Rarely accessed or modified data) for Azure.
	StorageClass *string `json:"storageClass,omitempty"`
}

// NewDatabaseClusterSpecBackupStoragesValueStorageProvider instantiates a new DatabaseClusterSpecBackupStoragesValueStorageProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatabaseClusterSpecBackupStoragesValueStorageProvider(credentialsSecret string) *DatabaseClusterSpecBackupStoragesValueStorageProvider {
	this := DatabaseClusterSpecBackupStoragesValueStorageProvider{}
	this.CredentialsSecret = credentialsSecret
	return &this
}

// NewDatabaseClusterSpecBackupStoragesValueStorageProviderWithDefaults instantiates a new DatabaseClusterSpecBackupStoragesValueStorageProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatabaseClusterSpecBackupStoragesValueStorageProviderWithDefaults() *DatabaseClusterSpecBackupStoragesValueStorageProvider {
	this := DatabaseClusterSpecBackupStoragesValueStorageProvider{}
	return &this
}

// GetBucket returns the Bucket field value if set, zero value otherwise.
func (o *DatabaseClusterSpecBackupStoragesValueStorageProvider) GetBucket() string {
	if o == nil || IsNil(o.Bucket) {
		var ret string
		return ret
	}
	return *o.Bucket
}

// GetBucketOk returns a tuple with the Bucket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseClusterSpecBackupStoragesValueStorageProvider) GetBucketOk() (*string, bool) {
	if o == nil || IsNil(o.Bucket) {
		return nil, false
	}
	return o.Bucket, true
}

// HasBucket returns a boolean if a field has been set.
func (o *DatabaseClusterSpecBackupStoragesValueStorageProvider) HasBucket() bool {
	if o != nil && !IsNil(o.Bucket) {
		return true
	}

	return false
}

// SetBucket gets a reference to the given string and assigns it to the Bucket field.
func (o *DatabaseClusterSpecBackupStoragesValueStorageProvider) SetBucket(v string) {
	o.Bucket = &v
}

// GetContainerName returns the ContainerName field value if set, zero value otherwise.
func (o *DatabaseClusterSpecBackupStoragesValueStorageProvider) GetContainerName() string {
	if o == nil || IsNil(o.ContainerName) {
		var ret string
		return ret
	}
	return *o.ContainerName
}

// GetContainerNameOk returns a tuple with the ContainerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseClusterSpecBackupStoragesValueStorageProvider) GetContainerNameOk() (*string, bool) {
	if o == nil || IsNil(o.ContainerName) {
		return nil, false
	}
	return o.ContainerName, true
}

// HasContainerName returns a boolean if a field has been set.
func (o *DatabaseClusterSpecBackupStoragesValueStorageProvider) HasContainerName() bool {
	if o != nil && !IsNil(o.ContainerName) {
		return true
	}

	return false
}

// SetContainerName gets a reference to the given string and assigns it to the ContainerName field.
func (o *DatabaseClusterSpecBackupStoragesValueStorageProvider) SetContainerName(v string) {
	o.ContainerName = &v
}

// GetCredentialsSecret returns the CredentialsSecret field value
func (o *DatabaseClusterSpecBackupStoragesValueStorageProvider) GetCredentialsSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CredentialsSecret
}

// GetCredentialsSecretOk returns a tuple with the CredentialsSecret field value
// and a boolean to check if the value has been set.
func (o *DatabaseClusterSpecBackupStoragesValueStorageProvider) GetCredentialsSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CredentialsSecret, true
}

// SetCredentialsSecret sets field value
func (o *DatabaseClusterSpecBackupStoragesValueStorageProvider) SetCredentialsSecret(v string) {
	o.CredentialsSecret = v
}

// GetEndpointUrl returns the EndpointUrl field value if set, zero value otherwise.
func (o *DatabaseClusterSpecBackupStoragesValueStorageProvider) GetEndpointUrl() string {
	if o == nil || IsNil(o.EndpointUrl) {
		var ret string
		return ret
	}
	return *o.EndpointUrl
}

// GetEndpointUrlOk returns a tuple with the EndpointUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseClusterSpecBackupStoragesValueStorageProvider) GetEndpointUrlOk() (*string, bool) {
	if o == nil || IsNil(o.EndpointUrl) {
		return nil, false
	}
	return o.EndpointUrl, true
}

// HasEndpointUrl returns a boolean if a field has been set.
func (o *DatabaseClusterSpecBackupStoragesValueStorageProvider) HasEndpointUrl() bool {
	if o != nil && !IsNil(o.EndpointUrl) {
		return true
	}

	return false
}

// SetEndpointUrl gets a reference to the given string and assigns it to the EndpointUrl field.
func (o *DatabaseClusterSpecBackupStoragesValueStorageProvider) SetEndpointUrl(v string) {
	o.EndpointUrl = &v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *DatabaseClusterSpecBackupStoragesValueStorageProvider) GetPrefix() string {
	if o == nil || IsNil(o.Prefix) {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseClusterSpecBackupStoragesValueStorageProvider) GetPrefixOk() (*string, bool) {
	if o == nil || IsNil(o.Prefix) {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *DatabaseClusterSpecBackupStoragesValueStorageProvider) HasPrefix() bool {
	if o != nil && !IsNil(o.Prefix) {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *DatabaseClusterSpecBackupStoragesValueStorageProvider) SetPrefix(v string) {
	o.Prefix = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *DatabaseClusterSpecBackupStoragesValueStorageProvider) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseClusterSpecBackupStoragesValueStorageProvider) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *DatabaseClusterSpecBackupStoragesValueStorageProvider) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *DatabaseClusterSpecBackupStoragesValueStorageProvider) SetRegion(v string) {
	o.Region = &v
}

// GetStorageClass returns the StorageClass field value if set, zero value otherwise.
func (o *DatabaseClusterSpecBackupStoragesValueStorageProvider) GetStorageClass() string {
	if o == nil || IsNil(o.StorageClass) {
		var ret string
		return ret
	}
	return *o.StorageClass
}

// GetStorageClassOk returns a tuple with the StorageClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseClusterSpecBackupStoragesValueStorageProvider) GetStorageClassOk() (*string, bool) {
	if o == nil || IsNil(o.StorageClass) {
		return nil, false
	}
	return o.StorageClass, true
}

// HasStorageClass returns a boolean if a field has been set.
func (o *DatabaseClusterSpecBackupStoragesValueStorageProvider) HasStorageClass() bool {
	if o != nil && !IsNil(o.StorageClass) {
		return true
	}

	return false
}

// SetStorageClass gets a reference to the given string and assigns it to the StorageClass field.
func (o *DatabaseClusterSpecBackupStoragesValueStorageProvider) SetStorageClass(v string) {
	o.StorageClass = &v
}

func (o DatabaseClusterSpecBackupStoragesValueStorageProvider) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatabaseClusterSpecBackupStoragesValueStorageProvider) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Bucket) {
		toSerialize["bucket"] = o.Bucket
	}
	if !IsNil(o.ContainerName) {
		toSerialize["containerName"] = o.ContainerName
	}
	toSerialize["credentialsSecret"] = o.CredentialsSecret
	if !IsNil(o.EndpointUrl) {
		toSerialize["endpointUrl"] = o.EndpointUrl
	}
	if !IsNil(o.Prefix) {
		toSerialize["prefix"] = o.Prefix
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !IsNil(o.StorageClass) {
		toSerialize["storageClass"] = o.StorageClass
	}
	return toSerialize, nil
}

type NullableDatabaseClusterSpecBackupStoragesValueStorageProvider struct {
	value *DatabaseClusterSpecBackupStoragesValueStorageProvider
	isSet bool
}

func (v NullableDatabaseClusterSpecBackupStoragesValueStorageProvider) Get() *DatabaseClusterSpecBackupStoragesValueStorageProvider {
	return v.value
}

func (v *NullableDatabaseClusterSpecBackupStoragesValueStorageProvider) Set(val *DatabaseClusterSpecBackupStoragesValueStorageProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableDatabaseClusterSpecBackupStoragesValueStorageProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableDatabaseClusterSpecBackupStoragesValueStorageProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatabaseClusterSpecBackupStoragesValueStorageProvider(val *DatabaseClusterSpecBackupStoragesValueStorageProvider) *NullableDatabaseClusterSpecBackupStoragesValueStorageProvider {
	return &NullableDatabaseClusterSpecBackupStoragesValueStorageProvider{value: val, isSet: true}
}

func (v NullableDatabaseClusterSpecBackupStoragesValueStorageProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatabaseClusterSpecBackupStoragesValueStorageProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


