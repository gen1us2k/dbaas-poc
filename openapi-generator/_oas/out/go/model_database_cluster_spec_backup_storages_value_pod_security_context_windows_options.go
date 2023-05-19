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

// checks if the DatabaseClusterSpecBackupStoragesValuePodSecurityContextWindowsOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatabaseClusterSpecBackupStoragesValuePodSecurityContextWindowsOptions{}

// DatabaseClusterSpecBackupStoragesValuePodSecurityContextWindowsOptions The Windows specific settings applied to all containers. If unspecified, the options within a container's SecurityContext will be used. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is linux.
type DatabaseClusterSpecBackupStoragesValuePodSecurityContextWindowsOptions struct {
	// GMSACredentialSpec is where the GMSA admission webhook (https://github.com/kubernetes-sigs/windows-gmsa) inlines the contents of the GMSA credential spec named by the GMSACredentialSpecName field.
	GmsaCredentialSpec *string `json:"gmsaCredentialSpec,omitempty"`
	// GMSACredentialSpecName is the name of the GMSA credential spec to use.
	GmsaCredentialSpecName *string `json:"gmsaCredentialSpecName,omitempty"`
	// HostProcess determines if a container should be run as a 'Host Process' container. This field is alpha-level and will only be honored by components that enable the WindowsHostProcessContainers feature flag. Setting this field without the feature flag will result in errors when validating the Pod. All of a Pod's containers must have the same effective HostProcess value (it is not allowed to have a mix of HostProcess containers and non-HostProcess containers).  In addition, if HostProcess is true then HostNetwork must also be set to true.
	HostProcess *bool `json:"hostProcess,omitempty"`
	// The UserName in Windows to run the entrypoint of the container process. Defaults to the user specified in image metadata if unspecified. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.
	RunAsUserName *string `json:"runAsUserName,omitempty"`
}

// NewDatabaseClusterSpecBackupStoragesValuePodSecurityContextWindowsOptions instantiates a new DatabaseClusterSpecBackupStoragesValuePodSecurityContextWindowsOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatabaseClusterSpecBackupStoragesValuePodSecurityContextWindowsOptions() *DatabaseClusterSpecBackupStoragesValuePodSecurityContextWindowsOptions {
	this := DatabaseClusterSpecBackupStoragesValuePodSecurityContextWindowsOptions{}
	return &this
}

// NewDatabaseClusterSpecBackupStoragesValuePodSecurityContextWindowsOptionsWithDefaults instantiates a new DatabaseClusterSpecBackupStoragesValuePodSecurityContextWindowsOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatabaseClusterSpecBackupStoragesValuePodSecurityContextWindowsOptionsWithDefaults() *DatabaseClusterSpecBackupStoragesValuePodSecurityContextWindowsOptions {
	this := DatabaseClusterSpecBackupStoragesValuePodSecurityContextWindowsOptions{}
	return &this
}

// GetGmsaCredentialSpec returns the GmsaCredentialSpec field value if set, zero value otherwise.
func (o *DatabaseClusterSpecBackupStoragesValuePodSecurityContextWindowsOptions) GetGmsaCredentialSpec() string {
	if o == nil || IsNil(o.GmsaCredentialSpec) {
		var ret string
		return ret
	}
	return *o.GmsaCredentialSpec
}

// GetGmsaCredentialSpecOk returns a tuple with the GmsaCredentialSpec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseClusterSpecBackupStoragesValuePodSecurityContextWindowsOptions) GetGmsaCredentialSpecOk() (*string, bool) {
	if o == nil || IsNil(o.GmsaCredentialSpec) {
		return nil, false
	}
	return o.GmsaCredentialSpec, true
}

// HasGmsaCredentialSpec returns a boolean if a field has been set.
func (o *DatabaseClusterSpecBackupStoragesValuePodSecurityContextWindowsOptions) HasGmsaCredentialSpec() bool {
	if o != nil && !IsNil(o.GmsaCredentialSpec) {
		return true
	}

	return false
}

// SetGmsaCredentialSpec gets a reference to the given string and assigns it to the GmsaCredentialSpec field.
func (o *DatabaseClusterSpecBackupStoragesValuePodSecurityContextWindowsOptions) SetGmsaCredentialSpec(v string) {
	o.GmsaCredentialSpec = &v
}

// GetGmsaCredentialSpecName returns the GmsaCredentialSpecName field value if set, zero value otherwise.
func (o *DatabaseClusterSpecBackupStoragesValuePodSecurityContextWindowsOptions) GetGmsaCredentialSpecName() string {
	if o == nil || IsNil(o.GmsaCredentialSpecName) {
		var ret string
		return ret
	}
	return *o.GmsaCredentialSpecName
}

// GetGmsaCredentialSpecNameOk returns a tuple with the GmsaCredentialSpecName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseClusterSpecBackupStoragesValuePodSecurityContextWindowsOptions) GetGmsaCredentialSpecNameOk() (*string, bool) {
	if o == nil || IsNil(o.GmsaCredentialSpecName) {
		return nil, false
	}
	return o.GmsaCredentialSpecName, true
}

