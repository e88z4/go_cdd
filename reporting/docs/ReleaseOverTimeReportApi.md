# \ReleaseOverTimeReportApi

All URIs are relative to *https://cddirector.io/cdd/reporting/0000/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetReleasesOverTimeReportUsingGET**](ReleaseOverTimeReportApi.md#GetReleasesOverTimeReportUsingGET) | **Get** /reports/releases-over-time-report | Get release over time report


# **GetReleasesOverTimeReportUsingGET**
> ReleasesOverTimeReportDto GetReleasesOverTimeReportUsingGET(ctx, optional)
Get release over time report

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetReleasesOverTimeReportUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetReleasesOverTimeReportUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **optional.Int64**| startDate | 
 **endDate** | **optional.Int64**| endDate | 
 **statusReason** | [**optional.Interface of []string**](string.md)| statusReasonFilter | 
 **pageSize** | **optional.Int32**| pageSize | 
 **pageNumber** | **optional.Int32**| pageNumber | 

### Return type

[**ReleasesOverTimeReportDto**](ReleasesOverTimeReportDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

