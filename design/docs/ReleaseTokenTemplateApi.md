# \ReleaseTokenTemplateApi

All URIs are relative to *https://cddirector.io/cdd/design/0000/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateReleaseTokenTemplateUsingPOST**](ReleaseTokenTemplateApi.md#CreateReleaseTokenTemplateUsingPOST) | **Post** /tokens/release-token-templates | Creates a release token template
[**DeleteReleaseTokenTemplateUsingDELETE**](ReleaseTokenTemplateApi.md#DeleteReleaseTokenTemplateUsingDELETE) | **Delete** /tokens/release-token-templates/{releaseTokenTemplateId} | Delete release token template
[**GetReleaseTokenTemplateUsingGET**](ReleaseTokenTemplateApi.md#GetReleaseTokenTemplateUsingGET) | **Get** /tokens/release-token-templates/{releaseTokenTemplateId} | Retrieves a release token template
[**GetReleaseTokensTemplateUsingGET**](ReleaseTokenTemplateApi.md#GetReleaseTokensTemplateUsingGET) | **Get** /tokens/release-token-templates | Retrieves all release tokens template
[**UpdateReleaseTemplateTokenUsingPUT**](ReleaseTokenTemplateApi.md#UpdateReleaseTemplateTokenUsingPUT) | **Put** /tokens/release-token-templates/{releaseTokenTemplateId} | Update a release template token


# **CreateReleaseTokenTemplateUsingPOST**
> ReleaseTokenTemplateDto CreateReleaseTokenTemplateUsingPOST(ctx, releaseTokenTemplateDto)
Creates a release token template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseTokenTemplateDto** | [**ReleaseTokenTemplateDto**](ReleaseTokenTemplateDto.md)| releaseTokenTemplateDto | 

### Return type

[**ReleaseTokenTemplateDto**](ReleaseTokenTemplateDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteReleaseTokenTemplateUsingDELETE**
> DeleteReleaseTokenTemplateUsingDELETE(ctx, releaseTokenTemplateId)
Delete release token template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseTokenTemplateId** | **int64**| releaseTokenTemplateId | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReleaseTokenTemplateUsingGET**
> ReleaseTokenTemplateDto GetReleaseTokenTemplateUsingGET(ctx, releaseTokenTemplateId)
Retrieves a release token template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseTokenTemplateId** | **int64**| releaseTemplateTokenId | 

### Return type

[**ReleaseTokenTemplateDto**](ReleaseTokenTemplateDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReleaseTokensTemplateUsingGET**
> ListHolderDtoReleaseTokenTemplateDto GetReleaseTokensTemplateUsingGET(ctx, optional)
Retrieves all release tokens template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetReleaseTokensTemplateUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetReleaseTokensTemplateUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **release** | **optional.Int64**| releaseId | 
 **filter** | **optional.String**| freeTextFilter | 
 **filterFields** | [**optional.Interface of []string**](string.md)| filterFieldDtos | 
 **pageSize** | **optional.Int32**| pageSize | 
 **pageNumber** | **optional.Int32**| pageNumber | 

### Return type

[**ListHolderDtoReleaseTokenTemplateDto**](ListHolderDto«ReleaseTokenTemplateDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateReleaseTemplateTokenUsingPUT**
> ReleaseTokenTemplateDto UpdateReleaseTemplateTokenUsingPUT(ctx, releaseTokenTemplateId, releaseTokenTemplateDto)
Update a release template token

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseTokenTemplateId** | **int64**| releaseTemplateTokenId | 
  **releaseTokenTemplateDto** | [**ReleaseTokenTemplateDto**](ReleaseTokenTemplateDto.md)| releaseTokenTemplateDto | 

### Return type

[**ReleaseTokenTemplateDto**](ReleaseTokenTemplateDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

