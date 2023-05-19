# DatabaseClusterSpecBackupContainerSecurityContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowPrivilegeEscalation** | Pointer to **bool** | AllowPrivilegeEscalation controls whether a process can gain more privileges than its parent process. This bool directly controls if the no_new_privs flag will be set on the container process. AllowPrivilegeEscalation is true always when the container is: 1) run as Privileged 2) has CAP_SYS_ADMIN Note that this field cannot be set when spec.os.name is windows. | [optional] 
**Capabilities** | Pointer to [**DatabaseClusterSpecBackupContainerSecurityContextCapabilities**](DatabaseClusterSpecBackupContainerSecurityContextCapabilities.md) |  | [optional] 
**Privileged** | Pointer to **bool** | Run container in privileged mode. Processes in privileged containers are essentially equivalent to root on the host. Defaults to false. Note that this field cannot be set when spec.os.name is windows. | [optional] 
**ProcMount** | Pointer to **string** | procMount denotes the type of proc mount to use for the containers. The default is DefaultProcMount which uses the container runtime defaults for readonly paths and masked paths. This requires the ProcMountType feature flag to be enabled. Note that this field cannot be set when spec.os.name is windows. | [optional] 
**ReadOnlyRootFilesystem** | Pointer to **bool** | Whether this container has a read-only root filesystem. Default is false. Note that this field cannot be set when spec.os.name is windows. | [optional] 
**RunAsGroup** | Pointer to **int64** | The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is windows. | [optional] 
**RunAsNonRoot** | Pointer to **bool** | Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. | [optional] 
**RunAsUser** | Pointer to **int64** | The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is windows. | [optional] 
**SeLinuxOptions** | Pointer to [**DatabaseClusterSpecBackupContainerSecurityContextSeLinuxOptions**](DatabaseClusterSpecBackupContainerSecurityContextSeLinuxOptions.md) |  | [optional] 
**SeccompProfile** | Pointer to [**DatabaseClusterSpecBackupContainerSecurityContextSeccompProfile**](DatabaseClusterSpecBackupContainerSecurityContextSeccompProfile.md) |  | [optional] 
**WindowsOptions** | Pointer to [**DatabaseClusterSpecBackupContainerSecurityContextWindowsOptions**](DatabaseClusterSpecBackupContainerSecurityContextWindowsOptions.md) |  | [optional] 

## Methods

### NewDatabaseClusterSpecBackupContainerSecurityContext

`func NewDatabaseClusterSpecBackupContainerSecurityContext() *DatabaseClusterSpecBackupContainerSecurityContext`

NewDatabaseClusterSpecBackupContainerSecurityContext instantiates a new DatabaseClusterSpecBackupContainerSecurityContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseClusterSpecBackupContainerSecurityContextWithDefaults

`func NewDatabaseClusterSpecBackupContainerSecurityContextWithDefaults() *DatabaseClusterSpecBackupContainerSecurityContext`

NewDatabaseClusterSpecBackupContainerSecurityContextWithDefaults instantiates a new DatabaseClusterSpecBackupContainerSecurityContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowPrivilegeEscalation

`func (o *DatabaseClusterSpecBackupContainerSecurityContext) GetAllowPrivilegeEscalation() bool`

GetAllowPrivilegeEscalation returns the AllowPrivilegeEscalation field if non-nil, zero value otherwise.

### GetAllowPrivilegeEscalationOk

`func (o *DatabaseClusterSpecBackupContainerSecurityContext) GetAllowPrivilegeEscalationOk() (*bool, bool)`

GetAllowPrivilegeEscalationOk returns a tuple with the AllowPrivilegeEscalation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPrivilegeEscalation

`func (o *DatabaseClusterSpecBackupContainerSecurityContext) SetAllowPrivilegeEscalation(v bool)`

SetAllowPrivilegeEscalation sets AllowPrivilegeEscalation field to given value.

### HasAllowPrivilegeEscalation

`func (o *DatabaseClusterSpecBackupContainerSecurityContext) HasAllowPrivilegeEscalation() bool`

HasAllowPrivilegeEscalation returns a boolean if a field has been set.

### GetCapabilities

