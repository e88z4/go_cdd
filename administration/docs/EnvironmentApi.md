# \EnvironmentApi

All URIs are relative to *https://cddirector.io/cdd/administration/0000/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEnvironmentAtApplicationUsingPOST**](EnvironmentApi.md#CreateEnvironmentAtApplicationUsingPOST) | **Post** /applications/{applicationId}/environments | Create an environment
[**DeleteEnvironmentUsingDELETE**](EnvironmentApi.md#DeleteEnvironmentUsingDELETE) | **Delete** /applications/{applicationId}/environments/{environmentId} | Delete an environment
[**GetAllEnvironmentsUsingGET**](EnvironmentApi.md#GetAllEnvironmentsUsingGET) | **Get** /environments | Get all environments
[**GetCalendarEnvironmentEntitiesUsingGET**](EnvironmentApi.md#GetCalendarEnvironmentEntitiesUsingGET) | **Get** /environments/planner | Get calendar environment entities : phases, maintenance windows and freeze periods
[**GetEnvironmentAtApplicationUsingGET**](EnvironmentApi.md#GetEnvironmentAtApplicationUsingGET) | **Get** /applications/{applicationId}/environments/{environmentId} | Get an environment at application
[**GetEnvironmentUsingGET**](EnvironmentApi.md#GetEnvironmentUsingGET) | **Get** /environments/{environmentId} | Get an environment
[**GetEnvironmentsAtApplicationUsingGET**](EnvironmentApi.md#GetEnvironmentsAtApplicationUsingGET) | **Get** /applications/{applicationId}/environments | Get application environments
[**PatchEnvironmentAtApplicationUsingPATCH**](EnvironmentApi.md#PatchEnvironmentAtApplicationUsingPATCH) | **Patch** /applications/{applicationId}/environments/{environmentId} | Patch an environment at application
[**PatchEnvironmentUsingPATCH**](EnvironmentApi.md#PatchEnvironmentUsingPATCH) | **Patch** /environments/{environmentId} | Patch an environment
[**UpdateEnvironmentAtApplicationUsingPUT**](EnvironmentApi.md#UpdateEnvironmentAtApplicationUsingPUT) | **Put** /applications/{applicationId}/environments/{environmentId} | Update an environment at application
[**UpdateEnvironmentUsingPUT**](EnvironmentApi.md#UpdateEnvironmentUsingPUT) | **Put** /environments/{environmentId} | Update an environment


# **CreateEnvironmentAtApplicationUsingPOST**
> EnvironmentDto CreateEnvironmentAtApplicationUsingPOST(ctx, applicationId, environmentDto)
Create an environment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int64**| applicationId | 
  **environmentDto** | [**EnvironmentDto**](EnvironmentDto.md)| environmentDto | 

### Return type

[**EnvironmentDto**](EnvironmentDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEnvironmentUsingDELETE**
> DeleteEnvironmentUsingDELETE(ctx, applicationId, environmentId)
Delete an environment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int64**| applicationId | 
  **environmentId** | **int64**| environmentId | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllEnvironmentsUsingGET**
> PagedResultDtoEnvironmentDto GetAllEnvironmentsUsingGET(ctx, optional)
Get all environments

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllEnvironmentsUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAllEnvironmentsUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **optional.String**| freeTextFilter | 
 **pageSize** | **optional.Int32**| pageSize | 
 **pageNumber** | **optional.Int32**| pageNumber | 
 **isProjectSpecific** | **optional.Bool**| isProjectSpecific | 

### Return type

[**PagedResultDtoEnvironmentDto**](PagedResultDto«EnvironmentDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCalendarEnvironmentEntitiesUsingGET**
> EnvironmentEntitiesDto GetCalendarEnvironmentEntitiesUsingGET(ctx, startDate, endDate, environment, optional)
Get calendar environment entities : phases, maintenance windows and freeze periods

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **startDate** | **int64**| startDate | 
  **endDate** | **int64**| endDate | 
  **environment** | [**[]int64**](int64.md)| environmentIds | 
 **optional** | ***GetCalendarEnvironmentEntitiesUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetCalendarEnvironmentEntitiesUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **embed** | [**optional.Interface of []string**](string.md)| embedFields | 

### Return type

[**EnvironmentEntitiesDto**](EnvironmentEntitiesDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEnvironmentAtApplicationUsingGET**
> EnvironmentDto GetEnvironmentAtApplicationUsingGET(ctx, applicationId, environmentId)
Get an environment at application

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int64**| applicationId | 
  **environmentId** | **int64**| environmentId | 

### Return type

[**EnvironmentDto**](EnvironmentDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEnvironmentUsingGET**
> EnvironmentDto GetEnvironmentUsingGET(ctx, environmentId)
Get an environment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **environmentId** | **int64**| environmentId | 

### Return type

[**EnvironmentDto**](EnvironmentDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEnvironmentsAtApplicationUsingGET**
> ListHolderDtoEnvironmentDto GetEnvironmentsAtApplicationUsingGET(ctx, applicationId, optional)
Get application environments

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int64**| applicationId | 
 **optional** | ***GetEnvironmentsAtApplicationUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetEnvironmentsAtApplicationUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| environmentName | 

### Return type

[**ListHolderDtoEnvironmentDto**](ListHolderDto«EnvironmentDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchEnvironmentAtApplicationUsingPATCH**
> EnvironmentDto PatchEnvironmentAtApplicationUsingPATCH(ctx, environmentId, environmentDto)
Patch an environment at application

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **environmentId** | **int64**| environmentId | 
  **environmentDto** | [**EnvironmentDto**](EnvironmentDto.md)| environmentDto | 

### Return type

[**EnvironmentDto**](EnvironmentDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchEnvironmentUsingPATCH**
> EnvironmentDto PatchEnvironmentUsingPATCH(ctx, environmentId, environmentDto)
Patch an environment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **environmentId** | **int64**| environmentId | 
  **environmentDto** | [**EnvironmentDto**](EnvironmentDto.md)| environmentDto | 

### Return type

[**EnvironmentDto**](EnvironmentDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEnvironmentAtApplicationUsingPUT**
> EnvironmentDto UpdateEnvironmentAtApplicationUsingPUT(ctx, environmentId, environmentDto)
Update an environment at application

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **environmentId** | **int64**| environmentId | 
  **environmentDto** | [**EnvironmentDto**](EnvironmentDto.md)| environmentDto | 

### Return type

[**EnvironmentDto**](EnvironmentDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEnvironmentUsingPUT**
> EnvironmentDto UpdateEnvironmentUsingPUT(ctx, environmentId, environmentDto)
Update an environment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **environmentId** | **int64**| environmentId | 
  **environmentDto** | [**EnvironmentDto**](EnvironmentDto.md)| environmentDto | 

### Return type

[**EnvironmentDto**](EnvironmentDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

