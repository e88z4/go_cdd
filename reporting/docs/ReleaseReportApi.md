# \ReleaseReportApi

All URIs are relative to *https://cddirector.io/cdd/reporting/0000/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetExecutedPhaseUsingGET**](ReleaseReportApi.md#GetExecutedPhaseUsingGET) | **Get** /reports/releases/{releaseId}/executed-phases/{phaseId} | Get executed tasks for executed phase
[**GetExecutedPhasesTokensUsingGET**](ReleaseReportApi.md#GetExecutedPhasesTokensUsingGET) | **Get** /reports/releases/executed-phases | Get tokens of executed phases
[**GetExecutedPhasesUsingGET1**](ReleaseReportApi.md#GetExecutedPhasesUsingGET1) | **Get** /reports/releases/{releaseId} | Get executed phases for release
[**GetExecutedTaskParametersUsingGET**](ReleaseReportApi.md#GetExecutedTaskParametersUsingGET) | **Get** /reports/releases/{releaseId}/executed-phases/{phaseId}/executed-tasks/{taskId} | Get executed parameters for executed task
[**GetExecutedTasksUsingGET**](ReleaseReportApi.md#GetExecutedTasksUsingGET) | **Get** /reports/releases/executed-phases/executed-tasks | Get executed tasks By filter


# **GetExecutedPhaseUsingGET**
> ListHolderDtoExecutedTaskReportDto GetExecutedPhaseUsingGET(ctx, releaseId, phaseId)
Get executed tasks for executed phase

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseId** | **int64**| releaseId | 
  **phaseId** | **int64**| executedPhaseId | 

### Return type

[**ListHolderDtoExecutedTaskReportDto**](ListHolderDto«ExecutedTaskReportDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetExecutedPhasesTokensUsingGET**
> ListHolderDtoExecutedPhaseReportDto GetExecutedPhasesTokensUsingGET(ctx, optional)
Get tokens of executed phases

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetExecutedPhasesTokensUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetExecutedPhasesTokensUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **release** | [**optional.Interface of []int64**](int64.md)| releaseIds | 
 **pageSize** | **optional.Int32**| pageSize | 
 **pageNumber** | **optional.Int32**| pageNumber | 

### Return type

[**ListHolderDtoExecutedPhaseReportDto**](ListHolderDto«ExecutedPhaseReportDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetExecutedPhasesUsingGET1**
> ListHolderDtoPhaseExecutionsReportDto GetExecutedPhasesUsingGET1(ctx, releaseId)
Get executed phases for release

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseId** | **int64**| releaseId | 

### Return type

[**ListHolderDtoPhaseExecutionsReportDto**](ListHolderDto«PhaseExecutionsReportDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetExecutedTaskParametersUsingGET**
> ExecutedTaskReportDto GetExecutedTaskParametersUsingGET(ctx, releaseId, phaseId, taskId)
Get executed parameters for executed task

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseId** | **int64**| releaseId | 
  **phaseId** | **int64**| executedPhaseId | 
  **taskId** | **int64**| executedTaskId | 

### Return type

[**ExecutedTaskReportDto**](ExecutedTaskReportDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetExecutedTasksUsingGET**
> ListHolderDtoExecutedTaskReportDto GetExecutedTasksUsingGET(ctx, optional)
Get executed tasks By filter

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetExecutedTasksUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetExecutedTasksUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **release** | [**optional.Interface of []int64**](int64.md)| releaseIds | 
 **phase** | [**optional.Interface of []int64**](int64.md)| phasesIds | 
 **pluginName** | **optional.String**| pluginName | 
 **pluginVendor** | **optional.String**| pluginVendor | 
 **pluginVersion** | **optional.String**| pluginVersion | 
 **taskTemplateName** | **optional.String**| taskTemplateName | 
 **endpointName** | **optional.String**| endpointName | 
 **sortField** | **optional.String**| sortField | 
 **sortDirection** | **optional.String**| sortDirection | 

### Return type

[**ListHolderDtoExecutedTaskReportDto**](ListHolderDto«ExecutedTaskReportDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

