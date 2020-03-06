# \FreezePeriodApi

All URIs are relative to *https://cddirector.io/cdd/administration/0000/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFreezePeriodUsingPOST**](FreezePeriodApi.md#CreateFreezePeriodUsingPOST) | **Post** /freeze-periods | Create freeze period entry
[**DeleteFreezePeriodUsingDELETE**](FreezePeriodApi.md#DeleteFreezePeriodUsingDELETE) | **Delete** /freeze-periods/{freezePeriodId} | Deletes freeze period entry
[**EnableDisableFreezePeriodUsingPATCH**](FreezePeriodApi.md#EnableDisableFreezePeriodUsingPATCH) | **Patch** /freeze-periods/{freezePeriodId} | enable/disable freeze period
[**GetFreezePeriodUsingGET**](FreezePeriodApi.md#GetFreezePeriodUsingGET) | **Get** /freeze-periods/{freezePeriodId} | Get freeze period entry
[**GetFreezePeriodsUsingGET**](FreezePeriodApi.md#GetFreezePeriodsUsingGET) | **Get** /freeze-periods | Retrieves all freeze periods
[**UpdateFreezePeriodUsingPUT**](FreezePeriodApi.md#UpdateFreezePeriodUsingPUT) | **Put** /freeze-periods/{freezePeriodId} | update freeze period


# **CreateFreezePeriodUsingPOST**
> FreezePeriodDto CreateFreezePeriodUsingPOST(ctx, freezePeriodDto)
Create freeze period entry

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **freezePeriodDto** | [**FreezePeriodDto**](FreezePeriodDto.md)| freezePeriodDto | 

### Return type

[**FreezePeriodDto**](FreezePeriodDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFreezePeriodUsingDELETE**
> DeleteFreezePeriodUsingDELETE(ctx, freezePeriodId)
Deletes freeze period entry

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **freezePeriodId** | **int64**| freezePeriodId | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnableDisableFreezePeriodUsingPATCH**
> FreezePeriodDto EnableDisableFreezePeriodUsingPATCH(ctx, freezePeriodId, freezePeriodDto)
enable/disable freeze period

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **freezePeriodId** | **int64**| freezePeriodId | 
  **freezePeriodDto** | [**FreezePeriodDto**](FreezePeriodDto.md)| freezePeriodDto | 

### Return type

[**FreezePeriodDto**](FreezePeriodDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFreezePeriodUsingGET**
> FreezePeriodDto GetFreezePeriodUsingGET(ctx, freezePeriodId)
Get freeze period entry

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **freezePeriodId** | **int64**| freezePeriodId | 

### Return type

[**FreezePeriodDto**](FreezePeriodDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFreezePeriodsUsingGET**
> ListHolderDtoFreezePeriodDto GetFreezePeriodsUsingGET(ctx, optional)
Retrieves all freeze periods

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetFreezePeriodsUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetFreezePeriodsUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **optional.String**| freeTextFilter | 
 **pageSize** | **optional.Int32**| pageSize | 
 **pageNumber** | **optional.Int32**| pageNumber | 

### Return type

[**ListHolderDtoFreezePeriodDto**](ListHolderDto«FreezePeriodDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFreezePeriodUsingPUT**
> FreezePeriodDto UpdateFreezePeriodUsingPUT(ctx, freezePeriodId, freezePeriodDto)
update freeze period

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **freezePeriodId** | **int64**| freezePeriodId | 
  **freezePeriodDto** | [**FreezePeriodDto**](FreezePeriodDto.md)| freezePeriodDto | 

### Return type

[**FreezePeriodDto**](FreezePeriodDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

