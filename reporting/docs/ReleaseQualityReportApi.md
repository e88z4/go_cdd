# \ReleaseQualityReportApi

All URIs are relative to *https://cddirector.io/cdd/reporting/0000/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetReleaseQualityMetricsPerApplicationVersionUsingGET**](ReleaseQualityReportApi.md#GetReleaseQualityMetricsPerApplicationVersionUsingGET) | **Get** /reports/releases/{releaseId}/applications/{applicationId}/application-versions/{applicationVersionId}/release-quality | Get release quality metrics for application version
[**GetReleaseQualityReportUsingGET**](ReleaseQualityReportApi.md#GetReleaseQualityReportUsingGET) | **Get** /reports/releases/{releaseId}/release-quality | Get release quality report


# **GetReleaseQualityMetricsPerApplicationVersionUsingGET**
> ReleaseQualityReportApplicationVersionDto GetReleaseQualityMetricsPerApplicationVersionUsingGET(ctx, releaseId, applicationId, applicationVersionId, optional)
Get release quality metrics for application version

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseId** | **int64**| releaseId | 
  **applicationId** | **int64**| applicationId | 
  **applicationVersionId** | **int64**| applicationVersionId | 
 **optional** | ***GetReleaseQualityMetricsPerApplicationVersionUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetReleaseQualityMetricsPerApplicationVersionUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **embed** | **optional.String**| embedField | 
 **isChangedFileExcluded** | **optional.Bool**| isChangedFileExcluded | 
 **changedFileStatus** | [**optional.Interface of []string**](string.md)| fileStatuses | 

### Return type

[**ReleaseQualityReportApplicationVersionDto**](ReleaseQualityReportApplicationVersionDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReleaseQualityReportUsingGET**
> ReleaseQualityReportDto GetReleaseQualityReportUsingGET(ctx, releaseId, optional)
Get release quality report

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseId** | **int64**| releaseId | 
 **optional** | ***GetReleaseQualityReportUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetReleaseQualityReportUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **isChangedFileExcluded** | **optional.Bool**| isChangedFileExcluded | 
 **changedFileStatus** | [**optional.Interface of []string**](string.md)| fileStatuses | 

### Return type

[**ReleaseQualityReportDto**](ReleaseQualityReportDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

