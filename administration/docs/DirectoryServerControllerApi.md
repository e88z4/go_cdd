# \DirectoryServerControllerApi

All URIs are relative to *https://cddirector.io/cdd/administration/0000/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDirectoryServerUsingPOST**](DirectoryServerControllerApi.md#CreateDirectoryServerUsingPOST) | **Post** /directory-servers | Creates a new directory server definition
[**DeleteDirectoryServerUsingDELETE**](DirectoryServerControllerApi.md#DeleteDirectoryServerUsingDELETE) | **Delete** /directory-servers/{directoryServerId} | Deletes a directory server definition
[**GetDirectoryServerUsingGET**](DirectoryServerControllerApi.md#GetDirectoryServerUsingGET) | **Get** /directory-servers/{directoryServerId} | Retrieves a directory server definition
[**GetDirectoryServersUsingGET**](DirectoryServerControllerApi.md#GetDirectoryServersUsingGET) | **Get** /directory-servers | Retrieves all directory server definitions
[**GetPotentialGroupsUsingGET**](DirectoryServerControllerApi.md#GetPotentialGroupsUsingGET) | **Get** /directory-user-groups | Searches for groups to be imported, inside a given directory server
[**GetPotentialUsersUsingGET**](DirectoryServerControllerApi.md#GetPotentialUsersUsingGET) | **Get** /directory-users | Searches for users to be imported, inside a given directory server
[**PatchDirectoryServerUsingPATCH**](DirectoryServerControllerApi.md#PatchDirectoryServerUsingPATCH) | **Patch** /directory-servers/{directoryServerId} | Patches a directory server definition
[**TestDirectoryServerConnectivityUsingPOST**](DirectoryServerControllerApi.md#TestDirectoryServerConnectivityUsingPOST) | **Post** /directory-servers/connectivity-tests | Connectivity Test for Directory Server
[**UpdateDirectoryServerUsingPUT**](DirectoryServerControllerApi.md#UpdateDirectoryServerUsingPUT) | **Put** /directory-servers/{directoryServerId} | Updates a directory server definition


# **CreateDirectoryServerUsingPOST**
> DirectoryServerDto CreateDirectoryServerUsingPOST(ctx, directoryServerDto)
Creates a new directory server definition

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **directoryServerDto** | [**DirectoryServerDto**](DirectoryServerDto.md)| directoryServerDto | 

### Return type

[**DirectoryServerDto**](DirectoryServerDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDirectoryServerUsingDELETE**
> DeleteDirectoryServerUsingDELETE(ctx, directoryServerId)
Deletes a directory server definition

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **directoryServerId** | **int64**| directoryServerId | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDirectoryServerUsingGET**
> DirectoryServerDto GetDirectoryServerUsingGET(ctx, directoryServerId)
Retrieves a directory server definition

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **directoryServerId** | **int64**| directoryServerId | 

### Return type

[**DirectoryServerDto**](DirectoryServerDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDirectoryServersUsingGET**
> ListHolderDtoDirectoryServerDto GetDirectoryServersUsingGET(ctx, )
Retrieves all directory server definitions

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ListHolderDtoDirectoryServerDto**](ListHolderDto«DirectoryServerDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPotentialGroupsUsingGET**
> ListHolderDtoUserGroupDto GetPotentialGroupsUsingGET(ctx, directoryServer, optional)
Searches for groups to be imported, inside a given directory server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **directoryServer** | **int64**| directoryServerId | 
 **optional** | ***GetPotentialGroupsUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetPotentialGroupsUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **optional.String**| filter | 
 **pageSize** | **optional.Int32**| pageSize | 
 **filterExistEntities** | **optional.Bool**| filterExistEntities | 

### Return type

[**ListHolderDtoUserGroupDto**](ListHolderDto«UserGroupDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPotentialUsersUsingGET**
> ListHolderDtoUserDto GetPotentialUsersUsingGET(ctx, directoryServer, optional)
Searches for users to be imported, inside a given directory server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **directoryServer** | **int64**| directoryServerId | 
 **optional** | ***GetPotentialUsersUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetPotentialUsersUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **optional.String**| filter | 
 **pageSize** | **optional.Int32**| pageSize | 
 **filterExistEntities** | **optional.Bool**| filterExistEntities | 

### Return type

[**ListHolderDtoUserDto**](ListHolderDto«UserDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchDirectoryServerUsingPATCH**
> DirectoryServerDto PatchDirectoryServerUsingPATCH(ctx, directoryServerId, directoryServerDto)
Patches a directory server definition

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **directoryServerId** | **int64**| directoryServerId | 
  **directoryServerDto** | [**DirectoryServerDto**](DirectoryServerDto.md)| directoryServerDto | 

### Return type

[**DirectoryServerDto**](DirectoryServerDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TestDirectoryServerConnectivityUsingPOST**
> DirectoryServerDto TestDirectoryServerConnectivityUsingPOST(ctx, directoryServerDto)
Connectivity Test for Directory Server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **directoryServerDto** | [**DirectoryServerDto**](DirectoryServerDto.md)| directoryServerDto | 

### Return type

[**DirectoryServerDto**](DirectoryServerDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDirectoryServerUsingPUT**
> DirectoryServerDto UpdateDirectoryServerUsingPUT(ctx, directoryServerId, directoryServerDto)
Updates a directory server definition

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **directoryServerId** | **int64**| directoryServerId | 
  **directoryServerDto** | [**DirectoryServerDto**](DirectoryServerDto.md)| directoryServerDto | 

### Return type

[**DirectoryServerDto**](DirectoryServerDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

