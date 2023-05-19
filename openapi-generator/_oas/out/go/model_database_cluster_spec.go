/*
Percona Everest schema

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the DatabaseClusterSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatabaseClusterSpec{}

// DatabaseClusterSpec DatabaseSpec defines the desired state of Database.
type DatabaseClusterSpec struct {
	Backup *DatabaseClusterSpecBackup `json:"backup,omitempty"`
	// ClusterSize is amount of nodes that required for the cluster. A database starts in cluster mode if clusterSize >= 3.
	ClusterSize int32 `json:"clusterSize"`
	// DatabaseConfig contains a config settings for the specified database.
	DatabaseConfig string `json:"databaseConfig"`
	// DatabaseVersion sets from version service and uses the recommended version by default.
	DatabaseImage string `json:"databaseImage"`
	// Database type stands for supported databases by the PMM API Now it's pxc or psmdb types but we can extend it.
	DatabaseType string `json:"databaseType"`
	DbInstance DatabaseClusterSpecDbInstance `json:"dbInstance"`
	LoadBalancer *DatabaseClusterSpecLoadBalancer `json:"loadBalancer,omitempty"`
	Monitoring *DatabaseClusterSpecMonitoring `json:"monitoring,omitempty"`
	// Pause represents is a cluster paused or not.
	Pause *bool `json:"pause,omitempty"`
	// SecretsName contains name of a secrets file for a database cluster.
	SecretsName *string `json:"secretsName,omitempty"`
}

// NewDatabaseClusterSpec instantiates a new DatabaseClusterSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatabaseClusterSpec(clusterSize int32, databaseConfig string, databaseImage string, databaseType string, dbInstance DatabaseClusterSpecDbInstance) *DatabaseClusterSpec {
	this := DatabaseClusterSpec{}
	this.ClusterSize = clusterSize
	this.DatabaseConfig = databaseConfig
	this.DatabaseImage = databaseImage
	this.DatabaseType = databaseType
	this.DbInstance = dbInstance
	return &this
}

// NewDatabaseClusterSpecWithDefaults instantiates a new DatabaseClusterSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatabaseClusterSpecWithDefaults() *DatabaseClusterSpec {
	this := DatabaseClusterSpec{}
	return &this
}

// GetBackup returns the Backup field value if set, zero value otherwise.
func (o *DatabaseClusterSpec) GetBackup() DatabaseClusterSpecBackup {
	if o == nil || IsNil(o.Backup) {
		var ret DatabaseClusterSpecBackup
		return ret
	}
	return *o.Backup
}

// GetBackupOk returns a tuple with the Backup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseClusterSpec) GetBackupOk() (*DatabaseClusterSpecBackup, bool) {
	if o == nil || IsNil(o.Backup) {
		return nil, false
	}
	return o.Backup, true
}

// HasBackup returns a boolean if a field has been set.
func (o *DatabaseClusterSpec) HasBackup() bool {
	if o != nil && !IsNil(o.Backup) {
		return true
	}

	return false
}

// SetBackup gets a reference to the given DatabaseClusterSpecBackup and assigns it to the Backup field.
func (o *DatabaseClusterSpec) SetBackup(v DatabaseClusterSpecBackup) {
	o.Backup = &v
}

// GetClusterSize returns the ClusterSize field value
func (o *DatabaseClusterSpec) GetClusterSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ClusterSize
}

// GetClusterSizeOk returns a tuple with the ClusterSize field value
// and a boolean to check if the value has been set.
func (o *DatabaseClusterSpec) GetClusterSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterSize, true
}

// SetClusterSize sets field value
func (o *DatabaseClusterSpec) SetClusterSize(v int32) {
	o.ClusterSize = v
}

// GetDatabaseConfig returns the DatabaseConfig field value
func (o *DatabaseClusterSpec) GetDatabaseConfig() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DatabaseConfig
}

// GetDatabaseConfigOk returns a tuple with the DatabaseConfig field value
// and a boolean to check if the value has been set.
func (o *DatabaseClusterSpec) GetDatabaseConfigOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatabaseConfig, true
}

// SetDatabaseConfig sets field value
func (o *DatabaseClusterSpec) SetDatabaseConfig(v string) {
	o.DatabaseConfig = v
}

// GetDatabaseImage returns the DatabaseImage field value
func (o *DatabaseClusterSpec) GetDatabaseImage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DatabaseImage
}

// GetDatabaseImageOk returns a tuple with the DatabaseImage field value
// and a boolean to check if the value has been set.
func (o *DatabaseClusterSpec) GetDatabaseImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatabaseImage, true
}

// SetDatabaseImage sets field value
func (o *DatabaseClusterSpec) SetDatabaseImage(v string) {
	o.DatabaseImage = v
}

// GetDatabaseType returns the DatabaseType field value
func (o *DatabaseClusterSpec) GetDatabaseType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DatabaseType
}

// GetDatabaseTypeOk returns a tuple with the DatabaseType field value
// and a boolean to check if the value has been set.
func (o *DatabaseClusterSpec) GetDatabaseTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatabaseType, true
}

// SetDatabaseType sets field value
func (o *DatabaseClusterSpec) SetDatabaseType(v string) {
	o.DatabaseType = v
}

// GetDbInstance returns the DbInstance field value
func (o *DatabaseClusterSpec) GetDbInstance() DatabaseClusterSpecDbInstance {
	if o == nil {
		var ret DatabaseClusterSpecDbInstance
		return ret
	}

	return o.DbInstance
}

// GetDbInstanceOk returns a tuple with the DbInstance field value
// and a boolean to check if the value has been set.
func (o *DatabaseClusterSpec) GetDbInstanceOk() (*DatabaseClusterSpecDbInstance, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DbInstance, true
}

// SetDbInstance sets field value
func (o *DatabaseClusterSpec) SetDbInstance(v DatabaseClusterSpecDbInstance) {
	o.DbInstance = v
}

// GetLoadBalancer returns the LoadBalancer field value if set, zero value otherwise.
func (o *DatabaseClusterSpec) GetLoadBalancer() DatabaseClusterSpecLoadBalancer {
	if o == nil || IsNil(o.LoadBalancer) {
		var ret DatabaseClusterSpecLoadBalancer
		return ret
	}
	return *o.LoadBalancer
}

// GetLoadBalancerOk returns a tuple with the LoadBalancer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseClusterSpec) GetLoadBalancerOk() (*DatabaseClusterSpecLoadBalancer, bool) {
	if o == nil || IsNil(o.LoadBalancer) {
		return nil, false
	}
	return o.LoadBalancer, true
}

// HasLoadBalancer returns a boolean if a field has been set.
func (o *DatabaseClusterSpec) HasLoadBalancer() bool {
	if o != nil && !IsNil(o.LoadBalancer) {
		return true
	}

	return false
}

// SetLoadBalancer gets a reference to the given DatabaseClusterSpecLoadBalancer and assigns it to the LoadBalancer field.
func (o *DatabaseClusterSpec) SetLoadBalancer(v DatabaseClusterSpecLoadBalancer) {
	o.LoadBalancer = &v
}

// GetMonitoring returns the Monitoring field value if set, zero value otherwise.
func (o *DatabaseClusterSpec) GetMonitoring() DatabaseClusterSpecMonitoring {
	if o == nil || IsNil(o.Monitoring) {
		var ret DatabaseClusterSpecMonitoring
		return ret
	}
	return *o.Monitoring
}

// GetMonitoringOk returns a tuple with the Monitoring field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseClusterSpec) GetMonitoringOk() (*DatabaseClusterSpecMonitoring, bool) {
	if o == nil || IsNil(o.Monitoring) {
		return nil, false
	}
	return o.Monitoring, true
}

// HasMonitoring returns a boolean if a field has been set.
func (o *DatabaseClusterSpec) HasMonitoring() bool {
	if o != nil && !IsNil(o.Monitoring) {
		return true
	}

	return false
}

// SetMonitoring gets a reference to the given DatabaseClusterSpecMonitoring and assigns it to the Monitoring field.
func (o *DatabaseClusterSpec) SetMonitoring(v DatabaseClusterSpecMonitoring) {
	o.Monitoring = &v
}

// GetPause returns the Pause field value if set, zero value otherwise.
func (o *DatabaseClusterSpec) GetPause() bool {
	if o == nil || IsNil(o.Pause) {
		var ret bool
		return ret
	}
	return *o.Pause
}

// GetPauseOk returns a tuple with the Pause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseClusterSpec) GetPauseOk() (*bool, bool) {
	if o == nil || IsNil(o.Pause) {
		return nil, false
	}
	return o.Pause, true
}

// HasPause returns a boolean if a field has been set.
func (o *DatabaseClusterSpec) HasPause() bool {
	if o != nil && !IsNil(o.Pause) {
		return true
	}

	return false
}

// SetPause gets a reference to the given bool and assigns it to the Pause field.
func (o *DatabaseClusterSpec) SetPause(v bool) {
	o.Pause = &v
}

// GetSecretsName returns the SecretsName field value if set, zero value otherwise.
func (o *DatabaseClusterSpec) GetSecretsName() string {
	if o == nil || IsNil(o.SecretsName) {
		var ret string
		return ret
	}
	return *o.SecretsName
}

// GetSecretsNameOk returns a tuple with the SecretsName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseClusterSpec) GetSecretsNameOk() (*string, bool) {
	if o == nil || IsNil(o.SecretsName) {
		return nil, false
	}
	return o.SecretsName, true
}

// HasSecretsName returns a boolean if a field has been set.
func (o *DatabaseClusterSpec) HasSecretsName() bool {
	if o != nil && !IsNil(o.SecretsName) {
		return true
	}

	return false
}

// SetSecretsName gets a reference to the given string and assigns it to the SecretsName field.
func (o *DatabaseClusterSpec) SetSecretsName(v string) {
	o.SecretsName = &v
}

func (o DatabaseClusterSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatabaseClusterSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Backup) {
		toSerialize["backup"] = o.Backup
	}
	toSerialize["clusterSize"] = o.ClusterSize
	toSerialize["databaseConfig"] = o.DatabaseConfig
	toSerialize["databaseImage"] = o.DatabaseImage
	toSerialize["databaseType"] = o.DatabaseType
	toSerialize["dbInstance"] = o.DbInstance
	if !IsNil(o.LoadBalancer) {
		toSerialize["loadBalancer"] = o.LoadBalancer
	}
	if !IsNil(o.Monitoring) {
		toSerialize["monitoring"] = o.Monitoring
	}
	if !IsNil(o.Pause) {
		toSerialize["pause"] = o.Pause
	}
	if !IsNil(o.SecretsName) {
		toSerialize["secretsName"] = o.SecretsName
	}
	return toSerialize, nil
}

type NullableDatabaseClusterSpec struct {
	value *DatabaseClusterSpec
	isSet bool
}

func (v NullableDatabaseClusterSpec) Get() *DatabaseClusterSpec {
	return v.value
}

func (v *NullableDatabaseClusterSpec) Set(val *DatabaseClusterSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableDatabaseClusterSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableDatabaseClusterSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatabaseClusterSpec(val *DatabaseClusterSpec) *NullableDatabaseClusterSpec {
	return &NullableDatabaseClusterSpec{value: val, isSet: true}
}

func (v NullableDatabaseClusterSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatabaseClusterSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


