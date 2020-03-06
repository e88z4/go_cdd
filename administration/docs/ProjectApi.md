# \ProjectApi

All URIs are relative to *https://cddirector.io/cdd/administration/0000/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPartiesToProjectUsingPOST**](ProjectApi.md#AddPartiesToProjectUsingPOST) | **Post** /projects/{projectId}/parties | Create an association between a specific project and a collection of parties
[**CreateProjectUsingPOST**](ProjectApi.md#CreateProjectUsingPOST) | **Post** /projects | Creates project
[**DeleteProjectUsingDELETE**](ProjectApi.md#DeleteProjectUsingDELETE) | **Delete** /projects/{projectId} | Delete a project
[**GetAllProjectUserGroupsUsingGET**](ProjectApi.md#GetAllProjectUserGroupsUsingGET) | **Get** /projects/user-groups | Retrieves all project user groups
[**GetAllProjectsUsingGET**](ProjectApi.md#GetAllProjectsUsingGET) | **Get** /projects | Retrieves all projects
[**GetProjectUsersUsingGET**](ProjectApi.md#GetProjectUsersUsingGET) | **Get** /projects/users | Retrieves all users for project
[**GetProjectUsingGET**](ProjectApi.md#GetProjectUsingGET) | **Get** /projects/{projectId} | Retrieve a project
[**RemovePartiesFromProjectUsingDELETE**](ProjectApi.md#RemovePartiesFromProjectUsingDELETE) | **Delete** /projects/{projectId}/parties | Delete multiple parties to project association
[**RemoveParyFromProjectUsingDELETE**](ProjectApi.md#RemoveParyFromProjectUsingDELETE) | **Delete** /projects/{projectId}/parties/{partyId} | Delete party to project association
[**UpdateProjectUsingPUT**](ProjectApi.md#UpdateProjectUsingPUT) | **Put** /projects/{projectId} | Updates a project


# **AddPartiesToProjectUsingPOST**
> ListHolderDtoPartyDto AddPartiesToProjectUsingPOST(ctx, projectId, parties)
Create an association between a specific project and a collection of parties

Every party must include party ID and Role ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| projectId | 
  **parties** | [**[]PartyDto**](PartyDto.md)| parties | 

### Return type

[**ListHolderDtoPartyDto**](ListHolderDto«PartyDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateProjectUsingPOST**
> ProjectDto CreateProjectUsingPOST(ctx, projectDto)
Creates project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectDto** | [**ProjectDto**](ProjectDto.md)| projectDto | 

### Return type

[**ProjectDto**](ProjectDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteProjectUsingDELETE**
> DeleteProjectUsingDELETE(ctx, projectId)
Delete a project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| projectId | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllProjectUserGroupsUsingGET**
> PagedResultDtoProjectUserGroupDto GetAllProjectUserGroupsUsingGET(ctx, optional)
Retrieves all project user groups

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllProjectUserGroupsUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAllProjectUserGroupsUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isProjectParty** | **optional.Bool**| isProjectParty | 
 **filter** | **optional.String**| freeTextFilter | 
 **pageSize** | **optional.Int32**| pageSize | 
 **pageNumber** | **optional.Int32**| pageNumber | 

### Return type

[**PagedResultDtoProjectUserGroupDto**](PagedResultDto«ProjectUserGroupDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllProjectsUsingGET**
> PagedResultDtoProjectDto GetAllProjectsUsingGET(ctx, optional)
Retrieves all projects

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllProjectsUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAllProjectsUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **optional.String**| freeTextFilter | 
 **pageSize** | **optional.Int32**| pageSize | 
 **pageNumber** | **optional.Int32**| pageNumber | 
 **currentUser** | **optional.Bool**| currentUser | 

### Return type

[**PagedResultDtoProjectDto**](PagedResultDto«ProjectDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProjectUsersUsingGET**
> PagedResultDtoProjectUserDto GetProjectUsersUsingGET(ctx, optional)
Retrieves all users for project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetProjectUsersUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetProjectUsersUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isProjectParty** | **optional.Bool**| isProjectParty | 
 **roles** | [**optional.Interface of []int64**](int64.md)| filterByProjectRoleIds | 
 **filter** | **optional.String**| freeTextFilter | 
 **pageSize** | **optional.Int32**| pageSize | 
 **pageNumber** | **optional.Int32**| pageNumber | 

### Return type

[**PagedResultDtoProjectUserDto**](PagedResultDto«ProjectUserDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProjectUsingGET**
> ProjectDto GetProjectUsingGET(ctx, projectId)
Retrieve a project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| projctId | 

### Return type

[**ProjectDto**](ProjectDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemovePartiesFromProjectUsingDELETE**
> RemovePartiesFromProjectUsingDELETE(ctx, projectId, party)
Delete multiple parties to project association

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| projectId | 
  **party** | [**[]int64**](int64.md)| parties | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveParyFromProjectUsingDELETE**
> RemoveParyFromProjectUsingDELETE(ctx, projectId, partyId)
Delete party to project association

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| projectId | 
  **partyId** | **int64**| partyId | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateProjectUsingPUT**
> ProjectDto UpdateProjectUsingPUT(ctx, projectId, projectDto)
Updates a project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| projectId | 
  **projectDto** | [**ProjectDto**](ProjectDto.md)| projectDto | 

### Return type

[**ProjectDto**](ProjectDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

