# \ContentSourceApi

All URIs are relative to *https://cddirector.io/cdd/design/0000/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateContentSourceUsingPOST**](ContentSourceApi.md#CreateContentSourceUsingPOST) | **Post** /releases/{releaseId}/application-versions/{applicationVersionId}/content-sources | Creates a content source
[**DeleteContentSourceUsingDELETE**](ContentSourceApi.md#DeleteContentSourceUsingDELETE) | **Delete** /releases/{releaseId}/application-versions/{applicationVersionId}/content-sources/{contentSourceId} | Delete a Content Source
[**GetContentSourceUsingGET**](ContentSourceApi.md#GetContentSourceUsingGET) | **Get** /releases/{releaseId}/application-versions/{applicationVersionId}/content-sources/{contentSourceId} | Get a Content Source
[**SyncContentSourceUsingPATCH**](ContentSourceApi.md#SyncContentSourceUsingPATCH) | **Patch** /releases/{releaseId}/application-versions/{applicationVersionId}/content-sources/{contentSourceId}/sync | Sync content source
[**UpdateContentSourceUsingPUT**](ContentSourceApi.md#UpdateContentSourceUsingPUT) | **Put** /releases/{releaseId}/application-versions/{applicationVersionId}/content-sources/{contentSourceId} | Edit a Content Source


# **CreateContentSourceUsingPOST**
> ContentSourceDto CreateContentSourceUsingPOST(ctx, releaseId, applicationVersionId, contentSourceDto)
Creates a content source

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseId** | **int64**| releaseId | 
  **applicationVersionId** | **int64**| applicationVersionId | 
  **contentSourceDto** | [**ContentSourceDto**](ContentSourceDto.md)| contentSourceDto | 

### Return type

[**ContentSourceDto**](ContentSourceDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteContentSourceUsingDELETE**
> DeleteContentSourceUsingDELETE(ctx, releaseId, applicationVersionId, contentSourceId)
Delete a Content Source

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseId** | **int64**| releaseId | 
  **applicationVersionId** | **int64**| applicationVersionId | 
  **contentSourceId** | **int64**| contentSourceId | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContentSourceUsingGET**
> ContentSourceDto GetContentSourceUsingGET(ctx, releaseId, applicationVersionId, contentSourceId, optional)
Get a Content Source

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseId** | **int64**| releaseId | 
  **applicationVersionId** | **int64**| applicationVersionId | 
  **contentSourceId** | **int64**| contentSourceId | 
 **optional** | ***GetContentSourceUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetContentSourceUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **embed** | [**optional.Interface of []string**](string.md)| embedFields | 
 **importContentItems** | **optional.Bool**| isImportContentItems | 
 **applicationVersionCommit** | **optional.Bool**| isAppVersionCommit | 

### Return type

[**ContentSourceDto**](ContentSourceDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SyncContentSourceUsingPATCH**
> ContentSourceDto SyncContentSourceUsingPATCH(ctx, releaseId, applicationVersionId, contentSourceId)
Sync content source

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseId** | **int64**| releaseId | 
  **applicationVersionId** | **int64**| applicationVersionId | 
  **contentSourceId** | **int64**| contentSourceId | 

### Return type

[**ContentSourceDto**](ContentSourceDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateContentSourceUsingPUT**
> ContentSourceDto UpdateContentSourceUsingPUT(ctx, releaseId, applicationVersionId, contentSourceId, contentSourceDto)
Edit a Content Source

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseId** | **int64**| releaseId | 
  **applicationVersionId** | **int64**| applicationVersionId | 
  **contentSourceId** | **int64**| contentSourceId | 
  **contentSourceDto** | [**ContentSourceDto**](ContentSourceDto.md)| contentSourceDto | 

### Return type

[**ContentSourceDto**](ContentSourceDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

