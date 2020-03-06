# \DbSettingsApi

All URIs are relative to *https://cddirector.io/cdd/administration/0000/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDbSettingsUsingGET**](DbSettingsApi.md#GetDbSettingsUsingGET) | **Get** /product/settings/database-server-settings | Get DB settings
[**UpdateDbSettingsUsingPATCH**](DbSettingsApi.md#UpdateDbSettingsUsingPATCH) | **Patch** /product/settings/database-server-settings | Update DB settings


# **GetDbSettingsUsingGET**
> DbPropertiesSettingsDto GetDbSettingsUsingGET(ctx, )
Get DB settings

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**DbPropertiesSettingsDto**](DbPropertiesSettingsDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDbSettingsUsingPATCH**
> DbPropertiesSettingsDto UpdateDbSettingsUsingPATCH(ctx, dbSettingsDto)
Update DB settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dbSettingsDto** | [**DbPropertiesSettingsDto**](DbPropertiesSettingsDto.md)| dbSettingsDto | 

### Return type

[**DbPropertiesSettingsDto**](DbPropertiesSettingsDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

