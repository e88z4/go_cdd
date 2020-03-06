# \ApiAccessTokenApi

All URIs are relative to *https://cddirector.io/cdd/administration/0000/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApiAccessTokenUsingPOST**](ApiAccessTokenApi.md#CreateApiAccessTokenUsingPOST) | **Post** /api-access-tokens | Create api access token
[**DeleteApiAccessTokenUsingDELETE**](ApiAccessTokenApi.md#DeleteApiAccessTokenUsingDELETE) | **Delete** /api-access-tokens/{apiAccessTokenId} | delete user api access token
[**DeleteApiAccessTokensUsingDELETE**](ApiAccessTokenApi.md#DeleteApiAccessTokensUsingDELETE) | **Delete** /api-access-tokens | delete all user api access tokens
[**GetApiAccessTokensUsingGET**](ApiAccessTokenApi.md#GetApiAccessTokensUsingGET) | **Get** /api-access-tokens/{apiAccessTokenId} | get user api access token
[**GetApiAccessTokensUsingGET1**](ApiAccessTokenApi.md#GetApiAccessTokensUsingGET1) | **Get** /api-access-tokens | get all user api access token


# **CreateApiAccessTokenUsingPOST**
> ApiAccessTokenDto CreateApiAccessTokenUsingPOST(ctx, )
Create api access token

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiAccessTokenDto**](ApiAccessTokenDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApiAccessTokenUsingDELETE**
> DeleteApiAccessTokenUsingDELETE(ctx, apiAccessTokenId)
delete user api access token

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiAccessTokenId** | **int64**| apiAccessTokenId | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApiAccessTokensUsingDELETE**
> DeleteApiAccessTokensUsingDELETE(ctx, )
delete all user api access tokens

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiAccessTokensUsingGET**
> ApiAccessTokenDto GetApiAccessTokensUsingGET(ctx, apiAccessTokenId)
get user api access token

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiAccessTokenId** | **int64**| apiAccessTokenId | 

### Return type

[**ApiAccessTokenDto**](ApiAccessTokenDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiAccessTokensUsingGET1**
> ListHolderDto GetApiAccessTokensUsingGET1(ctx, )
get all user api access token

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

