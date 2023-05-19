# DatabaseClusterSpecBackupStoragesValuePodSecurityContextWindowsOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GmsaCredentialSpec** | Pointer to **string** | GMSACredentialSpec is where the GMSA admission webhook (https://github.com/kubernetes-sigs/windows-gmsa) inlines the contents of the GMSA credential spec named by the GMSACredentialSpecName field. | [optional] 
**GmsaCredentialSpecName** | Pointer to **string** | GMSACredentialSpecName is the name of the GMSA credential spec to use. | [optional] 
**HostProcess** | Pointer to **bool** | HostProcess determines if a container should be run as a &#39;Host Process&#39; container. This field is alpha-level and will only be honored by components that enable the WindowsHostProcessContainers feature flag. Setting this field without the feature flag will result in errors when validating the Pod. All of a Pod&#39;s containers must have the same effective HostProcess value (it is not allowed to have a mix of HostProcess containers and non-HostProcess containers).  In addition, if HostProcess is true then HostNetwork must also be set to true. | [optional] 
**RunAsUserName** | Pointer to **string** | The UserName in Windows to run the entrypoint of the container process. Defaults to the user specified in image metadata if unspecified. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. | [optional] 

## Methods

### NewDatabaseClusterSpecBackupStoragesValuePodSecurityContextWindowsOptions

`func NewDatabaseClusterSpecBackupStoragesValuePodSecurityContextWindowsOptions() *DatabaseClusterSpecBackupStoragesValuePodSecurityContextWindowsOptions`

NewDatabaseClusterSpecBackupStoragesValuePodSecurityContextWindowsOptions instantiates a new DatabaseClusterSpecBackupStoragesValuePodSecurityContextWindowsOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseClusterSpecBackupStoragesValuePodSecurityContextWindowsOptionsWithDefaults

`func NewDatabaseClusterSpecBackupStoragesValuePodSecurityContextWindowsOptionsWithDefaults() *DatabaseClusterSpecBackupStoragesValuePodSecurityContextWindowsOptions`

NewDatabaseClusterSpecBackupStoragesValuePodSecurityContextWindowsOptionsWithDefaults instantiates a new DatabaseClusterSpecBackupStoragesValuePodSecurityContextWindowsOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGmsaCredentialSpec

`func (o *DatabaseClusterSpecBackupStoragesValuePodSecurityContextWindowsOptions) GetGmsaCredentialSpec() string`

GetGmsaCredentialSpec returns the GmsaCredentialSpec field if non-nil, zero value otherwise.

### GetGmsaCredentialSpecOk

`func (o *DatabaseClusterSpecBackupStoragesValuePodSecurityContextWindowsOptions) GetGmsaCredentialSpecOk() (*string, bool)`

GetGmsaCredentialSpecOk returns a tuple with the GmsaCredentialSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmsaCredentialSpec

`func (o *DatabaseClusterSpecBackupStoragesValuePodSecurityContextWindowsOptions) SetGmsaCredentialSpec(v string)`

SetGmsaCredentialSpec sets GmsaCredentialSpec field to given value.

### HasGmsaCredentialSpec

`func (o *DatabaseClusterSpecBackupStoragesValuePodSecurityContextWindowsOptions) HasGmsaCredentialSpec() bool`

HasGmsaCredentialSpec returns a boolean if a field has been set.

### GetGmsaCredentialSpecName

`func (o *DatabaseClusterSpecBackupStoragesValuePodSecurityContextWindowsOptions) GetGmsaCredentialSpecName() string`

GetGmsaCredentialSpecName returns the GmsaCredentialSpecName field if non-nil, zero value otherwise.

### GetGmsaCredentialSpecNameOk

`func (o *DatabaseClusterSpecBackupStoragesValuePodSecurityContextWindowsOptions) GetGmsaCredentialSpecNameOk() (*string, bool)`

GetGmsaCredentialSpecNameOk returns a tuple with the GmsaCredentialSpecName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmsaCredentialSpecName

`func (o *DatabaseClusterSpecBackupStoragesValuePodSecurityContextWindowsOptions) SetGmsaCredentialSpecName(v string)`

SetGmsaCredentialSpecName sets GmsaCredentialSpecName field to given value.

### HasGmsaCredentialSpecName

`func (o *DatabaseClusterSpecBackupStoragesValuePodSecurityContextWindowsOptions) HasGmsaCredentialSpecName() bool`

HasGmsaCredentialSpecName returns a boolean if a field has been set.

### GetHostProcess

`func (o *DatabaseClusterSpecBackupStoragesValuePodSecurityContextWindowsOptions) GetHostProcess() bool`

GetHostProcess returns the HostProcess field if non-nil, zero value otherwise.

### GetHostProcessOk

`func (o *DatabaseClusterSpecBackupStoragesValuePodSecurityContextWindowsOptions) GetHostProcessOk() (*bool, bool)`

GetHostProcessOk returns a tuple with the HostProcess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostProcess

`func (o *DatabaseClusterSpecBackupStoragesValuePodSecurityContextWindowsOptions) SetHostProcess(v bool)`

SetHostProcess sets HostProcess field to given value.

### HasHostProcess

`func (o *DatabaseClusterSpecBackupStoragesValuePodSecurityContextWindowsOptions) HasHostProcess() bool`

HasHostProcess returns a boolean if a field has been set.

### GetRunAsUserName

`func (o *DatabaseClusterSpecBackupStoragesValuePodSecurityContextWindowsOptions) GetRunAsUserName() string`

GetRunAsUserName returns the RunAsUserName field if non-nil, zero value otherwise.

### GetRunAsUserNameOk

`func (o *DatabaseClusterSpecBackupStoragesValuePodSecurityContextWindowsOptions) GetRunAsUserNameOk() (*string, bool)`

GetRunAsUserNameOk returns a tuple with the RunAsUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunAsUserName

`func (o *DatabaseClusterSpecBackupStoragesValuePodSecurityContextWindowsOptions) SetRunAsUserName(v string)`

SetRunAsUserName sets RunAsUserName field to given value.

### HasRunAsUserName

`func (o *DatabaseClusterSpecBackupStoragesValuePodSecurityContextWindowsOptions) HasRunAsUserName() bool`

HasRunAsUserName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


