# \ReleaseExecutionReportApi

All URIs are relative to *https://cddirector.io/cdd/reporting/0000/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetExecutedPhasesUsingGET**](ReleaseExecutionReportApi.md#GetExecutedPhasesUsingGET) | **Get** /reports/executed-releases/{releaseId} | Get executed release information


# **GetExecutedPhasesUsingGET**
> ReleaseExecutionReportDto GetExecutedPhasesUsingGET(ctx, releaseId)
Get executed release information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseId** | **int64**| releaseId | 

### Return type

[**ReleaseExecutionReportDto**](ReleaseExecutionReportDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

