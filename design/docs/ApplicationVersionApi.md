# \ApplicationVersionApi

All URIs are relative to *https://cddirector.io/cdd/design/0000/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignApplicationVersionUsingPOST**](ApplicationVersionApi.md#AssignApplicationVersionUsingPOST) | **Post** /releases/{releaseId}/application-versions | Assigns an application version to the release (and creates it if required)
[**GetApplicationVersionOfApplicationInReleaseUsingGET**](ApplicationVersionApi.md#GetApplicationVersionOfApplicationInReleaseUsingGET) | **Get** /releases/{releaseId}/applications/{applicationId}/application-versions | Retrieves application versions of application under release
[**GetApplicationVersionUsingGET**](ApplicationVersionApi.md#GetApplicationVersionUsingGET) | **Get** /releases/{releaseId}/application-versions/{applicationVersionId} | Retrieves an application version
[**GetApplicationVersionsOfApplicationUsingGET**](ApplicationVersionApi.md#GetApplicationVersionsOfApplicationUsingGET) | **Get** /applications/{applicationId}/application-versions | Retrieves application version of application
[**GetApplicationVersionsUsingGET**](ApplicationVersionApi.md#GetApplicationVersionsUsingGET) | **Get** /releases/{releaseId}/application-versions | Retrieves all application versions under the release
[**GetAvailableDependenciesUsingGET**](ApplicationVersionApi.md#GetAvailableDependenciesUsingGET) | **Get** /application-versions/{applicationVersionId}/available-dependencies | Retrieves available application version dependencies
[**UnAssignApplicationVersionUsingDELETE**](ApplicationVersionApi.md#UnAssignApplicationVersionUsingDELETE) | **Delete** /releases/{releaseId}/application-versions/{applicationVersionId} | Un assigns an application version from a release (will delete in case it has no references anymore)
[**UpdateApplicationVersionUsingPUT**](ApplicationVersionApi.md#UpdateApplicationVersionUsingPUT) | **Put** /releases/{releaseId}/application-versions/{applicationVersionId} | Updates an application version


# **AssignApplicationVersionUsingPOST**
> ApplicationVersionDto AssignApplicationVersionUsingPOST(ctx, releaseId, applicationVersionDto, optional)
Assigns an application version to the release (and creates it if required)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseId** | **int64**| releaseId | 
  **applicationVersionDto** | [**ApplicationVersionDto**](ApplicationVersionDto.md)| applicationVersionDto | 
 **optional** | ***AssignApplicationVersionUsingPOSTOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AssignApplicationVersionUsingPOSTOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **force** | **optional.Bool**| force | 

### Return type

[**ApplicationVersionDto**](ApplicationVersionDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApplicationVersionOfApplicationInReleaseUsingGET**
> ApplicationVersionDto GetApplicationVersionOfApplicationInReleaseUsingGET(ctx, releaseId, applicationId)
Retrieves application versions of application under release

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseId** | **int64**| releaseId | 
  **applicationId** | **int64**| applicationId | 

### Return type

[**ApplicationVersionDto**](ApplicationVersionDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApplicationVersionUsingGET**
> ApplicationVersionDto GetApplicationVersionUsingGET(ctx, releaseId, applicationVersionId, optional)
Retrieves an application version

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseId** | **int64**| releaseId | 
  **applicationVersionId** | **int64**| applicationVersionId | 
 **optional** | ***GetApplicationVersionUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetApplicationVersionUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **embed** | [**optional.Interface of []string**](string.md)| embedFields | 
 **importContentItems** | **optional.Bool**| isImportContentItems | 
 **applicationVersionCommit** | **optional.Bool**| isAppVersionCommit | 

### Return type

[**ApplicationVersionDto**](ApplicationVersionDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApplicationVersionsOfApplicationUsingGET**
> ListHolderDtoApplicationVersionDto GetApplicationVersionsOfApplicationUsingGET(ctx, applicationId, optional)
Retrieves application version of application

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int64**| applicationId | 
 **optional** | ***GetApplicationVersionsOfApplicationUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetApplicationVersionsOfApplicationUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sortField** | **optional.String**| sortField | 
 **sortDirection** | **optional.String**| sortDirection | 

### Return type

[**ListHolderDtoApplicationVersionDto**](ListHolderDto«ApplicationVersionDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApplicationVersionsUsingGET**
> ListHolderDtoApplicationVersionDto GetApplicationVersionsUsingGET(ctx, releaseId, optional)
Retrieves all application versions under the release

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseId** | **int64**| releaseId | 
 **optional** | ***GetApplicationVersionsUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetApplicationVersionsUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **embed** | [**optional.Interface of []string**](string.md)| embedFields | 
 **importContentItems** | **optional.Bool**| isImportContentItems | 
 **applicationVersionCommit** | **optional.Bool**| isAppVersionCommit | 

### Return type

[**ListHolderDtoApplicationVersionDto**](ListHolderDto«ApplicationVersionDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAvailableDependenciesUsingGET**
> PagedResultDtoApplicationVersionDto GetAvailableDependenciesUsingGET(ctx, applicationVersionId, optional)
Retrieves available application version dependencies

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationVersionId** | **int64**| excludeAppVersionId | 
 **optional** | ***GetAvailableDependenciesUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAvailableDependenciesUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **optional.String**| freeTextFilter | 
 **pageSize** | **optional.Int32**| pageSize | 
 **pageNumber** | **optional.Int32**| pageNumber | 

### Return type

[**PagedResultDtoApplicationVersionDto**](PagedResultDto«ApplicationVersionDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnAssignApplicationVersionUsingDELETE**
> UnAssignApplicationVersionUsingDELETE(ctx, applicationVersionId, releaseId)
Un assigns an application version from a release (will delete in case it has no references anymore)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationVersionId** | **int64**| applicationVersionId | 
  **releaseId** | **int64**| releaseId | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateApplicationVersionUsingPUT**
> ApplicationVersionDto UpdateApplicationVersionUsingPUT(ctx, applicationVersionId, releaseId, applicationVersionDto)
Updates an application version

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationVersionId** | **int64**| id | 
  **releaseId** | **int64**| releaseId | 
  **applicationVersionDto** | [**ApplicationVersionDto**](ApplicationVersionDto.md)| applicationVersionDto | 

### Return type

[**ApplicationVersionDto**](ApplicationVersionDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

