# \RolesApi

All URIs are relative to *https://cddirector.io/cdd/administration/0000/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRoleUsingPOST**](RolesApi.md#CreateRoleUsingPOST) | **Post** /roles | Create a role
[**DeleteRoleUsingDELETE**](RolesApi.md#DeleteRoleUsingDELETE) | **Delete** /roles/{roleId} | Delete a role
[**GetRoleUsingGET**](RolesApi.md#GetRoleUsingGET) | **Get** /roles/{roleId} | Get a role
[**GetRolesUsingGET**](RolesApi.md#GetRolesUsingGET) | **Get** /roles | Get roles
[**UpdateRoleUsingPUT**](RolesApi.md#UpdateRoleUsingPUT) | **Put** /roles/{roleId} | Update a role


# **CreateRoleUsingPOST**
> RoleDto CreateRoleUsingPOST(ctx, roleDto)
Create a role

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roleDto** | [**RoleDto**](RoleDto.md)| roleDto | 

### Return type

[**RoleDto**](RoleDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRoleUsingDELETE**
> DeleteRoleUsingDELETE(ctx, roleId)
Delete a role

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roleId** | **int64**| roleId | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRoleUsingGET**
> RoleDto GetRoleUsingGET(ctx, roleId)
Get a role

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roleId** | **int64**| roleId | 

### Return type

[**RoleDto**](RoleDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRolesUsingGET**
> ListHolderDtoRoleDto GetRolesUsingGET(ctx, optional)
Get roles

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetRolesUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetRolesUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isProjectSpecific** | **optional.Bool**| isProjectSpecific | 

### Return type

[**ListHolderDtoRoleDto**](ListHolderDto«RoleDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRoleUsingPUT**
> RoleDto UpdateRoleUsingPUT(ctx, roleId, roleDto)
Update a role

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roleId** | **int64**| roleId | 
  **roleDto** | [**RoleDto**](RoleDto.md)| roleDto | 

### Return type

[**RoleDto**](RoleDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

