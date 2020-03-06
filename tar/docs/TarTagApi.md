# \TarTagApi

All URIs are relative to *https://cddirector.io/cdd/tar/0000/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTagsUsingGET1**](TarTagApi.md#GetTagsUsingGET1) | **Get** /tags | Retrieves filtered tags


# **GetTagsUsingGET1**
> PagedResultDtostring GetTagsUsingGET1(ctx, optional)
Retrieves filtered tags

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetTagsUsingGET1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetTagsUsingGET1Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **optional.String**| freeTextFilter | 
 **pageSize** | **optional.Int32**| pageSize | 
 **pageNumber** | **optional.Int32**| pageNumber | 

### Return type

[**PagedResultDtostring**](PagedResultDto«string».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

