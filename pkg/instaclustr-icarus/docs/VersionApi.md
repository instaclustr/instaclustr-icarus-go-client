# {{classname}}

All URIs are relative to *http://localhost:4567*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VersionCassandraGet**](VersionApi.md#VersionCassandraGet) | **Get** /version/cassandra | returns version of Cassandra node
[**VersionGet**](VersionApi.md#VersionGet) | **Get** /version | returns version of Cassandra Sidecar itself
[**VersionSchemaGet**](VersionApi.md#VersionSchemaGet) | **Get** /version/schema | returns schema version this Cassandra node is on, same as calling StorageServiceMBean#getSchemaVersion
[**VersionSidecarGet**](VersionApi.md#VersionSidecarGet) | **Get** /version/sidecar | alias for /version endpoint, returns version of Cassandra Sidecar itself

# **VersionCassandraGet**
> CassandraVersion VersionCassandraGet(ctx, )
returns version of Cassandra node

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**CassandraVersion**](CassandraVersion.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VersionGet**
> SidecarVersion VersionGet(ctx, )
returns version of Cassandra Sidecar itself

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SidecarVersion**](SidecarVersion.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VersionSchemaGet**
> CassandraSchemaVersion VersionSchemaGet(ctx, )
returns schema version this Cassandra node is on, same as calling StorageServiceMBean#getSchemaVersion

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**CassandraSchemaVersion**](CassandraSchemaVersion.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VersionSidecarGet**
> SidecarVersion VersionSidecarGet(ctx, )
alias for /version endpoint, returns version of Cassandra Sidecar itself

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SidecarVersion**](SidecarVersion.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

