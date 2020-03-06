# \PermissionsApi

All URIs are relative to *https://cddirector.io/cdd/administration/0000/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPermissionsUsingGET**](PermissionsApi.md#GetPermissionsUsingGET) | **Get** /permissions | Get permissions


# **GetPermissionsUsingGET**
> ListHolderDtoPermissionDto GetPermissionsUsingGET(ctx, optional)
Get permissions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetPermissionsUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetPermissionsUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isProjectSpecific** | **optional.Bool**| isProjectSpecific | 

### Return type

[**ListHolderDtoPermissionDto**](ListHolderDto«PermissionDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

