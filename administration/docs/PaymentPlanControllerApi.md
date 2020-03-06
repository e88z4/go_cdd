# \PaymentPlanControllerApi

All URIs are relative to *https://cddirector.io/cdd/administration/0000/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFirstSubscriptionUsingGET**](PaymentPlanControllerApi.md#CreateFirstSubscriptionUsingGET) | **Get** /payment-plans/subscription | createFirstSubscription
[**GetCurrentPaymentPlanUsingGET**](PaymentPlanControllerApi.md#GetCurrentPaymentPlanUsingGET) | **Get** /payment-plans/current | getCurrentPaymentPlan
[**HandleChargifyWebhookUsingPOST**](PaymentPlanControllerApi.md#HandleChargifyWebhookUsingPOST) | **Post** /payment-plans | handleChargifyWebhook
[**RedirectToPurchasePlanUsingGET**](PaymentPlanControllerApi.md#RedirectToPurchasePlanUsingGET) | **Get** /payment-plans | redirectToPurchasePlan


# **CreateFirstSubscriptionUsingGET**
> CreateFirstSubscriptionUsingGET(ctx, id, ref)
createFirstSubscription

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| subscriptionId | 
  **ref** | **string**| tenantId | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCurrentPaymentPlanUsingGET**
> PaymentPlanDto GetCurrentPaymentPlanUsingGET(ctx, )
getCurrentPaymentPlan

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**PaymentPlanDto**](PaymentPlanDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HandleChargifyWebhookUsingPOST**
> HandleChargifyWebhookUsingPOST(ctx, payload)
handleChargifyWebhook

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **payload** | [**Object**](.md)| payload | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RedirectToPurchasePlanUsingGET**
> RedirectToPurchasePlanUsingGET(ctx, optional)
redirectToPurchasePlan

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RedirectToPurchasePlanUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RedirectToPurchasePlanUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **plan** | **optional.String**| paymentPlan | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

