# DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiGroup** | Pointer to **string** | APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required. | [optional] 
**Kind** | **string** | Kind is the type of resource being referenced | 
**Name** | **string** | Name is the name of resource being referenced | 
**Namespace** | Pointer to **string** | Namespace is the namespace of resource being referenced Note that when a namespace is specified, a gateway.networking.k8s.io/ReferenceGrant object is required in the referent namespace to allow that namespace&#39;s owner to accept the reference. See the ReferenceGrant documentation for details. (Alpha) This field requires the CrossNamespaceVolumeDataSource feature gate to be enabled. | [optional] 

## Methods

### NewDatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRef

`func NewDatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRef(kind string, name string, ) *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRef`

NewDatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRef instantiates a new DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRefWithDefaults

`func NewDatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRefWithDefaults() *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRef`

NewDatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRefWithDefaults instantiates a new DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiGroup

`func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRef) GetApiGroup() string`

GetApiGroup returns the ApiGroup field if non-nil, zero value otherwise.

### GetApiGroupOk

`func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRef) GetApiGroupOk() (*string, bool)`

GetApiGroupOk returns a tuple with the ApiGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiGroup

`func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRef) SetApiGroup(v string)`

SetApiGroup sets ApiGroup field to given value.

### HasApiGroup

`func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRef) HasApiGroup() bool`

HasApiGroup returns a boolean if a field has been set.

### GetKind

`func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRef) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRef) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRef) SetKind(v string)`

SetKind sets Kind field to given value.


### GetName

`func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRef) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRef) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRef) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRef) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRef) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRef) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimDataSourceRef) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


