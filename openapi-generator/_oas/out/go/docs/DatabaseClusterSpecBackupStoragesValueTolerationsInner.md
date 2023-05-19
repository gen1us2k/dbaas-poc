# DatabaseClusterSpecBackupStoragesValueTolerationsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Effect** | Pointer to **string** | Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute. | [optional] 
**Key** | Pointer to **string** | Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys. | [optional] 
**Operator** | Pointer to **string** | Operator represents a key&#39;s relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category. | [optional] 
**TolerationSeconds** | Pointer to **int64** | TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system. | [optional] 
**Value** | Pointer to **string** | Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string. | [optional] 

## Methods

### NewDatabaseClusterSpecBackupStoragesValueTolerationsInner

`func NewDatabaseClusterSpecBackupStoragesValueTolerationsInner() *DatabaseClusterSpecBackupStoragesValueTolerationsInner`

NewDatabaseClusterSpecBackupStoragesValueTolerationsInner instantiates a new DatabaseClusterSpecBackupStoragesValueTolerationsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseClusterSpecBackupStoragesValueTolerationsInnerWithDefaults

`func NewDatabaseClusterSpecBackupStoragesValueTolerationsInnerWithDefaults() *DatabaseClusterSpecBackupStoragesValueTolerationsInner`

NewDatabaseClusterSpecBackupStoragesValueTolerationsInnerWithDefaults instantiates a new DatabaseClusterSpecBackupStoragesValueTolerationsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEffect

`func (o *DatabaseClusterSpecBackupStoragesValueTolerationsInner) GetEffect() string`

GetEffect returns the Effect field if non-nil, zero value otherwise.

### GetEffectOk

`func (o *DatabaseClusterSpecBackupStoragesValueTolerationsInner) GetEffectOk() (*string, bool)`

GetEffectOk returns a tuple with the Effect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffect

`func (o *DatabaseClusterSpecBackupStoragesValueTolerationsInner) SetEffect(v string)`

SetEffect sets Effect field to given value.

### HasEffect

`func (o *DatabaseClusterSpecBackupStoragesValueTolerationsInner) HasEffect() bool`

HasEffect returns a boolean if a field has been set.

### GetKey

`func (o *DatabaseClusterSpecBackupStoragesValueTolerationsInner) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *DatabaseClusterSpecBackupStoragesValueTolerationsInner) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *DatabaseClusterSpecBackupStoragesValueTolerationsInner) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *DatabaseClusterSpecBackupStoragesValueTolerationsInner) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetOperator

`func (o *DatabaseClusterSpecBackupStoragesValueTolerationsInner) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *DatabaseClusterSpecBackupStoragesValueTolerationsInner) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *DatabaseClusterSpecBackupStoragesValueTolerationsInner) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *DatabaseClusterSpecBackupStoragesValueTolerationsInner) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetTolerationSeconds

`func (o *DatabaseClusterSpecBackupStoragesValueTolerationsInner) GetTolerationSeconds() int64`

GetTolerationSeconds returns the TolerationSeconds field if non-nil, zero value otherwise.

### GetTolerationSecondsOk

`func (o *DatabaseClusterSpecBackupStoragesValueTolerationsInner) GetTolerationSecondsOk() (*int64, bool)`

GetTolerationSecondsOk returns a tuple with the TolerationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTolerationSeconds

`func (o *DatabaseClusterSpecBackupStoragesValueTolerationsInner) SetTolerationSeconds(v int64)`

SetTolerationSeconds sets TolerationSeconds field to given value.

### HasTolerationSeconds

`func (o *DatabaseClusterSpecBackupStoragesValueTolerationsInner) HasTolerationSeconds() bool`

HasTolerationSeconds returns a boolean if a field has been set.

### GetValue

`func (o *DatabaseClusterSpecBackupStoragesValueTolerationsInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DatabaseClusterSpecBackupStoragesValueTolerationsInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DatabaseClusterSpecBackupStoragesValueTolerationsInner) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *DatabaseClusterSpecBackupStoragesValueTolerationsInner) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


