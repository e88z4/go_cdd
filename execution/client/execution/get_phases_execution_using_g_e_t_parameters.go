// Code generated by go-swagger; DO NOT EDIT.

package execution

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

// NewGetPhasesExecutionUsingGETParams creates a new GetPhasesExecutionUsingGETParams object
// with the default values initialized.
func NewGetPhasesExecutionUsingGETParams() *GetPhasesExecutionUsingGETParams {
	var ()
	return &GetPhasesExecutionUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPhasesExecutionUsingGETParamsWithTimeout creates a new GetPhasesExecutionUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPhasesExecutionUsingGETParamsWithTimeout(timeout time.Duration) *GetPhasesExecutionUsingGETParams {
	var ()
	return &GetPhasesExecutionUsingGETParams{

		timeout: timeout,
	}
}

// NewGetPhasesExecutionUsingGETParamsWithContext creates a new GetPhasesExecutionUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPhasesExecutionUsingGETParamsWithContext(ctx context.Context) *GetPhasesExecutionUsingGETParams {
	var ()
	return &GetPhasesExecutionUsingGETParams{

		Context: ctx,
	}
}

// NewGetPhasesExecutionUsingGETParamsWithHTTPClient creates a new GetPhasesExecutionUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPhasesExecutionUsingGETParamsWithHTTPClient(client *http.Client) *GetPhasesExecutionUsingGETParams {
	var ()
	return &GetPhasesExecutionUsingGETParams{
		HTTPClient: client,
	}
}

/*GetPhasesExecutionUsingGETParams contains all the parameters to send to the API endpoint
for the get phases execution using g e t operation typically these are written to a http.Request
*/
type GetPhasesExecutionUsingGETParams struct {

	/*ReleaseID
	  releaseId

	*/
	ReleaseID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get phases execution using g e t params
func (o *GetPhasesExecutionUsingGETParams) WithTimeout(timeout time.Duration) *GetPhasesExecutionUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get phases execution using g e t params
func (o *GetPhasesExecutionUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get phases execution using g e t params
func (o *GetPhasesExecutionUsingGETParams) WithContext(ctx context.Context) *GetPhasesExecutionUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get phases execution using g e t params
func (o *GetPhasesExecutionUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get phases execution using g e t params
func (o *GetPhasesExecutionUsingGETParams) WithHTTPClient(client *http.Client) *GetPhasesExecutionUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get phases execution using g e t params
func (o *GetPhasesExecutionUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithReleaseID adds the releaseID to the get phases execution using g e t params
func (o *GetPhasesExecutionUsingGETParams) WithReleaseID(releaseID int64) *GetPhasesExecutionUsingGETParams {
	o.SetReleaseID(releaseID)
	return o
}

// SetReleaseID adds the releaseId to the get phases execution using g e t params
func (o *GetPhasesExecutionUsingGETParams) SetReleaseID(releaseID int64) {
	o.ReleaseID = releaseID
}

// WriteToRequest writes these params to a swagger request
func (o *GetPhasesExecutionUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param releaseId
	if err := r.SetPathParam("releaseId", swag.FormatInt64(o.ReleaseID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
