# DatabaseClusterSpecBackup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Annotations** | Pointer to **map[string]string** |  | [optional] 
**ContainerSecurityContext** | Pointer to [**DatabaseClusterSpecBackupContainerSecurityContext**](DatabaseClusterSpecBackupContainerSecurityContext.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Image** | Pointer to **string** |  | [optional] 
**ImagePullPolicy** | Pointer to **string** | PullPolicy describes a policy for if/when to pull a container image | [optional] 
**ImagePullSecrets** | Pointer to [**[]DatabaseClusterSpecBackupImagePullSecretsInner**](DatabaseClusterSpecBackupImagePullSecretsInner.md) |  | [optional] 
**InitImage** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **map[string]string** |  | [optional] 
**Resources** | Pointer to [**DatabaseClusterSpecBackupResources**](DatabaseClusterSpecBackupResources.md) |  | [optional] 
**Schedule** | Pointer to [**[]DatabaseClusterSpecBackupScheduleInner**](DatabaseClusterSpecBackupScheduleInner.md) |  | [optional] 
**ServiceAccountName** | Pointer to **string** |  | [optional] 
**Storages** | Pointer to [**map[string]DatabaseClusterSpecBackupStoragesValue**](DatabaseClusterSpecBackupStoragesValue.md) |  | [optional] 

## Methods

### NewDatabaseClusterSpecBackup

`func NewDatabaseClusterSpecBackup() *DatabaseClusterSpecBackup`

NewDatabaseClusterSpecBackup instantiates a new DatabaseClusterSpecBackup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseClusterSpecBackupWithDefaults

`func NewDatabaseClusterSpecBackupWithDefaults() *DatabaseClusterSpecBackup`

NewDatabaseClusterSpecBackupWithDefaults instantiates a new DatabaseClusterSpecBackup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnnotations

`func (o *DatabaseClusterSpecBackup) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *DatabaseClusterSpecBackup) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *DatabaseClusterSpecBackup) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *DatabaseClusterSpecBackup) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetContainerSecurityContext

`func (o *DatabaseClusterSpecBackup) GetContainerSecurityContext() DatabaseClusterSpecBackupContainerSecurityContext`

GetContainerSecurityContext returns the ContainerSecurityContext field if non-nil, zero value otherwise.

### GetContainerSecurityContextOk

`func (o *DatabaseClusterSpecBackup) GetContainerSecurityContextOk() (*DatabaseClusterSpecBackupContainerSecurityContext, bool)`

GetContainerSecurityContextOk returns a tuple with the ContainerSecurityContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerSecurityContext

`func (o *DatabaseClusterSpecBackup) SetContainerSecurityContext(v DatabaseClusterSpecBackupContainerSecurityContext)`

SetContainerSecurityContext sets ContainerSecurityContext field to given value.

### HasContainerSecurityContext

`func (o *DatabaseClusterSpecBackup) HasContainerSecurityContext() bool`

HasContainerSecurityContext returns a boolean if a field has been set.

### GetEnabled

`func (o *DatabaseClusterSpecBackup) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DatabaseClusterSpecBackup) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DatabaseClusterSpecBackup) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *DatabaseClusterSpecBackup) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetImage

