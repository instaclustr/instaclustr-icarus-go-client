# {{classname}}

All URIs are relative to *http://localhost:4567*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OperationsGet**](OperationsApi.md#OperationsGet) | **Get** /operations | All operations of Icarus
[**OperationsOperationIdGet**](OperationsApi.md#OperationsOperationIdGet) | **Get** /operations/{operationId} | gets operation by its ID
[**OperationsPost**](OperationsApi.md#OperationsPost) | **Post** /operations | Submits an operation to this Sidecar

# **OperationsGet**
> []Object OperationsGet(ctx, optional)
All operations of Icarus

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OperationsApiOperationsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OperationsApiOperationsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | [**optional.Interface of []string**](string.md)| type of operations to filter on | 
 **status** | [**optional.Interface of []string**](string.md)| status of operations to filter on | 

### Return type

**[]Object**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationsOperationIdGet**
> InlineResponse200 OperationsOperationIdGet(ctx, operationId)
gets operation by its ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **operationId** | [**string**](.md)| ID of operation to return | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationsPost**
> InlineResponse200 OperationsPost(ctx, optional)
Submits an operation to this Sidecar

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OperationsApiOperationsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OperationsApiOperationsPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Body**](Body.md)|  | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

