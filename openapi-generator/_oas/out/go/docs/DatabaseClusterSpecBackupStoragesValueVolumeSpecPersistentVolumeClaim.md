# DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessModes** | Pointer to **[]string** | accessModes contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1 | [optional] 
**DataSource** | Pointer to [**DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSource**](DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSource.md) |  | [optional] 
**DataSourceRef** | Pointer to [**DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRef**](DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRef.md) |  | [optional] 
**Resources** | Pointer to [**DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimResources**](DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimResources.md) |  | [optional] 
**Selector** | Pointer to [**DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimSelector**](DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimSelector.md) |  | [optional] 
**StorageClassName** | Pointer to **string** | storageClassName is the name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1 | [optional] 
**VolumeMode** | Pointer to **string** | volumeMode defines what type of volume is required by the claim. Value of Filesystem is implied when not included in claim spec. | [optional] 
**VolumeName** | Pointer to **string** | volumeName is the binding reference to the PersistentVolume backing this claim. | [optional] 

## Methods

### NewDatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim

`func NewDatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim() *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim`

NewDatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim instantiates a new DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimWithDefaults

`func NewDatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimWithDefaults() *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim`

NewDatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimWithDefaults instantiates a new DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessModes

`func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim) GetAccessModes() []string`

GetAccessModes returns the AccessModes field if non-nil, zero value otherwise.

### GetAccessModesOk

`func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim) GetAccessModesOk() (*[]string, bool)`

GetAccessModesOk returns a tuple with the AccessModes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessModes

`func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim) SetAccessModes(v []string)`

SetAccessModes sets AccessModes field to given value.

### HasAccessModes

`func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim) HasAccessModes() bool`

HasAccessModes returns a boolean if a field has been set.

### GetDataSource

`func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim) GetDataSource() DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSource`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim) GetDataSourceOk() (*DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSource, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim) SetDataSource(v DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSource)`

SetDataSource sets DataSource field to given value.

### HasDataSource

`func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim) HasDataSource() bool`

HasDataSource returns a boolean if a field has been set.

### GetDataSourceRef

`func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim) GetDataSourceRef() DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRef`

GetDataSourceRef returns the DataSourceRef field if non-nil, zero value otherwise.

### GetDataSourceRefOk

`func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim) GetDataSourceRefOk() (*DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRef, bool)`

GetDataSourceRefOk returns a tuple with the DataSourceRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceRef

`func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim) SetDataSourceRef(v DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRef)`

SetDataSourceRef sets DataSourceRef field to given value.

### HasDataSourceRef

`func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim) HasDataSourceRef() bool`

HasDataSourceRef returns a boolean if a field has been set.

### GetResources

`func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim) GetResources() DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim) GetResourcesOk() (*DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim) SetResources(v DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimResources)`

SetResources sets Resources field to given value.

### HasResources

`func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetSelector

`func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim) GetSelector() DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimSelector`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim) GetSelectorOk() (*DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimSelector, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim) SetSelector(v DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimSelector)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim) HasSelector() bool`

HasSelector returns a boolean if a field has been set.

### GetStorageClassName

`func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim) GetStorageClassName() string`

GetStorageClassName returns the StorageClassName field if non-nil, zero value otherwise.

### GetStorageClassNameOk

`func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim) GetStorageClassNameOk() (*string, bool)`

GetStorageClassNameOk returns a tuple with the StorageClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClassName

`func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim) SetStorageClassName(v string)`

SetStorageClassName sets StorageClassName field to given value.

### HasStorageClassName

`func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim) HasStorageClassName() bool`

HasStorageClassName returns a boolean if a field has been set.

### GetVolumeMode

`func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim) GetVolumeMode() string`

GetVolumeMode returns the VolumeMode field if non-nil, zero value otherwise.

### GetVolumeModeOk

`func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim) GetVolumeModeOk() (*string, bool)`

GetVolumeModeOk returns a tuple with the VolumeMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeMode

`func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim) SetVolumeMode(v string)`

SetVolumeMode sets VolumeMode field to given value.

### HasVolumeMode

`func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim) HasVolumeMode() bool`

HasVolumeMode returns a boolean if a field has been set.

### GetVolumeName

`func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim) GetVolumeName() string`

GetVolumeName returns the VolumeName field if non-nil, zero value otherwise.

### GetVolumeNameOk

`func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim) GetVolumeNameOk() (*string, bool)`

GetVolumeNameOk returns a tuple with the VolumeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeName

`func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim) SetVolumeName(v string)`

SetVolumeName sets VolumeName field to given value.

### HasVolumeName

`func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim) HasVolumeName() bool`

HasVolumeName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


