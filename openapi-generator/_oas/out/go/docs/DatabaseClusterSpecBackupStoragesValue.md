# DatabaseClusterSpecBackupStoragesValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Affinity** | Pointer to [**DatabaseClusterSpecBackupStoragesValueAffinity**](DatabaseClusterSpecBackupStoragesValueAffinity.md) |  | [optional] 
**Annotations** | Pointer to **map[string]string** |  | [optional] 
**ContainerSecurityContext** | Pointer to [**DatabaseClusterSpecBackupContainerSecurityContext**](DatabaseClusterSpecBackupContainerSecurityContext.md) |  | [optional] 
**Labels** | Pointer to **map[string]string** |  | [optional] 
**NodeSelector** | Pointer to **map[string]string** |  | [optional] 
**PodSecurityContext** | Pointer to [**DatabaseClusterSpecBackupStoragesValuePodSecurityContext**](DatabaseClusterSpecBackupStoragesValuePodSecurityContext.md) |  | [optional] 
**PriorityClassName** | Pointer to **string** |  | [optional] 
**Resources** | Pointer to [**DatabaseClusterSpecBackupResources**](DatabaseClusterSpecBackupResources.md) |  | [optional] 
**RuntimeClassName** | Pointer to **string** |  | [optional] 
**SchedulerName** | Pointer to **string** |  | [optional] 
**StorageProvider** | Pointer to [**DatabaseClusterSpecBackupStoragesValueStorageProvider**](DatabaseClusterSpecBackupStoragesValueStorageProvider.md) |  | [optional] 
**Tolerations** | Pointer to [**[]DatabaseClusterSpecBackupStoragesValueTolerationsInner**](DatabaseClusterSpecBackupStoragesValueTolerationsInner.md) |  | [optional] 
**Type** | **string** | BackupStorageType represents backup storage type. | 
**VerifyTLS** | Pointer to **bool** |  | [optional] 
**VolumeSpec** | Pointer to [**DatabaseClusterSpecBackupStoragesValueVolumeSpec**](DatabaseClusterSpecBackupStoragesValueVolumeSpec.md) |  | [optional] 

## Methods

### NewDatabaseClusterSpecBackupStoragesValue

`func NewDatabaseClusterSpecBackupStoragesValue(type_ string, ) *DatabaseClusterSpecBackupStoragesValue`

NewDatabaseClusterSpecBackupStoragesValue instantiates a new DatabaseClusterSpecBackupStoragesValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseClusterSpecBackupStoragesValueWithDefaults

`func NewDatabaseClusterSpecBackupStoragesValueWithDefaults() *DatabaseClusterSpecBackupStoragesValue`

NewDatabaseClusterSpecBackupStoragesValueWithDefaults instantiates a new DatabaseClusterSpecBackupStoragesValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffinity

`func (o *DatabaseClusterSpecBackupStoragesValue) GetAffinity() DatabaseClusterSpecBackupStoragesValueAffinity`

GetAffinity returns the Affinity field if non-nil, zero value otherwise.

### GetAffinityOk

`func (o *DatabaseClusterSpecBackupStoragesValue) GetAffinityOk() (*DatabaseClusterSpecBackupStoragesValueAffinity, bool)`

GetAffinityOk returns a tuple with the Affinity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffinity

`func (o *DatabaseClusterSpecBackupStoragesValue) SetAffinity(v DatabaseClusterSpecBackupStoragesValueAffinity)`

SetAffinity sets Affinity field to given value.

### HasAffinity

`func (o *DatabaseClusterSpecBackupStoragesValue) HasAffinity() bool`

HasAffinity returns a boolean if a field has been set.

### GetAnnotations

`func (o *DatabaseClusterSpecBackupStoragesValue) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *DatabaseClusterSpecBackupStoragesValue) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *DatabaseClusterSpecBackupStoragesValue) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *DatabaseClusterSpecBackupStoragesValue) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetContainerSecurityContext

`func (o *DatabaseClusterSpecBackupStoragesValue) GetContainerSecurityContext() DatabaseClusterSpecBackupContainerSecurityContext`

GetContainerSecurityContext returns the ContainerSecurityContext field if non-nil, zero value otherwise.

### GetContainerSecurityContextOk

`func (o *DatabaseClusterSpecBackupStoragesValue) GetContainerSecurityContextOk() (*DatabaseClusterSpecBackupContainerSecurityContext, bool)`

GetContainerSecurityContextOk returns a tuple with the ContainerSecurityContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerSecurityContext

`func (o *DatabaseClusterSpecBackupStoragesValue) SetContainerSecurityContext(v DatabaseClusterSpecBackupContainerSecurityContext)`

SetContainerSecurityContext sets ContainerSecurityContext field to given value.

### HasContainerSecurityContext

`func (o *DatabaseClusterSpecBackupStoragesValue) HasContainerSecurityContext() bool`

HasContainerSecurityContext returns a boolean if a field has been set.

### GetLabels

`func (o *DatabaseClusterSpecBackupStoragesValue) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *DatabaseClusterSpecBackupStoragesValue) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *DatabaseClusterSpecBackupStoragesValue) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *DatabaseClusterSpecBackupStoragesValue) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetNodeSelector

