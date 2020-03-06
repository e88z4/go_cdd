# \WidgetApi

All URIs are relative to *https://cddirector.io/cdd/reporting/0000/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteWidgetUsingDELETE**](WidgetApi.md#DeleteWidgetUsingDELETE) | **Delete** /dashboard/widgets/{widgetId} | Deletes a widget
[**GetWidgetUsingGET**](WidgetApi.md#GetWidgetUsingGET) | **Get** /dashboard/widgets/{widgetId} | Retrieves a widget
[**GetWidgetsUsingGET**](WidgetApi.md#GetWidgetsUsingGET) | **Get** /dashboard/widgets | Retrieves all widgets
[**HandleUploadUsingPOST**](WidgetApi.md#HandleUploadUsingPOST) | **Post** /dashboard/widgets | Uploads a widget


# **DeleteWidgetUsingDELETE**
> DeleteWidgetUsingDELETE(ctx, widgetId)
Deletes a widget

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **widgetId** | **int64**| widgetId | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWidgetUsingGET**
> WidgetDto GetWidgetUsingGET(ctx, widgetId)
Retrieves a widget

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **widgetId** | **int64**| widgetId | 

### Return type

[**WidgetDto**](WidgetDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWidgetsUsingGET**
> ListHolderDtoWidgetDto GetWidgetsUsingGET(ctx, )
Retrieves all widgets

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ListHolderDtoWidgetDto**](ListHolderDto«WidgetDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HandleUploadUsingPOST**
> WidgetDto HandleUploadUsingPOST(ctx, file, tenantId)
Uploads a widget

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **file** | ***os.File**| file | 
  **tenantId** | **string**| tenantId | 

### Return type

[**WidgetDto**](WidgetDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

