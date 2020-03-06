# \ProductApi

All URIs are relative to *https://cddirector.io/cdd/administration/0000/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomerFeedbackUsingPOST**](ProductApi.md#CreateCustomerFeedbackUsingPOST) | **Post** /product/customer-feedback | Creates a customer feedback
[**CreatePortfolioLicensingAgreementUsingPOST**](ProductApi.md#CreatePortfolioLicensingAgreementUsingPOST) | **Post** /product/license | Creates portfolio license agreement info
[**GetLicenseActivitiesUsingGET**](ProductApi.md#GetLicenseActivitiesUsingGET) | **Get** /product/license/activities | Retrieves portfolio license agreement activities
[**GetPortfolioLicensingAgreementUsingGET**](ProductApi.md#GetPortfolioLicensingAgreementUsingGET) | **Get** /product/license | Retrieves portfolio license agreement info
[**PatchPortfolioLicensingAgreementUsingPATCH**](ProductApi.md#PatchPortfolioLicensingAgreementUsingPATCH) | **Patch** /product/license | Patch portfolio license agreement site id
[**UpdatePortfolioLicensingAgreementUsingPUT**](ProductApi.md#UpdatePortfolioLicensingAgreementUsingPUT) | **Put** /product/license | Updates portfolio license agreement info


# **CreateCustomerFeedbackUsingPOST**
> MessageDto CreateCustomerFeedbackUsingPOST(ctx, messageDto)
Creates a customer feedback

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **messageDto** | [**MessageDto**](MessageDto.md)| messageDto | 

### Return type

[**MessageDto**](MessageDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreatePortfolioLicensingAgreementUsingPOST**
> PortfolioLicensingAgreementDto CreatePortfolioLicensingAgreementUsingPOST(ctx, plaDto)
Creates portfolio license agreement info

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **plaDto** | [**PortfolioLicensingAgreementDto**](PortfolioLicensingAgreementDto.md)| plaDto | 

### Return type

[**PortfolioLicensingAgreementDto**](PortfolioLicensingAgreementDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLicenseActivitiesUsingGET**
> PagedResultDtoLicenseActivityDto GetLicenseActivitiesUsingGET(ctx, optional)
Retrieves portfolio license agreement activities

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetLicenseActivitiesUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetLicenseActivitiesUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **optional.Int64**| startDate | 
 **endDate** | **optional.Int64**| endDate | 
 **pageNumber** | **optional.Int32**| pageNumber | 
 **pageSize** | **optional.Int32**| pageSize | 

### Return type

[**PagedResultDtoLicenseActivityDto**](PagedResultDto«LicenseActivityDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPortfolioLicensingAgreementUsingGET**
> PortfolioLicensingAgreementDto GetPortfolioLicensingAgreementUsingGET(ctx, )
Retrieves portfolio license agreement info

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**PortfolioLicensingAgreementDto**](PortfolioLicensingAgreementDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchPortfolioLicensingAgreementUsingPATCH**
> PortfolioLicensingAgreementDto PatchPortfolioLicensingAgreementUsingPATCH(ctx, plaDto, optional)
Patch portfolio license agreement site id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **plaDto** | [**PortfolioLicensingAgreementDto**](PortfolioLicensingAgreementDto.md)| plaDto | 
 **optional** | ***PatchPortfolioLicensingAgreementUsingPATCHOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PatchPortfolioLicensingAgreementUsingPATCHOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **optional.String**| targetTenantId | 

### Return type

[**PortfolioLicensingAgreementDto**](PortfolioLicensingAgreementDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePortfolioLicensingAgreementUsingPUT**
> PortfolioLicensingAgreementDto UpdatePortfolioLicensingAgreementUsingPUT(ctx, plaDto)
Updates portfolio license agreement info

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **plaDto** | [**PortfolioLicensingAgreementDto**](PortfolioLicensingAgreementDto.md)| plaDto | 

### Return type

[**PortfolioLicensingAgreementDto**](PortfolioLicensingAgreementDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

