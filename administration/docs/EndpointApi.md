# \EndpointApi

All URIs are relative to *https://cddirector.io/cdd/administration/0000/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEndpointUsingPOST**](EndpointApi.md#CreateEndpointUsingPOST) | **Post** /endpoint-templates/{endpointTemplateId}/endpoints | Creates an endPoint
[**DeleteEndpointUsingDELETE**](EndpointApi.md#DeleteEndpointUsingDELETE) | **Delete** /endpoints/{endpointId} | Deletes an endPoint
[**GetAllEndpointTemplatesUsingGET**](EndpointApi.md#GetAllEndpointTemplatesUsingGET) | **Get** /endpoint-templates | Retrieves all endPoint templates
[**GetAllEndpointsUsingGET**](EndpointApi.md#GetAllEndpointsUsingGET) | **Get** /endpoints | Retrieves all endpoints
[**GetEndpointTemplateUsingGET**](EndpointApi.md#GetEndpointTemplateUsingGET) | **Get** /endpoint-templates/{endpointTemplateId} | Retrieves an endPoint template
[**GetEndpointUsingGET**](EndpointApi.md#GetEndpointUsingGET) | **Get** /endpoints/{endpointId} | Retrieves an endPoint
[**TestEndpointConnectivityUsingPOST**](EndpointApi.md#TestEndpointConnectivityUsingPOST) | **Post** /endpoints/connectivity-tests | Connectivity Test for a Plugin Endpoint
[**UpdateEndpointUsingPUT**](EndpointApi.md#UpdateEndpointUsingPUT) | **Put** /endpoints/{endpointId} | updates an endPoint


# **CreateEndpointUsingPOST**
> EndpointDto CreateEndpointUsingPOST(ctx, endpointTemplateId, endpointDto)
Creates an endPoint

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **endpointTemplateId** | **int64**| endpointTemplateId | 
  **endpointDto** | [**EndpointDto**](EndpointDto.md)| endpointDto | 

### Return type

[**EndpointDto**](EndpointDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEndpointUsingDELETE**
> DeleteEndpointUsingDELETE(ctx, endpointId)
Deletes an endPoint

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **endpointId** | **int64**| endpointId | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllEndpointTemplatesUsingGET**
> ListHolderDtoEndpointTemplateDto GetAllEndpointTemplatesUsingGET(ctx, )
Retrieves all endPoint templates

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ListHolderDtoEndpointTemplateDto**](ListHolderDto«EndpointTemplateDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllEndpointsUsingGET**
> ListHolderDtoEndpointDto GetAllEndpointsUsingGET(ctx, optional)
Retrieves all endpoints

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllEndpointsUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAllEndpointsUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pluginTemplateId** | **optional.Int64**| pluginServiceTemplateId | 
 **type_** | **optional.String**| pluginServiceTypeStr | 
 **application** | [**optional.Interface of []int64**](int64.md)| applicationIds | 
 **environment** | [**optional.Interface of []int64**](int64.md)| environmentIds | 
 **currentUser** | **optional.Bool**| userMustAllowedAllApplications | 
 **filter** | **optional.String**| freeTextFilter | 

### Return type

[**ListHolderDtoEndpointDto**](ListHolderDto«EndpointDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEndpointTemplateUsingGET**
> EndpointTemplateDto GetEndpointTemplateUsingGET(ctx, endpointTemplateId)
Retrieves an endPoint template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **endpointTemplateId** | **int64**| endpointTemplateId | 

### Return type

[**EndpointTemplateDto**](EndpointTemplateDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEndpointUsingGET**
> EndpointDto GetEndpointUsingGET(ctx, endpointId)
Retrieves an endPoint

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **endpointId** | **int64**| endpointId | 

### Return type

[**EndpointDto**](EndpointDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TestEndpointConnectivityUsingPOST**
> EndpointDto TestEndpointConnectivityUsingPOST(ctx, endpointDto)
Connectivity Test for a Plugin Endpoint

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **endpointDto** | [**EndpointDto**](EndpointDto.md)| endpointDto | 

### Return type

[**EndpointDto**](EndpointDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEndpointUsingPUT**
> EndpointDto UpdateEndpointUsingPUT(ctx, endpointId, endpointDto)
updates an endPoint

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **endpointId** | **int64**| endpointId | 
  **endpointDto** | [**EndpointDto**](EndpointDto.md)| endpointDto | 

### Return type

[**EndpointDto**](EndpointDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