`func (o *DatabaseClusterSpecBackup) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *DatabaseClusterSpecBackup) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *DatabaseClusterSpecBackup) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *DatabaseClusterSpecBackup) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetImagePullPolicy

`func (o *DatabaseClusterSpecBackup) GetImagePullPolicy() string`

GetImagePullPolicy returns the ImagePullPolicy field if non-nil, zero value otherwise.

### GetImagePullPolicyOk

`func (o *DatabaseClusterSpecBackup) GetImagePullPolicyOk() (*string, bool)`

GetImagePullPolicyOk returns a tuple with the ImagePullPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImagePullPolicy

`func (o *DatabaseClusterSpecBackup) SetImagePullPolicy(v string)`

SetImagePullPolicy sets ImagePullPolicy field to given value.

### HasImagePullPolicy

`func (o *DatabaseClusterSpecBackup) HasImagePullPolicy() bool`

HasImagePullPolicy returns a boolean if a field has been set.

### GetImagePullSecrets

`func (o *DatabaseClusterSpecBackup) GetImagePullSecrets() []DatabaseClusterSpecBackupImagePullSecretsInner`

GetImagePullSecrets returns the ImagePullSecrets field if non-nil, zero value otherwise.

### GetImagePullSecretsOk

`func (o *DatabaseClusterSpecBackup) GetImagePullSecretsOk() (*[]DatabaseClusterSpecBackupImagePullSecretsInner, bool)`

GetImagePullSecretsOk returns a tuple with the ImagePullSecrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImagePullSecrets

`func (o *DatabaseClusterSpecBackup) SetImagePullSecrets(v []DatabaseClusterSpecBackupImagePullSecretsInner)`

SetImagePullSecrets sets ImagePullSecrets field to given value.

### HasImagePullSecrets

`func (o *DatabaseClusterSpecBackup) HasImagePullSecrets() bool`

HasImagePullSecrets returns a boolean if a field has been set.

### GetInitImage

`func (o *DatabaseClusterSpecBackup) GetInitImage() string`

GetInitImage returns the InitImage field if non-nil, zero value otherwise.

### GetInitImageOk

`func (o *DatabaseClusterSpecBackup) GetInitImageOk() (*string, bool)`

GetInitImageOk returns a tuple with the InitImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitImage

`func (o *DatabaseClusterSpecBackup) SetInitImage(v string)`

SetInitImage sets InitImage field to given value.

### HasInitImage

`func (o *DatabaseClusterSpecBackup) HasInitImage() bool`

HasInitImage returns a boolean if a field has been set.

### GetLabels

`func (o *DatabaseClusterSpecBackup) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *DatabaseClusterSpecBackup) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *DatabaseClusterSpecBackup) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *DatabaseClusterSpecBackup) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetResources

`func (o *DatabaseClusterSpecBackup) GetResources() DatabaseClusterSpecBackupResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *DatabaseClusterSpecBackup) GetResourcesOk() (*DatabaseClusterSpecBackupResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *DatabaseClusterSpecBackup) SetResources(v DatabaseClusterSpecBackupResources)`

SetResources sets Resources field to given value.

### HasResources

`func (o *DatabaseClusterSpecBackup) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetSchedule

`func (o *DatabaseClusterSpecBackup) GetSchedule() []DatabaseClusterSpecBackupScheduleInner`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *DatabaseClusterSpecBackup) GetScheduleOk() (*[]DatabaseClusterSpecBackupScheduleInner, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *DatabaseClusterSpecBackup) SetSchedule(v []DatabaseClusterSpecBackupScheduleInner)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *DatabaseClusterSpecBackup) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetServiceAccountName

`func (o *DatabaseClusterSpecBackup) GetServiceAccountName() string`

GetServiceAccountName returns the ServiceAccountName field if non-nil, zero value otherwise.

### GetServiceAccountNameOk

`func (o *DatabaseClusterSpecBackup) GetServiceAccountNameOk() (*string, bool)`

GetServiceAccountNameOk returns a tuple with the ServiceAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountName

`func (o *DatabaseClusterSpecBackup) SetServiceAccountName(v string)`

SetServiceAccountName sets ServiceAccountName field to given value.

### HasServiceAccountName

`func (o *DatabaseClusterSpecBackup) HasServiceAccountName() bool`

HasServiceAccountName returns a boolean if a field has been set.

### GetStorages

`func (o *DatabaseClusterSpecBackup) GetStorages() map[string]DatabaseClusterSpecBackupStoragesValue`

GetStorages returns the Storages field if non-nil, zero value otherwise.

### GetStoragesOk

`func (o *DatabaseClusterSpecBackup) GetStoragesOk() (*map[string]DatabaseClusterSpecBackupStoragesValue, bool)`

GetStoragesOk returns a tuple with the Storages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorages

`func (o *DatabaseClusterSpecBackup) SetStorages(v map[string]DatabaseClusterSpecBackupStoragesValue)`

SetStorages sets Storages field to given value.

### HasStorages

`func (o *DatabaseClusterSpecBackup) HasStorages() bool`

HasStorages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


