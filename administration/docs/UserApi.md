# \UserApi

All URIs are relative to *https://cddirector.io/cdd/administration/0000/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangeUserPasswordUsingPOST**](UserApi.md#ChangeUserPasswordUsingPOST) | **Post** /users/current/password/reset | Submit a password reset request (auth0)
[**ChangeUserPasswordUsingPUT**](UserApi.md#ChangeUserPasswordUsingPUT) | **Put** /users/{userId}/password | Change a own user password
[**CreateApiAccessTokenUsingPOST1**](UserApi.md#CreateApiAccessTokenUsingPOST1) | **Post** /users/{userId}/api-access-tokens | Create api access token for appropriate user
[**CreateUsersUsingPOST**](UserApi.md#CreateUsersUsingPOST) | **Post** /users | Creates a user
[**DeleteUserUsingDELETE**](UserApi.md#DeleteUserUsingDELETE) | **Delete** /users/{userId} | Deletes a user
[**GetApiAccessTokenUsingGET**](UserApi.md#GetApiAccessTokenUsingGET) | **Get** /users/{userId}/api-access-tokens | Get api access token for appropriate user
[**GetUserUsingGET**](UserApi.md#GetUserUsingGET) | **Get** /users/{userId} | Retrieves a user
[**GetUsersUsingGET**](UserApi.md#GetUsersUsingGET) | **Get** /users | Retrieves all users
[**PatchUserUsingPATCH**](UserApi.md#PatchUserUsingPATCH) | **Patch** /users/{userId} | Patches a user
[**UpdateUserUsingPUT**](UserApi.md#UpdateUserUsingPUT) | **Put** /users/{userId} | Updates a user


# **ChangeUserPasswordUsingPOST**
> UserDto ChangeUserPasswordUsingPOST(ctx, )
Submit a password reset request (auth0)

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**UserDto**](UserDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChangeUserPasswordUsingPUT**
> UserDto ChangeUserPasswordUsingPUT(ctx, userId, changeUserPasswordDto)
Change a own user password

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **int64**| userId | 
  **changeUserPasswordDto** | [**ChangeUserPasswordDto**](ChangeUserPasswordDto.md)| changeUserPasswordDto | 

### Return type

[**UserDto**](UserDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateApiAccessTokenUsingPOST1**
> ApiAccessTokenDto CreateApiAccessTokenUsingPOST1(ctx, userId)
Create api access token for appropriate user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **int64**| userId | 

### Return type

[**ApiAccessTokenDto**](ApiAccessTokenDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateUsersUsingPOST**
> ListHolderDtoUserDto CreateUsersUsingPOST(ctx, userListDto)
Creates a user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userListDto** | [**[]UserDto**](UserDto.md)| userListDto | 

### Return type

[**ListHolderDtoUserDto**](ListHolderDto«UserDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteUserUsingDELETE**
> DeleteUserUsingDELETE(ctx, userId)
Deletes a user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **int64**| userId | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiAccessTokenUsingGET**
> ApiAccessTokenDto GetApiAccessTokenUsingGET(ctx, userId)
Get api access token for appropriate user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **int64**| userId | 

### Return type

[**ApiAccessTokenDto**](ApiAccessTokenDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserUsingGET**
> UserDto GetUserUsingGET(ctx, userId)
Retrieves a user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **int64**| userId | 

### Return type

[**UserDto**](UserDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsersUsingGET**
> PagedResultDtoUserDto GetUsersUsingGET(ctx, optional)
Retrieves all users

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetUsersUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetUsersUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isSystem** | **optional.Bool**| isSystem | 
 **sortField** | **optional.String**| primarySortField | 
 **sortDirection** | **optional.String**| sortDirection | 
 **filter** | **optional.String**| freeTextFilter | 
 **pageSize** | **optional.Int32**| pageSize | 
 **pageNumber** | **optional.Int32**| pageNumber | 

### Return type

[**PagedResultDtoUserDto**](PagedResultDto«UserDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchUserUsingPATCH**
> UserDto PatchUserUsingPATCH(ctx, userId, userDto)
Patches a user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **int64**| userId | 
  **userDto** | [**UserDto**](UserDto.md)| userDto | 

### Return type

[**UserDto**](UserDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUserUsingPUT**
> UserDto UpdateUserUsingPUT(ctx, userId, userDto)
Updates a user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **int64**| userId | 
  **userDto** | [**UserDto**](UserDto.md)| userDto | 

### Return type

[**UserDto**](UserDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

