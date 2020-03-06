# \FileSourceApi

All URIs are relative to *https://cddirector.io/cdd/design/0000/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFileSourceUsingPOST**](FileSourceApi.md#CreateFileSourceUsingPOST) | **Post** /applications/application-versions/file-sources | Create a file source
[**DeleteFileSourceUsingDELETE**](FileSourceApi.md#DeleteFileSourceUsingDELETE) | **Delete** /applications/application-versions/file-sources/{fileSourceId} | Delete a file source
[**DslManifestWebhookUsingPOST**](FileSourceApi.md#DslManifestWebhookUsingPOST) | **Post** /applications/application-versions/file-sources/dsl-manifests/webhooks/{webhookTokenName} | Automatic release creation by editing release manifest files in a commit source
[**GetFileSourceUsingGET**](FileSourceApi.md#GetFileSourceUsingGET) | **Get** /applications/application-versions/file-sources/{fileSourceId} | Retrieve a file source
[**GetFileSourcesUsingGET**](FileSourceApi.md#GetFileSourcesUsingGET) | **Get** /applications/application-versions/file-sources | Retrieve all file source
[**PatchFileSourceUsingPATCH**](FileSourceApi.md#PatchFileSourceUsingPATCH) | **Patch** /applications/application-versions/file-sources/{fileSourceId} | Patch a file source
[**RegenerateWebhookUrlUsingPOST**](FileSourceApi.md#RegenerateWebhookUrlUsingPOST) | **Post** /applications/application-versions/file-sources/{fileSourceId}/webhooks | File source - Regenerate Webhook URL
[**UpdateFileSourceUsingPUT**](FileSourceApi.md#UpdateFileSourceUsingPUT) | **Put** /applications/application-versions/file-sources/{fileSourceId} | Update a file source


# **CreateFileSourceUsingPOST**
> FileSourceDto CreateFileSourceUsingPOST(ctx, fileSourceDto)
Create a file source

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fileSourceDto** | [**FileSourceDto**](FileSourceDto.md)| fileSourceDto | 

### Return type

[**FileSourceDto**](FileSourceDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFileSourceUsingDELETE**
> FileSourceDto DeleteFileSourceUsingDELETE(ctx, fileSourceId)
Delete a file source

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fileSourceId** | **int64**| fileSourceId | 

### Return type

[**FileSourceDto**](FileSourceDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DslManifestWebhookUsingPOST**
> DslManifestWebhookUsingPOST(ctx, webhookTokenName, jsonNode, optional)
Automatic release creation by editing release manifest files in a commit source

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **webhookTokenName** | **string**| webhookToken | 
  **jsonNode** | [**JsonNode**](JsonNode.md)| jsonNode | 
 **optional** | ***DslManifestWebhookUsingPOSTOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DslManifestWebhookUsingPOSTOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dslManifest** | **optional.String**| fileName | 
 **branchName** | **optional.String**| branchName | 
 **commitId** | **optional.String**| commitId | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFileSourceUsingGET**
> FileSourceDto GetFileSourceUsingGET(ctx, fileSourceId)
Retrieve a file source

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fileSourceId** | **int64**| fileSourceId | 

### Return type

[**FileSourceDto**](FileSourceDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFileSourcesUsingGET**
> ListHolderDtoFileSourceDto GetFileSourcesUsingGET(ctx, )
Retrieve all file source

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ListHolderDtoFileSourceDto**](ListHolderDto«FileSourceDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchFileSourceUsingPATCH**
> FileSourceDto PatchFileSourceUsingPATCH(ctx, fileSourceId, fileSourceDto)
Patch a file source

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fileSourceId** | **int64**| fileSourceId | 
  **fileSourceDto** | [**FileSourceDto**](FileSourceDto.md)| fileSourceDto | 

### Return type

[**FileSourceDto**](FileSourceDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegenerateWebhookUrlUsingPOST**
> FileSourceDto RegenerateWebhookUrlUsingPOST(ctx, fileSourceId)
File source - Regenerate Webhook URL

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fileSourceId** | **int64**| fileSourceId | 

### Return type

[**FileSourceDto**](FileSourceDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFileSourceUsingPUT**
> FileSourceDto UpdateFileSourceUsingPUT(ctx, fileSourceId, fileSourceDto)
Update a file source

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fileSourceId** | **int64**| fileSourceId | 
  **fileSourceDto** | [**FileSourceDto**](FileSourceDto.md)| fileSourceDto | 

### Return type

[**FileSourceDto**](FileSourceDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

