# \ReleaseScoringApi

All URIs are relative to *https://cddirector.io/cdd/design/0000/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetReleaseScoreUsingGET**](ReleaseScoringApi.md#GetReleaseScoreUsingGET) | **Get** /releases/{releaseId}/score | Retrieves a release score


# **GetReleaseScoreUsingGET**
> ReleaseScoreDto GetReleaseScoreUsingGET(ctx, releaseId)
Retrieves a release score

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseId** | **int64**| releaseId | 

### Return type

[**ReleaseScoreDto**](ReleaseScoreDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

