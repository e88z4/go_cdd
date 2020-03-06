# \ReleaseTokenApi

All URIs are relative to *https://cddirector.io/cdd/design/0000/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateReleaseTokenUsingPOST**](ReleaseTokenApi.md#CreateReleaseTokenUsingPOST) | **Post** /releases/{releaseId}/tokens | Creates a release token
[**DeleteReleaseTokenUsingDELETE**](ReleaseTokenApi.md#DeleteReleaseTokenUsingDELETE) | **Delete** /releases/{releaseId}/tokens/{tokenId} | Deletes a release token
[**GetReleaseTokenUsagesUsingGET**](ReleaseTokenApi.md#GetReleaseTokenUsagesUsingGET) | **Get** /releases/{releaseId}/tokens/{tokenId}/usages | Retrieves a release token usages
[**GetReleaseTokenUsingGET**](ReleaseTokenApi.md#GetReleaseTokenUsingGET) | **Get** /releases/{releaseId}/tokens/{tokenId} | Retrieves a release token
[**GetReleaseTokensOfReleaseUsingGET**](ReleaseTokenApi.md#GetReleaseTokensOfReleaseUsingGET) | **Get** /releases/{releaseId}/tokens | Retrieves all tokens of a release
[**PatchReleaseTokenByNameUsingPATCH**](ReleaseTokenApi.md#PatchReleaseTokenByNameUsingPATCH) | **Patch** /releases/{releaseName}/{releaseVersionName}/tokens/{tokenName} | Update a release token
[**PatchReleaseTokenValuesUsingPATCH**](ReleaseTokenApi.md#PatchReleaseTokenValuesUsingPATCH) | **Patch** /releases/{releaseName}/{releaseVersionName}/tokens | Update a release tokens value
[**UpdateReleaseTokenUsingPUT**](ReleaseTokenApi.md#UpdateReleaseTokenUsingPUT) | **Put** /releases/{releaseId}/tokens/{tokenId} | Updates a release token


# **CreateReleaseTokenUsingPOST**
> ReleaseTokenDto CreateReleaseTokenUsingPOST(ctx, releaseId, releaseTokenDto)
Creates a release token

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseId** | **int64**| releaseId | 
  **releaseTokenDto** | [**ReleaseTokenDto**](ReleaseTokenDto.md)| releaseTokenDto | 

### Return type

[**ReleaseTokenDto**](ReleaseTokenDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteReleaseTokenUsingDELETE**
> IdentifiableDto DeleteReleaseTokenUsingDELETE(ctx, releaseId, tokenId)
Deletes a release token

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseId** | **int64**| releaseId | 
  **tokenId** | **int64**| tokenId | 

### Return type

[**IdentifiableDto**](IdentifiableDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReleaseTokenUsagesUsingGET**
> ListHolderDtoEntityFullLocationDto GetReleaseTokenUsagesUsingGET(ctx, releaseId, tokenId)
Retrieves a release token usages

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseId** | **int64**| releaseId | 
  **tokenId** | **int64**| tokenId | 

### Return type

[**ListHolderDtoEntityFullLocationDto**](ListHolderDto«EntityFullLocationDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReleaseTokenUsingGET**
> ReleaseTokenDto GetReleaseTokenUsingGET(ctx, releaseId, tokenId)
Retrieves a release token

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseId** | **int64**| releaseId | 
  **tokenId** | **int64**| tokenId | 

### Return type

[**ReleaseTokenDto**](ReleaseTokenDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReleaseTokensOfReleaseUsingGET**
> ListHolderDtoReleaseTokenDto GetReleaseTokensOfReleaseUsingGET(ctx, releaseId, optional)
Retrieves all tokens of a release

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseId** | **int64**| releaseId | 
 **optional** | ***GetReleaseTokensOfReleaseUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetReleaseTokensOfReleaseUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **optional.String**| freeTextFilter | 

### Return type

[**ListHolderDtoReleaseTokenDto**](ListHolderDto«ReleaseTokenDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchReleaseTokenByNameUsingPATCH**
> ReleaseTokenDto PatchReleaseTokenByNameUsingPATCH(ctx, releaseName, releaseVersionName, tokenName, releaseTokenDto)
Update a release token

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseName** | **string**| releaseName | 
  **releaseVersionName** | **string**| releaseVersion | 
  **tokenName** | **string**| tokenName | 
  **releaseTokenDto** | [**ReleaseTokenDto**](ReleaseTokenDto.md)| releaseTokenDto | 

### Return type

[**ReleaseTokenDto**](ReleaseTokenDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchReleaseTokenValuesUsingPATCH**
> ListHolderDtoReleaseTokenDto PatchReleaseTokenValuesUsingPATCH(ctx, releaseName, releaseVersionName, releaseTokenDtos)
Update a release tokens value

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseName** | **string**| releaseName | 
  **releaseVersionName** | **string**| releaseVersion | 
  **releaseTokenDtos** | [**[]ReleaseTokenDto**](ReleaseTokenDto.md)| releaseTokenDtos | 

### Return type

[**ListHolderDtoReleaseTokenDto**](ListHolderDto«ReleaseTokenDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateReleaseTokenUsingPUT**
> ReleaseTokenDto UpdateReleaseTokenUsingPUT(ctx, releaseId, tokenId, releaseTokenDto)
Updates a release token

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseId** | **int64**| releaseId | 
  **tokenId** | **int64**| releaseTokenId | 
  **releaseTokenDto** | [**ReleaseTokenDto**](ReleaseTokenDto.md)| releaseTokenDto | 

### Return type

[**ReleaseTokenDto**](ReleaseTokenDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

