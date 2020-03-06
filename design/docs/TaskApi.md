# \TaskApi

All URIs are relative to *https://cddirector.io/cdd/design/0000/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangeTaskTransitionsUsingPATCH**](TaskApi.md#ChangeTaskTransitionsUsingPATCH) | **Patch** /releases/{releaseId}/phases/{phaseId}/tasks/{taskId}/transitions | Change task transitions
[**CreateTaskUsingPOST**](TaskApi.md#CreateTaskUsingPOST) | **Post** /releases/{releaseId}/phases/{phaseId}/tasks | Creates a task
[**DeleteTaskUsingDELETE**](TaskApi.md#DeleteTaskUsingDELETE) | **Delete** /releases/{releaseId}/phases/{phaseId}/tasks/{taskId} | Deletes a task
[**DeleteTasksUsingDELETE**](TaskApi.md#DeleteTasksUsingDELETE) | **Delete** /releases/{releaseId}/phases/{phaseId}/tasks | Delete multiple tasks
[**GetAllTasksOfPhaseUsingGET**](TaskApi.md#GetAllTasksOfPhaseUsingGET) | **Get** /releases/{releaseId}/phases/{phaseId}/tasks | Retrieves all tasks of a phase
[**GetTaskUsingGET**](TaskApi.md#GetTaskUsingGET) | **Get** /releases/{releaseId}/phases/{phaseId}/tasks/{taskId} | Retrieves a task
[**PatchTaskUsingPATCH**](TaskApi.md#PatchTaskUsingPATCH) | **Patch** /releases/{releaseId}/phases/{phaseId}/tasks/{taskId} | Changes a task
[**PatchTasksUsingPATCH**](TaskApi.md#PatchTasksUsingPATCH) | **Patch** /releases/{releaseId}/phases/{phaseId}/tasks | Change multiple tasks
[**UpdateTaskUsingPUT**](TaskApi.md#UpdateTaskUsingPUT) | **Put** /releases/{releaseId}/phases/{phaseId}/tasks/{taskId} | Updates a task


# **ChangeTaskTransitionsUsingPATCH**
> TaskDto ChangeTaskTransitionsUsingPATCH(ctx, releaseId, phaseId, taskId, taskDto)
Change task transitions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseId** | **int64**| releaseId | 
  **phaseId** | **int64**| phaseId | 
  **taskId** | **int64**| taskId | 
  **taskDto** | [**TaskDto**](TaskDto.md)| taskDto | 

### Return type

[**TaskDto**](TaskDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateTaskUsingPOST**
> TaskDto CreateTaskUsingPOST(ctx, releaseId, phaseId, taskDto)
Creates a task

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseId** | **int64**| containingReleaseId | 
  **phaseId** | **int64**| containingPhaseId | 
  **taskDto** | [**TaskDto**](TaskDto.md)| taskDto | 

### Return type

[**TaskDto**](TaskDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTaskUsingDELETE**
> TaskDto DeleteTaskUsingDELETE(ctx, releaseId, phaseId, taskId)
Deletes a task

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseId** | **int64**| containingReleaseId | 
  **phaseId** | **int64**| containingPhaseId | 
  **taskId** | **int64**| taskId | 

### Return type

[**TaskDto**](TaskDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTasksUsingDELETE**
> DeleteTasksUsingDELETE(ctx, releaseId, phaseId, task)
Delete multiple tasks

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseId** | **int64**| containingReleaseId | 
  **phaseId** | **int64**| containingPhaseId | 
  **task** | [**[]int64**](int64.md)| taskIds | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllTasksOfPhaseUsingGET**
> ListHolderDtoTaskDto GetAllTasksOfPhaseUsingGET(ctx, releaseId, phaseId)
Retrieves all tasks of a phase

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseId** | **int64**| containingReleaseId | 
  **phaseId** | **int64**| phaseId | 

### Return type

[**ListHolderDtoTaskDto**](ListHolderDto«TaskDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTaskUsingGET**
> TaskDto GetTaskUsingGET(ctx, releaseId, phaseId, taskId)
Retrieves a task

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseId** | **int64**| containingReleaseId | 
  **phaseId** | **int64**| containingPhaseId | 
  **taskId** | **int64**| taskId | 

### Return type

[**TaskDto**](TaskDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchTaskUsingPATCH**
> TaskDto PatchTaskUsingPATCH(ctx, releaseId, phaseId, taskId, taskDto)
Changes a task

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseId** | **int64**| containingReleaseId | 
  **phaseId** | **int64**| containingPhaseId | 
  **taskId** | **int64**| taskId | 
  **taskDto** | [**TaskDto**](TaskDto.md)| taskDto | 

### Return type

[**TaskDto**](TaskDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchTasksUsingPATCH**
> PatchTasksUsingPATCH(ctx, releaseId, phaseId, taskDTOs)
Change multiple tasks

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseId** | **int64**| containingReleaseId | 
  **phaseId** | **int64**| containingPhaseId | 
  **taskDTOs** | [**ListHolderDtoTaskDto**](ListHolderDtoTaskDto.md)| taskDTOs | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTaskUsingPUT**
> TaskDto UpdateTaskUsingPUT(ctx, releaseId, phaseId, taskId, taskDto)
Updates a task

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseId** | **int64**| containingReleaseId | 
  **phaseId** | **int64**| containingPhaseId | 
  **taskId** | **int64**| taskId | 
  **taskDto** | [**TaskDto**](TaskDto.md)| taskDto | 

### Return type

[**TaskDto**](TaskDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