`func (o *DatabaseClusterSpecBackupStoragesValue) GetNodeSelector() map[string]string`

GetNodeSelector returns the NodeSelector field if non-nil, zero value otherwise.

### GetNodeSelectorOk

`func (o *DatabaseClusterSpecBackupStoragesValue) GetNodeSelectorOk() (*map[string]string, bool)`

GetNodeSelectorOk returns a tuple with the NodeSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeSelector

`func (o *DatabaseClusterSpecBackupStoragesValue) SetNodeSelector(v map[string]string)`

SetNodeSelector sets NodeSelector field to given value.

### HasNodeSelector

`func (o *DatabaseClusterSpecBackupStoragesValue) HasNodeSelector() bool`

HasNodeSelector returns a boolean if a field has been set.

### GetPodSecurityContext

`func (o *DatabaseClusterSpecBackupStoragesValue) GetPodSecurityContext() DatabaseClusterSpecBackupStoragesValuePodSecurityContext`

GetPodSecurityContext returns the PodSecurityContext field if non-nil, zero value otherwise.

### GetPodSecurityContextOk

`func (o *DatabaseClusterSpecBackupStoragesValue) GetPodSecurityContextOk() (*DatabaseClusterSpecBackupStoragesValuePodSecurityContext, bool)`

GetPodSecurityContextOk returns a tuple with the PodSecurityContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodSecurityContext

`func (o *DatabaseClusterSpecBackupStoragesValue) SetPodSecurityContext(v DatabaseClusterSpecBackupStoragesValuePodSecurityContext)`

SetPodSecurityContext sets PodSecurityContext field to given value.

### HasPodSecurityContext

`func (o *DatabaseClusterSpecBackupStoragesValue) HasPodSecurityContext() bool`

HasPodSecurityContext returns a boolean if a field has been set.

### GetPriorityClassName

`func (o *DatabaseClusterSpecBackupStoragesValue) GetPriorityClassName() string`

GetPriorityClassName returns the PriorityClassName field if non-nil, zero value otherwise.

### GetPriorityClassNameOk

`func (o *DatabaseClusterSpecBackupStoragesValue) GetPriorityClassNameOk() (*string, bool)`

GetPriorityClassNameOk returns a tuple with the PriorityClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityClassName

`func (o *DatabaseClusterSpecBackupStoragesValue) SetPriorityClassName(v string)`

SetPriorityClassName sets PriorityClassName field to given value.

### HasPriorityClassName

`func (o *DatabaseClusterSpecBackupStoragesValue) HasPriorityClassName() bool`

HasPriorityClassName returns a boolean if a field has been set.

### GetResources

`func (o *DatabaseClusterSpecBackupStoragesValue) GetResources() DatabaseClusterSpecBackupResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *DatabaseClusterSpecBackupStoragesValue) GetResourcesOk() (*DatabaseClusterSpecBackupResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *DatabaseClusterSpecBackupStoragesValue) SetResources(v DatabaseClusterSpecBackupResources)`

SetResources sets Resources field to given value.

### HasResources

