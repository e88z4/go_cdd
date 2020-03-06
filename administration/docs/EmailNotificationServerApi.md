# \EmailNotificationServerApi

All URIs are relative to *https://cddirector.io/cdd/administration/0000/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEmailNotificationServerUsingPOST**](EmailNotificationServerApi.md#CreateEmailNotificationServerUsingPOST) | **Post** /email-notification-servers | Creates a new email notification server definition
[**DeleteEmailNotificationServerUsingDELETE**](EmailNotificationServerApi.md#DeleteEmailNotificationServerUsingDELETE) | **Delete** /email-notification-servers/{emailNotificationServerId} | Deletes an email notification definition
[**DeleteEmailNotificationServersUsingDELETE**](EmailNotificationServerApi.md#DeleteEmailNotificationServersUsingDELETE) | **Delete** /email-notification-servers | Deletes all email notification definitions
[**GetEmailNotificationServerUsingGET**](EmailNotificationServerApi.md#GetEmailNotificationServerUsingGET) | **Get** /email-notification-servers/{emailNotificationServerId} | Retrieves an email notification server definition
[**GetEmailNotificationServersUsingGET**](EmailNotificationServerApi.md#GetEmailNotificationServersUsingGET) | **Get** /email-notification-servers | Retrieves all email notification server definitions
[**PatchEmailNotificationServerUsingPATCH**](EmailNotificationServerApi.md#PatchEmailNotificationServerUsingPATCH) | **Patch** /email-notification-servers/{emailNotificationServerId} | Patches an email notification server definition
[**TestEmailNotificationServerConnectivityUsingPOST**](EmailNotificationServerApi.md#TestEmailNotificationServerConnectivityUsingPOST) | **Post** /email-notification-servers/connectivity-tests | Connectivity Test for email notification server
[**UpdateEmailNotificationServerUsingPUT**](EmailNotificationServerApi.md#UpdateEmailNotificationServerUsingPUT) | **Put** /email-notification-servers/{emailNotificationServerId} | Updates an email notification server definition


# **CreateEmailNotificationServerUsingPOST**
> EmailNotificationServerDto CreateEmailNotificationServerUsingPOST(ctx, emailNotificationServerDto)
Creates a new email notification server definition

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **emailNotificationServerDto** | [**EmailNotificationServerDto**](EmailNotificationServerDto.md)| emailNotificationServerDto | 

### Return type

[**EmailNotificationServerDto**](EmailNotificationServerDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEmailNotificationServerUsingDELETE**
> DeleteEmailNotificationServerUsingDELETE(ctx, emailNotificationServerId)
Deletes an email notification definition

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **emailNotificationServerId** | **int64**| emailNotificationServerId | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEmailNotificationServersUsingDELETE**
> DeleteEmailNotificationServersUsingDELETE(ctx, )
Deletes all email notification definitions

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEmailNotificationServerUsingGET**
> EmailNotificationServerDto GetEmailNotificationServerUsingGET(ctx, emailNotificationServerId)
Retrieves an email notification server definition

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **emailNotificationServerId** | **int64**| emailNotificationServerId | 

### Return type

[**EmailNotificationServerDto**](EmailNotificationServerDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEmailNotificationServersUsingGET**
> ListHolderDtoEmailNotificationServerDto GetEmailNotificationServersUsingGET(ctx, )
Retrieves all email notification server definitions

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ListHolderDtoEmailNotificationServerDto**](ListHolderDto«EmailNotificationServerDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchEmailNotificationServerUsingPATCH**
> EmailNotificationServerDto PatchEmailNotificationServerUsingPATCH(ctx, emailNotificationServerId, emailNotificationServerDto)
Patches an email notification server definition

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **emailNotificationServerId** | **int64**| emailNotificationServerId | 
  **emailNotificationServerDto** | [**EmailNotificationServerDto**](EmailNotificationServerDto.md)| emailNotificationServerDto | 

### Return type

[**EmailNotificationServerDto**](EmailNotificationServerDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TestEmailNotificationServerConnectivityUsingPOST**
> EmailNotificationServerDto TestEmailNotificationServerConnectivityUsingPOST(ctx, emailNotificationServerDto)
Connectivity Test for email notification server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **emailNotificationServerDto** | [**EmailNotificationServerDto**](EmailNotificationServerDto.md)| emailNotificationServerDto | 

### Return type

[**EmailNotificationServerDto**](EmailNotificationServerDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEmailNotificationServerUsingPUT**
> EmailNotificationServerDto UpdateEmailNotificationServerUsingPUT(ctx, emailNotificationServerId, emailNotificationServerDto)
Updates an email notification server definition

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **emailNotificationServerId** | **int64**| emailNotificationServerId | 
  **emailNotificationServerDto** | [**EmailNotificationServerDto**](EmailNotificationServerDto.md)| emailNotificationServerDto | 

### Return type

[**EmailNotificationServerDto**](EmailNotificationServerDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

