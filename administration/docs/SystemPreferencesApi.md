# \SystemPreferencesApi

All URIs are relative to *https://cddirector.io/cdd/administration/0000/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSupportedDisplayNamesUsingGET**](SystemPreferencesApi.md#GetSupportedDisplayNamesUsingGET) | **Get** /product/settings/system-preferences/display-names | Get System Preference Display Names
[**GetSystemPreferencesDtoUsingGET**](SystemPreferencesApi.md#GetSystemPreferencesDtoUsingGET) | **Get** /product/settings/system-preferences | Get System Preferences
[**TestMicroservicesConnectivityUsingPOST**](SystemPreferencesApi.md#TestMicroservicesConnectivityUsingPOST) | **Post** /adaptive-testing/connectivity-tests | Connectivity Test for Adaptive Testing Services
[**UpdateSystemPreferencesDtoUsingPATCH**](SystemPreferencesApi.md#UpdateSystemPreferencesDtoUsingPATCH) | **Patch** /product/settings/system-preferences | Set System Preferences


# **GetSupportedDisplayNamesUsingGET**
> []string GetSupportedDisplayNamesUsingGET(ctx, )
Get System Preference Display Names

### Required Parameters
This endpoint does not need any parameter.

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSystemPreferencesDtoUsingGET**
> SystemPreferencesDto GetSystemPreferencesDtoUsingGET(ctx, optional)
Get System Preferences

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetSystemPreferencesDtoUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetSystemPreferencesDtoUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **systemPreference** | [**optional.Interface of []string**](string.md)| systemPreference | 

### Return type

[**SystemPreferencesDto**](SystemPreferencesDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TestMicroservicesConnectivityUsingPOST**
> ConnectableEntityDto TestMicroservicesConnectivityUsingPOST(ctx, servicesDto)
Connectivity Test for Adaptive Testing Services

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **servicesDto** | [**MicroservicesDto**](MicroservicesDto.md)| servicesDto | 

### Return type

[**ConnectableEntityDto**](ConnectableEntityDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSystemPreferencesDtoUsingPATCH**
> SystemPreferencesDto UpdateSystemPreferencesDtoUsingPATCH(ctx, preferenceDtos)
Set System Preferences

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **preferenceDtos** | [**[]ShallowSystemPreferenceDto**](ShallowSystemPreferenceDto.md)| preferenceDtos | 

### Return type

[**SystemPreferencesDto**](SystemPreferencesDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