`func (o *DatabaseClusterSpecBackupContainerSecurityContext) GetCapabilities() DatabaseClusterSpecBackupContainerSecurityContextCapabilities`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *DatabaseClusterSpecBackupContainerSecurityContext) GetCapabilitiesOk() (*DatabaseClusterSpecBackupContainerSecurityContextCapabilities, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *DatabaseClusterSpecBackupContainerSecurityContext) SetCapabilities(v DatabaseClusterSpecBackupContainerSecurityContextCapabilities)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *DatabaseClusterSpecBackupContainerSecurityContext) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetPrivileged

`func (o *DatabaseClusterSpecBackupContainerSecurityContext) GetPrivileged() bool`

GetPrivileged returns the Privileged field if non-nil, zero value otherwise.

### GetPrivilegedOk

`func (o *DatabaseClusterSpecBackupContainerSecurityContext) GetPrivilegedOk() (*bool, bool)`

GetPrivilegedOk returns a tuple with the Privileged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileged

`func (o *DatabaseClusterSpecBackupContainerSecurityContext) SetPrivileged(v bool)`

SetPrivileged sets Privileged field to given value.

### HasPrivileged

`func (o *DatabaseClusterSpecBackupContainerSecurityContext) HasPrivileged() bool`

HasPrivileged returns a boolean if a field has been set.

### GetProcMount

`func (o *DatabaseClusterSpecBackupContainerSecurityContext) GetProcMount() string`

GetProcMount returns the ProcMount field if non-nil, zero value otherwise.

### GetProcMountOk

`func (o *DatabaseClusterSpecBackupContainerSecurityContext) GetProcMountOk() (*string, bool)`

GetProcMountOk returns a tuple with the ProcMount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcMount

`func (o *DatabaseClusterSpecBackupContainerSecurityContext) SetProcMount(v string)`

SetProcMount sets ProcMount field to given value.

### HasProcMount

`func (o *DatabaseClusterSpecBackupContainerSecurityContext) HasProcMount() bool`

HasProcMount returns a boolean if a field has been set.

### GetReadOnlyRootFilesystem

`func (o *DatabaseClusterSpecBackupContainerSecurityContext) GetReadOnlyRootFilesystem() bool`

GetReadOnlyRootFilesystem returns the ReadOnlyRootFilesystem field if non-nil, zero value otherwise.

### GetReadOnlyRootFilesystemOk

`func (o *DatabaseClusterSpecBackupContainerSecurityContext) GetReadOnlyRootFilesystemOk() (*bool, bool)`

GetReadOnlyRootFilesystemOk returns a tuple with the ReadOnlyRootFilesystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnlyRootFilesystem

`func (o *DatabaseClusterSpecBackupContainerSecurityContext) SetReadOnlyRootFilesystem(v bool)`

SetReadOnlyRootFilesystem sets ReadOnlyRootFilesystem field to given value.

### HasReadOnlyRootFilesystem

`func (o *DatabaseClusterSpecBackupContainerSecurityContext) HasReadOnlyRootFilesystem() bool`

HasReadOnlyRootFilesystem returns a boolean if a field has been set.

### GetRunAsGroup

`func (o *DatabaseClusterSpecBackupContainerSecurityContext) GetRunAsGroup() int64`

GetRunAsGroup returns the RunAsGroup field if non-nil, zero value otherwise.

### GetRunAsGroupOk

`func (o *DatabaseClusterSpecBackupContainerSecurityContext) GetRunAsGroupOk() (*int64, bool)`

GetRunAsGroupOk returns a tuple with the RunAsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunAsGroup

`func (o *DatabaseClusterSpecBackupContainerSecurityContext) SetRunAsGroup(v int64)`

SetRunAsGroup sets RunAsGroup field to given value.

### HasRunAsGroup

`func (o *DatabaseClusterSpecBackupContainerSecurityContext) HasRunAsGroup() bool`

HasRunAsGroup returns a boolean if a field has been set.

### GetRunAsNonRoot

`func (o *DatabaseClusterSpecBackupContainerSecurityContext) GetRunAsNonRoot() bool`

GetRunAsNonRoot returns the RunAsNonRoot field if non-nil, zero value otherwise.

### GetRunAsNonRootOk

`func (o *DatabaseClusterSpecBackupContainerSecurityContext) GetRunAsNonRootOk() (*bool, bool)`

GetRunAsNonRootOk returns a tuple with the RunAsNonRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunAsNonRoot

`func (o *DatabaseClusterSpecBackupContainerSecurityContext) SetRunAsNonRoot(v bool)`

