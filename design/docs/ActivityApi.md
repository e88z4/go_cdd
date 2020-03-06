# \ActivityApi

All URIs are relative to *https://cddirector.io/cdd/design/0000/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateActivityUsingPOST**](ActivityApi.md#CreateActivityUsingPOST) | **Post** /releases/{releaseId}/activities | Creates an activity
[**GetActivityCategoriesUsingGET**](ActivityApi.md#GetActivityCategoriesUsingGET) | **Get** /activities/categories | Retrieves activity categories
[**GetActivityGeneratorsUsingGET**](ActivityApi.md#GetActivityGeneratorsUsingGET) | **Get** /activities/generators | Retrieves activity generators
[**GetAllActivitiesContentTemplatesUsingGET**](ActivityApi.md#GetAllActivitiesContentTemplatesUsingGET) | **Get** /activities/content-templates/{languageName} | Retrieves activity content templates
[**GetAllActivitiesOfReleaseUsingGET**](ActivityApi.md#GetAllActivitiesOfReleaseUsingGET) | **Get** /releases/{releaseId}/activities | Retrieves all activities of a release
[**GetAllActivitiesUsingGET**](ActivityApi.md#GetAllActivitiesUsingGET) | **Get** /activities | Retrieves all activities
[**PatchActivityUsingPATCH**](ActivityApi.md#PatchActivityUsingPATCH) | **Patch** /releases/{releaseId}/activities/{activityId} | Changes an activity status


# **CreateActivityUsingPOST**
> ActivityDto CreateActivityUsingPOST(ctx, releaseId, activityDto)
Creates an activity

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseId** | **int64**| releaseId | 
  **activityDto** | [**ActivityDto**](ActivityDto.md)| activityDto | 

### Return type

[**ActivityDto**](ActivityDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetActivityCategoriesUsingGET**
> ListHolderDtoActivityCategoryDto GetActivityCategoriesUsingGET(ctx, optional)
Retrieves activity categories

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetActivityCategoriesUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetActivityCategoriesUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **optional.String**| filter | 

### Return type

[**ListHolderDtoActivityCategoryDto**](ListHolderDto«ActivityCategoryDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetActivityGeneratorsUsingGET**
> ListHolderDtoStringWrapperDto GetActivityGeneratorsUsingGET(ctx, startDate, optional)
Retrieves activity generators

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **startDate** | **int64**| startDate | 
 **optional** | ***GetActivityGeneratorsUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetActivityGeneratorsUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **optional.String**| filter | 
 **category** | **optional.String**| category | 
 **creator** | [**optional.Interface of []string**](string.md)| creatorNamesFilter | 
 **contentTemplate** | **optional.String**| contentTemplateFilter | 
 **creatorEmailAddress** | [**optional.Interface of []string**](string.md)| creatorEmailsFilter | 
 **endDate** | **optional.Int64**| endDate | 
 **pageSize** | **optional.Int32**| limit | 

### Return type

[**ListHolderDtoStringWrapperDto**](ListHolderDto«StringWrapperDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllActivitiesContentTemplatesUsingGET**
> ListHolderDtoActivityContentTemplateDto GetAllActivitiesContentTemplatesUsingGET(ctx, languageName, optional)
Retrieves activity content templates

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **languageName** | **string**| lang | 
 **optional** | ***GetAllActivitiesContentTemplatesUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAllActivitiesContentTemplatesUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **optional.String**| filter | 
 **category** | **optional.String**| categoryFilter | 

### Return type

[**ListHolderDtoActivityContentTemplateDto**](ListHolderDto«ActivityContentTemplateDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllActivitiesOfReleaseUsingGET**
> PagedResultDtoActivityDto GetAllActivitiesOfReleaseUsingGET(ctx, releaseId, optional)
Retrieves all activities of a release

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseId** | **int64**| releaseId | 
 **optional** | ***GetAllActivitiesOfReleaseUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAllActivitiesOfReleaseUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **optional.String**| freeTextFilter | 
 **pageSize** | **optional.Int32**| pageSize | 
 **pageNumber** | **optional.Int32**| pageNumber | 
 **phase** | **optional.Int64**| phaseId | 
 **task** | [**optional.Interface of []int64**](int64.md)| taskIds | 
 **type_** | [**optional.Interface of []string**](string.md)| filterByType | 
 **currentUser** | **optional.Bool**| showMyActivities | 
 **status** | [**optional.Interface of []string**](string.md)| filterByStatus | 

### Return type

[**PagedResultDtoActivityDto**](PagedResultDto«ActivityDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllActivitiesUsingGET**
> PagedResultDtoActivityDto GetAllActivitiesUsingGET(ctx, startDate, optional)
Retrieves all activities

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **startDate** | **int64**| startDate | 
 **optional** | ***GetAllActivitiesUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAllActivitiesUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **optional.Int32**| pageSize | 
 **pageNumber** | **optional.Int32**| pageNumber | 
 **endDate** | **optional.Int64**| endDate | 
 **type_** | [**optional.Interface of []string**](string.md)| typesFilter | 
 **creator** | [**optional.Interface of []string**](string.md)| creatorNamesFilter | 
 **creatorEmailAddress** | [**optional.Interface of []string**](string.md)| creatorEmailsFilter | 
 **category** | **optional.String**| categoryFilter | 
 **contentTemplate** | **optional.String**| contentTemplateFilter | 
 **generator** | **optional.String**| contentGeneratorFilter | 
 **project** | **optional.Int64**| projectId | 

### Return type

[**PagedResultDtoActivityDto**](PagedResultDto«ActivityDto».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchActivityUsingPATCH**
> ActivityDto PatchActivityUsingPATCH(ctx, releaseId, activityId, activityDto)
Changes an activity status

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **releaseId** | **int64**| releaseId | 
  **activityId** | **int64**| activityId | 
  **activityDto** | [**ActivityDto**](ActivityDto.md)| activityDto | 

### Return type

[**ActivityDto**](ActivityDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

