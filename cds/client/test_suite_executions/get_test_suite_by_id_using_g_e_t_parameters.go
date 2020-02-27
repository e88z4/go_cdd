// Code generated by go-swagger; DO NOT EDIT.

package test_suite_executions

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

// NewGetTestSuiteByIDUsingGETParams creates a new GetTestSuiteByIDUsingGETParams object
// with the default values initialized.
func NewGetTestSuiteByIDUsingGETParams() *GetTestSuiteByIDUsingGETParams {
	var ()
	return &GetTestSuiteByIDUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTestSuiteByIDUsingGETParamsWithTimeout creates a new GetTestSuiteByIDUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTestSuiteByIDUsingGETParamsWithTimeout(timeout time.Duration) *GetTestSuiteByIDUsingGETParams {
	var ()
	return &GetTestSuiteByIDUsingGETParams{

		timeout: timeout,
	}
}

// NewGetTestSuiteByIDUsingGETParamsWithContext creates a new GetTestSuiteByIDUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTestSuiteByIDUsingGETParamsWithContext(ctx context.Context) *GetTestSuiteByIDUsingGETParams {
	var ()
	return &GetTestSuiteByIDUsingGETParams{

		Context: ctx,
	}
}

// NewGetTestSuiteByIDUsingGETParamsWithHTTPClient creates a new GetTestSuiteByIDUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTestSuiteByIDUsingGETParamsWithHTTPClient(client *http.Client) *GetTestSuiteByIDUsingGETParams {
	var ()
	return &GetTestSuiteByIDUsingGETParams{
		HTTPClient: client,
	}
}

/*GetTestSuiteByIDUsingGETParams contains all the parameters to send to the API endpoint
for the get test suite by Id using g e t operation typically these are written to a http.Request
*/
type GetTestSuiteByIDUsingGETParams struct {

	/*ApplicationID
	  applicationId

	*/
	ApplicationID int64
	/*ApplicationVersionID
	  applicationVersionId

	*/
	ApplicationVersionID int64
	/*TestSuiteExecutionsHexID
	  testSuiteExecutionId

	*/
	TestSuiteExecutionsHexID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get test suite by Id using g e t params
func (o *GetTestSuiteByIDUsingGETParams) WithTimeout(timeout time.Duration) *GetTestSuiteByIDUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get test suite by Id using g e t params
func (o *GetTestSuiteByIDUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get test suite by Id using g e t params
func (o *GetTestSuiteByIDUsingGETParams) WithContext(ctx context.Context) *GetTestSuiteByIDUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get test suite by Id using g e t params
func (o *GetTestSuiteByIDUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get test suite by Id using g e t params
func (o *GetTestSuiteByIDUsingGETParams) WithHTTPClient(client *http.Client) *GetTestSuiteByIDUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get test suite by Id using g e t params
func (o *GetTestSuiteByIDUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApplicationID adds the applicationID to the get test suite by Id using g e t params
func (o *GetTestSuiteByIDUsingGETParams) WithApplicationID(applicationID int64) *GetTestSuiteByIDUsingGETParams {
	o.SetApplicationID(applicationID)
	return o
}

// SetApplicationID adds the applicationId to the get test suite by Id using g e t params
func (o *GetTestSuiteByIDUsingGETParams) SetApplicationID(applicationID int64) {
	o.ApplicationID = applicationID
}

// WithApplicationVersionID adds the applicationVersionID to the get test suite by Id using g e t params
func (o *GetTestSuiteByIDUsingGETParams) WithApplicationVersionID(applicationVersionID int64) *GetTestSuiteByIDUsingGETParams {
	o.SetApplicationVersionID(applicationVersionID)
	return o
}

// SetApplicationVersionID adds the applicationVersionId to the get test suite by Id using g e t params
func (o *GetTestSuiteByIDUsingGETParams) SetApplicationVersionID(applicationVersionID int64) {
	o.ApplicationVersionID = applicationVersionID
}

// WithTestSuiteExecutionsHexID adds the testSuiteExecutionsHexID to the get test suite by Id using g e t params
func (o *GetTestSuiteByIDUsingGETParams) WithTestSuiteExecutionsHexID(testSuiteExecutionsHexID string) *GetTestSuiteByIDUsingGETParams {
	o.SetTestSuiteExecutionsHexID(testSuiteExecutionsHexID)
	return o
}

// SetTestSuiteExecutionsHexID adds the testSuiteExecutionsHexId to the get test suite by Id using g e t params
func (o *GetTestSuiteByIDUsingGETParams) SetTestSuiteExecutionsHexID(testSuiteExecutionsHexID string) {
	o.TestSuiteExecutionsHexID = testSuiteExecutionsHexID
}

// WriteToRequest writes these params to a swagger request
func (o *GetTestSuiteByIDUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param testSuiteExecutionsHexId
	if err := r.SetPathParam("testSuiteExecutionsHexId", o.TestSuiteExecutionsHexID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}