`func (o *DatabaseClusterSpecBackupStoragesValue) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetRuntimeClassName

`func (o *DatabaseClusterSpecBackupStoragesValue) GetRuntimeClassName() string`

GetRuntimeClassName returns the RuntimeClassName field if non-nil, zero value otherwise.

### GetRuntimeClassNameOk

`func (o *DatabaseClusterSpecBackupStoragesValue) GetRuntimeClassNameOk() (*string, bool)`

GetRuntimeClassNameOk returns a tuple with the RuntimeClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeClassName

`func (o *DatabaseClusterSpecBackupStoragesValue) SetRuntimeClassName(v string)`

SetRuntimeClassName sets RuntimeClassName field to given value.

### HasRuntimeClassName

`func (o *DatabaseClusterSpecBackupStoragesValue) HasRuntimeClassName() bool`

HasRuntimeClassName returns a boolean if a field has been set.

### GetSchedulerName

`func (o *DatabaseClusterSpecBackupStoragesValue) GetSchedulerName() string`

GetSchedulerName returns the SchedulerName field if non-nil, zero value otherwise.

### GetSchedulerNameOk

`func (o *DatabaseClusterSpecBackupStoragesValue) GetSchedulerNameOk() (*string, bool)`

GetSchedulerNameOk returns a tuple with the SchedulerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedulerName

`func (o *DatabaseClusterSpecBackupStoragesValue) SetSchedulerName(v string)`

SetSchedulerName sets SchedulerName field to given value.

### HasSchedulerName

`func (o *DatabaseClusterSpecBackupStoragesValue) HasSchedulerName() bool`

HasSchedulerName returns a boolean if a field has been set.

### GetStorageProvider

`func (o *DatabaseClusterSpecBackupStoragesValue) GetStorageProvider() DatabaseClusterSpecBackupStoragesValueStorageProvider`

GetStorageProvider returns the StorageProvider field if non-nil, zero value otherwise.

### GetStorageProviderOk

`func (o *DatabaseClusterSpecBackupStoragesValue) GetStorageProviderOk() (*DatabaseClusterSpecBackupStoragesValueStorageProvider, bool)`

GetStorageProviderOk returns a tuple with the StorageProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProvider

`func (o *DatabaseClusterSpecBackupStoragesValue) SetStorageProvider(v DatabaseClusterSpecBackupStoragesValueStorageProvider)`

SetStorageProvider sets StorageProvider field to given value.

### HasStorageProvider

`func (o *DatabaseClusterSpecBackupStoragesValue) HasStorageProvider() bool`

HasStorageProvider returns a boolean if a field has been set.

### GetTolerations

`func (o *DatabaseClusterSpecBackupStoragesValue) GetTolerations() []DatabaseClusterSpecBackupStoragesValueTolerationsInner`

GetTolerations returns the Tolerations field if non-nil, zero value otherwise.

### GetTolerationsOk

`func (o *DatabaseClusterSpecBackupStoragesValue) GetTolerationsOk() (*[]DatabaseClusterSpecBackupStoragesValueTolerationsInner, bool)`

GetTolerationsOk returns a tuple with the Tolerations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTolerations

`func (o *DatabaseClusterSpecBackupStoragesValue) SetTolerations(v []DatabaseClusterSpecBackupStoragesValueTolerationsInner)`

SetTolerations sets Tolerations field to given value.

### HasTolerations

`func (o *DatabaseClusterSpecBackupStoragesValue) HasTolerations() bool`

HasTolerations returns a boolean if a field has been set.

### GetType

`func (o *DatabaseClusterSpecBackupStoragesValue) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DatabaseClusterSpecBackupStoragesValue) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DatabaseClusterSpecBackupStoragesValue) SetType(v string)`

SetType sets Type field to given value.


### GetVerifyTLS

`func (o *DatabaseClusterSpecBackupStoragesValue) GetVerifyTLS() bool`

GetVerifyTLS returns the VerifyTLS field if non-nil, zero value otherwise.

### GetVerifyTLSOk

`func (o *DatabaseClusterSpecBackupStoragesValue) GetVerifyTLSOk() (*bool, bool)`

GetVerifyTLSOk returns a tuple with the VerifyTLS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyTLS

`func (o *DatabaseClusterSpecBackupStoragesValue) SetVerifyTLS(v bool)`

SetVerifyTLS sets VerifyTLS field to given value.

### HasVerifyTLS

`func (o *DatabaseClusterSpecBackupStoragesValue) HasVerifyTLS() bool`

HasVerifyTLS returns a boolean if a field has been set.

### GetVolumeSpec

`func (o *DatabaseClusterSpecBackupStoragesValue) GetVolumeSpec() DatabaseClusterSpecBackupStoragesValueVolumeSpec`

GetVolumeSpec returns the VolumeSpec field if non-nil, zero value otherwise.

### GetVolumeSpecOk

`func (o *DatabaseClusterSpecBackupStoragesValue) GetVolumeSpecOk() (*DatabaseClusterSpecBackupStoragesValueVolumeSpec, bool)`

GetVolumeSpecOk returns a tuple with the VolumeSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeSpec

`func (o *DatabaseClusterSpecBackupStoragesValue) SetVolumeSpec(v DatabaseClusterSpecBackupStoragesValueVolumeSpec)`

SetVolumeSpec sets VolumeSpec field to given value.

### HasVolumeSpec

`func (o *DatabaseClusterSpecBackupStoragesValue) HasVolumeSpec() bool`

HasVolumeSpec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


