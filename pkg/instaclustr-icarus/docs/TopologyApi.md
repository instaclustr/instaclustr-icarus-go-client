# {{classname}}

All URIs are relative to *http://localhost:4567*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TopologyDcGet**](TopologyApi.md#TopologyDcGet) | **Get** /topology/{dc} | returns topology of a datacenter of a cluster
[**TopologyGet**](TopologyApi.md#TopologyGet) | **Get** /topology | returns topology of a cluster as seen from this node

# **TopologyDcGet**
> ClusterTopology TopologyDcGet(ctx, dc)
returns topology of a datacenter of a cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dc** | **string**| name of datacenter to resolve a topology of | 

### Return type

[**ClusterTopology**](ClusterTopology.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TopologyGet**
> ClusterTopology TopologyGet(ctx, )
returns topology of a cluster as seen from this node

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ClusterTopology**](ClusterTopology.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

