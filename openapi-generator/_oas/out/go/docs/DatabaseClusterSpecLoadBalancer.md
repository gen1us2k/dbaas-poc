# DatabaseClusterSpecLoadBalancer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Annotations** | Pointer to **map[string]string** |  | [optional] 
**Configuration** | Pointer to **string** |  | [optional] 
**ExposeType** | Pointer to **string** | Service Type string describes ingress methods for a service | [optional] 
**Image** | Pointer to **string** |  | [optional] 
**LoadBalancerSourceRanges** | Pointer to **[]string** |  | [optional] 
**Resources** | Pointer to [**DatabaseClusterSpecBackupResources**](DatabaseClusterSpecBackupResources.md) |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 
**TrafficPolicy** | Pointer to **string** | ServiceExternalTrafficPolicyType describes how nodes distribute service traffic they receive on one of the Service&#39;s \&quot;externally-facing\&quot; addresses (NodePorts, ExternalIPs, and LoadBalancer IPs). | [optional] 
**Type** | Pointer to **string** | LoadBalancerType contains supported loadbalancers. It can be proxysql or haproxy for PXC clusters, mongos for PSMDB clusters or pgbouncer for Postgresql clusters. | [optional] 

## Methods

### NewDatabaseClusterSpecLoadBalancer

`func NewDatabaseClusterSpecLoadBalancer() *DatabaseClusterSpecLoadBalancer`

NewDatabaseClusterSpecLoadBalancer instantiates a new DatabaseClusterSpecLoadBalancer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseClusterSpecLoadBalancerWithDefaults

`func NewDatabaseClusterSpecLoadBalancerWithDefaults() *DatabaseClusterSpecLoadBalancer`

NewDatabaseClusterSpecLoadBalancerWithDefaults instantiates a new DatabaseClusterSpecLoadBalancer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnnotations

`func (o *DatabaseClusterSpecLoadBalancer) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *DatabaseClusterSpecLoadBalancer) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *DatabaseClusterSpecLoadBalancer) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *DatabaseClusterSpecLoadBalancer) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetConfiguration

`func (o *DatabaseClusterSpecLoadBalancer) GetConfiguration() string`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *DatabaseClusterSpecLoadBalancer) GetConfigurationOk() (*string, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *DatabaseClusterSpecLoadBalancer) SetConfiguration(v string)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *DatabaseClusterSpecLoadBalancer) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetExposeType

`func (o *DatabaseClusterSpecLoadBalancer) GetExposeType() string`

GetExposeType returns the ExposeType field if non-nil, zero value otherwise.

### GetExposeTypeOk

`func (o *DatabaseClusterSpecLoadBalancer) GetExposeTypeOk() (*string, bool)`

GetExposeTypeOk returns a tuple with the ExposeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposeType

`func (o *DatabaseClusterSpecLoadBalancer) SetExposeType(v string)`

SetExposeType sets ExposeType field to given value.

### HasExposeType

`func (o *DatabaseClusterSpecLoadBalancer) HasExposeType() bool`

HasExposeType returns a boolean if a field has been set.

### GetImage

`func (o *DatabaseClusterSpecLoadBalancer) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *DatabaseClusterSpecLoadBalancer) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *DatabaseClusterSpecLoadBalancer) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *DatabaseClusterSpecLoadBalancer) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetLoadBalancerSourceRanges

`func (o *DatabaseClusterSpecLoadBalancer) GetLoadBalancerSourceRanges() []string`

GetLoadBalancerSourceRanges returns the LoadBalancerSourceRanges field if non-nil, zero value otherwise.

### GetLoadBalancerSourceRangesOk

`func (o *DatabaseClusterSpecLoadBalancer) GetLoadBalancerSourceRangesOk() (*[]string, bool)`

GetLoadBalancerSourceRangesOk returns a tuple with the LoadBalancerSourceRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerSourceRanges

`func (o *DatabaseClusterSpecLoadBalancer) SetLoadBalancerSourceRanges(v []string)`

SetLoadBalancerSourceRanges sets LoadBalancerSourceRanges field to given value.

### HasLoadBalancerSourceRanges

`func (o *DatabaseClusterSpecLoadBalancer) HasLoadBalancerSourceRanges() bool`

HasLoadBalancerSourceRanges returns a boolean if a field has been set.

### GetResources

`func (o *DatabaseClusterSpecLoadBalancer) GetResources() DatabaseClusterSpecBackupResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *DatabaseClusterSpecLoadBalancer) GetResourcesOk() (*DatabaseClusterSpecBackupResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *DatabaseClusterSpecLoadBalancer) SetResources(v DatabaseClusterSpecBackupResources)`

SetResources sets Resources field to given value.

### HasResources

`func (o *DatabaseClusterSpecLoadBalancer) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetSize

`func (o *DatabaseClusterSpecLoadBalancer) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *DatabaseClusterSpecLoadBalancer) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *DatabaseClusterSpecLoadBalancer) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *DatabaseClusterSpecLoadBalancer) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetTrafficPolicy

`func (o *DatabaseClusterSpecLoadBalancer) GetTrafficPolicy() string`

GetTrafficPolicy returns the TrafficPolicy field if non-nil, zero value otherwise.

### GetTrafficPolicyOk

`func (o *DatabaseClusterSpecLoadBalancer) GetTrafficPolicyOk() (*string, bool)`

GetTrafficPolicyOk returns a tuple with the TrafficPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficPolicy

`func (o *DatabaseClusterSpecLoadBalancer) SetTrafficPolicy(v string)`

SetTrafficPolicy sets TrafficPolicy field to given value.

### HasTrafficPolicy

`func (o *DatabaseClusterSpecLoadBalancer) HasTrafficPolicy() bool`

HasTrafficPolicy returns a boolean if a field has been set.

### GetType

`func (o *DatabaseClusterSpecLoadBalancer) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DatabaseClusterSpecLoadBalancer) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DatabaseClusterSpecLoadBalancer) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DatabaseClusterSpecLoadBalancer) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


