# DatabaseClusterSpecMonitoring

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainerSecurityContext** | Pointer to [**DatabaseClusterSpecBackupContainerSecurityContext**](DatabaseClusterSpecBackupContainerSecurityContext.md) |  | [optional] 
**ImagePullPolicy** | Pointer to **string** | PullPolicy describes a policy for if/when to pull a container image | [optional] 
**Pmm** | Pointer to [**DatabaseClusterSpecMonitoringPmm**](DatabaseClusterSpecMonitoringPmm.md) |  | [optional] 
**Resources** | Pointer to [**DatabaseClusterSpecBackupResources**](DatabaseClusterSpecBackupResources.md) |  | [optional] 
**RuntimeClassName** | Pointer to **string** |  | [optional] 

## Methods

### NewDatabaseClusterSpecMonitoring

`func NewDatabaseClusterSpecMonitoring() *DatabaseClusterSpecMonitoring`

NewDatabaseClusterSpecMonitoring instantiates a new DatabaseClusterSpecMonitoring object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseClusterSpecMonitoringWithDefaults

`func NewDatabaseClusterSpecMonitoringWithDefaults() *DatabaseClusterSpecMonitoring`

NewDatabaseClusterSpecMonitoringWithDefaults instantiates a new DatabaseClusterSpecMonitoring object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainerSecurityContext

`func (o *DatabaseClusterSpecMonitoring) GetContainerSecurityContext() DatabaseClusterSpecBackupContainerSecurityContext`

GetContainerSecurityContext returns the ContainerSecurityContext field if non-nil, zero value otherwise.

### GetContainerSecurityContextOk

`func (o *DatabaseClusterSpecMonitoring) GetContainerSecurityContextOk() (*DatabaseClusterSpecBackupContainerSecurityContext, bool)`

GetContainerSecurityContextOk returns a tuple with the ContainerSecurityContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerSecurityContext

`func (o *DatabaseClusterSpecMonitoring) SetContainerSecurityContext(v DatabaseClusterSpecBackupContainerSecurityContext)`

SetContainerSecurityContext sets ContainerSecurityContext field to given value.

### HasContainerSecurityContext

`func (o *DatabaseClusterSpecMonitoring) HasContainerSecurityContext() bool`

HasContainerSecurityContext returns a boolean if a field has been set.

### GetImagePullPolicy

`func (o *DatabaseClusterSpecMonitoring) GetImagePullPolicy() string`

GetImagePullPolicy returns the ImagePullPolicy field if non-nil, zero value otherwise.

### GetImagePullPolicyOk

`func (o *DatabaseClusterSpecMonitoring) GetImagePullPolicyOk() (*string, bool)`

GetImagePullPolicyOk returns a tuple with the ImagePullPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImagePullPolicy

`func (o *DatabaseClusterSpecMonitoring) SetImagePullPolicy(v string)`

SetImagePullPolicy sets ImagePullPolicy field to given value.

### HasImagePullPolicy

`func (o *DatabaseClusterSpecMonitoring) HasImagePullPolicy() bool`

HasImagePullPolicy returns a boolean if a field has been set.

### GetPmm

`func (o *DatabaseClusterSpecMonitoring) GetPmm() DatabaseClusterSpecMonitoringPmm`

GetPmm returns the Pmm field if non-nil, zero value otherwise.

### GetPmmOk

`func (o *DatabaseClusterSpecMonitoring) GetPmmOk() (*DatabaseClusterSpecMonitoringPmm, bool)`

GetPmmOk returns a tuple with the Pmm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPmm

`func (o *DatabaseClusterSpecMonitoring) SetPmm(v DatabaseClusterSpecMonitoringPmm)`

SetPmm sets Pmm field to given value.

### HasPmm

`func (o *DatabaseClusterSpecMonitoring) HasPmm() bool`

HasPmm returns a boolean if a field has been set.

### GetResources

`func (o *DatabaseClusterSpecMonitoring) GetResources() DatabaseClusterSpecBackupResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *DatabaseClusterSpecMonitoring) GetResourcesOk() (*DatabaseClusterSpecBackupResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *DatabaseClusterSpecMonitoring) SetResources(v DatabaseClusterSpecBackupResources)`

SetResources sets Resources field to given value.

### HasResources

`func (o *DatabaseClusterSpecMonitoring) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetRuntimeClassName

`func (o *DatabaseClusterSpecMonitoring) GetRuntimeClassName() string`

GetRuntimeClassName returns the RuntimeClassName field if non-nil, zero value otherwise.

### GetRuntimeClassNameOk

`func (o *DatabaseClusterSpecMonitoring) GetRuntimeClassNameOk() (*string, bool)`

GetRuntimeClassNameOk returns a tuple with the RuntimeClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeClassName

`func (o *DatabaseClusterSpecMonitoring) SetRuntimeClassName(v string)`

SetRuntimeClassName sets RuntimeClassName field to given value.

### HasRuntimeClassName

`func (o *DatabaseClusterSpecMonitoring) HasRuntimeClassName() bool`

HasRuntimeClassName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


