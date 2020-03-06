# \ApiApi

All URIs are relative to *https://cddirector.io/cdd/administration/0000/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiDocumentationUsingGET**](ApiApi.md#GetApiDocumentationUsingGET) | **Get** /api | Retrieves the API documentation


# **GetApiDocumentationUsingGET**
> Json GetApiDocumentationUsingGET(ctx, optional)
Retrieves the API documentation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetApiDocumentationUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetApiDocumentationUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **group** | **optional.String**| groupName | 

### Return type

[**Json**](Json.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

