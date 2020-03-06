# \GitHubIntegrationControllerApi

All URIs are relative to *https://cddirector.io/cdd/design/0000/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGitHubNotificationUsingPOST**](GitHubIntegrationControllerApi.md#GetGitHubNotificationUsingPOST) | **Post** /applications/{applicationName}/application-versions/{applicationVersionName} | getGitHubNotification


# **GetGitHubNotificationUsingPOST**
> BuildNotificationResponseDto GetGitHubNotificationUsingPOST(ctx, gitHubCommitDto, applicationName, applicationVersionName)
getGitHubNotification

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gitHubCommitDto** | [**GitHubCommitDto**](GitHubCommitDto.md)| gitHubCommitDto | 
  **applicationName** | **string**| applicationName | 
  **applicationVersionName** | **string**| applicationVersionName | 

### Return type

[**BuildNotificationResponseDto**](BuildNotificationResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

