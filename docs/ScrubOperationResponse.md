# ScrubOperationResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type_** | **string** |  | [default to null]
**Keyspace** | **string** | keyspace to scrub  | [default to null]
**Tables** | **[]string** | tables to scrub, empty or not provided will scrub all tables in respective keyspace  | [optional] [default to null]
**Jobs** | **int32** | number of sstables to scrub simultanously, set to 0 to use all available compaction threads  | [optional] [default to null]
**DisableSnapshot** | **bool** | scrubbed CFs will be snapshotted first, defaults to false  | [optional] [default to null]
**SkipCorrupted** | **bool** | skip corrupted partitions even when scrubbing counter tables, defaults to false  | [optional] [default to null]
**NoValidate** | **bool** | do not validate columns using column validator, defaults to false  | [optional] [default to null]
**ReinsertOverflowedTTL** | **bool** | Rewrites rows with overflowed expiration date affected by CASSANDRA-14092 with the maximum supported expiration date of 2038-01-19T03:14:06+00:00. The rows are rewritten with the original timestamp incremented by one millisecond to override/supersede any potential tombstone that may have been generated during compaction of the affected rows.  | [optional] [default to null]
**Id** | **string** | unique identifier of an operation, a random id is assigned to each operation after a request is submitted, from caller&#x27;s perspective, an id is sent back as a response to his request so he can further query state of that operation, referencing id, by operations/{id} endpoint  | [default to null]
**State** | **string** | state of an operation, operation might be in various states, PENDING - this operation is pending for being submitted. RUNNING - this operation is actively doing its job, COMPLETED - this operation has finished successfully, CANCELLED - this operation was interrupted while being run, FAILED - this operation has finished errorneously  | [default to null]
**Progress** | **string** | float from 0.0 to 1.0, 1.0 telling that operation is completed, either successfully or with errors.  | [default to null]
**CreationTime** | [**time.Time**](time.Time.md) | timestamp telling when this operation was created on Sidecar&#x27;s side  | [default to null]
**StartTime** | [**time.Time**](time.Time.md) | timestamp telling when this operation was started by Sidecar, if an operation is created, it does not necessarily mean that it will be started right away, in most cases it is the case but if e.g. ExecutorService is full on its working thread, an execution of an operation is postponed and start time is updated only after that  | [optional] [default to null]
**CompletionTime** | [**time.Time**](time.Time.md) | timestamp telling when an operation has finished, irrelevant of its result, an operation can be failed and it would still have this field populated.  | [optional] [default to null]
**FailureCause** | [***interface{}**](interface{}.md) | This field contains serialized java.lang.Throwable in case this operation has failed  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

