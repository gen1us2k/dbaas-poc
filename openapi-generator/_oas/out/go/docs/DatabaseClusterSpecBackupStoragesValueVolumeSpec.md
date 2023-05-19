# DatabaseClusterSpecBackupStoragesValueVolumeSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmptyDir** | Pointer to [**DatabaseClusterSpecBackupStoragesValueVolumeSpecEmptyDir**](DatabaseClusterSpecBackupStoragesValueVolumeSpecEmptyDir.md) |  | [optional] 
**HostPath** | Pointer to [**DatabaseClusterSpecBackupStoragesValueVolumeSpecHostPath**](DatabaseClusterSpecBackupStoragesValueVolumeSpecHostPath.md) |  | [optional] 
**PersistentVolumeClaim** | Pointer to [**DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim**](DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim.md) |  | [optional] 

## Methods

### NewDatabaseClusterSpecBackupStoragesValueVolumeSpec

`func NewDatabaseClusterSpecBackupStoragesValueVolumeSpec() *DatabaseClusterSpecBackupStoragesValueVolumeSpec`

NewDatabaseClusterSpecBackupStoragesValueVolumeSpec instantiates a new DatabaseClusterSpecBackupStoragesValueVolumeSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseClusterSpecBackupStoragesValueVolumeSpecWithDefaults

`func NewDatabaseClusterSpecBackupStoragesValueVolumeSpecWithDefaults() *DatabaseClusterSpecBackupStoragesValueVolumeSpec`

NewDatabaseClusterSpecBackupStoragesValueVolumeSpecWithDefaults instantiates a new DatabaseClusterSpecBackupStoragesValueVolumeSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmptyDir

`func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpec) GetEmptyDir() DatabaseClusterSpecBackupStoragesValueVolumeSpecEmptyDir`

GetEmptyDir returns the EmptyDir field if non-nil, zero value otherwise.

### GetEmptyDirOk

`func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpec) GetEmptyDirOk() (*DatabaseClusterSpecBackupStoragesValueVolumeSpecEmptyDir, bool)`

GetEmptyDirOk returns a tuple with the EmptyDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmptyDir

`func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpec) SetEmptyDir(v DatabaseClusterSpecBackupStoragesValueVolumeSpecEmptyDir)`

SetEmptyDir sets EmptyDir field to given value.

### HasEmptyDir

`func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpec) HasEmptyDir() bool`

HasEmptyDir returns a boolean if a field has been set.

### GetHostPath

`func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpec) GetHostPath() DatabaseClusterSpecBackupStoragesValueVolumeSpecHostPath`

GetHostPath returns the HostPath field if non-nil, zero value otherwise.

### GetHostPathOk

`func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpec) GetHostPathOk() (*DatabaseClusterSpecBackupStoragesValueVolumeSpecHostPath, bool)`

GetHostPathOk returns a tuple with the HostPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostPath

`func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpec) SetHostPath(v DatabaseClusterSpecBackupStoragesValueVolumeSpecHostPath)`

SetHostPath sets HostPath field to given value.

### HasHostPath

`func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpec) HasHostPath() bool`

HasHostPath returns a boolean if a field has been set.

### GetPersistentVolumeClaim

`func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpec) GetPersistentVolumeClaim() DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim`

GetPersistentVolumeClaim returns the PersistentVolumeClaim field if non-nil, zero value otherwise.

### GetPersistentVolumeClaimOk

`func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpec) GetPersistentVolumeClaimOk() (*DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim, bool)`

GetPersistentVolumeClaimOk returns a tuple with the PersistentVolumeClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentVolumeClaim

`func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpec) SetPersistentVolumeClaim(v DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaim)`

SetPersistentVolumeClaim sets PersistentVolumeClaim field to given value.

### HasPersistentVolumeClaim

`func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpec) HasPersistentVolumeClaim() bool`

HasPersistentVolumeClaim returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


