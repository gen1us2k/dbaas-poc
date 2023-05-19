# DatabaseClusterSpecDbInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpu** | Pointer to [**DatabaseClusterSpecBackupResourcesLimitsValue**](DatabaseClusterSpecBackupResourcesLimitsValue.md) |  | [optional] 
**DiskSize** | Pointer to [**DatabaseClusterSpecBackupResourcesLimitsValue**](DatabaseClusterSpecBackupResourcesLimitsValue.md) |  | [optional] 
**Memory** | Pointer to [**DatabaseClusterSpecBackupResourcesLimitsValue**](DatabaseClusterSpecBackupResourcesLimitsValue.md) |  | [optional] 
**StorageClassName** | Pointer to **string** |  | [optional] 

## Methods

### NewDatabaseClusterSpecDbInstance

`func NewDatabaseClusterSpecDbInstance() *DatabaseClusterSpecDbInstance`

NewDatabaseClusterSpecDbInstance instantiates a new DatabaseClusterSpecDbInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseClusterSpecDbInstanceWithDefaults

`func NewDatabaseClusterSpecDbInstanceWithDefaults() *DatabaseClusterSpecDbInstance`

NewDatabaseClusterSpecDbInstanceWithDefaults instantiates a new DatabaseClusterSpecDbInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpu

`func (o *DatabaseClusterSpecDbInstance) GetCpu() DatabaseClusterSpecBackupResourcesLimitsValue`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *DatabaseClusterSpecDbInstance) GetCpuOk() (*DatabaseClusterSpecBackupResourcesLimitsValue, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *DatabaseClusterSpecDbInstance) SetCpu(v DatabaseClusterSpecBackupResourcesLimitsValue)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *DatabaseClusterSpecDbInstance) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetDiskSize

`func (o *DatabaseClusterSpecDbInstance) GetDiskSize() DatabaseClusterSpecBackupResourcesLimitsValue`

GetDiskSize returns the DiskSize field if non-nil, zero value otherwise.

### GetDiskSizeOk

`func (o *DatabaseClusterSpecDbInstance) GetDiskSizeOk() (*DatabaseClusterSpecBackupResourcesLimitsValue, bool)`

GetDiskSizeOk returns a tuple with the DiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSize

`func (o *DatabaseClusterSpecDbInstance) SetDiskSize(v DatabaseClusterSpecBackupResourcesLimitsValue)`

SetDiskSize sets DiskSize field to given value.

### HasDiskSize

`func (o *DatabaseClusterSpecDbInstance) HasDiskSize() bool`

HasDiskSize returns a boolean if a field has been set.

### GetMemory

`func (o *DatabaseClusterSpecDbInstance) GetMemory() DatabaseClusterSpecBackupResourcesLimitsValue`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *DatabaseClusterSpecDbInstance) GetMemoryOk() (*DatabaseClusterSpecBackupResourcesLimitsValue, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *DatabaseClusterSpecDbInstance) SetMemory(v DatabaseClusterSpecBackupResourcesLimitsValue)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *DatabaseClusterSpecDbInstance) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetStorageClassName

`func (o *DatabaseClusterSpecDbInstance) GetStorageClassName() string`

GetStorageClassName returns the StorageClassName field if non-nil, zero value otherwise.

### GetStorageClassNameOk

`func (o *DatabaseClusterSpecDbInstance) GetStorageClassNameOk() (*string, bool)`

GetStorageClassNameOk returns a tuple with the StorageClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClassName

`func (o *DatabaseClusterSpecDbInstance) SetStorageClassName(v string)`

SetStorageClassName sets StorageClassName field to given value.

### HasStorageClassName

`func (o *DatabaseClusterSpecDbInstance) HasStorageClassName() bool`

HasStorageClassName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


