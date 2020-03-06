# \PermissionsApi

All URIs are relative to *https://cddirector.io/cdd/design/0000/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPhasePermissionsUsingGET**](PermissionsApi.md#GetPhasePermissionsUsingGET) | **Get** /releases/{releaseId}/phases/{phaseId}/permissions | Get phase editing permissions
[**GetReleasePermissionsUsingGET**](PermissionsApi.md#GetReleasePermissionsUsingGET) | **Get** /releases/{releaseId}/permissions | Get release editing permissions
[**GetTaskPermissionsUsingGET**](PermissionsApi.md#GetTaskPermissionsUsingGET) | **Get** /releases/{releaseId}/phases/{phaseId}/tasks/{taskId}/permissions | Get task editing permissions


# **GetPhasePermissionsUsingGET**
> EntityPermissionsDto GetPhasePermissionsUsingGET(ctx, releaseId, phaseId)
Get phase editing permissions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseId** | **int64**| releaseId | 
  **phaseId** | **int64**| phaseId | 

### Return type

[**EntityPermissionsDto**](EntityPermissionsDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReleasePermissionsUsingGET**
> EntityPermissionsDto GetReleasePermissionsUsingGET(ctx, releaseId)
Get release editing permissions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseId** | **int64**| releaseId | 

### Return type

[**EntityPermissionsDto**](EntityPermissionsDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTaskPermissionsUsingGET**
> EntityPermissionsDto GetTaskPermissionsUsingGET(ctx, releaseId, phaseId, taskId)
Get task editing permissions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseId** | **int64**| releaseId | 
  **phaseId** | **int64**| phaseId | 
  **taskId** | **int64**| taskId | 

### Return type

[**EntityPermissionsDto**](EntityPermissionsDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

