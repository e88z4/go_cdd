# \EnvironmentReportApi

All URIs are relative to *https://cddirector.io/cdd/reporting/0000/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEnvironmentReportUsingGET**](EnvironmentReportApi.md#GetEnvironmentReportUsingGET) | **Get** /reports/environments/{environmentId} | Get environment report
[**GetEnvironmentsReportUsingGET**](EnvironmentReportApi.md#GetEnvironmentsReportUsingGET) | **Get** /reports/environments | Get environments report


# **GetEnvironmentReportUsingGET**
> EnvironmentReportDto GetEnvironmentReportUsingGET(ctx, environmentId)
Get environment report

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **environmentId** | **int64**| environmentId | 

### Return type

[**EnvironmentReportDto**](EnvironmentReportDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEnvironmentsReportUsingGET**
> ListHolderDtoEnvironmentReportDto GetEnvironmentsReportUsingGET(ctx, optional)
Get environments report

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetEnvironmentsReportUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetEnvironmentsReportUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | [**optional.Interface of []int64**](int64.md)| environmentIds | 

### Return type

[**ListHolderDtoEnvironmentReportDto**](ListHolderDto«EnvironmentReportDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

