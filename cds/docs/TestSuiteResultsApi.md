# \TestSuiteResultsApi

All URIs are relative to *https://cddirector.io/cdd/cds/0000/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTestSuiteResultUsingGET**](TestSuiteResultsApi.md#GetTestSuiteResultUsingGET) | **Get** /applications/{applicationId}/application-versions/{applicationVersionId}/test-suite-results | Retrieve test suite result


# **GetTestSuiteResultUsingGET**
> PagedResultDtoTestSuiteResultDto GetTestSuiteResultUsingGET(ctx, applicationId, applicationVersionId, optional)
Retrieve test suite result

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int64**| applicationId | 
  **applicationVersionId** | **int64**| applicationVersionId | 
 **optional** | ***GetTestSuiteResultUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetTestSuiteResultUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **testSuiteExecutionId** | **optional.String**| testSuiteResultId | 
 **sortField** | **optional.String**| sortField | 
 **testCaseResultStatus** | [**optional.Interface of []string**](string.md)| testCaseStatuses | 
 **embed** | [**optional.Interface of []string**](string.md)| embed | 
 **sortDirection** | **optional.String**| sortDirection | 
 **filter** | **optional.String**| filter | 
 **endpointId** | **optional.Int64**| endpointId | 
 **pluginId** | **optional.Int64**| pluginId | 
 **tag** | **optional.String**| tags | 
 **environment** | **optional.Int64**| environmentId | 
 **pageNumber** | **optional.Int32**| pageNumber | 
 **pageSize** | **optional.Int32**| pageSize | 

### Return type

[**PagedResultDtoTestSuiteResultDto**](PagedResultDto«TestSuiteResultDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

