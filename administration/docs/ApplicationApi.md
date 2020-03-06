# \ApplicationApi

All URIs are relative to *https://cddirector.io/cdd/administration/0000/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApplicationUsingPOST**](ApplicationApi.md#CreateApplicationUsingPOST) | **Post** /applications | Creates an application
[**DeleteApplicationUsingDELETE**](ApplicationApi.md#DeleteApplicationUsingDELETE) | **Delete** /applications/{applicationId} | Deletes an application
[**GetAllApplicationsUsingGET**](ApplicationApi.md#GetAllApplicationsUsingGET) | **Get** /applications | Retrieves all applications
[**GetApplicationReportUsingGET**](ApplicationApi.md#GetApplicationReportUsingGET) | **Get** /applications/{applicationId}/environments-version | Get an application report
[**GetApplicationUsingGET**](ApplicationApi.md#GetApplicationUsingGET) | **Get** /applications/{applicationId} | Retrieves an application
[**PatchApplicationUsingPATCH**](ApplicationApi.md#PatchApplicationUsingPATCH) | **Patch** /applications/{applicationId} | Patches an application
[**UpdateApplicationUsingPUT**](ApplicationApi.md#UpdateApplicationUsingPUT) | **Put** /applications/{applicationId} | Updates an application


# **CreateApplicationUsingPOST**
> ApplicationDto CreateApplicationUsingPOST(ctx, applicationDto)
Creates an application

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationDto** | [**ApplicationDto**](ApplicationDto.md)| applicationDto | 

### Return type

[**ApplicationDto**](ApplicationDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApplicationUsingDELETE**
> DeleteApplicationUsingDELETE(ctx, applicationId)
Deletes an application

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int64**| applicationId | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllApplicationsUsingGET**
> PagedResultDtoApplicationDto GetAllApplicationsUsingGET(ctx, optional)
Retrieves all applications

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllApplicationsUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAllApplicationsUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **optional.String**| freeTextFilter | 
 **embed** | [**optional.Interface of []string**](string.md)| embedFields | 
 **name** | **optional.String**| nameFilter | 
 **pageSize** | **optional.Int32**| pageSize | 
 **pageNumber** | **optional.Int32**| pageNumber | 

### Return type

[**PagedResultDtoApplicationDto**](PagedResultDto«ApplicationDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApplicationReportUsingGET**
> ApplicationReportDto GetApplicationReportUsingGET(ctx, applicationId)
Get an application report

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int64**| applicationId | 

### Return type

[**ApplicationReportDto**](ApplicationReportDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApplicationUsingGET**
> ApplicationDto GetApplicationUsingGET(ctx, applicationId)
Retrieves an application

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int64**| applicationId | 

### Return type

[**ApplicationDto**](ApplicationDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchApplicationUsingPATCH**
> ApplicationDto PatchApplicationUsingPATCH(ctx, applicationId, applicationDto)
Patches an application

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int64**| applicationId | 
  **applicationDto** | [**ApplicationDto**](ApplicationDto.md)| applicationDto | 

### Return type

[**ApplicationDto**](ApplicationDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateApplicationUsingPUT**
> ApplicationDto UpdateApplicationUsingPUT(ctx, applicationId, applicationDto)
Updates an application

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int64**| applicationId | 
  **applicationDto** | [**ApplicationDto**](ApplicationDto.md)| applicationDto | 

### Return type

[**ApplicationDto**](ApplicationDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

