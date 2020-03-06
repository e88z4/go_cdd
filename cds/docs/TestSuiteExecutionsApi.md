# \TestSuiteExecutionsApi

All URIs are relative to *https://cddirector.io/cdd/cds/0000/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTestSuiteByIdUsingGET**](TestSuiteExecutionsApi.md#GetTestSuiteByIdUsingGET) | **Get** /applications/{applicationId}/application-versions/{applicationVersionId}/test-suite-executions/{testSuiteExecutionsHexId} | Retrieve test suite executions by id
[**GetTestSuiteExecutionsUsingGET**](TestSuiteExecutionsApi.md#GetTestSuiteExecutionsUsingGET) | **Get** /applications/{applicationId}/application-versions/{applicationVersionId}/test-suite-executions | Retrieve all test suite executions


# **GetTestSuiteByIdUsingGET**
> TestSuitesExecutionDto GetTestSuiteByIdUsingGET(ctx, applicationId, applicationVersionId, testSuiteExecutionsHexId)
Retrieve test suite executions by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int64**| applicationId | 
  **applicationVersionId** | **int64**| applicationVersionId | 
  **testSuiteExecutionsHexId** | **string**| testSuiteExecutionId | 

### Return type

[**TestSuitesExecutionDto**](TestSuitesExecutionDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTestSuiteExecutionsUsingGET**
> TestSuitesExecutionPagedDto GetTestSuiteExecutionsUsingGET(ctx, applicationId, applicationVersionId, optional)
Retrieve all test suite executions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int64**| applicationId | 
  **applicationVersionId** | **int64**| applicationVersionId | 
 **optional** | ***GetTestSuiteExecutionsUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetTestSuiteExecutionsUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **testSource** | **optional.Int64**| testSourceId | 
 **sortField** | **optional.String**| sortField | 
 **sortDirection** | **optional.String**| sortDirection | 
 **filter** | **optional.String**| filter | 
 **endpointId** | **optional.Int64**| endpointId | 
 **tag** | **optional.String**| tags | 
 **environment** | **optional.Int64**| environmentId | 
 **release** | **optional.Int64**| releaseId | 
 **phase** | [**optional.Interface of []int64**](int64.md)| phaseIds | 
 **embed** | [**optional.Interface of []string**](string.md)| embed | 
 **testSuiteExecutionStatus** | [**optional.Interface of []string**](string.md)| testSuiteExecutionStatuses | 
 **pageNumber** | **optional.Int32**| pageNumber | 
 **pageSize** | **optional.Int32**| pageSize | 

### Return type

[**TestSuitesExecutionPagedDto**](TestSuitesExecutionPagedDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

