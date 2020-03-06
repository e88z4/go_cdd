# \ContentReportApi

All URIs are relative to *https://cddirector.io/cdd/reporting/0000/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetContentItemsReportUsingGET**](ContentReportApi.md#GetContentItemsReportUsingGET) | **Get** /reports/content-items-report | Get content items for report
[**GetContentItemsUsingGET1**](ContentReportApi.md#GetContentItemsUsingGET1) | **Get** /reports/content-items | Get Content Items by Filter


# **GetContentItemsReportUsingGET**
> ListHolderDtoContentItemReportDto GetContentItemsReportUsingGET(ctx, optional)
Get content items for report

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetContentItemsReportUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetContentItemsReportUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentItem** | [**optional.Interface of []int64**](int64.md)| contentItemsIds | 

### Return type

[**ListHolderDtoContentItemReportDto**](ListHolderDto«ContentItemReportDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContentItemsUsingGET1**
> ListHolderDtoContentItemDto GetContentItemsUsingGET1(ctx, optional)
Get Content Items by Filter

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetContentItemsUsingGET1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetContentItemsUsingGET1Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **releaseId** | **optional.Int64**| releaseId | 
 **applicationId** | **optional.Int64**| applicationId | 
 **filter** | **optional.String**| freeTextFilter | 

### Return type

[**ListHolderDtoContentItemDto**](ListHolderDto«ContentItemDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

