# \TenantApi

All URIs are relative to *https://cddirector.io/cdd/administration/0000/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteTenantUsingDELETE**](TenantApi.md#DeleteTenantUsingDELETE) | **Delete** /tenants/{tenantId} | Deletes a tenant
[**GetAllTenantsUsingGET**](TenantApi.md#GetAllTenantsUsingGET) | **Get** /tenants | Retrieves all tenants
[**GetTenantUsingGET**](TenantApi.md#GetTenantUsingGET) | **Get** /tenants/{tenantId} | Retrieves tenant
[**PatchTenantUsingPATCH1**](TenantApi.md#PatchTenantUsingPATCH1) | **Patch** /tenants/{tenantId} | Patches a tenant


# **DeleteTenantUsingDELETE**
> DeleteTenantUsingDELETE(ctx, tenantId)
Deletes a tenant

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tenantId** | **int64**| tenantId | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllTenantsUsingGET**
> ListHolderDtoTenantDto GetAllTenantsUsingGET(ctx, optional)
Retrieves all tenants

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllTenantsUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAllTenantsUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **currentUser** | **optional.Bool**| myTenants | 

### Return type

[**ListHolderDtoTenantDto**](ListHolderDto«TenantDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTenantUsingGET**
> TenantDto GetTenantUsingGET(ctx, tenantId)
Retrieves tenant

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tenantId** | **int64**| tenantId | 

### Return type

[**TenantDto**](TenantDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchTenantUsingPATCH1**
> TenantDto PatchTenantUsingPATCH1(ctx, tenantId, tenantDto)
Patches a tenant

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tenantId** | **int64**| tenantId | 
  **tenantDto** | [**TenantDto**](TenantDto.md)| tenantDto | 

### Return type

[**TenantDto**](TenantDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

