// Code generated by go-swagger; DO NOT EDIT.

package application

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

// NewGetApplicationReportUsingGETParams creates a new GetApplicationReportUsingGETParams object
// with the default values initialized.
func NewGetApplicationReportUsingGETParams() *GetApplicationReportUsingGETParams {
	var ()
	return &GetApplicationReportUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetApplicationReportUsingGETParamsWithTimeout creates a new GetApplicationReportUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetApplicationReportUsingGETParamsWithTimeout(timeout time.Duration) *GetApplicationReportUsingGETParams {
	var ()
	return &GetApplicationReportUsingGETParams{

		timeout: timeout,
	}
}

// NewGetApplicationReportUsingGETParamsWithContext creates a new GetApplicationReportUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetApplicationReportUsingGETParamsWithContext(ctx context.Context) *GetApplicationReportUsingGETParams {
	var ()
	return &GetApplicationReportUsingGETParams{

		Context: ctx,
	}
}

// NewGetApplicationReportUsingGETParamsWithHTTPClient creates a new GetApplicationReportUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetApplicationReportUsingGETParamsWithHTTPClient(client *http.Client) *GetApplicationReportUsingGETParams {
	var ()
	return &GetApplicationReportUsingGETParams{
		HTTPClient: client,
	}
}

/*GetApplicationReportUsingGETParams contains all the parameters to send to the API endpoint
for the get application report using g e t operation typically these are written to a http.Request
*/
type GetApplicationReportUsingGETParams struct {

	/*ApplicationID
	  applicationId

	*/
	ApplicationID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get application report using g e t params
func (o *GetApplicationReportUsingGETParams) WithTimeout(timeout time.Duration) *GetApplicationReportUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get application report using g e t params
func (o *GetApplicationReportUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get application report using g e t params
func (o *GetApplicationReportUsingGETParams) WithContext(ctx context.Context) *GetApplicationReportUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get application report using g e t params
func (o *GetApplicationReportUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get application report using g e t params
func (o *GetApplicationReportUsingGETParams) WithHTTPClient(client *http.Client) *GetApplicationReportUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get application report using g e t params
func (o *GetApplicationReportUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApplicationID adds the applicationID to the get application report using g e t params
func (o *GetApplicationReportUsingGETParams) WithApplicationID(applicationID int64) *GetApplicationReportUsingGETParams {
	o.SetApplicationID(applicationID)
	return o
}

// SetApplicationID adds the applicationId to the get application report using g e t params
func (o *GetApplicationReportUsingGETParams) SetApplicationID(applicationID int64) {
	o.ApplicationID = applicationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetApplicationReportUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param applicationId
	if err := r.SetPathParam("applicationId", swag.FormatInt64(o.ApplicationID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
