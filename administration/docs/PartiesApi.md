# \PartiesApi

All URIs are relative to *https://cddirector.io/cdd/administration/0000/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPartiesUsingGET**](PartiesApi.md#GetPartiesUsingGET) | **Get** /parties | Retrieves all users and user-groups


# **GetPartiesUsingGET**
> PagedResultDtoPartyDto GetPartiesUsingGET(ctx, optional)
Retrieves all users and user-groups

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetPartiesUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetPartiesUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **optional.String**| freeTextFilter | 
 **inProjectContext** | **optional.Bool**| inProjectContext | 
 **pageSize** | **optional.Int32**| pageSize | 
 **pageNumber** | **optional.Int32**| pageNumber | 

### Return type

[**PagedResultDtoPartyDto**](PagedResultDto«PartyDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

