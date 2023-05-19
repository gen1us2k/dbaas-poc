# DatabaseCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources | [optional] 
**Kind** | Pointer to **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**Spec** | Pointer to [**DatabaseClusterSpec**](DatabaseClusterSpec.md) |  | [optional] 
**Status** | Pointer to [**DatabaseClusterStatus**](DatabaseClusterStatus.md) |  | [optional] 

## Methods

### NewDatabaseCluster

`func NewDatabaseCluster() *DatabaseCluster`

NewDatabaseCluster instantiates a new DatabaseCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseClusterWithDefaults

`func NewDatabaseClusterWithDefaults() *DatabaseCluster`

NewDatabaseClusterWithDefaults instantiates a new DatabaseCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *DatabaseCluster) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *DatabaseCluster) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *DatabaseCluster) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *DatabaseCluster) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *DatabaseCluster) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *DatabaseCluster) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *DatabaseCluster) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *DatabaseCluster) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *DatabaseCluster) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DatabaseCluster) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DatabaseCluster) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DatabaseCluster) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *DatabaseCluster) GetSpec() DatabaseClusterSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *DatabaseCluster) GetSpecOk() (*DatabaseClusterSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *DatabaseCluster) SetSpec(v DatabaseClusterSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *DatabaseCluster) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *DatabaseCluster) GetStatus() DatabaseClusterStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DatabaseCluster) GetStatusOk() (*DatabaseClusterStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DatabaseCluster) SetStatus(v DatabaseClusterStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DatabaseCluster) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


