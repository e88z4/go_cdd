# \EnvironmentTokenApi

All URIs are relative to *https://cddirector.io/cdd/design/0000/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEnvironmentTokenUsingPOST**](EnvironmentTokenApi.md#CreateEnvironmentTokenUsingPOST) | **Post** /tokens/environment-tokens | Creates an environment token
[**DeleteEnvironmentTokenUsingDELETE**](EnvironmentTokenApi.md#DeleteEnvironmentTokenUsingDELETE) | **Delete** /tokens/environment-tokens/{environmentTokenId} | Deletes an environment token
[**GetEnvironmentTokenUsingGET**](EnvironmentTokenApi.md#GetEnvironmentTokenUsingGET) | **Get** /tokens/environment-tokens/{environmentTokenId} | Retrieves an environment token
[**GetEnvironmentTokensUsingGET**](EnvironmentTokenApi.md#GetEnvironmentTokensUsingGET) | **Get** /tokens/environment-tokens | Retrieves all environment tokens
[**UpdateEnvironmentTokenUsingPUT**](EnvironmentTokenApi.md#UpdateEnvironmentTokenUsingPUT) | **Put** /tokens/environment-tokens/{environmentTokenId} | Update an environment token


# **CreateEnvironmentTokenUsingPOST**
> EnvironmentTokenDto CreateEnvironmentTokenUsingPOST(ctx, environmentTokenDto)
Creates an environment token

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **environmentTokenDto** | [**EnvironmentTokenDto**](EnvironmentTokenDto.md)| environmentTokenDto | 

### Return type

[**EnvironmentTokenDto**](EnvironmentTokenDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEnvironmentTokenUsingDELETE**
> DeleteEnvironmentTokenUsingDELETE(ctx, environmentTokenId)
Deletes an environment token

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **environmentTokenId** | **int64**| environmentTokenId | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEnvironmentTokenUsingGET**
> EnvironmentTokenDto GetEnvironmentTokenUsingGET(ctx, environmentTokenId)
Retrieves an environment token

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **environmentTokenId** | **int64**| environmentTokenId | 

### Return type

[**EnvironmentTokenDto**](EnvironmentTokenDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEnvironmentTokensUsingGET**
> ListHolderDtoEnvironmentTokenDto GetEnvironmentTokensUsingGET(ctx, optional)
Retrieves all environment tokens

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetEnvironmentTokensUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetEnvironmentTokensUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **optional.String**| freeTextFilter | 
 **pageSize** | **optional.Int32**| pageSize | 
 **pageNumber** | **optional.Int32**| pageNumber | 

### Return type

[**ListHolderDtoEnvironmentTokenDto**](ListHolderDto«EnvironmentTokenDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEnvironmentTokenUsingPUT**
> EnvironmentTokenDto UpdateEnvironmentTokenUsingPUT(ctx, environmentTokenId, environmentTokenDto)
Update an environment token

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **environmentTokenId** | **int64**| environmentTokenId | 
  **environmentTokenDto** | [**EnvironmentTokenDto**](EnvironmentTokenDto.md)| environmentTokenDto | 

### Return type

[**EnvironmentTokenDto**](EnvironmentTokenDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

