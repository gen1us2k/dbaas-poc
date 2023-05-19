# DatabaseClusterSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Backup** | Pointer to [**DatabaseClusterSpecBackup**](DatabaseClusterSpecBackup.md) |  | [optional] 
**ClusterSize** | **int32** | ClusterSize is amount of nodes that required for the cluster. A database starts in cluster mode if clusterSize &gt;&#x3D; 3. | 
**DatabaseConfig** | **string** | DatabaseConfig contains a config settings for the specified database. | 
**DatabaseImage** | **string** | DatabaseVersion sets from version service and uses the recommended version by default. | 
**DatabaseType** | **string** | Database type stands for supported databases by the PMM API Now it&#39;s pxc or psmdb types but we can extend it. | 
**DbInstance** | [**DatabaseClusterSpecDbInstance**](DatabaseClusterSpecDbInstance.md) |  | 
**LoadBalancer** | Pointer to [**DatabaseClusterSpecLoadBalancer**](DatabaseClusterSpecLoadBalancer.md) |  | [optional] 
**Monitoring** | Pointer to [**DatabaseClusterSpecMonitoring**](DatabaseClusterSpecMonitoring.md) |  | [optional] 
**Pause** | Pointer to **bool** | Pause represents is a cluster paused or not. | [optional] 
**SecretsName** | Pointer to **string** | SecretsName contains name of a secrets file for a database cluster. | [optional] 

## Methods

### NewDatabaseClusterSpec

`func NewDatabaseClusterSpec(clusterSize int32, databaseConfig string, databaseImage string, databaseType string, dbInstance DatabaseClusterSpecDbInstance, ) *DatabaseClusterSpec`

NewDatabaseClusterSpec instantiates a new DatabaseClusterSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseClusterSpecWithDefaults

`func NewDatabaseClusterSpecWithDefaults() *DatabaseClusterSpec`

NewDatabaseClusterSpecWithDefaults instantiates a new DatabaseClusterSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackup

`func (o *DatabaseClusterSpec) GetBackup() DatabaseClusterSpecBackup`

GetBackup returns the Backup field if non-nil, zero value otherwise.

### GetBackupOk

`func (o *DatabaseClusterSpec) GetBackupOk() (*DatabaseClusterSpecBackup, bool)`

GetBackupOk returns a tuple with the Backup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackup

`func (o *DatabaseClusterSpec) SetBackup(v DatabaseClusterSpecBackup)`

SetBackup sets Backup field to given value.

### HasBackup

`func (o *DatabaseClusterSpec) HasBackup() bool`

HasBackup returns a boolean if a field has been set.

### GetClusterSize

`func (o *DatabaseClusterSpec) GetClusterSize() int32`

GetClusterSize returns the ClusterSize field if non-nil, zero value otherwise.

### GetClusterSizeOk

`func (o *DatabaseClusterSpec) GetClusterSizeOk() (*int32, bool)`

GetClusterSizeOk returns a tuple with the ClusterSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterSize

`func (o *DatabaseClusterSpec) SetClusterSize(v int32)`

SetClusterSize sets ClusterSize field to given value.


### GetDatabaseConfig

`func (o *DatabaseClusterSpec) GetDatabaseConfig() string`

GetDatabaseConfig returns the DatabaseConfig field if non-nil, zero value otherwise.

### GetDatabaseConfigOk

`func (o *DatabaseClusterSpec) GetDatabaseConfigOk() (*string, bool)`

GetDatabaseConfigOk returns a tuple with the DatabaseConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseConfig

`func (o *DatabaseClusterSpec) SetDatabaseConfig(v string)`

SetDatabaseConfig sets DatabaseConfig field to given value.


### GetDatabaseImage

`func (o *DatabaseClusterSpec) GetDatabaseImage() string`

GetDatabaseImage returns the DatabaseImage field if non-nil, zero value otherwise.

### GetDatabaseImageOk

`func (o *DatabaseClusterSpec) GetDatabaseImageOk() (*string, bool)`

