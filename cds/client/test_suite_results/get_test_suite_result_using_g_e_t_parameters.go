// Code generated by go-swagger; DO NOT EDIT.

package test_suite_results

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetTestSuiteResultUsingGETParams creates a new GetTestSuiteResultUsingGETParams object
// with the default values initialized.
func NewGetTestSuiteResultUsingGETParams() *GetTestSuiteResultUsingGETParams {
	var ()
	return &GetTestSuiteResultUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTestSuiteResultUsingGETParamsWithTimeout creates a new GetTestSuiteResultUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTestSuiteResultUsingGETParamsWithTimeout(timeout time.Duration) *GetTestSuiteResultUsingGETParams {
	var ()
	return &GetTestSuiteResultUsingGETParams{

		timeout: timeout,
	}
}

// NewGetTestSuiteResultUsingGETParamsWithContext creates a new GetTestSuiteResultUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTestSuiteResultUsingGETParamsWithContext(ctx context.Context) *GetTestSuiteResultUsingGETParams {
	var ()
	return &GetTestSuiteResultUsingGETParams{

		Context: ctx,
	}
}

// NewGetTestSuiteResultUsingGETParamsWithHTTPClient creates a new GetTestSuiteResultUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTestSuiteResultUsingGETParamsWithHTTPClient(client *http.Client) *GetTestSuiteResultUsingGETParams {
	var ()
	return &GetTestSuiteResultUsingGETParams{
		HTTPClient: client,
	}
}

