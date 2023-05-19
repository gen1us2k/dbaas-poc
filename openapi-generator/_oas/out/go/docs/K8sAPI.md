# \K8sAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDatabaseCluster**](K8sAPI.md#CreateDatabaseCluster) | **Post** /kubernetes/{kubernetesName}/database-cluster | Create a database cluster on given kubernetes cluster
[**ListDatabases**](K8sAPI.md#ListDatabases) | **Get** /kubernetes/{kubernetesName}/database-cluster | List of created database clusters on provided kubernetes cluster
[**ListKubernetesClusters**](K8sAPI.md#ListKubernetesClusters) | **Get** /kubernetes | List of registered kubernetes clusters
[**RegisterKubernetes**](K8sAPI.md#RegisterKubernetes) | **Post** /kubernetes | Register Kubernetes cluster in the control plane



## CreateDatabaseCluster

> DatabaseCluster CreateDatabaseCluster(ctx, kubernetesName).DatabaseCluster(databaseCluster).Execute()

Create a database cluster on given kubernetes cluster



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    kubernetesName := "kubernetesName_example" // string | Name of kubernetes Cluster
    databaseCluster := *openapiclient.NewDatabaseCluster() // DatabaseCluster | Register a new kubernetes cluster in the control plane

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.K8sAPI.CreateDatabaseCluster(context.Background(), kubernetesName).DatabaseCluster(databaseCluster).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `K8sAPI.CreateDatabaseCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDatabaseCluster`: DatabaseCluster
    fmt.Fprintf(os.Stdout, "Response from `K8sAPI.CreateDatabaseCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kubernetesName** | **string** | Name of kubernetes Cluster | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDatabaseClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **databaseCluster** | [**DatabaseCluster**](DatabaseCluster.md) | Register a new kubernetes cluster in the control plane | 

### Return type

[**DatabaseCluster**](DatabaseCluster.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDatabases

> []DatabaseCluster ListDatabases(ctx, kubernetesName).Execute()

List of created database clusters on provided kubernetes cluster



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    kubernetesName := "kubernetesName_example" // string | Name of kubernetes Cluster

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.K8sAPI.ListDatabases(context.Background(), kubernetesName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `K8sAPI.ListDatabases``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDatabases`: []DatabaseCluster
    fmt.Fprintf(os.Stdout, "Response from `K8sAPI.ListDatabases`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kubernetesName** | **string** | Name of kubernetes Cluster | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDatabasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]DatabaseCluster**](DatabaseCluster.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKubernetesClusters

> KubernetesCluster ListKubernetesClusters(ctx).Execute()

List of registered kubernetes clusters



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.K8sAPI.ListKubernetesClusters(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `K8sAPI.ListKubernetesClusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListKubernetesClusters`: KubernetesCluster
    fmt.Fprintf(os.Stdout, "Response from `K8sAPI.ListKubernetesClusters`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListKubernetesClustersRequest struct via the builder pattern


### Return type

[**KubernetesCluster**](KubernetesCluster.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterKubernetes

> KubernetesCluster RegisterKubernetes(ctx).KubernetesCluster(kubernetesCluster).Execute()

Register Kubernetes cluster in the control plane



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    kubernetesCluster := *openapiclient.NewKubernetesCluster() // KubernetesCluster | Register a new kubernetes cluster in the control plane

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.K8sAPI.RegisterKubernetes(context.Background()).KubernetesCluster(kubernetesCluster).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `K8sAPI.RegisterKubernetes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegisterKubernetes`: KubernetesCluster
    fmt.Fprintf(os.Stdout, "Response from `K8sAPI.RegisterKubernetes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterKubernetesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kubernetesCluster** | [**KubernetesCluster**](KubernetesCluster.md) | Register a new kubernetes cluster in the control plane | 

### Return type

[**KubernetesCluster**](KubernetesCluster.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

