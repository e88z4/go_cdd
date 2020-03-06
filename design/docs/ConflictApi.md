# \ConflictApi

All URIs are relative to *https://cddirector.io/cdd/design/0000/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApplicationVersionConflictsUsingGET**](ConflictApi.md#GetApplicationVersionConflictsUsingGET) | **Get** /releases/{releaseId}/application-version-conflicts | Retrieves application version conflicts
[**GetConflictsUsingGET**](ConflictApi.md#GetConflictsUsingGET) | **Get** /releases/{releaseId}/conflicts | Retrieves all conflicts


# **GetApplicationVersionConflictsUsingGET**
> ReleaseApplicationVersionConflictsDto GetApplicationVersionConflictsUsingGET(ctx, releaseId)
Retrieves application version conflicts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseId** | **int64**| releaseId | 

### Return type

[**ReleaseApplicationVersionConflictsDto**](ReleaseApplicationVersionConflictsDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConflictsUsingGET**
> ListHolderDtoConflictDto GetConflictsUsingGET(ctx, startDate, endDate, environments, optional)
Retrieves all conflicts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **startDate** | **int64**| startDate | 
  **endDate** | **int64**| endDate | 
  **environments** | [**[]int64**](int64.md)| environmentIds | 
 **optional** | ***GetConflictsUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetConflictsUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **phaseId** | **optional.Int64**| phaseId | 

### Return type

[**ListHolderDtoConflictDto**](ListHolderDto«ConflictDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

