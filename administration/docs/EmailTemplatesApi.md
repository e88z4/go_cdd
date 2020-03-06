# \EmailTemplatesApi

All URIs are relative to *https://cddirector.io/cdd/administration/0000/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEmailNotificationMessageUsingGET**](EmailTemplatesApi.md#GetEmailNotificationMessageUsingGET) | **Get** /email-notification-messages/{emailNotificationMessageName}/{languageTag} | Get an email message template
[**GetEmailNotificationMessagesUsingGET**](EmailTemplatesApi.md#GetEmailNotificationMessagesUsingGET) | **Get** /email-notification-messages/{languageTag} | Get list of email message templates
[**GetEmailNotificationTemplateUsingGET**](EmailTemplatesApi.md#GetEmailNotificationTemplateUsingGET) | **Get** /email-notification-templates/{emailNotificationTemplateName}/{languageTag} | Get an email template
[**GetEmailNotificationTemplatesUsingGET**](EmailTemplatesApi.md#GetEmailNotificationTemplatesUsingGET) | **Get** /email-notification-templates/{languageTag} | Get list of email templates
[**PatchEmailNotificationMessageUsingPATCH**](EmailTemplatesApi.md#PatchEmailNotificationMessageUsingPATCH) | **Patch** /email-notification-messages/{emailNotificationMessageName}/{languageTag} | Patches an email notification
[**PatchEmailNotificationTemplateUsingPATCH**](EmailTemplatesApi.md#PatchEmailNotificationTemplateUsingPATCH) | **Patch** /email-notification-templates/{emailNotificationTemplateName}/{languageTag} | Patches an email template
[**TryEmailNotificationMessageUsingPOST**](EmailTemplatesApi.md#TryEmailNotificationMessageUsingPOST) | **Post** /email-notification-messages/{emailNotificationMessageName}/{languageTag}/test-email | Try email notification message
[**TryEmailNotificationTemplateUsingPOST**](EmailTemplatesApi.md#TryEmailNotificationTemplateUsingPOST) | **Post** /email-notification-templates/{emailNotificationTemplateName}/{languageTag}/test-email | Try email notification template


# **GetEmailNotificationMessageUsingGET**
> EmailNotificationMessageDto GetEmailNotificationMessageUsingGET(ctx, emailNotificationMessageName, languageTag)
Get an email message template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **emailNotificationMessageName** | **string**| msgName | 
  **languageTag** | **string**| language | 

### Return type

[**EmailNotificationMessageDto**](EmailNotificationMessageDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEmailNotificationMessagesUsingGET**
> ListHolderDtoEmailNotificationDto GetEmailNotificationMessagesUsingGET(ctx, languageTag)
Get list of email message templates

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **languageTag** | **string**| language | 

### Return type

[**ListHolderDtoEmailNotificationDto**](ListHolderDto«EmailNotificationDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEmailNotificationTemplateUsingGET**
> EmailNotificationTemplateDto GetEmailNotificationTemplateUsingGET(ctx, emailNotificationTemplateName, languageTag)
Get an email template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **emailNotificationTemplateName** | **string**| templateName | 
  **languageTag** | **string**| language | 

### Return type

[**EmailNotificationTemplateDto**](EmailNotificationTemplateDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEmailNotificationTemplatesUsingGET**
> ListHolderDtoEmailNotificationTemplateDto GetEmailNotificationTemplatesUsingGET(ctx, languageTag)
Get list of email templates

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **languageTag** | **string**| language | 

### Return type

[**ListHolderDtoEmailNotificationTemplateDto**](ListHolderDto«EmailNotificationTemplateDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchEmailNotificationMessageUsingPATCH**
> EmailNotificationMessageDto PatchEmailNotificationMessageUsingPATCH(ctx, emailNotificationMessageName, languageTag, emailNotificationMessageDto)
Patches an email notification

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **emailNotificationMessageName** | **string**| msgName | 
  **languageTag** | **string**| language | 
  **emailNotificationMessageDto** | [**EmailNotificationMessageDto**](EmailNotificationMessageDto.md)| emailNotificationMessageDto | 

### Return type

[**EmailNotificationMessageDto**](EmailNotificationMessageDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchEmailNotificationTemplateUsingPATCH**
> EmailNotificationTemplateDto PatchEmailNotificationTemplateUsingPATCH(ctx, emailNotificationTemplateName, languageTag, emailNotificationTemplateDto)
Patches an email template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **emailNotificationTemplateName** | **string**| templateName | 
  **languageTag** | **string**| language | 
  **emailNotificationTemplateDto** | [**EmailNotificationTemplateDto**](EmailNotificationTemplateDto.md)| emailNotificationTemplateDto | 

### Return type

[**EmailNotificationTemplateDto**](EmailNotificationTemplateDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TryEmailNotificationMessageUsingPOST**
> EmailNotificationMessageDto TryEmailNotificationMessageUsingPOST(ctx, emailNotificationMessageName, languageTag, emailNotificationMessageDto)
Try email notification message

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **emailNotificationMessageName** | **string**| msgName | 
  **languageTag** | **string**| language | 
  **emailNotificationMessageDto** | [**EmailNotificationMessageDto**](EmailNotificationMessageDto.md)| emailNotificationMessageDto | 

### Return type

[**EmailNotificationMessageDto**](EmailNotificationMessageDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TryEmailNotificationTemplateUsingPOST**
> EmailNotificationTemplateDto TryEmailNotificationTemplateUsingPOST(ctx, emailNotificationTemplateName, languageTag, emailNotificationTemplateDto)
Try email notification template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **emailNotificationTemplateName** | **string**| templateName | 
  **languageTag** | **string**| language | 
  **emailNotificationTemplateDto** | [**EmailNotificationTemplateDto**](EmailNotificationTemplateDto.md)| emailNotificationTemplateDto | 

### Return type

[**EmailNotificationTemplateDto**](EmailNotificationTemplateDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

