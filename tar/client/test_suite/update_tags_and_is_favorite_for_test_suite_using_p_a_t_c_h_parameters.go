// Code generated by go-swagger; DO NOT EDIT.

package test_suite

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

	"github.com/e88z4/go_cdd/tar/models"
)

// NewUpdateTagsAndIsFavoriteForTestSuiteUsingPATCHParams creates a new UpdateTagsAndIsFavoriteForTestSuiteUsingPATCHParams object
// with the default values initialized.
func NewUpdateTagsAndIsFavoriteForTestSuiteUsingPATCHParams() *UpdateTagsAndIsFavoriteForTestSuiteUsingPATCHParams {
	var ()
	return &UpdateTagsAndIsFavoriteForTestSuiteUsingPATCHParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateTagsAndIsFavoriteForTestSuiteUsingPATCHParamsWithTimeout creates a new UpdateTagsAndIsFavoriteForTestSuiteUsingPATCHParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateTagsAndIsFavoriteForTestSuiteUsingPATCHParamsWithTimeout(timeout time.Duration) *UpdateTagsAndIsFavoriteForTestSuiteUsingPATCHParams {
	var ()
	return &UpdateTagsAndIsFavoriteForTestSuiteUsingPATCHParams{

		timeout: timeout,
	}
}

// NewUpdateTagsAndIsFavoriteForTestSuiteUsingPATCHParamsWithContext creates a new UpdateTagsAndIsFavoriteForTestSuiteUsingPATCHParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateTagsAndIsFavoriteForTestSuiteUsingPATCHParamsWithContext(ctx context.Context) *UpdateTagsAndIsFavoriteForTestSuiteUsingPATCHParams {
	var ()
	return &UpdateTagsAndIsFavoriteForTestSuiteUsingPATCHParams{

		Context: ctx,
	}
}

// NewUpdateTagsAndIsFavoriteForTestSuiteUsingPATCHParamsWithHTTPClient creates a new UpdateTagsAndIsFavoriteForTestSuiteUsingPATCHParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateTagsAndIsFavoriteForTestSuiteUsingPATCHParamsWithHTTPClient(client *http.Client) *UpdateTagsAndIsFavoriteForTestSuiteUsingPATCHParams {
	var ()
	return &UpdateTagsAndIsFavoriteForTestSuiteUsingPATCHParams{
		HTTPClient: client,
	}
}

/*UpdateTagsAndIsFavoriteForTestSuiteUsingPATCHParams contains all the parameters to send to the API endpoint
for the update tags and is favorite for test suite using p a t c h operation typically these are written to a http.Request
*/
type UpdateTagsAndIsFavoriteForTestSuiteUsingPATCHParams struct {

	/*ApplicationID
	  applicationId

	*/
	ApplicationID int64
	/*ApplicationVersionID
	  applicationVersionId

	*/
	ApplicationVersionID int64
	/*TestSuiteDto
	  testSuiteDto

	*/
	TestSuiteDto *models.TestSuiteDto
	/*TestSuiteHexID
	  testSuiteId

	*/
	TestSuiteHexID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update tags and is favorite for test suite using p a t c h params
func (o *UpdateTagsAndIsFavoriteForTestSuiteUsingPATCHParams) WithTimeout(timeout time.Duration) *UpdateTagsAndIsFavoriteForTestSuiteUsingPATCHParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update tags and is favorite for test suite using p a t c h params
func (o *UpdateTagsAndIsFavoriteForTestSuiteUsingPATCHParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update tags and is favorite for test suite using p a t c h params
func (o *UpdateTagsAndIsFavoriteForTestSuiteUsingPATCHParams) WithContext(ctx context.Context) *UpdateTagsAndIsFavoriteForTestSuiteUsingPATCHParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update tags and is favorite for test suite using p a t c h params
func (o *UpdateTagsAndIsFavoriteForTestSuiteUsingPATCHParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update tags and is favorite for test suite using p a t c h params
func (o *UpdateTagsAndIsFavoriteForTestSuiteUsingPATCHParams) WithHTTPClient(client *http.Client) *UpdateTagsAndIsFavoriteForTestSuiteUsingPATCHParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update tags and is favorite for test suite using p a t c h params
func (o *UpdateTagsAndIsFavoriteForTestSuiteUsingPATCHParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApplicationID adds the applicationID to the update tags and is favorite for test suite using p a t c h params
func (o *UpdateTagsAndIsFavoriteForTestSuiteUsingPATCHParams) WithApplicationID(applicationID int64) *UpdateTagsAndIsFavoriteForTestSuiteUsingPATCHParams {
	o.SetApplicationID(applicationID)
	return o
}

// SetApplicationID adds the applicationId to the update tags and is favorite for test suite using p a t c h params
func (o *UpdateTagsAndIsFavoriteForTestSuiteUsingPATCHParams) SetApplicationID(applicationID int64) {
	o.ApplicationID = applicationID
}

// WithApplicationVersionID adds the applicationVersionID to the update tags and is favorite for test suite using p a t c h params
func (o *UpdateTagsAndIsFavoriteForTestSuiteUsingPATCHParams) WithApplicationVersionID(applicationVersionID int64) *UpdateTagsAndIsFavoriteForTestSuiteUsingPATCHParams {
	o.SetApplicationVersionID(applicationVersionID)
	return o
}

// SetApplicationVersionID adds the applicationVersionId to the update tags and is favorite for test suite using p a t c h params
func (o *UpdateTagsAndIsFavoriteForTestSuiteUsingPATCHParams) SetApplicationVersionID(applicationVersionID int64) {
	o.ApplicationVersionID = applicationVersionID
}

// WithTestSuiteDto adds the testSuiteDto to the update tags and is favorite for test suite using p a t c h params
func (o *UpdateTagsAndIsFavoriteForTestSuiteUsingPATCHParams) WithTestSuiteDto(testSuiteDto *models.TestSuiteDto) *UpdateTagsAndIsFavoriteForTestSuiteUsingPATCHParams {
	o.SetTestSuiteDto(testSuiteDto)
	return o
}

// SetTestSuiteDto adds the testSuiteDto to the update tags and is favorite for test suite using p a t c h params
func (o *UpdateTagsAndIsFavoriteForTestSuiteUsingPATCHParams) SetTestSuiteDto(testSuiteDto *models.TestSuiteDto) {
	o.TestSuiteDto = testSuiteDto
}

// WithTestSuiteHexID adds the testSuiteHexID to the update tags and is favorite for test suite using p a t c h params
func (o *UpdateTagsAndIsFavoriteForTestSuiteUsingPATCHParams) WithTestSuiteHexID(testSuiteHexID string) *UpdateTagsAndIsFavoriteForTestSuiteUsingPATCHParams {
	o.SetTestSuiteHexID(testSuiteHexID)
	return o
}

// SetTestSuiteHexID adds the testSuiteHexId to the update tags and is favorite for test suite using p a t c h params
func (o *UpdateTagsAndIsFavoriteForTestSuiteUsingPATCHParams) SetTestSuiteHexID(testSuiteHexID string) {
	o.TestSuiteHexID = testSuiteHexID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateTagsAndIsFavoriteForTestSuiteUsingPATCHParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.TestSuiteDto != nil {
		if err := r.SetBodyParam(o.TestSuiteDto); err != nil {
			return err
		}
	}

	// path param testSuiteHexId
	if err := r.SetPathParam("testSuiteHexId", o.TestSuiteHexID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}