// HasGmsaCredentialSpecName returns a boolean if a field has been set.
func (o *DatabaseClusterSpecBackupStoragesValuePodSecurityContextWindowsOptions) HasGmsaCredentialSpecName() bool {
	if o != nil && !IsNil(o.GmsaCredentialSpecName) {
		return true
	}

	return false
}

// SetGmsaCredentialSpecName gets a reference to the given string and assigns it to the GmsaCredentialSpecName field.
func (o *DatabaseClusterSpecBackupStoragesValuePodSecurityContextWindowsOptions) SetGmsaCredentialSpecName(v string) {
	o.GmsaCredentialSpecName = &v
}

// GetHostProcess returns the HostProcess field value if set, zero value otherwise.
func (o *DatabaseClusterSpecBackupStoragesValuePodSecurityContextWindowsOptions) GetHostProcess() bool {
	if o == nil || IsNil(o.HostProcess) {
		var ret bool
		return ret
	}
	return *o.HostProcess
}

// GetHostProcessOk returns a tuple with the HostProcess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseClusterSpecBackupStoragesValuePodSecurityContextWindowsOptions) GetHostProcessOk() (*bool, bool) {
	if o == nil || IsNil(o.HostProcess) {
		return nil, false
	}
	return o.HostProcess, true
}

// HasHostProcess returns a boolean if a field has been set.
func (o *DatabaseClusterSpecBackupStoragesValuePodSecurityContextWindowsOptions) HasHostProcess() bool {
	if o != nil && !IsNil(o.HostProcess) {
		return true
	}

	return false
}

// SetHostProcess gets a reference to the given bool and assigns it to the HostProcess field.
func (o *DatabaseClusterSpecBackupStoragesValuePodSecurityContextWindowsOptions) SetHostProcess(v bool) {
	o.HostProcess = &v
}

// GetRunAsUserName returns the RunAsUserName field value if set, zero value otherwise.
func (o *DatabaseClusterSpecBackupStoragesValuePodSecurityContextWindowsOptions) GetRunAsUserName() string {
	if o == nil || IsNil(o.RunAsUserName) {
		var ret string
		return ret
	}
	return *o.RunAsUserName
}

// GetRunAsUserNameOk returns a tuple with the RunAsUserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseClusterSpecBackupStoragesValuePodSecurityContextWindowsOptions) GetRunAsUserNameOk() (*string, bool) {
	if o == nil || IsNil(o.RunAsUserName) {
		return nil, false
	}
	return o.RunAsUserName, true
}

// HasRunAsUserName returns a boolean if a field has been set.
func (o *DatabaseClusterSpecBackupStoragesValuePodSecurityContextWindowsOptions) HasRunAsUserName() bool {
	if o != nil && !IsNil(o.RunAsUserName) {
		return true
	}

	return false
}

// SetRunAsUserName gets a reference to the given string and assigns it to the RunAsUserName field.
func (o *DatabaseClusterSpecBackupStoragesValuePodSecurityContextWindowsOptions) SetRunAsUserName(v string) {
	o.RunAsUserName = &v
}

func (o DatabaseClusterSpecBackupStoragesValuePodSecurityContextWindowsOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatabaseClusterSpecBackupStoragesValuePodSecurityContextWindowsOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GmsaCredentialSpec) {
		toSerialize["gmsaCredentialSpec"] = o.GmsaCredentialSpec
	}
	if !IsNil(o.GmsaCredentialSpecName) {
		toSerialize["gmsaCredentialSpecName"] = o.GmsaCredentialSpecName
	}
	if !IsNil(o.HostProcess) {
		toSerialize["hostProcess"] = o.HostProcess
	}
	if !IsNil(o.RunAsUserName) {
		toSerialize["runAsUserName"] = o.RunAsUserName
	}
	return toSerialize, nil
}

type NullableDatabaseClusterSpecBackupStoragesValuePodSecurityContextWindowsOptions struct {
	value *DatabaseClusterSpecBackupStoragesValuePodSecurityContextWindowsOptions
	isSet bool
}

func (v NullableDatabaseClusterSpecBackupStoragesValuePodSecurityContextWindowsOptions) Get() *DatabaseClusterSpecBackupStoragesValuePodSecurityContextWindowsOptions {
	return v.value
}

func (v *NullableDatabaseClusterSpecBackupStoragesValuePodSecurityContextWindowsOptions) Set(val *DatabaseClusterSpecBackupStoragesValuePodSecurityContextWindowsOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableDatabaseClusterSpecBackupStoragesValuePodSecurityContextWindowsOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableDatabaseClusterSpecBackupStoragesValuePodSecurityContextWindowsOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatabaseClusterSpecBackupStoragesValuePodSecurityContextWindowsOptions(val *DatabaseClusterSpecBackupStoragesValuePodSecurityContextWindowsOptions) *NullableDatabaseClusterSpecBackupStoragesValuePodSecurityContextWindowsOptions {
	return &NullableDatabaseClusterSpecBackupStoragesValuePodSecurityContextWindowsOptions{value: val, isSet: true}
}

func (v NullableDatabaseClusterSpecBackupStoragesValuePodSecurityContextWindowsOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatabaseClusterSpecBackupStoragesValuePodSecurityContextWindowsOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

