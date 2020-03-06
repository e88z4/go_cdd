# \SamlDomainApi

All URIs are relative to *https://cddirector.io/cdd/administration/0000/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSamlDomainUsingPOST**](SamlDomainApi.md#CreateSamlDomainUsingPOST) | **Post** /saml-domains | Create a SAML domain record
[**DeleteSamlDomainUsingDELETE**](SamlDomainApi.md#DeleteSamlDomainUsingDELETE) | **Delete** /saml-domains/{samlDomainId} | Delete a specific SAML domain record
[**GetSamlDomainUsingGET**](SamlDomainApi.md#GetSamlDomainUsingGET) | **Get** /saml-domains/{samlDomainId} | Get all SAML domain records
[**GetSamlDomainsUsingGET**](SamlDomainApi.md#GetSamlDomainsUsingGET) | **Get** /saml-domains | Get all SAML domain records
[**PatchAllSamlDomainsUsingPATCH**](SamlDomainApi.md#PatchAllSamlDomainsUsingPATCH) | **Patch** /saml-domains | Patch all the SAML domain records within a specific tenant
[**PatchSamlDomainUsingPATCH**](SamlDomainApi.md#PatchSamlDomainUsingPATCH) | **Patch** /saml-domains/{samlDomainId} | Patch a specific SAML domain record
[**UpdateSamlDomainUsingPUT**](SamlDomainApi.md#UpdateSamlDomainUsingPUT) | **Put** /saml-domains/{samlDomainId} | Update the domain of a specific SAML domain record


# **CreateSamlDomainUsingPOST**
> SamlDomainDto CreateSamlDomainUsingPOST(ctx, samlDomainDto)
Create a SAML domain record

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **samlDomainDto** | [**SamlDomainDto**](SamlDomainDto.md)| samlDomainDto | 

### Return type

[**SamlDomainDto**](SamlDomainDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSamlDomainUsingDELETE**
> DeleteSamlDomainUsingDELETE(ctx, samlDomainId)
Delete a specific SAML domain record

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **samlDomainId** | **int64**| samlDomainId | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSamlDomainUsingGET**
> SamlDomainDto GetSamlDomainUsingGET(ctx, samlDomainId)
Get all SAML domain records

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **samlDomainId** | **int64**| samlDomainId | 

### Return type

[**SamlDomainDto**](SamlDomainDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSamlDomainsUsingGET**
> ListHolderDtoSamlDomainDto GetSamlDomainsUsingGET(ctx, )
Get all SAML domain records

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ListHolderDtoSamlDomainDto**](ListHolderDto«SamlDomainDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchAllSamlDomainsUsingPATCH**
> PatchAllSamlDomainsUsingPATCH(ctx, tenantId, samlDomainDto, optional)
Patch all the SAML domain records within a specific tenant

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tenantId** | **string**| targetTenantId | 
  **samlDomainDto** | [**SamlDomainDto**](SamlDomainDto.md)| samlDomainDto | 
 **optional** | ***PatchAllSamlDomainsUsingPATCHOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PatchAllSamlDomainsUsingPATCHOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **domainName** | **optional.String**| domainName | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchSamlDomainUsingPATCH**
> SamlDomainDto PatchSamlDomainUsingPATCH(ctx, samlDomainId, samlDomainDto, optional)
Patch a specific SAML domain record

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **samlDomainId** | **int64**| samlDomainId | 
  **samlDomainDto** | [**SamlDomainDto**](SamlDomainDto.md)| samlDomainDto | 
 **optional** | ***PatchSamlDomainUsingPATCHOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PatchSamlDomainUsingPATCHOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **tenantId** | **optional.String**| targetTenantId | 

### Return type

[**SamlDomainDto**](SamlDomainDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSamlDomainUsingPUT**
> SamlDomainDto UpdateSamlDomainUsingPUT(ctx, samlDomainId, samlDomainDto)
Update the domain of a specific SAML domain record

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **samlDomainId** | **int64**| samlDomainId | 
  **samlDomainDto** | [**SamlDomainDto**](SamlDomainDto.md)| samlDomainDto | 

### Return type

[**SamlDomainDto**](SamlDomainDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

