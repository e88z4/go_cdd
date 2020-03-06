# \ApplicationSourceApi

All URIs are relative to *https://cddirector.io/cdd/administration/0000/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApplicationSourceUsingPOST**](ApplicationSourceApi.md#CreateApplicationSourceUsingPOST) | **Post** /application-sources | Creates an application source
[**DeleteApplicationSourceUsingDELETE**](ApplicationSourceApi.md#DeleteApplicationSourceUsingDELETE) | **Delete** /application-sources/{applicationSourceId} | Deletes an application source
[**GetAllApplicationSourcesUsingGET**](ApplicationSourceApi.md#GetAllApplicationSourcesUsingGET) | **Get** /application-sources | Retrieves all application sources
[**SyncApplicationSourceUsingPATCH**](ApplicationSourceApi.md#SyncApplicationSourceUsingPATCH) | **Patch** /application-sources/{applicationSourceId}/sync | Syncs an application source


# **CreateApplicationSourceUsingPOST**
> ApplicationSourceDto CreateApplicationSourceUsingPOST(ctx, applicationSourceDto)
Creates an application source

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationSourceDto** | [**ApplicationSourceDto**](ApplicationSourceDto.md)| applicationSourceDto | 

### Return type

[**ApplicationSourceDto**](ApplicationSourceDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApplicationSourceUsingDELETE**
> DeleteApplicationSourceUsingDELETE(ctx, applicationSourceId)
Deletes an application source

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationSourceId** | **int64**| applicationSourceId | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllApplicationSourcesUsingGET**
> ListHolderDto GetAllApplicationSourcesUsingGET(ctx, )
Retrieves all application sources

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ListHolderDto**](ListHolderDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SyncApplicationSourceUsingPATCH**
> ApplicationSourceDto SyncApplicationSourceUsingPATCH(ctx, applicationSourceId)
Syncs an application source

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationSourceId** | **int64**| applicationSourceId | 

### Return type

[**ApplicationSourceDto**](ApplicationSourceDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

