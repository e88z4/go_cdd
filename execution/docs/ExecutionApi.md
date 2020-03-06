# \ExecutionApi

All URIs are relative to *https://cddirector.io/cdd/execution/0000/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangePhaseExecutionUsingPATCH**](ExecutionApi.md#ChangePhaseExecutionUsingPATCH) | **Patch** /releases-execution/{releaseId}/phases-execution/{phaseId} | Changes a phase execution
[**ChangePhaseExecutionUsingPATCH1**](ExecutionApi.md#ChangePhaseExecutionUsingPATCH1) | **Patch** /releases-execution/{releaseName}/{releaseVersionName}/phases-execution/{phaseName} | Change a phase execution
[**ChangeReleaseExecutionUsingPATCH**](ExecutionApi.md#ChangeReleaseExecutionUsingPATCH) | **Patch** /releases-execution/{releaseId} | Changes a release execution
[**ChangeTaskExecutionUsingPATCH**](ExecutionApi.md#ChangeTaskExecutionUsingPATCH) | **Patch** /releases-execution/{releaseId}/phases-execution/{phaseId}/tasks-execution/{taskId} | Change a task execution
[**CreateReleaseManifestUsingPOST**](ExecutionApi.md#CreateReleaseManifestUsingPOST) | **Post** /releases-execution/{releaseId}/manifests | Creates a release execution manifest
[**DeleteReleaseManifestUsingDELETE**](ExecutionApi.md#DeleteReleaseManifestUsingDELETE) | **Delete** /releases-execution/{releaseId}/manifests/{manifestId} | Deletes a release execution manifest
[**GetAllPhaseTasksExecutionUsingGET**](ExecutionApi.md#GetAllPhaseTasksExecutionUsingGET) | **Get** /releases-execution/{releaseId}/phases-execution/{phaseId}/tasks-execution | Get all phase tasks execution
[**GetPhaseExecutionUsingGET**](ExecutionApi.md#GetPhaseExecutionUsingGET) | **Get** /releases-execution/{releaseId}/phases-execution/{phaseId} | Retrieves a phase execution
[**GetPhasesExecutionUsingGET**](ExecutionApi.md#GetPhasesExecutionUsingGET) | **Get** /releases-execution/{releaseId}/phases-execution | Retrieves all phases execution
[**GetReleaseExecutionUsingGET**](ExecutionApi.md#GetReleaseExecutionUsingGET) | **Get** /releases-execution/{releaseId} | Retrieves a release execution
[**GetReleaseManifestUsingGET**](ExecutionApi.md#GetReleaseManifestUsingGET) | **Get** /releases-execution/{releaseId}/manifests/{manifestId} | Gets a release execution manifest
[**GetTaskExecutionUsingGET**](ExecutionApi.md#GetTaskExecutionUsingGET) | **Get** /releases-execution/{releaseId}/phases-execution/{phaseId}/tasks-execution/{taskId} | Get a task execution


# **ChangePhaseExecutionUsingPATCH**
> PhaseExecutionDto ChangePhaseExecutionUsingPATCH(ctx, releaseId, phaseId, phaseExceution)
Changes a phase execution

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseId** | **int64**| releaseId | 
  **phaseId** | **int64**| phaseId | 
  **phaseExceution** | [**PhaseExecutionDto**](PhaseExecutionDto.md)| phaseExceution | 

### Return type

[**PhaseExecutionDto**](PhaseExecutionDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChangePhaseExecutionUsingPATCH1**
> PhaseExecutionDto ChangePhaseExecutionUsingPATCH1(ctx, releaseName, releaseVersionName, phaseName, phaseExecution)
Change a phase execution

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseName** | **string**| releaseName | 
  **releaseVersionName** | **string**| releaseVersion | 
  **phaseName** | **string**| phaseName | 
  **phaseExecution** | [**PhaseExecutionDto**](PhaseExecutionDto.md)| phaseExecution | 

### Return type

[**PhaseExecutionDto**](PhaseExecutionDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChangeReleaseExecutionUsingPATCH**
> ReleaseExecutionDto ChangeReleaseExecutionUsingPATCH(ctx, releaseId, releaseExecution, optional)
Changes a release execution

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseId** | **int64**| releaseId | 
  **releaseExecution** | [**ReleaseExecutionDto**](ReleaseExecutionDto.md)| releaseExecution | 
 **optional** | ***ChangeReleaseExecutionUsingPATCHOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChangeReleaseExecutionUsingPATCHOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **force** | **optional.Bool**| isForce | 

### Return type

[**ReleaseExecutionDto**](ReleaseExecutionDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChangeTaskExecutionUsingPATCH**
> TaskExecutionDto ChangeTaskExecutionUsingPATCH(ctx, releaseId, phaseId, taskId, taskExceution)
Change a task execution

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseId** | **int64**| releaseId | 
  **phaseId** | **int64**| phaseId | 
  **taskId** | **int64**| taskId | 
  **taskExceution** | [**TaskExecutionDto**](TaskExecutionDto.md)| taskExceution | 

### Return type

[**TaskExecutionDto**](TaskExecutionDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateReleaseManifestUsingPOST**
> ReleaseManifestDto CreateReleaseManifestUsingPOST(ctx, releaseId, releaseManifestDto)
Creates a release execution manifest

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseId** | **int64**| releaseId | 
  **releaseManifestDto** | [**ReleaseManifestDto**](ReleaseManifestDto.md)| releaseManifestDto | 

### Return type

[**ReleaseManifestDto**](ReleaseManifestDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteReleaseManifestUsingDELETE**
> IdentifiableDto DeleteReleaseManifestUsingDELETE(ctx, releaseId, manifestId)
Deletes a release execution manifest

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseId** | **int64**| releaseId | 
  **manifestId** | **int64**| manifestId | 

### Return type

[**IdentifiableDto**](IdentifiableDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllPhaseTasksExecutionUsingGET**
> ListHolderDtoTaskExecutionDto GetAllPhaseTasksExecutionUsingGET(ctx, releaseId, phaseId)
Get all phase tasks execution

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseId** | **int64**| releaseId | 
  **phaseId** | **int64**| phaseId | 

### Return type

[**ListHolderDtoTaskExecutionDto**](ListHolderDto«TaskExecutionDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPhaseExecutionUsingGET**
> PhaseExecutionDto GetPhaseExecutionUsingGET(ctx, releaseId, phaseId)
Retrieves a phase execution

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseId** | **int64**| releaseId | 
  **phaseId** | **int64**| phaseId | 

### Return type

[**PhaseExecutionDto**](PhaseExecutionDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPhasesExecutionUsingGET**
> ListHolderDtoPhaseExecutionDto GetPhasesExecutionUsingGET(ctx, releaseId)
Retrieves all phases execution

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseId** | **int64**| releaseId | 

### Return type

[**ListHolderDtoPhaseExecutionDto**](ListHolderDto«PhaseExecutionDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReleaseExecutionUsingGET**
> ReleaseExecutionDto GetReleaseExecutionUsingGET(ctx, releaseId)
Retrieves a release execution

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseId** | **int64**| releaseId | 

### Return type

[**ReleaseExecutionDto**](ReleaseExecutionDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReleaseManifestUsingGET**
> IdentifiableDto GetReleaseManifestUsingGET(ctx, releaseId, manifestId)
Gets a release execution manifest

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseId** | **int64**| releaseId | 
  **manifestId** | **int64**| manifestId | 

### Return type

[**IdentifiableDto**](IdentifiableDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTaskExecutionUsingGET**
> TaskExecutionDto GetTaskExecutionUsingGET(ctx, releaseId, phaseId, taskId)
Get a task execution

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseId** | **int64**| releaseId | 
  **phaseId** | **int64**| phaseId | 
  **taskId** | **int64**| taskId | 

### Return type

[**TaskExecutionDto**](TaskExecutionDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

