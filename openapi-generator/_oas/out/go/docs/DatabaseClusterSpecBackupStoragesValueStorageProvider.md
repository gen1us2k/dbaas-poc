# DatabaseClusterSpecBackupStoragesValueStorageProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | Pointer to **string** |  | [optional] 
**ContainerName** | Pointer to **string** | A container name is a valid DNS name that conforms to the Azure naming rules. | [optional] 
**CredentialsSecret** | **string** |  | 
**EndpointUrl** | Pointer to **string** |  | [optional] 
**Prefix** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**StorageClass** | Pointer to **string** | STANDARD, NEARLINE, COLDLINE, ARCHIVE for GCP Hot (Frequently accessed or modified data), Cool (Infrequently accessed or modified data), Archive (Rarely accessed or modified data) for Azure. | [optional] 

## Methods

### NewDatabaseClusterSpecBackupStoragesValueStorageProvider

`func NewDatabaseClusterSpecBackupStoragesValueStorageProvider(credentialsSecret string, ) *DatabaseClusterSpecBackupStoragesValueStorageProvider`

NewDatabaseClusterSpecBackupStoragesValueStorageProvider instantiates a new DatabaseClusterSpecBackupStoragesValueStorageProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseClusterSpecBackupStoragesValueStorageProviderWithDefaults

`func NewDatabaseClusterSpecBackupStoragesValueStorageProviderWithDefaults() *DatabaseClusterSpecBackupStoragesValueStorageProvider`

NewDatabaseClusterSpecBackupStoragesValueStorageProviderWithDefaults instantiates a new DatabaseClusterSpecBackupStoragesValueStorageProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucket

`func (o *DatabaseClusterSpecBackupStoragesValueStorageProvider) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *DatabaseClusterSpecBackupStoragesValueStorageProvider) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *DatabaseClusterSpecBackupStoragesValueStorageProvider) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *DatabaseClusterSpecBackupStoragesValueStorageProvider) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetContainerName

`func (o *DatabaseClusterSpecBackupStoragesValueStorageProvider) GetContainerName() string`

GetContainerName returns the ContainerName field if non-nil, zero value otherwise.

### GetContainerNameOk

`func (o *DatabaseClusterSpecBackupStoragesValueStorageProvider) GetContainerNameOk() (*string, bool)`

GetContainerNameOk returns a tuple with the ContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerName

`func (o *DatabaseClusterSpecBackupStoragesValueStorageProvider) SetContainerName(v string)`

SetContainerName sets ContainerName field to given value.

### HasContainerName

`func (o *DatabaseClusterSpecBackupStoragesValueStorageProvider) HasContainerName() bool`

HasContainerName returns a boolean if a field has been set.

### GetCredentialsSecret

`func (o *DatabaseClusterSpecBackupStoragesValueStorageProvider) GetCredentialsSecret() string`

GetCredentialsSecret returns the CredentialsSecret field if non-nil, zero value otherwise.

### GetCredentialsSecretOk

`func (o *DatabaseClusterSpecBackupStoragesValueStorageProvider) GetCredentialsSecretOk() (*string, bool)`

GetCredentialsSecretOk returns a tuple with the CredentialsSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialsSecret

`func (o *DatabaseClusterSpecBackupStoragesValueStorageProvider) SetCredentialsSecret(v string)`

SetCredentialsSecret sets CredentialsSecret field to given value.


### GetEndpointUrl

`func (o *DatabaseClusterSpecBackupStoragesValueStorageProvider) GetEndpointUrl() string`

GetEndpointUrl returns the EndpointUrl field if non-nil, zero value otherwise.

### GetEndpointUrlOk

`func (o *DatabaseClusterSpecBackupStoragesValueStorageProvider) GetEndpointUrlOk() (*string, bool)`

GetEndpointUrlOk returns a tuple with the EndpointUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointUrl

`func (o *DatabaseClusterSpecBackupStoragesValueStorageProvider) SetEndpointUrl(v string)`

SetEndpointUrl sets EndpointUrl field to given value.

### HasEndpointUrl

`func (o *DatabaseClusterSpecBackupStoragesValueStorageProvider) HasEndpointUrl() bool`

HasEndpointUrl returns a boolean if a field has been set.

### GetPrefix

`func (o *DatabaseClusterSpecBackupStoragesValueStorageProvider) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *DatabaseClusterSpecBackupStoragesValueStorageProvider) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *DatabaseClusterSpecBackupStoragesValueStorageProvider) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *DatabaseClusterSpecBackupStoragesValueStorageProvider) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetRegion

`func (o *DatabaseClusterSpecBackupStoragesValueStorageProvider) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *DatabaseClusterSpecBackupStoragesValueStorageProvider) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *DatabaseClusterSpecBackupStoragesValueStorageProvider) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *DatabaseClusterSpecBackupStoragesValueStorageProvider) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetStorageClass

`func (o *DatabaseClusterSpecBackupStoragesValueStorageProvider) GetStorageClass() string`

GetStorageClass returns the StorageClass field if non-nil, zero value otherwise.

### GetStorageClassOk

`func (o *DatabaseClusterSpecBackupStoragesValueStorageProvider) GetStorageClassOk() (*string, bool)`

GetStorageClassOk returns a tuple with the StorageClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClass

`func (o *DatabaseClusterSpecBackupStoragesValueStorageProvider) SetStorageClass(v string)`

SetStorageClass sets StorageClass field to given value.

### HasStorageClass

`func (o *DatabaseClusterSpecBackupStoragesValueStorageProvider) HasStorageClass() bool`

HasStorageClass returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