GetDatabaseImageOk returns a tuple with the DatabaseImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseImage

`func (o *DatabaseClusterSpec) SetDatabaseImage(v string)`

SetDatabaseImage sets DatabaseImage field to given value.


### GetDatabaseType

`func (o *DatabaseClusterSpec) GetDatabaseType() string`

GetDatabaseType returns the DatabaseType field if non-nil, zero value otherwise.

### GetDatabaseTypeOk

`func (o *DatabaseClusterSpec) GetDatabaseTypeOk() (*string, bool)`

GetDatabaseTypeOk returns a tuple with the DatabaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseType

`func (o *DatabaseClusterSpec) SetDatabaseType(v string)`

SetDatabaseType sets DatabaseType field to given value.


### GetDbInstance

`func (o *DatabaseClusterSpec) GetDbInstance() DatabaseClusterSpecDbInstance`

GetDbInstance returns the DbInstance field if non-nil, zero value otherwise.

### GetDbInstanceOk

`func (o *DatabaseClusterSpec) GetDbInstanceOk() (*DatabaseClusterSpecDbInstance, bool)`

GetDbInstanceOk returns a tuple with the DbInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbInstance

`func (o *DatabaseClusterSpec) SetDbInstance(v DatabaseClusterSpecDbInstance)`

SetDbInstance sets DbInstance field to given value.


### GetLoadBalancer

`func (o *DatabaseClusterSpec) GetLoadBalancer() DatabaseClusterSpecLoadBalancer`

GetLoadBalancer returns the LoadBalancer field if non-nil, zero value otherwise.

### GetLoadBalancerOk

`func (o *DatabaseClusterSpec) GetLoadBalancerOk() (*DatabaseClusterSpecLoadBalancer, bool)`

GetLoadBalancerOk returns a tuple with the LoadBalancer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancer

`func (o *DatabaseClusterSpec) SetLoadBalancer(v DatabaseClusterSpecLoadBalancer)`

SetLoadBalancer sets LoadBalancer field to given value.

### HasLoadBalancer

`func (o *DatabaseClusterSpec) HasLoadBalancer() bool`

HasLoadBalancer returns a boolean if a field has been set.

### GetMonitoring

`func (o *DatabaseClusterSpec) GetMonitoring() DatabaseClusterSpecMonitoring`

GetMonitoring returns the Monitoring field if non-nil, zero value otherwise.

### GetMonitoringOk

`func (o *DatabaseClusterSpec) GetMonitoringOk() (*DatabaseClusterSpecMonitoring, bool)`

GetMonitoringOk returns a tuple with the Monitoring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoring

`func (o *DatabaseClusterSpec) SetMonitoring(v DatabaseClusterSpecMonitoring)`

SetMonitoring sets Monitoring field to given value.

### HasMonitoring

`func (o *DatabaseClusterSpec) HasMonitoring() bool`

HasMonitoring returns a boolean if a field has been set.

### GetPause

`func (o *DatabaseClusterSpec) GetPause() bool`

GetPause returns the Pause field if non-nil, zero value otherwise.

### GetPauseOk

`func (o *DatabaseClusterSpec) GetPauseOk() (*bool, bool)`

GetPauseOk returns a tuple with the Pause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPause

`func (o *DatabaseClusterSpec) SetPause(v bool)`

SetPause sets Pause field to given value.

### HasPause

`func (o *DatabaseClusterSpec) HasPause() bool`

HasPause returns a boolean if a field has been set.

### GetSecretsName

`func (o *DatabaseClusterSpec) GetSecretsName() string`

GetSecretsName returns the SecretsName field if non-nil, zero value otherwise.

### GetSecretsNameOk

`func (o *DatabaseClusterSpec) GetSecretsNameOk() (*string, bool)`

GetSecretsNameOk returns a tuple with the SecretsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretsName

`func (o *DatabaseClusterSpec) SetSecretsName(v string)`

SetSecretsName sets SecretsName field to given value.

### HasSecretsName

`func (o *DatabaseClusterSpec) HasSecretsName() bool`

HasSecretsName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


