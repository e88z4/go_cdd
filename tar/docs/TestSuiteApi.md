# \TestSuiteApi

All URIs are relative to *https://cddirector.io/cdd/tar/0000/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddTagsAndIsFavoriteToTestSuitesUsingPATCH**](TestSuiteApi.md#AddTagsAndIsFavoriteToTestSuitesUsingPATCH) | **Patch** /applications/{applicationId}/application-versions/{applicationVersionId}/test-suites | Add tags to test suites
[**DeleteTestSuitesUsingDELETE**](TestSuiteApi.md#DeleteTestSuitesUsingDELETE) | **Delete** /applications/{applicationId}/application-versions/{applicationVersionId}/test-suites | Delete test suites
[**GetTestSuitesUsingGET**](TestSuiteApi.md#GetTestSuitesUsingGET) | **Get** /applications/{applicationId}/application-versions/{applicationVersionId}/test-suites | Retrieve all test suites
[**UpdateTagsAndIsFavoriteForTestSuiteUsingPATCH**](TestSuiteApi.md#UpdateTagsAndIsFavoriteForTestSuiteUsingPATCH) | **Patch** /applications/{applicationId}/application-versions/{applicationVersionId}/test-suites/{testSuiteHexId} | Update Test Suite (Tags or IsFavorite)


# **AddTagsAndIsFavoriteToTestSuitesUsingPATCH**
> AddTagsAndIsFavoriteToTestSuitesUsingPATCH(ctx, applicationId, applicationVersionId, testSuiteIds, optional)
Add tags to test suites

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int64**| applicationId | 
  **applicationVersionId** | **int64**| applicationVersionId | 
  **testSuiteIds** | [**StringsDtoList**](StringsDtoList.md)| testSuiteIds | 
 **optional** | ***AddTagsAndIsFavoriteToTestSuitesUsingPATCHOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AddTagsAndIsFavoriteToTestSuitesUsingPATCHOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **tag** | [**optional.Interface of []string**](string.md)| tags | 
 **isFavorite** | **optional.Bool**| isFavorite | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTestSuitesUsingDELETE**
> DeletedTestsDto DeleteTestSuitesUsingDELETE(ctx, applicationId, applicationVersionId, testSuite)
Delete test suites

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int64**| applicationId | 
  **applicationVersionId** | **int64**| applicationVersionId | 
  **testSuite** | [**[]string**](string.md)| ids | 

### Return type

[**DeletedTestsDto**](DeletedTestsDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTestSuitesUsingGET**
> PagedResultDtoTestSuiteDto GetTestSuitesUsingGET(ctx, applicationId, applicationVersionId, optional)
Retrieve all test suites

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int64**| applicationId | 
  **applicationVersionId** | **int64**| applicationVersionId | 
 **optional** | ***GetTestSuitesUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetTestSuitesUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **testSource** | **optional.Int64**| testSourceId | 
 **sortField** | **optional.String**| sortField | 
 **sortDirection** | **optional.String**| sortDirection | 
 **filter** | **optional.String**| filter | 
 **endpointId** | **optional.Int64**| endpointId | 
 **tag** | **optional.String**| tagsCsv | 
 **pageNumber** | **optional.Int32**| pageNumber | 
 **pageSize** | **optional.Int32**| pageSize | 
 **embed** | **optional.String**| embedFilter | 

### Return type

[**PagedResultDtoTestSuiteDto**](PagedResultDto«TestSuiteDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTagsAndIsFavoriteForTestSuiteUsingPATCH**
> TestSuiteDto UpdateTagsAndIsFavoriteForTestSuiteUsingPATCH(ctx, applicationId, applicationVersionId, testSuiteHexId, testSuiteDto)
Update Test Suite (Tags or IsFavorite)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int64**| applicationId | 
  **applicationVersionId** | **int64**| applicationVersionId | 
  **testSuiteHexId** | **string**| testSuiteId | 
  **testSuiteDto** | [**TestSuiteDto**](TestSuiteDto.md)| testSuiteDto | 

### Return type

[**TestSuiteDto**](TestSuiteDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

