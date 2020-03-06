# \TrackApi

All URIs are relative to *https://cddirector.io/cdd/design/0000/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignReleaseToTrackUsingPOST**](TrackApi.md#AssignReleaseToTrackUsingPOST) | **Post** /tracks/{trackId}/releases | Associate release to track
[**ChangeReleaseTrackOrderUsingPATCH**](TrackApi.md#ChangeReleaseTrackOrderUsingPATCH) | **Patch** /tracks/{trackId}/releases/{releaseId}/transitions | Changes the release order within a track
[**CreateTrackUsingPOST**](TrackApi.md#CreateTrackUsingPOST) | **Post** /tracks | Creates track
[**DeleteTrackUsingDELETE**](TrackApi.md#DeleteTrackUsingDELETE) | **Delete** /tracks/{trackId} | Deletes a track
[**GetAllTracksUsingGET**](TrackApi.md#GetAllTracksUsingGET) | **Get** /tracks | Retrieves all tracks
[**GetReleaseTrackAssociationUsingGET**](TrackApi.md#GetReleaseTrackAssociationUsingGET) | **Get** /tracks/{trackId}/releases/{releaseId} | Get single release in track
[**GetReleaseTrackAssociationsUsingGET**](TrackApi.md#GetReleaseTrackAssociationsUsingGET) | **Get** /tracks/{trackId}/releases | Get releases in track
[**GetTrackUsingGET**](TrackApi.md#GetTrackUsingGET) | **Get** /tracks/{trackId} | Retrieves a track
[**PatchAllReleaseInTrackUsingPATCH**](TrackApi.md#PatchAllReleaseInTrackUsingPATCH) | **Patch** /tracks/{trackId}/releases | Patch releases in track
[**PatchReleaseInTrackUsingPATCH**](TrackApi.md#PatchReleaseInTrackUsingPATCH) | **Patch** /tracks/{trackId}/releases/{releaseId} | Patch release in track
[**UnAssignReleaseFromTrackUsingDELETE**](TrackApi.md#UnAssignReleaseFromTrackUsingDELETE) | **Delete** /tracks/{trackId}/releases/{releaseId} | Disassociate release to track
[**UpdateReleaseTrackAssignmentUsingPUT**](TrackApi.md#UpdateReleaseTrackAssignmentUsingPUT) | **Put** /tracks/{trackId}/releases/{releaseId} | Update release to track assignment properties
[**UpdateTrackUsingPUT**](TrackApi.md#UpdateTrackUsingPUT) | **Put** /tracks/{trackId} | Updates a track


# **AssignReleaseToTrackUsingPOST**
> ReleaseInTrackDto AssignReleaseToTrackUsingPOST(ctx, trackId, releaseInTrackDto)
Associate release to track

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **trackId** | **int64**| trackId | 
  **releaseInTrackDto** | [**ReleaseInTrackDto**](ReleaseInTrackDto.md)| releaseInTrackDto | 

### Return type

[**ReleaseInTrackDto**](ReleaseInTrackDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChangeReleaseTrackOrderUsingPATCH**
> ReleaseInTrackDto ChangeReleaseTrackOrderUsingPATCH(ctx, trackId, releaseId, releaseInTrackDto)
Changes the release order within a track

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **trackId** | **int64**| trackId | 
  **releaseId** | **int64**| releaseInTrackId | 
  **releaseInTrackDto** | [**ReleaseInTrackDto**](ReleaseInTrackDto.md)| releaseInTrackDto | 

### Return type

[**ReleaseInTrackDto**](ReleaseInTrackDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateTrackUsingPOST**
> TrackDto CreateTrackUsingPOST(ctx, trackDto)
Creates track

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **trackDto** | [**TrackDto**](TrackDto.md)| trackDto | 

### Return type

[**TrackDto**](TrackDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTrackUsingDELETE**
> DeleteTrackUsingDELETE(ctx, trackId)
Deletes a track

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **trackId** | **int64**| trackId | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllTracksUsingGET**
> PagedResultDto GetAllTracksUsingGET(ctx, optional)
Retrieves all tracks

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllTracksUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAllTracksUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **optional.String**| freeTextFilter | 
 **filterFields** | [**optional.Interface of []string**](string.md)| filterFields | 
 **pageSize** | **optional.Int32**| pageSize | 
 **pageNumber** | **optional.Int32**| pageNumber | 
 **sortDirection** | **optional.String**| sortDirection | 
 **sortField** | **optional.String**| sortField | 
 **status** | [**optional.Interface of []string**](string.md)| filterByStatus | 
 **embed** | [**optional.Interface of []string**](string.md)| embedFields | 

### Return type

[**PagedResultDto**](PagedResultDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReleaseTrackAssociationUsingGET**
> ReleaseInTrackDto GetReleaseTrackAssociationUsingGET(ctx, trackId, releaseId)
Get single release in track

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **trackId** | **int64**| trackId | 
  **releaseId** | **int64**| releaseId | 

### Return type

[**ReleaseInTrackDto**](ReleaseInTrackDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReleaseTrackAssociationsUsingGET**
> ListHolderDtoReleaseInTrackDto GetReleaseTrackAssociationsUsingGET(ctx, trackId)
Get releases in track

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **trackId** | **int64**| trackId | 

### Return type

[**ListHolderDtoReleaseInTrackDto**](ListHolderDto«ReleaseInTrackDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTrackUsingGET**
> TrackDto GetTrackUsingGET(ctx, trackId, optional)
Retrieves a track

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **trackId** | **int64**| trackId | 
 **optional** | ***GetTrackUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetTrackUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **embed** | [**optional.Interface of []string**](string.md)| embedFields | 

### Return type

[**TrackDto**](TrackDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchAllReleaseInTrackUsingPATCH**
> ListHolderDtoReleaseInTrackDto PatchAllReleaseInTrackUsingPATCH(ctx, trackId, releasesInTrackDto)
Patch releases in track

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **trackId** | **int64**| trackId | 
  **releasesInTrackDto** | [**[]ReleaseInTrackDto**](ReleaseInTrackDto.md)| releasesInTrackDto | 

### Return type

[**ListHolderDtoReleaseInTrackDto**](ListHolderDto«ReleaseInTrackDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchReleaseInTrackUsingPATCH**
> ReleaseInTrackDto PatchReleaseInTrackUsingPATCH(ctx, trackId, releaseId, releaseInTrackDto)
Patch release in track

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **trackId** | **int64**| trackId | 
  **releaseId** | **int64**| releaseId | 
  **releaseInTrackDto** | [**ReleaseInTrackDto**](ReleaseInTrackDto.md)| releaseInTrackDto | 

### Return type

[**ReleaseInTrackDto**](ReleaseInTrackDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnAssignReleaseFromTrackUsingDELETE**
> UnAssignReleaseFromTrackUsingDELETE(ctx, trackId, releaseId)
Disassociate release to track

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **trackId** | **int64**| trackId | 
  **releaseId** | **int64**| releaseInTrackId | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateReleaseTrackAssignmentUsingPUT**
> ReleaseInTrackDto UpdateReleaseTrackAssignmentUsingPUT(ctx, trackId, releaseId, releaseInTrackDto)
Update release to track assignment properties

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **trackId** | **int64**| trackId | 
  **releaseId** | **int64**| releaseInTrackId | 
  **releaseInTrackDto** | [**ReleaseInTrackDto**](ReleaseInTrackDto.md)| releaseInTrackDto | 

### Return type

[**ReleaseInTrackDto**](ReleaseInTrackDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTrackUsingPUT**
> TrackDto UpdateTrackUsingPUT(ctx, trackId, trackDto)
Updates a track

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **trackId** | **int64**| trackId | 
  **trackDto** | [**TrackDto**](TrackDto.md)| trackDto | 

### Return type

[**TrackDto**](TrackDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