SetRunAsNonRoot sets RunAsNonRoot field to given value.

### HasRunAsNonRoot

`func (o *DatabaseClusterSpecBackupContainerSecurityContext) HasRunAsNonRoot() bool`

HasRunAsNonRoot returns a boolean if a field has been set.

### GetRunAsUser

`func (o *DatabaseClusterSpecBackupContainerSecurityContext) GetRunAsUser() int64`

GetRunAsUser returns the RunAsUser field if non-nil, zero value otherwise.

### GetRunAsUserOk

`func (o *DatabaseClusterSpecBackupContainerSecurityContext) GetRunAsUserOk() (*int64, bool)`

GetRunAsUserOk returns a tuple with the RunAsUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunAsUser

`func (o *DatabaseClusterSpecBackupContainerSecurityContext) SetRunAsUser(v int64)`

SetRunAsUser sets RunAsUser field to given value.

### HasRunAsUser

`func (o *DatabaseClusterSpecBackupContainerSecurityContext) HasRunAsUser() bool`

HasRunAsUser returns a boolean if a field has been set.

### GetSeLinuxOptions

`func (o *DatabaseClusterSpecBackupContainerSecurityContext) GetSeLinuxOptions() DatabaseClusterSpecBackupContainerSecurityContextSeLinuxOptions`

GetSeLinuxOptions returns the SeLinuxOptions field if non-nil, zero value otherwise.

### GetSeLinuxOptionsOk

`func (o *DatabaseClusterSpecBackupContainerSecurityContext) GetSeLinuxOptionsOk() (*DatabaseClusterSpecBackupContainerSecurityContextSeLinuxOptions, bool)`

GetSeLinuxOptionsOk returns a tuple with the SeLinuxOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeLinuxOptions

`func (o *DatabaseClusterSpecBackupContainerSecurityContext) SetSeLinuxOptions(v DatabaseClusterSpecBackupContainerSecurityContextSeLinuxOptions)`

SetSeLinuxOptions sets SeLinuxOptions field to given value.

### HasSeLinuxOptions

`func (o *DatabaseClusterSpecBackupContainerSecurityContext) HasSeLinuxOptions() bool`

HasSeLinuxOptions returns a boolean if a field has been set.

### GetSeccompProfile

`func (o *DatabaseClusterSpecBackupContainerSecurityContext) GetSeccompProfile() DatabaseClusterSpecBackupContainerSecurityContextSeccompProfile`

GetSeccompProfile returns the SeccompProfile field if non-nil, zero value otherwise.

### GetSeccompProfileOk

`func (o *DatabaseClusterSpecBackupContainerSecurityContext) GetSeccompProfileOk() (*DatabaseClusterSpecBackupContainerSecurityContextSeccompProfile, bool)`

GetSeccompProfileOk returns a tuple with the SeccompProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeccompProfile

`func (o *DatabaseClusterSpecBackupContainerSecurityContext) SetSeccompProfile(v DatabaseClusterSpecBackupContainerSecurityContextSeccompProfile)`

SetSeccompProfile sets SeccompProfile field to given value.

### HasSeccompProfile

`func (o *DatabaseClusterSpecBackupContainerSecurityContext) HasSeccompProfile() bool`

HasSeccompProfile returns a boolean if a field has been set.

### GetWindowsOptions

`func (o *DatabaseClusterSpecBackupContainerSecurityContext) GetWindowsOptions() DatabaseClusterSpecBackupContainerSecurityContextWindowsOptions`

GetWindowsOptions returns the WindowsOptions field if non-nil, zero value otherwise.

### GetWindowsOptionsOk

`func (o *DatabaseClusterSpecBackupContainerSecurityContext) GetWindowsOptionsOk() (*DatabaseClusterSpecBackupContainerSecurityContextWindowsOptions, bool)`

GetWindowsOptionsOk returns a tuple with the WindowsOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsOptions

`func (o *DatabaseClusterSpecBackupContainerSecurityContext) SetWindowsOptions(v DatabaseClusterSpecBackupContainerSecurityContextWindowsOptions)`

SetWindowsOptions sets WindowsOptions field to given value.

### HasWindowsOptions

`func (o *DatabaseClusterSpecBackupContainerSecurityContext) HasWindowsOptions() bool`

HasWindowsOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


