# \ReleaseApi

All URIs are relative to *https://cddirector.io/cdd/design/0000/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignApplicationToReleaseUsingPOST**](ReleaseApi.md#AssignApplicationToReleaseUsingPOST) | **Post** /releases/{releaseId}/applications | Assigns an application to a release
[**CreateReleaseUsingPOST**](ReleaseApi.md#CreateReleaseUsingPOST) | **Post** /releases | Creates or duplicates a release
[**DeleteReleaseUsingDELETE**](ReleaseApi.md#DeleteReleaseUsingDELETE) | **Delete** /releases/{releaseId} | Deletes a release
[**DeleteReleasesUsingDELETE**](ReleaseApi.md#DeleteReleasesUsingDELETE) | **Delete** /releases | Deletes multiple releases
[**GetAllEnvironmentsOfReleaseUsingGET**](ReleaseApi.md#GetAllEnvironmentsOfReleaseUsingGET) | **Get** /releases/{releaseId}/environments | Retrieves release environments
[**GetAllReleasesUsingGET**](ReleaseApi.md#GetAllReleasesUsingGET) | **Get** /releases | Retrieves all releases
[**GetReleaseStatusReasonsUsingGET**](ReleaseApi.md#GetReleaseStatusReasonsUsingGET) | **Get** /releases/release-status-reasons | Retrieves a release
[**GetReleaseUsingGET**](ReleaseApi.md#GetReleaseUsingGET) | **Get** /releases/{releaseId} | Retrieves a release
[**PatchReleaseUsingPATCH**](ReleaseApi.md#PatchReleaseUsingPATCH) | **Patch** /releases/{releaseId} | Patches a release
[**UnassignApplicationFromReleaseUsingDELETE**](ReleaseApi.md#UnassignApplicationFromReleaseUsingDELETE) | **Delete** /releases/{releaseId}/applications/{applicationId} | Unassigns an application from a release
[**UpdateReleaseUsingPUT**](ReleaseApi.md#UpdateReleaseUsingPUT) | **Put** /releases/{releaseId} | Updates a release


# **AssignApplicationToReleaseUsingPOST**
> ListHolderDtoNamedIdentifiableDto AssignApplicationToReleaseUsingPOST(ctx, releaseId, applicationIdsDto)
Assigns an application to a release

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseId** | **int64**| releaseId | 
  **applicationIdsDto** | [**[]IdentifiableDto**](IdentifiableDto.md)| applicationIdsDto | 

### Return type

[**ListHolderDtoNamedIdentifiableDto**](ListHolderDto«NamedIdentifiableDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateReleaseUsingPOST**
> ReleaseDto CreateReleaseUsingPOST(ctx, releaseDto)
Creates or duplicates a release

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseDto** | [**ReleaseDto**](ReleaseDto.md)| releaseDto | 

### Return type

[**ReleaseDto**](ReleaseDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteReleaseUsingDELETE**
> DeleteReleaseUsingDELETE(ctx, releaseId)
Deletes a release

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseId** | **int64**| releaseId | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteReleasesUsingDELETE**
> DeleteReleasesUsingDELETE(ctx, release)
Deletes multiple releases

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **release** | [**[]int64**](int64.md)| releaseIds | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllEnvironmentsOfReleaseUsingGET**
> ListHolderDtoEnvironmentDto GetAllEnvironmentsOfReleaseUsingGET(ctx, releaseId, optional)
Retrieves release environments

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseId** | **int64**| releaseId | 
 **optional** | ***GetAllEnvironmentsOfReleaseUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAllEnvironmentsOfReleaseUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **optional.String**| filter | 

### Return type

[**ListHolderDtoEnvironmentDto**](ListHolderDto«EnvironmentDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllReleasesUsingGET**
> PagedResultWithPermissionsDtoReleaseDto GetAllReleasesUsingGET(ctx, optional)
Retrieves all releases

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllReleasesUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAllReleasesUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **optional.String**| freeTextFilter | 
 **pageSize** | **optional.Int32**| pageSize | 
 **pageNumber** | **optional.Int32**| pageNumber | 
 **status** | [**optional.Interface of []string**](string.md)| filterByStatus | 
 **application** | [**optional.Interface of []int64**](int64.md)| filterByApplicationIds | 
 **embed** | [**optional.Interface of []string**](string.md)| embedFields | 
 **currentUser** | **optional.Bool**| showMyReleases | 
 **owner** | [**optional.Interface of []int64**](int64.md)| ownerList | 
 **member** | [**optional.Interface of []int64**](int64.md)| memberList | 
 **startDate** | **optional.Int64**| startDate | 
 **endDate** | **optional.Int64**| endDate | 
 **taskStatus** | [**optional.Interface of []string**](string.md)| embeddedTaskStatusFilter | 
 **hasContentItems** | **optional.Bool**| hasContentFilter | 
 **notPartOfTrack** | **optional.Bool**| notPartOfTrack | 
 **isProjectSpecific** | **optional.Bool**| isProjectSpecific | 

### Return type

[**PagedResultWithPermissionsDtoReleaseDto**](PagedResultWithPermissionsDto«ReleaseDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReleaseStatusReasonsUsingGET**
> ListHolderDtoStringIdentifiableDto GetReleaseStatusReasonsUsingGET(ctx, )
Retrieves a release

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ListHolderDtoStringIdentifiableDto**](ListHolderDto«StringIdentifiableDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReleaseUsingGET**
> ReleaseDto GetReleaseUsingGET(ctx, releaseId)
Retrieves a release

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseId** | **int64**| releaseId | 

### Return type

[**ReleaseDto**](ReleaseDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchReleaseUsingPATCH**
> ReleaseDto PatchReleaseUsingPATCH(ctx, releaseId, releaseDto)
Patches a release

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseId** | **int64**| releaseId | 
  **releaseDto** | [**ReleaseDto**](ReleaseDto.md)| releaseDto | 

### Return type

[**ReleaseDto**](ReleaseDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnassignApplicationFromReleaseUsingDELETE**
> ListHolderDtoNamedIdentifiableDto UnassignApplicationFromReleaseUsingDELETE(ctx, releaseId, applicationId)
Unassigns an application from a release

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseId** | **int64**| releaseId | 
  **applicationId** | **int64**| applicationId | 

### Return type

[**ListHolderDtoNamedIdentifiableDto**](ListHolderDto«NamedIdentifiableDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateReleaseUsingPUT**
> ReleaseDto UpdateReleaseUsingPUT(ctx, releaseId, releaseDto)
Updates a release

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseId** | **int64**| releaseId | 
  **releaseDto** | [**ReleaseDto**](ReleaseDto.md)| releaseDto | 

### Return type

[**ReleaseDto**](ReleaseDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

