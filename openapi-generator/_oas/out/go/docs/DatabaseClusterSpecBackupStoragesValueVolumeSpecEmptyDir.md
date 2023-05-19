# DatabaseClusterSpecBackupStoragesValueVolumeSpecEmptyDir

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Medium** | Pointer to **string** | medium represents what type of storage medium should back this directory. The default is \&quot;\&quot; which means to use the node&#39;s default medium. Must be an empty string (default) or Memory. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir | [optional] 
**SizeLimit** | Pointer to [**DatabaseClusterSpecBackupStoragesValueVolumeSpecEmptyDirSizeLimit**](DatabaseClusterSpecBackupStoragesValueVolumeSpecEmptyDirSizeLimit.md) |  | [optional] 

## Methods

### NewDatabaseClusterSpecBackupStoragesValueVolumeSpecEmptyDir

`func NewDatabaseClusterSpecBackupStoragesValueVolumeSpecEmptyDir() *DatabaseClusterSpecBackupStoragesValueVolumeSpecEmptyDir`

NewDatabaseClusterSpecBackupStoragesValueVolumeSpecEmptyDir instantiates a new DatabaseClusterSpecBackupStoragesValueVolumeSpecEmptyDir object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseClusterSpecBackupStoragesValueVolumeSpecEmptyDirWithDefaults

`func NewDatabaseClusterSpecBackupStoragesValueVolumeSpecEmptyDirWithDefaults() *DatabaseClusterSpecBackupStoragesValueVolumeSpecEmptyDir`

NewDatabaseClusterSpecBackupStoragesValueVolumeSpecEmptyDirWithDefaults instantiates a new DatabaseClusterSpecBackupStoragesValueVolumeSpecEmptyDir object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMedium

`func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecEmptyDir) GetMedium() string`

GetMedium returns the Medium field if non-nil, zero value otherwise.

### GetMediumOk

`func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecEmptyDir) GetMediumOk() (*string, bool)`

GetMediumOk returns a tuple with the Medium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedium

`func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecEmptyDir) SetMedium(v string)`

SetMedium sets Medium field to given value.

### HasMedium

`func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecEmptyDir) HasMedium() bool`

HasMedium returns a boolean if a field has been set.

### GetSizeLimit

`func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecEmptyDir) GetSizeLimit() DatabaseClusterSpecBackupStoragesValueVolumeSpecEmptyDirSizeLimit`

GetSizeLimit returns the SizeLimit field if non-nil, zero value otherwise.

### GetSizeLimitOk

`func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecEmptyDir) GetSizeLimitOk() (*DatabaseClusterSpecBackupStoragesValueVolumeSpecEmptyDirSizeLimit, bool)`

GetSizeLimitOk returns a tuple with the SizeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeLimit

`func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecEmptyDir) SetSizeLimit(v DatabaseClusterSpecBackupStoragesValueVolumeSpecEmptyDirSizeLimit)`

SetSizeLimit sets SizeLimit field to given value.

### HasSizeLimit

`func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecEmptyDir) HasSizeLimit() bool`

HasSizeLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


