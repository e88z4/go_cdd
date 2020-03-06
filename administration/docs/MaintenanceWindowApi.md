# \MaintenanceWindowApi

All URIs are relative to *https://cddirector.io/cdd/administration/0000/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMaintenanceWindowUsingPOST**](MaintenanceWindowApi.md#CreateMaintenanceWindowUsingPOST) | **Post** /maintenance-windows | Creates a maintenance window
[**DeleteMaintenanceWindowUsingDELETE**](MaintenanceWindowApi.md#DeleteMaintenanceWindowUsingDELETE) | **Delete** /maintenance-windows/{maintenanceWindowId} | Deletes a maintenance window
[**GetMaintenanceWindowUsingGET**](MaintenanceWindowApi.md#GetMaintenanceWindowUsingGET) | **Get** /maintenance-windows/{maintenanceWindowId} | Retrieves a maintenance window
[**GetMaintenanceWindowsUsingGET**](MaintenanceWindowApi.md#GetMaintenanceWindowsUsingGET) | **Get** /maintenance-windows | Retrieves all maintenance windows
[**PatchMaintenanceWindowUsingPATCH**](MaintenanceWindowApi.md#PatchMaintenanceWindowUsingPATCH) | **Patch** /maintenance-windows/{maintenanceWindowId} | Changes a maintenance window
[**UpdateMaintenanceWindowUsingPUT**](MaintenanceWindowApi.md#UpdateMaintenanceWindowUsingPUT) | **Put** /maintenance-windows/{maintenanceWindowId} | Updates a maintenance window


# **CreateMaintenanceWindowUsingPOST**
> MaintenanceWindowDto CreateMaintenanceWindowUsingPOST(ctx, maintenanceWindowDto)
Creates a maintenance window

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **maintenanceWindowDto** | [**MaintenanceWindowDto**](MaintenanceWindowDto.md)| maintenanceWindowDto | 

### Return type

[**MaintenanceWindowDto**](MaintenanceWindowDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMaintenanceWindowUsingDELETE**
> DeleteMaintenanceWindowUsingDELETE(ctx, maintenanceWindowId)
Deletes a maintenance window

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **maintenanceWindowId** | **int64**| maintenanceWindowId | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMaintenanceWindowUsingGET**
> MaintenanceWindowDto GetMaintenanceWindowUsingGET(ctx, maintenanceWindowId)
Retrieves a maintenance window

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **maintenanceWindowId** | **int64**| maintenanceWindowId | 

### Return type

[**MaintenanceWindowDto**](MaintenanceWindowDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMaintenanceWindowsUsingGET**
> ListHolderDtoMaintenanceWindowDto GetMaintenanceWindowsUsingGET(ctx, optional)
Retrieves all maintenance windows

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetMaintenanceWindowsUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetMaintenanceWindowsUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **optional.String**| freeTextFilter | 

### Return type

[**ListHolderDtoMaintenanceWindowDto**](ListHolderDto«MaintenanceWindowDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchMaintenanceWindowUsingPATCH**
> MaintenanceWindowDto PatchMaintenanceWindowUsingPATCH(ctx, maintenanceWindowId, maintenanceWindowDto)
Changes a maintenance window

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **maintenanceWindowId** | **int64**| maintenanceWindowId | 
  **maintenanceWindowDto** | [**MaintenanceWindowDto**](MaintenanceWindowDto.md)| maintenanceWindowDto | 

### Return type

[**MaintenanceWindowDto**](MaintenanceWindowDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMaintenanceWindowUsingPUT**
> MaintenanceWindowDto UpdateMaintenanceWindowUsingPUT(ctx, maintenanceWindowId, maintenanceWindowDto)
Updates a maintenance window

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **maintenanceWindowId** | **int64**| maintenanceWindowId | 
  **maintenanceWindowDto** | [**MaintenanceWindowDto**](MaintenanceWindowDto.md)| maintenanceWindowDto | 

### Return type

[**MaintenanceWindowDto**](MaintenanceWindowDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

