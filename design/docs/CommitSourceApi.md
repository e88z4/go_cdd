# \CommitSourceApi

All URIs are relative to *https://cddirector.io/cdd/design/0000/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCommitSourceUsingPOST**](CommitSourceApi.md#CreateCommitSourceUsingPOST) | **Post** /applications/{applicationId}/application-versions/{applicationVersionId}/commit-sources | Create a commit source
[**DeleteCommitSourceUsingDELETE**](CommitSourceApi.md#DeleteCommitSourceUsingDELETE) | **Delete** /applications/{applicationId}/application-versions/{applicationVersionId}/commit-sources/{commitSourceId} | Delete a commit source
[**GetCommitSourceUsingGET**](CommitSourceApi.md#GetCommitSourceUsingGET) | **Get** /applications/{applicationId}/application-versions/{applicationVersionId}/commit-sources/{commitSourceId} | Retrieve a commit source
[**GetCommitSourcesUsingGET**](CommitSourceApi.md#GetCommitSourcesUsingGET) | **Get** /applications/{applicationId}/application-versions/{applicationVersionId}/commit-sources | Retrieve all commit sources
[**GetCommitsUsingGET**](CommitSourceApi.md#GetCommitsUsingGET) | **Get** /applications/{applicationId}/application-versions/{applicationVersionId}/application-version-commits/{applicationVersionCommitId} | Retrieve an application version commit
[**GetCommitsUsingGET1**](CommitSourceApi.md#GetCommitsUsingGET1) | **Get** /applications/{applicationId}/application-versions/{applicationVersionId}/application-version-commits | Retrieve all application version commits
[**UpdateCommitSourceUsingPUT**](CommitSourceApi.md#UpdateCommitSourceUsingPUT) | **Put** /applications/{applicationId}/application-versions/{applicationVersionId}/commit-sources/{commitSourceId} | Update a commit source


# **CreateCommitSourceUsingPOST**
> CommitSourceDto CreateCommitSourceUsingPOST(ctx, applicationId, applicationVersionId, commitSourceDto)
Create a commit source

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int64**| applicationId | 
  **applicationVersionId** | **int64**| applicationVersionId | 
  **commitSourceDto** | [**CommitSourceDto**](CommitSourceDto.md)| commitSourceDto | 

### Return type

[**CommitSourceDto**](CommitSourceDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCommitSourceUsingDELETE**
> CommitSourceDto DeleteCommitSourceUsingDELETE(ctx, applicationId, applicationVersionId, commitSourceId)
Delete a commit source

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int64**| applicationId | 
  **applicationVersionId** | **int64**| applicationVersionId | 
  **commitSourceId** | **int64**| commitSourceId | 

### Return type

[**CommitSourceDto**](CommitSourceDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCommitSourceUsingGET**
> CommitSourceDto GetCommitSourceUsingGET(ctx, applicationId, applicationVersionId, commitSourceId)
Retrieve a commit source

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int64**| applicationId | 
  **applicationVersionId** | **int64**| applicationVersionId | 
  **commitSourceId** | **int64**| commitSourceId | 

### Return type

[**CommitSourceDto**](CommitSourceDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCommitSourcesUsingGET**
> ListHolderDtoCommitSourceDto GetCommitSourcesUsingGET(ctx, applicationId, applicationVersionId)
Retrieve all commit sources

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int64**| applicationId | 
  **applicationVersionId** | **int64**| applicationVersionId | 

### Return type

[**ListHolderDtoCommitSourceDto**](ListHolderDto«CommitSourceDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCommitsUsingGET**
> ApplicationVersionCommitDto GetCommitsUsingGET(ctx, applicationId, applicationVersionId, applicationVersionCommitId)
Retrieve an application version commit

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int64**| applicationId | 
  **applicationVersionId** | **int64**| applicationVersionId | 
  **applicationVersionCommitId** | **int64**| applicationVersionCommitId | 

### Return type

[**ApplicationVersionCommitDto**](ApplicationVersionCommitDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCommitsUsingGET1**
> ListHolderDtoApplicationVersionCommitDto GetCommitsUsingGET1(ctx, applicationId, applicationVersionId)
Retrieve all application version commits

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int64**| applicationId | 
  **applicationVersionId** | **int64**| applicationVersionId | 

### Return type

[**ListHolderDtoApplicationVersionCommitDto**](ListHolderDto«ApplicationVersionCommitDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCommitSourceUsingPUT**
> CommitSourceDto UpdateCommitSourceUsingPUT(ctx, applicationId, applicationVersionId, commitSourceId, commitSourceDto)
Update a commit source

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int64**| applicationId | 
  **applicationVersionId** | **int64**| applicationVersionId | 
  **commitSourceId** | **int64**| commitSourceId | 
  **commitSourceDto** | [**CommitSourceDto**](CommitSourceDto.md)| commitSourceDto | 

### Return type

[**CommitSourceDto**](CommitSourceDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

