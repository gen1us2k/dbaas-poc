# DatabaseClusterSpecBackupStoragesValueAffinity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeAffinity** | Pointer to [**DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinity**](DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinity.md) |  | [optional] 
**PodAffinity** | Pointer to [**DatabaseClusterSpecBackupStoragesValueAffinityPodAffinity**](DatabaseClusterSpecBackupStoragesValueAffinityPodAffinity.md) |  | [optional] 
**PodAntiAffinity** | Pointer to [**DatabaseClusterSpecBackupStoragesValueAffinityPodAntiAffinity**](DatabaseClusterSpecBackupStoragesValueAffinityPodAntiAffinity.md) |  | [optional] 

## Methods

### NewDatabaseClusterSpecBackupStoragesValueAffinity

`func NewDatabaseClusterSpecBackupStoragesValueAffinity() *DatabaseClusterSpecBackupStoragesValueAffinity`

NewDatabaseClusterSpecBackupStoragesValueAffinity instantiates a new DatabaseClusterSpecBackupStoragesValueAffinity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseClusterSpecBackupStoragesValueAffinityWithDefaults

`func NewDatabaseClusterSpecBackupStoragesValueAffinityWithDefaults() *DatabaseClusterSpecBackupStoragesValueAffinity`

NewDatabaseClusterSpecBackupStoragesValueAffinityWithDefaults instantiates a new DatabaseClusterSpecBackupStoragesValueAffinity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeAffinity

`func (o *DatabaseClusterSpecBackupStoragesValueAffinity) GetNodeAffinity() DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinity`

GetNodeAffinity returns the NodeAffinity field if non-nil, zero value otherwise.

### GetNodeAffinityOk

`func (o *DatabaseClusterSpecBackupStoragesValueAffinity) GetNodeAffinityOk() (*DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinity, bool)`

GetNodeAffinityOk returns a tuple with the NodeAffinity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeAffinity

`func (o *DatabaseClusterSpecBackupStoragesValueAffinity) SetNodeAffinity(v DatabaseClusterSpecBackupStoragesValueAffinityNodeAffinity)`

SetNodeAffinity sets NodeAffinity field to given value.

### HasNodeAffinity

`func (o *DatabaseClusterSpecBackupStoragesValueAffinity) HasNodeAffinity() bool`

HasNodeAffinity returns a boolean if a field has been set.

### GetPodAffinity

`func (o *DatabaseClusterSpecBackupStoragesValueAffinity) GetPodAffinity() DatabaseClusterSpecBackupStoragesValueAffinityPodAffinity`

GetPodAffinity returns the PodAffinity field if non-nil, zero value otherwise.

### GetPodAffinityOk

`func (o *DatabaseClusterSpecBackupStoragesValueAffinity) GetPodAffinityOk() (*DatabaseClusterSpecBackupStoragesValueAffinityPodAffinity, bool)`

GetPodAffinityOk returns a tuple with the PodAffinity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodAffinity

`func (o *DatabaseClusterSpecBackupStoragesValueAffinity) SetPodAffinity(v DatabaseClusterSpecBackupStoragesValueAffinityPodAffinity)`

SetPodAffinity sets PodAffinity field to given value.

### HasPodAffinity

`func (o *DatabaseClusterSpecBackupStoragesValueAffinity) HasPodAffinity() bool`

HasPodAffinity returns a boolean if a field has been set.

### GetPodAntiAffinity

`func (o *DatabaseClusterSpecBackupStoragesValueAffinity) GetPodAntiAffinity() DatabaseClusterSpecBackupStoragesValueAffinityPodAntiAffinity`

GetPodAntiAffinity returns the PodAntiAffinity field if non-nil, zero value otherwise.

### GetPodAntiAffinityOk

`func (o *DatabaseClusterSpecBackupStoragesValueAffinity) GetPodAntiAffinityOk() (*DatabaseClusterSpecBackupStoragesValueAffinityPodAntiAffinity, bool)`

GetPodAntiAffinityOk returns a tuple with the PodAntiAffinity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodAntiAffinity

`func (o *DatabaseClusterSpecBackupStoragesValueAffinity) SetPodAntiAffinity(v DatabaseClusterSpecBackupStoragesValueAffinityPodAntiAffinity)`

SetPodAntiAffinity sets PodAntiAffinity field to given value.

### HasPodAntiAffinity

`func (o *DatabaseClusterSpecBackupStoragesValueAffinity) HasPodAntiAffinity() bool`

HasPodAntiAffinity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


