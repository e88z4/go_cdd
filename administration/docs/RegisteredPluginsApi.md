# \RegisteredPluginsApi

All URIs are relative to *https://cddirector.io/cdd/administration/0000/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAutoRegisterPluginUsingPOST**](RegisteredPluginsApi.md#CreateAutoRegisterPluginUsingPOST) | **Post** /registered-plugins | Create registered plugin
[**DeleteAutoRegisterPluginUsingDELETE**](RegisteredPluginsApi.md#DeleteAutoRegisterPluginUsingDELETE) | **Delete** /registered-plugins/{registeredPluginId} | Deletes a registered plugin
[**GetAutoRegisterPluginsUsingGET**](RegisteredPluginsApi.md#GetAutoRegisterPluginsUsingGET) | **Get** /registered-plugins | Retrieves all registered plugins
[**PatchRegisteredPluginsUsingPATCH**](RegisteredPluginsApi.md#PatchRegisteredPluginsUsingPATCH) | **Patch** /registered-plugins | Patch registered plugins


# **CreateAutoRegisterPluginUsingPOST**
> RegisteredPluginDto CreateAutoRegisterPluginUsingPOST(ctx, registeredPluginDto)
Create registered plugin

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **registeredPluginDto** | [**RegisteredPluginDto**](RegisteredPluginDto.md)| registeredPluginDto | 

### Return type

[**RegisteredPluginDto**](RegisteredPluginDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAutoRegisterPluginUsingDELETE**
> DeleteAutoRegisterPluginUsingDELETE(ctx, registeredPluginId)
Deletes a registered plugin

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **registeredPluginId** | **int64**| autoRegisterPluginId | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAutoRegisterPluginsUsingGET**
> ListHolderDto GetAutoRegisterPluginsUsingGET(ctx, )
Retrieves all registered plugins

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ListHolderDto**](ListHolderDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchRegisteredPluginsUsingPATCH**
> RegisteredPluginDtoList PatchRegisteredPluginsUsingPATCH(ctx, registeredPlugins)
Patch registered plugins

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **registeredPlugins** | [**RegisteredPluginDtoList**](RegisteredPluginDtoList.md)| registeredPlugins | 

### Return type

[**RegisteredPluginDtoList**](RegisteredPluginDtoList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

