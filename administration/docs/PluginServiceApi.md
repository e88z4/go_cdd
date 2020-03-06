# \PluginServiceApi

All URIs are relative to *https://cddirector.io/cdd/administration/0000/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePluginUsingDELETE**](PluginServiceApi.md#DeletePluginUsingDELETE) | **Delete** /plugins/{pluginId} | Deletes a plugin
[**EditPluginManifestUrlUsingPUT**](PluginServiceApi.md#EditPluginManifestUrlUsingPUT) | **Put** /plugins/{pluginId} | Edit a plugin manifest URL
[**GetAllPluginServiceTemplatesUsingGET**](PluginServiceApi.md#GetAllPluginServiceTemplatesUsingGET) | **Get** /plugin-service-templates | Retrieves all plugin service templates
[**GetAllPluginsUsingGET**](PluginServiceApi.md#GetAllPluginsUsingGET) | **Get** /plugins | Retrieves all plugins
[**GetPluginIconUsingGET**](PluginServiceApi.md#GetPluginIconUsingGET) | **Get** /plugins/{pluginId}/icon | Retrieves a plugin&#39;s icon
[**GetPluginServiceTemplateParameterUsingPOST**](PluginServiceApi.md#GetPluginServiceTemplateParameterUsingPOST) | **Post** /plugin-service-templates/{pluginServiceTemplateId}/plugin-service-template-parameters/{pluginServiceTemplateParameterId} | Retrieves a plugin service template parameter dynamic values
[**GetPluginServiceTemplateUsingGET**](PluginServiceApi.md#GetPluginServiceTemplateUsingGET) | **Get** /plugin-service-templates/{pluginServiceTemplateId} | Retrieves a plugin service template
[**GetPluginServiceUsingGET**](PluginServiceApi.md#GetPluginServiceUsingGET) | **Get** /plugin-services/{pluginServiceId} | Retrieves a plugin service
[**GetPluginUsingGET**](PluginServiceApi.md#GetPluginUsingGET) | **Get** /plugins/{pluginId} | Retrieves a plugin
[**RegisterPluginUsingPOST**](PluginServiceApi.md#RegisterPluginUsingPOST) | **Post** /plugins | Registers a plugin by URL
[**SyncPluginUsingPATCH**](PluginServiceApi.md#SyncPluginUsingPATCH) | **Patch** /plugins/{pluginId}/sync | Sync a plugin


# **DeletePluginUsingDELETE**
> DeletePluginUsingDELETE(ctx, pluginId)
Deletes a plugin

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pluginId** | **int64**| pluginId | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditPluginManifestUrlUsingPUT**
> PluginDto EditPluginManifestUrlUsingPUT(ctx, pluginId, pluginDto)
Edit a plugin manifest URL

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pluginId** | **int64**| pluginId | 
  **pluginDto** | [**PluginDto**](PluginDto.md)| pluginDto | 

### Return type

[**PluginDto**](PluginDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllPluginServiceTemplatesUsingGET**
> ListHolderDtoPluginServiceTemplateDto GetAllPluginServiceTemplatesUsingGET(ctx, type_)
Retrieves all plugin service templates

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **type_** | **string**| pluginServiceTypeStr | 

### Return type

[**ListHolderDtoPluginServiceTemplateDto**](ListHolderDto«PluginServiceTemplateDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllPluginsUsingGET**
> ListHolderDtoPluginDto GetAllPluginsUsingGET(ctx, )
Retrieves all plugins

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ListHolderDtoPluginDto**](ListHolderDto«PluginDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPluginIconUsingGET**
> GetPluginIconUsingGET(ctx, pluginId)
Retrieves a plugin's icon

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pluginId** | **int64**| pluginId | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPluginServiceTemplateParameterUsingPOST**
> PagedResultDtoPairDto GetPluginServiceTemplateParameterUsingPOST(ctx, pluginServiceTemplateId, pluginServiceTemplateParameterId, dynamicValuesDto, optional)
Retrieves a plugin service template parameter dynamic values

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pluginServiceTemplateId** | **int64**| pluginServiceTemplateId | 
  **pluginServiceTemplateParameterId** | **int64**| pluginServiceTemplateParameterId | 
  **dynamicValuesDto** | [**DynamicValuesInputDto**](DynamicValuesInputDto.md)| dynamicValuesDto | 
 **optional** | ***GetPluginServiceTemplateParameterUsingPOSTOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetPluginServiceTemplateParameterUsingPOSTOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **filter** | **optional.String**| freeTextFilter | 
 **pageSize** | **optional.Int32**| pageSize | 
 **pageNumber** | **optional.Int32**| pageNumber | 

### Return type

[**PagedResultDtoPairDto**](PagedResultDto«PairDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPluginServiceTemplateUsingGET**
> PluginServiceTemplateDto GetPluginServiceTemplateUsingGET(ctx, pluginServiceTemplateId)
Retrieves a plugin service template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pluginServiceTemplateId** | **int64**| pluginServiceTemplateId | 

### Return type

[**PluginServiceTemplateDto**](PluginServiceTemplateDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPluginServiceUsingGET**
> PluginServiceDto GetPluginServiceUsingGET(ctx, pluginServiceId)
Retrieves a plugin service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pluginServiceId** | **int64**| pluginServiceId | 

### Return type

[**PluginServiceDto**](PluginServiceDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPluginUsingGET**
> PluginDto GetPluginUsingGET(ctx, pluginId)
Retrieves a plugin

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pluginId** | **int64**| pluginId | 

### Return type

[**PluginDto**](PluginDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegisterPluginUsingPOST**
> PluginDto RegisterPluginUsingPOST(ctx, pluginUrlDto)
Registers a plugin by URL

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pluginUrlDto** | [**PluginUrlDto**](PluginUrlDto.md)| pluginUrlDto | 

### Return type

[**PluginDto**](PluginDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SyncPluginUsingPATCH**
> PluginDto SyncPluginUsingPATCH(ctx, pluginId)
Sync a plugin

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pluginId** | **int64**| pluginId | 

### Return type

[**PluginDto**](PluginDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

