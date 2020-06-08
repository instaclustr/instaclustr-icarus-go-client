# RebuildOperationRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type_** | **string** |  | [default to null]
**Keyspace** | **string** | specific keyspace to rebuild, if not specified, all keyspaces are rebuilt  | [optional] [default to null]
**SourceDC** | **string** | name of DC from which to select sources for streaming, by default, pick any DC  | [optional] [default to null]
**SpecificTokens** | [**[]TokenRange**](TokenRange.md) | rebuild specific token ranges  | [optional] [default to null]
**SpecificSources** | **[]string** | specify hosts that this node should stream from when specificTokens are used  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

