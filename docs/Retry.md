# Retry

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Interval** | **int32** | Time gap between retries, linear strategy will have always this gap constant, exponential strategy will make the gap bigger exponentially (power of 2) on each attempt  | [optional] [default to null]
**Strategy** | **string** | Strategy how retry should be driven, might be either &#x27;LINEAR&#x27; or &#x27;EXPONENTIAL&#x27;  | [optional] [default to null]
**MaxAttempts** | **int32** | Number of repetitions of an upload / download operation in case it fails before giving up completely.  | [optional] [default to null]
**Enabled** | **bool** | Defaults to false if not specified. If false, retry mechanism on upload / download operations in case they fail will not be used.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

