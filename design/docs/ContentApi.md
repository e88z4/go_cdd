# \ContentApi

All URIs are relative to *https://cddirector.io/cdd/design/0000/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateManualContentUsingPOST**](ContentApi.md#CreateManualContentUsingPOST) | **Post** /releases/{releaseId}/application-versions/{applicationVersionId}/content-sources/{contentSourceId}/content-items | Creates a manual content
[**DeleteContentUsingDELETE**](ContentApi.md#DeleteContentUsingDELETE) | **Delete** /releases/{releaseId}/application-versions/{applicationVersionId}/content-sources/{contentSourceId}/content-items/{contentId} | Deletes a content
[**GetContentItemsEmbedUsingGET**](ContentApi.md#GetContentItemsEmbedUsingGET) | **Get** /releases/application-versions/{applicationVersionId}/content-sources/content-items | Get files by Content Items with  and test suites
[**GetContentItemsUsingGET**](ContentApi.md#GetContentItemsUsingGET) | **Get** /releases/application-versions/content-sources/content-items | Get Content Items by Filter
[**PatchContentUsingPATCH**](ContentApi.md#PatchContentUsingPATCH) | **Patch** /releases/{releaseId}/application-versions/{applicationVersionId}/content-sources/{contentSourceId}/content-items/{contentId} | Patch a content item
[**UpdateContentUsingPUT**](ContentApi.md#UpdateContentUsingPUT) | **Put** /releases/{releaseId}/application-versions/{applicationVersionId}/content-sources/{contentSourceId}/content-items/{contentId} | Updates a content


# **CreateManualContentUsingPOST**
> ContentItemDto CreateManualContentUsingPOST(ctx, releaseId, applicationVersionId, contentSourceId, contentItemDto)
Creates a manual content

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseId** | **int64**| releaseId | 
  **applicationVersionId** | **int64**| applicationVersionId | 
  **contentSourceId** | **int64**| contentSourceId | 
  **contentItemDto** | [**ContentItemDto**](ContentItemDto.md)| contentItemDto | 

### Return type

[**ContentItemDto**](ContentItemDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteContentUsingDELETE**
> ContentItemDto DeleteContentUsingDELETE(ctx, releaseId, applicationVersionId, contentSourceId, contentId)
Deletes a content

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseId** | **int64**| releaseId | 
  **applicationVersionId** | **int64**| applicationVersionId | 
  **contentSourceId** | **int64**| contentSourceId | 
  **contentId** | **int64**| contentId | 

### Return type

[**ContentItemDto**](ContentItemDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContentItemsEmbedUsingGET**
> ContentItemPagedDto GetContentItemsEmbedUsingGET(ctx, applicationVersionId, optional)
Get files by Content Items with  and test suites

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationVersionId** | **int64**| applicationVersionId | 
 **optional** | ***GetContentItemsEmbedUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetContentItemsEmbedUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageNumber** | **optional.Int32**| pageNumber | 
 **pageSize** | **optional.Int32**| pageSize | 
 **embed** | [**optional.Interface of []string**](string.md)| embed | 

### Return type

[**ContentItemPagedDto**](ContentItemPagedDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContentItemsUsingGET**
> PagedResultDtoContentItemDto GetContentItemsUsingGET(ctx, optional)
Get Content Items by Filter

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetContentItemsUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetContentItemsUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **optional.String**| freeTextFilter | 
 **pageNumber** | **optional.Int32**| pageNumber | 
 **pageSize** | **optional.Int32**| pageSize | 
 **track** | **optional.Int64**| byTrackId | 
 **parentContentItem** | [**optional.Interface of []int64**](int64.md)| byParentContentItemIds | 
 **importContentItems** | **optional.Bool**| onlyImportedContentItems | 
 **applicationVersionCommit** | **optional.Bool**| appVersionCommit | 
 **contentItemConflict** | [**optional.Interface of []string**](string.md)| contentItemConflicts | 

### Return type

[**PagedResultDtoContentItemDto**](PagedResultDto«ContentItemDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchContentUsingPATCH**
> ContentItemDto PatchContentUsingPATCH(ctx, releaseId, applicationVersionId, contentSourceId, contentId, contentItemDto)
Patch a content item

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseId** | **int64**| releaseId | 
  **applicationVersionId** | **int64**| applicationVersionId | 
  **contentSourceId** | **int64**| contentSourceId | 
  **contentId** | **int64**| contentId | 
  **contentItemDto** | [**ContentItemDto**](ContentItemDto.md)| contentItemDto | 

### Return type

[**ContentItemDto**](ContentItemDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateContentUsingPUT**
> ContentItemDto UpdateContentUsingPUT(ctx, releaseId, applicationVersionId, contentSourceId, contentId, contentItemDto)
Updates a content

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseId** | **int64**| releaseId | 
  **applicationVersionId** | **int64**| applicationVersionId | 
  **contentSourceId** | **int64**| contentSourceId | 
  **contentId** | **int64**| contentId | 
  **contentItemDto** | [**ContentItemDto**](ContentItemDto.md)| contentItemDto | 

### Return type

[**ContentItemDto**](ContentItemDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

