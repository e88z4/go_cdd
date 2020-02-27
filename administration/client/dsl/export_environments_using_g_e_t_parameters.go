// Code generated by go-swagger; DO NOT EDIT.

package dsl

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

// NewExportEnvironmentsUsingGETParams creates a new ExportEnvironmentsUsingGETParams object
// with the default values initialized.
func NewExportEnvironmentsUsingGETParams() *ExportEnvironmentsUsingGETParams {
	var ()
	return &ExportEnvironmentsUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewExportEnvironmentsUsingGETParamsWithTimeout creates a new ExportEnvironmentsUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewExportEnvironmentsUsingGETParamsWithTimeout(timeout time.Duration) *ExportEnvironmentsUsingGETParams {
	var ()
	return &ExportEnvironmentsUsingGETParams{

		timeout: timeout,
	}
}

// NewExportEnvironmentsUsingGETParamsWithContext creates a new ExportEnvironmentsUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewExportEnvironmentsUsingGETParamsWithContext(ctx context.Context) *ExportEnvironmentsUsingGETParams {
	var ()
	return &ExportEnvironmentsUsingGETParams{

		Context: ctx,
	}
}

// NewExportEnvironmentsUsingGETParamsWithHTTPClient creates a new ExportEnvironmentsUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewExportEnvironmentsUsingGETParamsWithHTTPClient(client *http.Client) *ExportEnvironmentsUsingGETParams {
	var ()
	return &ExportEnvironmentsUsingGETParams{
		HTTPClient: client,
	}
}

/*ExportEnvironmentsUsingGETParams contains all the parameters to send to the API endpoint
for the export environments using g e t operation typically these are written to a http.Request
*/
type ExportEnvironmentsUsingGETParams struct {

	/*DispositionType
	  dispositionTypeDto

	*/
	DispositionType *string
	/*Environment
	  environmentIds

	*/
	Environment []int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the export environments using g e t params
func (o *ExportEnvironmentsUsingGETParams) WithTimeout(timeout time.Duration) *ExportEnvironmentsUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the export environments using g e t params
func (o *ExportEnvironmentsUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the export environments using g e t params
func (o *ExportEnvironmentsUsingGETParams) WithContext(ctx context.Context) *ExportEnvironmentsUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the export environments using g e t params
func (o *ExportEnvironmentsUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the export environments using g e t params
func (o *ExportEnvironmentsUsingGETParams) WithHTTPClient(client *http.Client) *ExportEnvironmentsUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the export environments using g e t params
func (o *ExportEnvironmentsUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDispositionType adds the dispositionType to the export environments using g e t params
func (o *ExportEnvironmentsUsingGETParams) WithDispositionType(dispositionType *string) *ExportEnvironmentsUsingGETParams {
	o.SetDispositionType(dispositionType)
	return o
}

// SetDispositionType adds the dispositionType to the export environments using g e t params
func (o *ExportEnvironmentsUsingGETParams) SetDispositionType(dispositionType *string) {
	o.DispositionType = dispositionType
}

// WithEnvironment adds the environment to the export environments using g e t params
func (o *ExportEnvironmentsUsingGETParams) WithEnvironment(environment []int64) *ExportEnvironmentsUsingGETParams {
	o.SetEnvironment(environment)
	return o
}

// SetEnvironment adds the environment to the export environments using g e t params
func (o *ExportEnvironmentsUsingGETParams) SetEnvironment(environment []int64) {
	o.Environment = environment
}

// WriteToRequest writes these params to a swagger request
func (o *ExportEnvironmentsUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DispositionType != nil {

		// query param disposition_type
		var qrDispositionType string
		if o.DispositionType != nil {
			qrDispositionType = *o.DispositionType
		}
		qDispositionType := qrDispositionType
		if qDispositionType != "" {
			if err := r.SetQueryParam("disposition_type", qDispositionType); err != nil {
				return err
			}
		}

	}

	var valuesEnvironment []string
	for _, v := range o.Environment {
		valuesEnvironment = append(valuesEnvironment, swag.FormatInt64(v))
	}

	joinedEnvironment := swag.JoinByFormat(valuesEnvironment, "multi")
	// query array param environment
	if err := r.SetQueryParam("environment", joinedEnvironment...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
