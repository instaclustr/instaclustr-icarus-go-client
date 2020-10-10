# RefreshOperationResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type_** | **string** |  | [default to null]
**Keyspace** | **string** | keyspace to refresh  | [default to null]
**Table** | **string** | table to refresh  | [default to null]
**Id** | **string** | unique identifier of an operation, a random id is assigned to each operation after a request is submitted, from caller&#x27;s perspective, an id is sent back as a response to his request so he can further query state of that operation, referencing id, by operations/{id} endpoint  | [default to null]
**State** | **string** | state of an operation, operation might be in various states, PENDING - this operation is pending for being submitted. RUNNING - this operation is actively doing its job, COMPLETED - this operation has finished successfully, CANCELLED - this operation was interrupted while being run, FAILED - this operation has finished errorneously  | [default to null]
**Progress** | **string** | float from 0.0 to 1.0, 1.0 telling that operation is completed, either successfully or with errors.  | [default to null]
**CreationTime** | [**time.Time**](time.Time.md) | timestamp telling when this operation was created on Sidecar&#x27;s side  | [default to null]
**StartTime** | [**time.Time**](time.Time.md) | timestamp telling when this operation was started by Sidecar, if an operation is created, it does not necessarily mean that it will be started right away, in most cases it is the case but if e.g. ExecutorService is full on its working thread, an execution of an operation is postponed and start time is updated only after that  | [optional] [default to null]
**CompletionTime** | [**time.Time**](time.Time.md) | timestamp telling when an operation has finished, irrelevant of its result, an operation can be failed and it would still have this field populated.  | [optional] [default to null]
**FailureCause** | [***interface{}**](interface{}.md) | This field contains serialized java.lang.Throwable in case this operation has failed  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

