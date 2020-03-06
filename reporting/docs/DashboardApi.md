# \DashboardApi

All URIs are relative to *https://cddirector.io/cdd/reporting/0000/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDashboardPageUsingDELETE**](DashboardApi.md#DeleteDashboardPageUsingDELETE) | **Delete** /dashboard/pages/{pageId} | Deletes a dashboard page
[**GetDashboardPageUsingGET**](DashboardApi.md#GetDashboardPageUsingGET) | **Get** /dashboard/pages/{pageId} | Retrieves a dashboard page
[**GetDashboardPagesUsingGET**](DashboardApi.md#GetDashboardPagesUsingGET) | **Get** /dashboard/pages | Retrieves all dashboard pages
[**SaveDashboardPageUsingPOST**](DashboardApi.md#SaveDashboardPageUsingPOST) | **Post** /dashboard/pages | Saves a dashboard page
[**SortDashboardPagesUsingPATCH**](DashboardApi.md#SortDashboardPagesUsingPATCH) | **Patch** /dashboard/pages | Sorts dashboard pages


# **DeleteDashboardPageUsingDELETE**
> DeleteDashboardPageUsingDELETE(ctx, pageId)
Deletes a dashboard page

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **int64**| dashboardPageId | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDashboardPageUsingGET**
> DashboardPageDto GetDashboardPageUsingGET(ctx, pageId)
Retrieves a dashboard page

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **int64**| dashboardPageId | 

### Return type

[**DashboardPageDto**](DashboardPageDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDashboardPagesUsingGET**
> ListHolderDtoDashboardPageDto GetDashboardPagesUsingGET(ctx, optional)
Retrieves all dashboard pages

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetDashboardPagesUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetDashboardPagesUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isDefault** | **optional.Bool**| isDefault | 
 **auth** | [**optional.Interface of Authentication**](Authentication.md)| auth | 

### Return type

[**ListHolderDtoDashboardPageDto**](ListHolderDto«DashboardPageDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SaveDashboardPageUsingPOST**
> DashboardPageDto SaveDashboardPageUsingPOST(ctx, dashboardDto, optional)
Saves a dashboard page

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dashboardDto** | [**DashboardPageDto**](DashboardPageDto.md)| dashboardDto | 
 **optional** | ***SaveDashboardPageUsingPOSTOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SaveDashboardPageUsingPOSTOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **auth** | [**optional.Interface of Authentication**](Authentication.md)| auth | 

### Return type

[**DashboardPageDto**](DashboardPageDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SortDashboardPagesUsingPATCH**
> SortDashboardPagesUsingPATCH(ctx, pageIds)
Sorts dashboard pages

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageIds** | **[]int64**| pageIds | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

