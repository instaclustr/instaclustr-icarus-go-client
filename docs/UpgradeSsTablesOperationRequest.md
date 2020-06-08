# UpgradeSsTablesOperationRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type_** | **string** |  | [default to null]
**Keyspace** | **string** | keyspace to upgrade SSTables of  | [default to null]
**Tables** | **[]string** | an array of tables to upgrade SSTables of, empty or not provided array will default to upgrading of SSTables of all tables in respective keyspace  | [optional] [default to null]
**Jobs** | **int32** | the number of threads to use - 0 means use all available, it never uses more than concurrent_compactor threads  | [optional] [default to null]
**IncludeAllSStables** | **bool** | include all sstables, even those already on the current version, defaults to false | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

