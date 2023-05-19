# DatabaseClusterStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Ready** | Pointer to **int32** |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** | AppState is used to represent cluster&#39;s state. | [optional] 

## Methods

### NewDatabaseClusterStatus

`func NewDatabaseClusterStatus() *DatabaseClusterStatus`

NewDatabaseClusterStatus instantiates a new DatabaseClusterStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseClusterStatusWithDefaults

`func NewDatabaseClusterStatusWithDefaults() *DatabaseClusterStatus`

NewDatabaseClusterStatusWithDefaults instantiates a new DatabaseClusterStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *DatabaseClusterStatus) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *DatabaseClusterStatus) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *DatabaseClusterStatus) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *DatabaseClusterStatus) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetMessage

`func (o *DatabaseClusterStatus) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DatabaseClusterStatus) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DatabaseClusterStatus) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *DatabaseClusterStatus) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetReady

`func (o *DatabaseClusterStatus) GetReady() int32`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *DatabaseClusterStatus) GetReadyOk() (*int32, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *DatabaseClusterStatus) SetReady(v int32)`

SetReady sets Ready field to given value.

### HasReady

`func (o *DatabaseClusterStatus) HasReady() bool`

HasReady returns a boolean if a field has been set.

### GetSize

`func (o *DatabaseClusterStatus) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *DatabaseClusterStatus) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *DatabaseClusterStatus) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *DatabaseClusterStatus) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetStatus

`func (o *DatabaseClusterStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DatabaseClusterStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DatabaseClusterStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DatabaseClusterStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


