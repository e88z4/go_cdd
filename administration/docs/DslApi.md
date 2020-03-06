# \DslApi

All URIs are relative to *https://cddirector.io/cdd/administration/0000/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportApplicationsUsingGET**](DslApi.md#ExportApplicationsUsingGET) | **Get** /applications/dsl-manifest | Export Applications And Environments
[**ExportEndpointsUsingGET**](DslApi.md#ExportEndpointsUsingGET) | **Get** /endpoints/dsl-manifest | Export Endpoints
[**ExportEnvironmentsUsingGET**](DslApi.md#ExportEnvironmentsUsingGET) | **Get** /environments/dsl-manifest | Export Environments
[**ExportReleaseUsingGET**](DslApi.md#ExportReleaseUsingGET) | **Get** /releases/{releaseId}/dsl-manifest | Export Release
[**ExportTestSourceUsingGET**](DslApi.md#ExportTestSourceUsingGET) | **Get** /applications/{applicationId}/application-versions/{applicationVersionId}/test-sources/{testSourceId}/dsl-manifest | Export test source
[**ExportTestSourcesUsingGET**](DslApi.md#ExportTestSourcesUsingGET) | **Get** /applications/{applicationId}/application-versions/{applicationVersionId}/test-sources/dsl-manifest | Export test sources
[**ImportEntitiesUsingPOST**](DslApi.md#ImportEntitiesUsingPOST) | **Post** /dsl-manifest/attachment | Import from file
[**ImportEntitiesUsingPOST1**](DslApi.md#ImportEntitiesUsingPOST1) | **Post** /dsl-manifests | Import


# **ExportApplicationsUsingGET**
> DslTemplateDto ExportApplicationsUsingGET(ctx, optional)
Export Applications And Environments

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ExportApplicationsUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExportApplicationsUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dispositionType** | **optional.String**| dispositionTypeDto | 
 **application** | [**optional.Interface of []int64**](int64.md)| applicationIds | 
 **embed** | [**optional.Interface of []string**](string.md)| embedFields | 

### Return type

[**DslTemplateDto**](DslTemplateDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExportEndpointsUsingGET**
> DslTemplateDto ExportEndpointsUsingGET(ctx, optional)
Export Endpoints

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ExportEndpointsUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExportEndpointsUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dispositionType** | **optional.String**| dispositionTypeDto | 
 **endpoint** | [**optional.Interface of []int64**](int64.md)| endpointIds | 

### Return type

[**DslTemplateDto**](DslTemplateDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExportEnvironmentsUsingGET**
> DslTemplateDto ExportEnvironmentsUsingGET(ctx, optional)
Export Environments

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ExportEnvironmentsUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExportEnvironmentsUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dispositionType** | **optional.String**| dispositionTypeDto | 
 **environment** | [**optional.Interface of []int64**](int64.md)| environmentIds | 

### Return type

[**DslTemplateDto**](DslTemplateDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExportReleaseUsingGET**
> DslTemplateDto ExportReleaseUsingGET(ctx, releaseId, optional)
Export Release

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseId** | **int64**| releaseId | 
 **optional** | ***ExportReleaseUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExportReleaseUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dispositionType** | **optional.String**| dispositionTypeDto | 

### Return type

[**DslTemplateDto**](DslTemplateDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExportTestSourceUsingGET**
> DslTemplateDto ExportTestSourceUsingGET(ctx, applicationVersionId, testSourceId, optional)
Export test source

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationVersionId** | **int64**| applicationVersionId | 
  **testSourceId** | **int64**| testSourceId | 
 **optional** | ***ExportTestSourceUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExportTestSourceUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dispositionType** | **optional.String**| dispositionTypeDto | 

### Return type

[**DslTemplateDto**](DslTemplateDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExportTestSourcesUsingGET**
> DslTemplateDto ExportTestSourcesUsingGET(ctx, applicationVersionId, optional)
Export test sources

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationVersionId** | **int64**| applicationVersionId | 
 **optional** | ***ExportTestSourcesUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExportTestSourcesUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dispositionType** | **optional.String**| dispositionTypeDto | 

### Return type

[**DslTemplateDto**](DslTemplateDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImportEntitiesUsingPOST**
> DslResponseDto ImportEntitiesUsingPOST(ctx, file, optional)
Import from file

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **file** | ***os.File**| file | 
 **optional** | ***ImportEntitiesUsingPOSTOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ImportEntitiesUsingPOSTOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kind** | [**optional.Interface of []string**](string.md)| availableKinds | 

### Return type

[**DslResponseDto**](DslResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImportEntitiesUsingPOST1**
> DslResponseDto ImportEntitiesUsingPOST1(ctx, dslEntities, optional)
Import

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dslEntities** | [**DslTemplateDto**](DslTemplateDto.md)| dslEntities | 
 **optional** | ***ImportEntitiesUsingPOST1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ImportEntitiesUsingPOST1Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kind** | [**optional.Interface of []string**](string.md)| availableKinds | 

### Return type

[**DslResponseDto**](DslResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