/*GetTestSuiteResultUsingGETParams contains all the parameters to send to the API endpoint
for the get test suite result using g e t operation typically these are written to a http.Request
*/
type GetTestSuiteResultUsingGETParams struct {

	/*ApplicationID
	  applicationId

	*/
	ApplicationID int64
	/*ApplicationVersionID
	  applicationVersionId

	*/
	ApplicationVersionID int64
	/*Embed
	  embed

	*/
	Embed []string
	/*EndpointID
	  endpointId

	*/
	EndpointID *int64
	/*Environment
	  environmentId

	*/
	Environment *int64
	/*Filter
	  filter

	*/
	Filter *string
	/*PageNumber
	  pageNumber

	*/
	PageNumber *int32
	/*PageSize
	  pageSize

	*/
	PageSize *int32
	/*PluginID
	  pluginId

	*/
	PluginID *int64
	/*SortDirection
	  sortDirection

	*/
	SortDirection *string
	/*SortField
	  sortField

	*/
	SortField *string
	/*Tag
	  tags

	*/
	Tag *string
	/*TestCaseResultStatus
	  testCaseStatuses

	*/
	TestCaseResultStatus []string
	/*TestSuiteExecutionID
	  testSuiteResultId

	*/
	TestSuiteExecutionID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get test suite result using g e t params
func (o *GetTestSuiteResultUsingGETParams) WithTimeout(timeout time.Duration) *GetTestSuiteResultUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get test suite result using g e t params
func (o *GetTestSuiteResultUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get test suite result using g e t params
func (o *GetTestSuiteResultUsingGETParams) WithContext(ctx context.Context) *GetTestSuiteResultUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get test suite result using g e t params
func (o *GetTestSuiteResultUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get test suite result using g e t params
func (o *GetTestSuiteResultUsingGETParams) WithHTTPClient(client *http.Client) *GetTestSuiteResultUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get test suite result using g e t params
func (o *GetTestSuiteResultUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApplicationID adds the applicationID to the get test suite result using g e t params
func (o *GetTestSuiteResultUsingGETParams) WithApplicationID(applicationID int64) *GetTestSuiteResultUsingGETParams {
	o.SetApplicationID(applicationID)
	return o
}

// SetApplicationID adds the applicationId to the get test suite result using g e t params
func (o *GetTestSuiteResultUsingGETParams) SetApplicationID(applicationID int64) {
	o.ApplicationID = applicationID
}

// WithApplicationVersionID adds the applicationVersionID to the get test suite result using g e t params
func (o *GetTestSuiteResultUsingGETParams) WithApplicationVersionID(applicationVersionID int64) *GetTestSuiteResultUsingGETParams {
	o.SetApplicationVersionID(applicationVersionID)
	return o
}

// SetApplicationVersionID adds the applicationVersionId to the get test suite result using g e t params
func (o *GetTestSuiteResultUsingGETParams) SetApplicationVersionID(applicationVersionID int64) {
	o.ApplicationVersionID = applicationVersionID
}

// WithEmbed adds the embed to the get test suite result using g e t params
func (o *GetTestSuiteResultUsingGETParams) WithEmbed(embed []string) *GetTestSuiteResultUsingGETParams {
	o.SetEmbed(embed)
	return o
}

// SetEmbed adds the embed to the get test suite result using g e t params
func (o *GetTestSuiteResultUsingGETParams) SetEmbed(embed []string) {
	o.Embed = embed
}

// WithEndpointID adds the endpointID to the get test suite result using g e t params
func (o *GetTestSuiteResultUsingGETParams) WithEndpointID(endpointID *int64) *GetTestSuiteResultUsingGETParams {
	o.SetEndpointID(endpointID)
	return o
}

// SetEndpointID adds the endpointId to the get test suite result using g e t params
func (o *GetTestSuiteResultUsingGETParams) SetEndpointID(endpointID *int64) {
	o.EndpointID = endpointID
}

// WithEnvironment adds the environment to the get test suite result using g e t params
func (o *GetTestSuiteResultUsingGETParams) WithEnvironment(environment *int64) *GetTestSuiteResultUsingGETParams {
	o.SetEnvironment(environment)
	return o
}

// SetEnvironment adds the environment to the get test suite result using g e t params
func (o *GetTestSuiteResultUsingGETParams) SetEnvironment(environment *int64) {
	o.Environment = environment
}

// WithFilter adds the filter to the get test suite result using g e t params
func (o *GetTestSuiteResultUsingGETParams) WithFilter(filter *string) *GetTestSuiteResultUsingGETParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the get test suite result using g e t params
func (o *GetTestSuiteResultUsingGETParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithPageNumber adds the pageNumber to the get test suite result using g e t params
func (o *GetTestSuiteResultUsingGETParams) WithPageNumber(pageNumber *int32) *GetTestSuiteResultUsingGETParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get test suite result using g e t params
func (o *GetTestSuiteResultUsingGETParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get test suite result using g e t params
func (o *GetTestSuiteResultUsingGETParams) WithPageSize(pageSize *int32) *GetTestSuiteResultUsingGETParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get test suite result using g e t params
func (o *GetTestSuiteResultUsingGETParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithPluginID adds the pluginID to the get test suite result using g e t params
func (o *GetTestSuiteResultUsingGETParams) WithPluginID(pluginID *int64) *GetTestSuiteResultUsingGETParams {
	o.SetPluginID(pluginID)
	return o
}

// SetPluginID adds the pluginId to the get test suite result using g e t params
func (o *GetTestSuiteResultUsingGETParams) SetPluginID(pluginID *int64) {
	o.PluginID = pluginID
}

// WithSortDirection adds the sortDirection to the get test suite result using g e t params
func (o *GetTestSuiteResultUsingGETParams) WithSortDirection(sortDirection *string) *GetTestSuiteResultUsingGETParams {
	o.SetSortDirection(sortDirection)
	return o
}

// SetSortDirection adds the sortDirection to the get test suite result using g e t params
func (o *GetTestSuiteResultUsingGETParams) SetSortDirection(sortDirection *string) {
	o.SortDirection = sortDirection
}

// WithSortField adds the sortField to the get test suite result using g e t params
func (o *GetTestSuiteResultUsingGETParams) WithSortField(sortField *string) *GetTestSuiteResultUsingGETParams {
	o.SetSortField(sortField)
	return o
}

// SetSortField adds the sortField to the get test suite result using g e t params
func (o *GetTestSuiteResultUsingGETParams) SetSortField(sortField *string) {
	o.SortField = sortField
}

// WithTag adds the tag to the get test suite result using g e t params
func (o *GetTestSuiteResultUsingGETParams) WithTag(tag *string) *GetTestSuiteResultUsingGETParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the get test suite result using g e t params
func (o *GetTestSuiteResultUsingGETParams) SetTag(tag *string) {
	o.Tag = tag
}

// WithTestCaseResultStatus adds the testCaseResultStatus to the get test suite result using g e t params
func (o *GetTestSuiteResultUsingGETParams) WithTestCaseResultStatus(testCaseResultStatus []string) *GetTestSuiteResultUsingGETParams {
	o.SetTestCaseResultStatus(testCaseResultStatus)
	return o
}

// SetTestCaseResultStatus adds the testCaseResultStatus to the get test suite result using g e t params
func (o *GetTestSuiteResultUsingGETParams) SetTestCaseResultStatus(testCaseResultStatus []string) {
	o.TestCaseResultStatus = testCaseResultStatus
}

// WithTestSuiteExecutionID adds the testSuiteExecutionID to the get test suite result using g e t params
func (o *GetTestSuiteResultUsingGETParams) WithTestSuiteExecutionID(testSuiteExecutionID *string) *GetTestSuiteResultUsingGETParams {
	o.SetTestSuiteExecutionID(testSuiteExecutionID)
	return o
}

// SetTestSuiteExecutionID adds the testSuiteExecutionId to the get test suite result using g e t params
func (o *GetTestSuiteResultUsingGETParams) SetTestSuiteExecutionID(testSuiteExecutionID *string) {
	o.TestSuiteExecutionID = testSuiteExecutionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetTestSuiteResultUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param applicationId
	if err := r.SetPathParam("applicationId", swag.FormatInt64(o.ApplicationID)); err != nil {
		return err
	}

	// path param applicationVersionId
	if err := r.SetPathParam("applicationVersionId", swag.FormatInt64(o.ApplicationVersionID)); err != nil {
		return err
	}

	valuesEmbed := o.Embed

	joinedEmbed := swag.JoinByFormat(valuesEmbed, "multi")
	// query array param embed
	if err := r.SetQueryParam("embed", joinedEmbed...); err != nil {
		return err
	}

	if o.EndpointID != nil {

		// query param endpoint_id
		var qrEndpointID int64
		if o.EndpointID != nil {
			qrEndpointID = *o.EndpointID
		}
		qEndpointID := swag.FormatInt64(qrEndpointID)
		if qEndpointID != "" {
			if err := r.SetQueryParam("endpoint_id", qEndpointID); err != nil {
				return err
			}
		}

	}

	if o.Environment != nil {

		// query param environment
		var qrEnvironment int64
		if o.Environment != nil {
			qrEnvironment = *o.Environment
		}
		qEnvironment := swag.FormatInt64(qrEnvironment)
		if qEnvironment != "" {
			if err := r.SetQueryParam("environment", qEnvironment); err != nil {
				return err
			}
		}

	}

	if o.Filter != nil {

		// query param filter
		var qrFilter string
		if o.Filter != nil {
			qrFilter = *o.Filter
		}
		qFilter := qrFilter
		if qFilter != "" {
			if err := r.SetQueryParam("filter", qFilter); err != nil {
				return err
			}
		}

	}

	if o.PageNumber != nil {

		// query param page_number
		var qrPageNumber int32
		if o.PageNumber != nil {
			qrPageNumber = *o.PageNumber
		}
		qPageNumber := swag.FormatInt32(qrPageNumber)
		if qPageNumber != "" {
			if err := r.SetQueryParam("page_number", qPageNumber); err != nil {
				return err
			}
		}

	}

	if o.PageSize != nil {

		// query param page_size
		var qrPageSize int32
		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt32(qrPageSize)
		if qPageSize != "" {
			if err := r.SetQueryParam("page_size", qPageSize); err != nil {
				return err
			}
		}

	}

	if o.PluginID != nil {

		// query param plugin_id
		var qrPluginID int64
		if o.PluginID != nil {
			qrPluginID = *o.PluginID
		}
		qPluginID := swag.FormatInt64(qrPluginID)
		if qPluginID != "" {
			if err := r.SetQueryParam("plugin_id", qPluginID); err != nil {
				return err
			}
		}

	}

	if o.SortDirection != nil {

		// query param sort_direction
		var qrSortDirection string
		if o.SortDirection != nil {
			qrSortDirection = *o.SortDirection
		}
		qSortDirection := qrSortDirection
		if qSortDirection != "" {
			if err := r.SetQueryParam("sort_direction", qSortDirection); err != nil {
				return err
			}
		}

	}

	if o.SortField != nil {

		// query param sort_field
		var qrSortField string
		if o.SortField != nil {
			qrSortField = *o.SortField
		}
		qSortField := qrSortField
		if qSortField != "" {
			if err := r.SetQueryParam("sort_field", qSortField); err != nil {
				return err
			}
		}

	}

	if o.Tag != nil {

		// query param tag
		var qrTag string
		if o.Tag != nil {
			qrTag = *o.Tag
		}
		qTag := qrTag
		if qTag != "" {
			if err := r.SetQueryParam("tag", qTag); err != nil {
				return err
			}
		}

	}

	valuesTestCaseResultStatus := o.TestCaseResultStatus

	joinedTestCaseResultStatus := swag.JoinByFormat(valuesTestCaseResultStatus, "multi")
	// query array param test_case_result_status
	if err := r.SetQueryParam("test_case_result_status", joinedTestCaseResultStatus...); err != nil {
		return err
	}

	if o.TestSuiteExecutionID != nil {

		// query param test_suite_execution_id
		var qrTestSuiteExecutionID string
		if o.TestSuiteExecutionID != nil {
			qrTestSuiteExecutionID = *o.TestSuiteExecutionID
		}
		qTestSuiteExecutionID := qrTestSuiteExecutionID
		if qTestSuiteExecutionID != "" {
			if err := r.SetQueryParam("test_suite_execution_id", qTestSuiteExecutionID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}