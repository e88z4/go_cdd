# \PhaseApi

All URIs are relative to *https://cddirector.io/cdd/design/0000/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangePhaseTransitionsUsingPATCH**](PhaseApi.md#ChangePhaseTransitionsUsingPATCH) | **Patch** /releases/{releaseId}/phases/{phaseId}/transitions | Change phase transition
[**CreatePhaseUsingPOST**](PhaseApi.md#CreatePhaseUsingPOST) | **Post** /releases/{releaseId}/phases | Creates or duplicates a phase
[**DeletePhaseUsingDELETE**](PhaseApi.md#DeletePhaseUsingDELETE) | **Delete** /releases/{releaseId}/phases/{phaseId} | Deletes a phase
[**GetAllAppVersionsForPhaseUsingGET**](PhaseApi.md#GetAllAppVersionsForPhaseUsingGET) | **Get** /releases/{releaseId}/phases/{phaseId}/application-versions | Retrieves all application versions of a phase
[**GetAllPhasesOfReleaseUsingGET**](PhaseApi.md#GetAllPhasesOfReleaseUsingGET) | **Get** /releases/{releaseId}/phases | Retrieves all phases
[**GetPhaseUsingGET**](PhaseApi.md#GetPhaseUsingGET) | **Get** /releases/{releaseId}/phases/{phaseId} | Retrieves a phase
[**PatchPhaseUsingPATCH**](PhaseApi.md#PatchPhaseUsingPATCH) | **Patch** /releases/{releaseId}/phases/{phaseId} | Changes a phase
[**UpdatePhaseUsingPUT**](PhaseApi.md#UpdatePhaseUsingPUT) | **Put** /releases/{releaseId}/phases/{phaseId} | Updates a phase


# **ChangePhaseTransitionsUsingPATCH**
> PhaseDto ChangePhaseTransitionsUsingPATCH(ctx, releaseId, phaseId, phaseDto)
Change phase transition

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseId** | **int64**| releaseId | 
  **phaseId** | **int64**| phaseId | 
  **phaseDto** | [**PhaseDto**](PhaseDto.md)| phaseDto | 

### Return type

[**PhaseDto**](PhaseDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreatePhaseUsingPOST**
> PhaseDto CreatePhaseUsingPOST(ctx, releaseId, phaseDto)
Creates or duplicates a phase

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseId** | **int64**| containingReleaseId | 
  **phaseDto** | [**PhaseDto**](PhaseDto.md)| phaseDto | 

### Return type

[**PhaseDto**](PhaseDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePhaseUsingDELETE**
> IdentifiableDto DeletePhaseUsingDELETE(ctx, releaseId, phaseId)
Deletes a phase

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseId** | **int64**| containingReleaseId | 
  **phaseId** | **int64**| phaseId | 

### Return type

[**IdentifiableDto**](IdentifiableDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllAppVersionsForPhaseUsingGET**
> ListHolderDto GetAllAppVersionsForPhaseUsingGET(ctx, releaseId, phaseId)
Retrieves all application versions of a phase

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseId** | **int64**| releaseId | 
  **phaseId** | **int64**| phaseId | 

### Return type

[**ListHolderDto**](ListHolderDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllPhasesOfReleaseUsingGET**
> ListHolderDtoPhaseDto GetAllPhasesOfReleaseUsingGET(ctx, releaseId)
Retrieves all phases

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseId** | **int64**| releaseId | 

### Return type

[**ListHolderDtoPhaseDto**](ListHolderDto«PhaseDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPhaseUsingGET**
> PhaseDto GetPhaseUsingGET(ctx, releaseId, phaseId)
Retrieves a phase

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseId** | **int64**| containingReleaseId | 
  **phaseId** | **int64**| phaseId | 

### Return type

[**PhaseDto**](PhaseDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchPhaseUsingPATCH**
> PhaseDto PatchPhaseUsingPATCH(ctx, releaseId, phaseId, phaseDto)
Changes a phase

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseId** | **int64**| containingReleaseId | 
  **phaseId** | **int64**| phaseId | 
  **phaseDto** | [**PhaseDto**](PhaseDto.md)| phaseDto | 

### Return type

[**PhaseDto**](PhaseDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePhaseUsingPUT**
> PhaseDto UpdatePhaseUsingPUT(ctx, releaseId, phaseId, phaseDto)
Updates a phase

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseId** | **int64**| containingReleaseId | 
  **phaseId** | **int64**| phaseId | 
  **phaseDto** | [**PhaseDto**](PhaseDto.md)| phaseDto | 

### Return type

[**PhaseDto**](PhaseDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

