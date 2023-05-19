# DatabaseClusterSpecBackupResourcesClaimsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container. | 

## Methods

### NewDatabaseClusterSpecBackupResourcesClaimsInner

`func NewDatabaseClusterSpecBackupResourcesClaimsInner(name string, ) *DatabaseClusterSpecBackupResourcesClaimsInner`

NewDatabaseClusterSpecBackupResourcesClaimsInner instantiates a new DatabaseClusterSpecBackupResourcesClaimsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseClusterSpecBackupResourcesClaimsInnerWithDefaults

`func NewDatabaseClusterSpecBackupResourcesClaimsInnerWithDefaults() *DatabaseClusterSpecBackupResourcesClaimsInner`

NewDatabaseClusterSpecBackupResourcesClaimsInnerWithDefaults instantiates a new DatabaseClusterSpecBackupResourcesClaimsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DatabaseClusterSpecBackupResourcesClaimsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DatabaseClusterSpecBackupResourcesClaimsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DatabaseClusterSpecBackupResourcesClaimsInner) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


