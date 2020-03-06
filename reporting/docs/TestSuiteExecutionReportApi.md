# \TestSuiteExecutionReportApi

All URIs are relative to *https://cddirector.io/cdd/reporting/0000/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApplicationTestingInsightReportUsingGET**](TestSuiteExecutionReportApi.md#GetApplicationTestingInsightReportUsingGET) | **Get** /reports/applications/{applicationId}/application-versions/{applicationVersionId}/application-testing-insights/phases/{phaseId}/plugins/{pluginId} | Get application insight report per plugin and phase
[**GetTestSuiteExecutionReportUsingGET**](TestSuiteExecutionReportApi.md#GetTestSuiteExecutionReportUsingGET) | **Get** /reports/applications/{applicationId}/application-versions/{applicationVersionId}/application-testing-insights | Get executed phases for release


# **GetApplicationTestingInsightReportUsingGET**
> PhasePluginTestExecutionDto GetApplicationTestingInsightReportUsingGET(ctx, applicationId, applicationVersionId, phaseId, pluginId)
Get application insight report per plugin and phase

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int64**| applicationId | 
  **applicationVersionId** | **int64**| applicationVersionId | 
  **phaseId** | **int64**| phaseId | 
  **pluginId** | **int64**| pluginId | 

### Return type

[**PhasePluginTestExecutionDto**](PhasePluginTestExecutionDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTestSuiteExecutionReportUsingGET**
> ListHolderDtoPhaseTestExecutionReportDto GetTestSuiteExecutionReportUsingGET(ctx, applicationId, applicationVersionId)
Get executed phases for release

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int64**| applicationId | 
  **applicationVersionId** | **int64**| applicationVersionId | 

### Return type

[**ListHolderDtoPhaseTestExecutionReportDto**](ListHolderDto«PhaseTestExecutionReportDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

