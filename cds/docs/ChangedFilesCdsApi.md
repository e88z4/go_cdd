# \ChangedFilesCdsApi

All URIs are relative to *https://cddirector.io/cdd/cds/0000/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetChangedFilesUsingGET**](ChangedFilesCdsApi.md#GetChangedFilesUsingGET) | **Get** /applications/{applicationId}/application-versions/{applicationVersionId}/changed-files | Retrieve changed files per Application version from CDS


# **GetChangedFilesUsingGET**
> PagedResultDtoChangedFileDto GetChangedFilesUsingGET(ctx, applicationId, applicationVersionId, optional)
Retrieve changed files per Application version from CDS

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int64**| applicationId | 
  **applicationVersionId** | **int64**| applicationVersionId | 
 **optional** | ***GetChangedFilesUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetChangedFilesUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **isAccessed** | **optional.Bool**| isAccessed | 
 **isExcluded** | **optional.Bool**| isExcluded | 
 **status** | [**optional.Interface of []string**](string.md)| fileStatuses | 
 **pageNumber** | **optional.Int32**| pageNumber | 
 **pageSize** | **optional.Int32**| pageSize | 

### Return type

[**PagedResultDtoChangedFileDto**](PagedResultDto«ChangedFileDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

