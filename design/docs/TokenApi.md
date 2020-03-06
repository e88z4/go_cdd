# \TokenApi

All URIs are relative to *https://cddirector.io/cdd/design/0000/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTokensUsingGET**](TokenApi.md#GetTokensUsingGET) | **Get** /tokens | getTokens


# **GetTokensUsingGET**
> TokenListHolderDto GetTokensUsingGET(ctx, optional)
getTokens

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetTokensUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetTokensUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **release** | **optional.Int64**| releaseId | 
 **phase** | **optional.Int64**| phaseId | 
 **application** | **optional.Int64**| applicationId | 
 **filter** | **optional.String**| filter | 
 **limit** | **optional.Int32**| limit | 

### Return type

[**TokenListHolderDto**](TokenListHolderDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

