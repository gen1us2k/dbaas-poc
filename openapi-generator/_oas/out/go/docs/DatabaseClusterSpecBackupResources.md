# DatabaseClusterSpecBackupResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Claims** | Pointer to [**[]DatabaseClusterSpecBackupResourcesClaimsInner**](DatabaseClusterSpecBackupResourcesClaimsInner.md) | Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.   This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.   This field is immutable. It can only be set for containers. | [optional] 
**Limits** | Pointer to [**map[string]DatabaseClusterSpecBackupResourcesLimitsValue**](DatabaseClusterSpecBackupResourcesLimitsValue.md) | Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/ | [optional] 
**Requests** | Pointer to [**map[string]DatabaseClusterSpecBackupResourcesLimitsValue**](DatabaseClusterSpecBackupResourcesLimitsValue.md) | Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/ | [optional] 

## Methods

### NewDatabaseClusterSpecBackupResources

`func NewDatabaseClusterSpecBackupResources() *DatabaseClusterSpecBackupResources`

NewDatabaseClusterSpecBackupResources instantiates a new DatabaseClusterSpecBackupResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseClusterSpecBackupResourcesWithDefaults

`func NewDatabaseClusterSpecBackupResourcesWithDefaults() *DatabaseClusterSpecBackupResources`

NewDatabaseClusterSpecBackupResourcesWithDefaults instantiates a new DatabaseClusterSpecBackupResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClaims

`func (o *DatabaseClusterSpecBackupResources) GetClaims() []DatabaseClusterSpecBackupResourcesClaimsInner`

GetClaims returns the Claims field if non-nil, zero value otherwise.

### GetClaimsOk

`func (o *DatabaseClusterSpecBackupResources) GetClaimsOk() (*[]DatabaseClusterSpecBackupResourcesClaimsInner, bool)`

GetClaimsOk returns a tuple with the Claims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaims

`func (o *DatabaseClusterSpecBackupResources) SetClaims(v []DatabaseClusterSpecBackupResourcesClaimsInner)`

SetClaims sets Claims field to given value.

### HasClaims

`func (o *DatabaseClusterSpecBackupResources) HasClaims() bool`

HasClaims returns a boolean if a field has been set.

### GetLimits

`func (o *DatabaseClusterSpecBackupResources) GetLimits() map[string]DatabaseClusterSpecBackupResourcesLimitsValue`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *DatabaseClusterSpecBackupResources) GetLimitsOk() (*map[string]DatabaseClusterSpecBackupResourcesLimitsValue, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *DatabaseClusterSpecBackupResources) SetLimits(v map[string]DatabaseClusterSpecBackupResourcesLimitsValue)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *DatabaseClusterSpecBackupResources) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### GetRequests

`func (o *DatabaseClusterSpecBackupResources) GetRequests() map[string]DatabaseClusterSpecBackupResourcesLimitsValue`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *DatabaseClusterSpecBackupResources) GetRequestsOk() (*map[string]DatabaseClusterSpecBackupResourcesLimitsValue, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *DatabaseClusterSpecBackupResources) SetRequests(v map[string]DatabaseClusterSpecBackupResourcesLimitsValue)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *DatabaseClusterSpecBackupResources) HasRequests() bool`

HasRequests returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


