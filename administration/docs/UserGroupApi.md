# \UserGroupApi

All URIs are relative to *https://cddirector.io/cdd/administration/0000/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUserGroupsUsingPOST**](UserGroupApi.md#CreateUserGroupsUsingPOST) | **Post** /user-groups | Creates a user group
[**DeleteUserGroupUsingDELETE**](UserGroupApi.md#DeleteUserGroupUsingDELETE) | **Delete** /user-groups/{userGroupId} | Deletes a user group
[**GetAllUserGroupsUsingGET**](UserGroupApi.md#GetAllUserGroupsUsingGET) | **Get** /user-groups | Retrieves all user groups
[**GetUserGroupUsingGET**](UserGroupApi.md#GetUserGroupUsingGET) | **Get** /user-groups/{userGroupId} | Retrieves a user group
[**UpdateUserGroupUsingPUT**](UserGroupApi.md#UpdateUserGroupUsingPUT) | **Put** /user-groups/{userGroupId} | Updates a user group
[**UpdateUserGroupsUsingPUT**](UserGroupApi.md#UpdateUserGroupsUsingPUT) | **Put** /user-groups | Updates user groups


# **CreateUserGroupsUsingPOST**
> ListHolderDtoUserGroupDto CreateUserGroupsUsingPOST(ctx, userGroupDto)
Creates a user group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userGroupDto** | [**[]UserGroupDto**](UserGroupDto.md)| userGroupDto | 

### Return type

[**ListHolderDtoUserGroupDto**](ListHolderDto«UserGroupDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteUserGroupUsingDELETE**
> DeleteUserGroupUsingDELETE(ctx, userGroupId)
Deletes a user group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userGroupId** | **int64**| userId | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllUserGroupsUsingGET**
> PagedResultDtoUserGroupDto GetAllUserGroupsUsingGET(ctx, optional)
Retrieves all user groups

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllUserGroupsUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAllUserGroupsUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **optional.String**| freeTextFilter | 
 **pageSize** | **optional.Int32**| pageSize | 
 **pageNumber** | **optional.Int32**| pageNumber | 

### Return type

[**PagedResultDtoUserGroupDto**](PagedResultDto«UserGroupDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserGroupUsingGET**
> UserGroupDto GetUserGroupUsingGET(ctx, userGroupId)
Retrieves a user group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userGroupId** | **int64**| userGroupId | 

### Return type

[**UserGroupDto**](UserGroupDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUserGroupUsingPUT**
> UserGroupDto UpdateUserGroupUsingPUT(ctx, userGroupId, userGroupDto)
Updates a user group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userGroupId** | **int64**| userGroupId | 
  **userGroupDto** | [**UserGroupDto**](UserGroupDto.md)| userGroupDto | 

### Return type

[**UserGroupDto**](UserGroupDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUserGroupsUsingPUT**
> ListHolderDtoUserGroupDto UpdateUserGroupsUsingPUT(ctx, usersIdsDto, groups)
Updates user groups

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **usersIdsDto** | [**[]IdentifiableDto**](IdentifiableDto.md)| usersIdsDto | 
  **groups** | [**[]int64**](int64.md)| userGroupsIds | 

### Return type

[**ListHolderDtoUserGroupDto**](ListHolderDto«UserGroupDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

