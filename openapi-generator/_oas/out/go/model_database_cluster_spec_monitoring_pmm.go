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

// checks if the DatabaseClusterSpecMonitoringPmm type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatabaseClusterSpecMonitoringPmm{}

// DatabaseClusterSpecMonitoringPmm PMMSpec contains PMM settings.
type DatabaseClusterSpecMonitoringPmm struct {
	Image *string `json:"image,omitempty"`
	Login *string `json:"login,omitempty"`
	Password *string `json:"password,omitempty"`
	PublicAddress *string `json:"publicAddress,omitempty"`
	ServerHost *string `json:"serverHost,omitempty"`
	ServerUser *string `json:"serverUser,omitempty"`
}

// NewDatabaseClusterSpecMonitoringPmm instantiates a new DatabaseClusterSpecMonitoringPmm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatabaseClusterSpecMonitoringPmm() *DatabaseClusterSpecMonitoringPmm {
	this := DatabaseClusterSpecMonitoringPmm{}
	return &this
}

// NewDatabaseClusterSpecMonitoringPmmWithDefaults instantiates a new DatabaseClusterSpecMonitoringPmm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatabaseClusterSpecMonitoringPmmWithDefaults() *DatabaseClusterSpecMonitoringPmm {
	this := DatabaseClusterSpecMonitoringPmm{}
	return &this
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *DatabaseClusterSpecMonitoringPmm) GetImage() string {
	if o == nil || IsNil(o.Image) {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseClusterSpecMonitoringPmm) GetImageOk() (*string, bool) {
	if o == nil || IsNil(o.Image) {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *DatabaseClusterSpecMonitoringPmm) HasImage() bool {
	if o != nil && !IsNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *DatabaseClusterSpecMonitoringPmm) SetImage(v string) {
	o.Image = &v
}

// GetLogin returns the Login field value if set, zero value otherwise.
func (o *DatabaseClusterSpecMonitoringPmm) GetLogin() string {
	if o == nil || IsNil(o.Login) {
		var ret string
		return ret
	}
	return *o.Login
}

// GetLoginOk returns a tuple with the Login field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseClusterSpecMonitoringPmm) GetLoginOk() (*string, bool) {
	if o == nil || IsNil(o.Login) {
		return nil, false
	}
	return o.Login, true
}

// HasLogin returns a boolean if a field has been set.
func (o *DatabaseClusterSpecMonitoringPmm) HasLogin() bool {
	if o != nil && !IsNil(o.Login) {
		return true
	}

	return false
}

// SetLogin gets a reference to the given string and assigns it to the Login field.
func (o *DatabaseClusterSpecMonitoringPmm) SetLogin(v string) {
	o.Login = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *DatabaseClusterSpecMonitoringPmm) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseClusterSpecMonitoringPmm) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *DatabaseClusterSpecMonitoringPmm) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *DatabaseClusterSpecMonitoringPmm) SetPassword(v string) {
	o.Password = &v
}

// GetPublicAddress returns the PublicAddress field value if set, zero value otherwise.
func (o *DatabaseClusterSpecMonitoringPmm) GetPublicAddress() string {
	if o == nil || IsNil(o.PublicAddress) {
		var ret string
		return ret
	}
	return *o.PublicAddress
}

// GetPublicAddressOk returns a tuple with the PublicAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseClusterSpecMonitoringPmm) GetPublicAddressOk() (*string, bool) {
	if o == nil || IsNil(o.PublicAddress) {
		return nil, false
	}
	return o.PublicAddress, true
}

// HasPublicAddress returns a boolean if a field has been set.
func (o *DatabaseClusterSpecMonitoringPmm) HasPublicAddress() bool {
	if o != nil && !IsNil(o.PublicAddress) {
		return true
	}

	return false
}

// SetPublicAddress gets a reference to the given string and assigns it to the PublicAddress field.
func (o *DatabaseClusterSpecMonitoringPmm) SetPublicAddress(v string) {
	o.PublicAddress = &v
}

// GetServerHost returns the ServerHost field value if set, zero value otherwise.
func (o *DatabaseClusterSpecMonitoringPmm) GetServerHost() string {
	if o == nil || IsNil(o.ServerHost) {
		var ret string
		return ret
	}
	return *o.ServerHost
}

// GetServerHostOk returns a tuple with the ServerHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseClusterSpecMonitoringPmm) GetServerHostOk() (*string, bool) {
	if o == nil || IsNil(o.ServerHost) {
		return nil, false
	}
	return o.ServerHost, true
}

// HasServerHost returns a boolean if a field has been set.
func (o *DatabaseClusterSpecMonitoringPmm) HasServerHost() bool {
	if o != nil && !IsNil(o.ServerHost) {
		return true
	}

	return false
}

// SetServerHost gets a reference to the given string and assigns it to the ServerHost field.
func (o *DatabaseClusterSpecMonitoringPmm) SetServerHost(v string) {
	o.ServerHost = &v
}

// GetServerUser returns the ServerUser field value if set, zero value otherwise.
func (o *DatabaseClusterSpecMonitoringPmm) GetServerUser() string {
	if o == nil || IsNil(o.ServerUser) {
		var ret string
		return ret
	}
	return *o.ServerUser
}

// GetServerUserOk returns a tuple with the ServerUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseClusterSpecMonitoringPmm) GetServerUserOk() (*string, bool) {
	if o == nil || IsNil(o.ServerUser) {
		return nil, false
	}
	return o.ServerUser, true
}

// HasServerUser returns a boolean if a field has been set.
func (o *DatabaseClusterSpecMonitoringPmm) HasServerUser() bool {
	if o != nil && !IsNil(o.ServerUser) {
		return true
	}

	return false
}

// SetServerUser gets a reference to the given string and assigns it to the ServerUser field.
func (o *DatabaseClusterSpecMonitoringPmm) SetServerUser(v string) {
	o.ServerUser = &v
}

func (o DatabaseClusterSpecMonitoringPmm) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatabaseClusterSpecMonitoringPmm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	if !IsNil(o.Login) {
		toSerialize["login"] = o.Login
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.PublicAddress) {
		toSerialize["publicAddress"] = o.PublicAddress
	}
	if !IsNil(o.ServerHost) {
		toSerialize["serverHost"] = o.ServerHost
	}
	if !IsNil(o.ServerUser) {
		toSerialize["serverUser"] = o.ServerUser
	}
	return toSerialize, nil
}

type NullableDatabaseClusterSpecMonitoringPmm struct {
	value *DatabaseClusterSpecMonitoringPmm
	isSet bool
}

func (v NullableDatabaseClusterSpecMonitoringPmm) Get() *DatabaseClusterSpecMonitoringPmm {
	return v.value
}

func (v *NullableDatabaseClusterSpecMonitoringPmm) Set(val *DatabaseClusterSpecMonitoringPmm) {
	v.value = val
	v.isSet = true
}

func (v NullableDatabaseClusterSpecMonitoringPmm) IsSet() bool {
	return v.isSet
}

func (v *NullableDatabaseClusterSpecMonitoringPmm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatabaseClusterSpecMonitoringPmm(val *DatabaseClusterSpecMonitoringPmm) *NullableDatabaseClusterSpecMonitoringPmm {
	return &NullableDatabaseClusterSpecMonitoringPmm{value: val, isSet: true}
}

func (v NullableDatabaseClusterSpecMonitoringPmm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatabaseClusterSpecMonitoringPmm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


