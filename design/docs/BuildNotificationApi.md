# \BuildNotificationApi

All URIs are relative to *https://cddirector.io/cdd/design/0000/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OnBuildNotificationUsingPOST**](BuildNotificationApi.md#OnBuildNotificationUsingPOST) | **Post** /applications/application-versions/application-version-builds | Send build notification


# **OnBuildNotificationUsingPOST**
> BuildNotificationResponseDto OnBuildNotificationUsingPOST(ctx, buildNotificationDto)
Send build notification

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **buildNotificationDto** | [**BuildNotificationDto**](BuildNotificationDto.md)| buildNotificationDto | 

### Return type

[**BuildNotificationResponseDto**](BuildNotificationResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

