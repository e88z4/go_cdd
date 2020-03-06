# \TestSourceApi

All URIs are relative to *https://cddirector.io/cdd/design/0000/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTestSourceUsingPOST**](TestSourceApi.md#CreateTestSourceUsingPOST) | **Post** /applications/{applicationId}/application-versions/{applicationVersionId}/test-sources | Create a test source and import tests to TAR
[**DeleteTestSourceUsingDELETE**](TestSourceApi.md#DeleteTestSourceUsingDELETE) | **Delete** /applications/{applicationId}/application-versions/{applicationVersionId}/test-sources/{testSourceId} | Delete a test source
[**DeleteTestSourcesUsingDELETE**](TestSourceApi.md#DeleteTestSourcesUsingDELETE) | **Delete** /applications/{applicationId}/application-versions/{applicationVersionId}/test-sources | Delete Test Sources
[**DuplicateTestSuitesBasedOnAppVersionUsingPATCH**](TestSourceApi.md#DuplicateTestSuitesBasedOnAppVersionUsingPATCH) | **Patch** /applications/{applicationId}/application-versions/{applicationVersionId} | Assign all test suites of application version to a different application version
[**GetTestSourceUsingGET**](TestSourceApi.md#GetTestSourceUsingGET) | **Get** /applications/{applicationId}/application-versions/{applicationVersionId}/test-sources/{testSourceId} | Retrieve a test source
[**GetTestSourcesUsingGET**](TestSourceApi.md#GetTestSourcesUsingGET) | **Get** /applications/{applicationId}/application-versions/{applicationVersionId}/test-sources | Retrieve all test sources
[**SyncTestSuitesUsingPATCH**](TestSourceApi.md#SyncTestSuitesUsingPATCH) | **Patch** /applications/{applicationId}/application-versions/{applicationVersionId}/test-sources | Sync test suites
[**UpdateTestSourceUsingPUT**](TestSourceApi.md#UpdateTestSourceUsingPUT) | **Put** /applications/{applicationId}/application-versions/{applicationVersionId}/test-sources/{testSourceId} | update a test source


# **CreateTestSourceUsingPOST**
> TestSourceDto CreateTestSourceUsingPOST(ctx, applicationId, applicationVersionId, testSourceDto)
Create a test source and import tests to TAR

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int64**| applicationId | 
  **applicationVersionId** | **int64**| applicationVersionId | 
  **testSourceDto** | [**TestSourceDto**](TestSourceDto.md)| testSourceDto | 

### Return type

[**TestSourceDto**](TestSourceDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTestSourceUsingDELETE**
> DeletedTestsDto DeleteTestSourceUsingDELETE(ctx, applicationId, applicationVersionId, testSourceId)
Delete a test source

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int64**| applicationId | 
  **applicationVersionId** | **int64**| applicationVersionId | 
  **testSourceId** | **int64**| testSourceId | 

### Return type

[**DeletedTestsDto**](DeletedTestsDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTestSourcesUsingDELETE**
> DeletedTestsDto DeleteTestSourcesUsingDELETE(ctx, applicationId, applicationVersionId, optional)
Delete Test Sources

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int64**| applicationId | 
  **applicationVersionId** | **int64**| applicationVersionId | 
 **optional** | ***DeleteTestSourcesUsingDELETEOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteTestSourcesUsingDELETEOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **testSource** | [**optional.Interface of []int64**](int64.md)| ids | 

### Return type

[**DeletedTestsDto**](DeletedTestsDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DuplicateTestSuitesBasedOnAppVersionUsingPATCH**
> DuplicateTestSuitesBasedOnAppVersionUsingPATCH(ctx, applicationId, applicationVersionId, basedOnDto)
Assign all test suites of application version to a different application version

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int64**| applicationId | 
  **applicationVersionId** | **int64**| applicationVersionId | 
  **basedOnDto** | [**BasedOnDto**](BasedOnDto.md)| basedOnDto | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTestSourceUsingGET**
> TestSourceDto GetTestSourceUsingGET(ctx, applicationId, applicationVersionId, testSourceId)
Retrieve a test source

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int64**| applicationId | 
  **applicationVersionId** | **int64**| applicationVersionId | 
  **testSourceId** | **int64**| testSourceId | 

### Return type

[**TestSourceDto**](TestSourceDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTestSourcesUsingGET**
> PagedResultDtoTestSourceDto GetTestSourcesUsingGET(ctx, applicationId, applicationVersionId, optional)
Retrieve all test sources

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int64**| applicationId | 
  **applicationVersionId** | **int64**| applicationVersionId | 
 **optional** | ***GetTestSourcesUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetTestSourcesUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sortField** | **optional.String**| sortField | 
 **sortDirection** | **optional.String**| sortDirection | 
 **filter** | **optional.String**| filter | 
 **endpointId** | **optional.Int64**| endpointId | 
 **tag** | **optional.String**| tag | 
 **pageNumber** | **optional.Int32**| pageNumber | 
 **pageSize** | **optional.Int32**| pageSize | 

### Return type

[**PagedResultDtoTestSourceDto**](PagedResultDto«TestSourceDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SyncTestSuitesUsingPATCH**
> SyncTestSuitesUsingPATCH(ctx, applicationId, applicationVersionId, statusDto, optional)
Sync test suites

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int64**| applicationId | 
  **applicationVersionId** | **int64**| applicationVersionId | 
  **statusDto** | [**TestSuitesStatusDto**](TestSuitesStatusDto.md)| statusDto | 
 **optional** | ***SyncTestSuitesUsingPATCHOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SyncTestSuitesUsingPATCHOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **testSource** | [**optional.Interface of []int64**](int64.md)| testSourceIds | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTestSourceUsingPUT**
> TestSourceDto UpdateTestSourceUsingPUT(ctx, applicationId, applicationVersionId, testSourceId, testSourceDto)
update a test source

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int64**| applicationId | 
  **applicationVersionId** | **int64**| applicationVersionId | 
  **testSourceId** | **int64**| testSourceId | 
  **testSourceDto** | [**TestSourceDto**](TestSourceDto.md)| testSourceDto | 

### Return type

[**TestSourceDto**](TestSourceDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

