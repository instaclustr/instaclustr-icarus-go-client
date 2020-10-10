# CleanupOperationRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type_** | **string** |  | [default to null]
**Keyspace** | **string** | keyspace to cleanup  | [default to null]
**Tables** | **[]string** | tables to cleanup, when not specified, all tables in a keyspace will be cleaned up  | [optional] [default to null]
**Jobs** | **int32** | number of jobs to use, never uses more that concurrent_compactor threads  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

