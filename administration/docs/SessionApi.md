# \SessionApi

All URIs are relative to *https://cddirector.io/cdd/administration/0000/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSessionUsingGET**](SessionApi.md#GetSessionUsingGET) | **Get** /sessions/current | Get session
[**PatchTenantUsingPATCH**](SessionApi.md#PatchTenantUsingPATCH) | **Patch** /sessions/current | Patch session


# **GetSessionUsingGET**
> SessionDto GetSessionUsingGET(ctx, )
Get session

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SessionDto**](SessionDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchTenantUsingPATCH**
> SessionDto PatchTenantUsingPATCH(ctx, sessionDto)
Patch session

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sessionDto** | [**SessionDto**](SessionDto.md)| sessionDto | 

### Return type

[**SessionDto**](SessionDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

