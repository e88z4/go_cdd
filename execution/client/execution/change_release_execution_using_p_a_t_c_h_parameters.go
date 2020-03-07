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

	"github.com/e88z4/go_cdd/execution/models"
)

// NewChangeReleaseExecutionUsingPATCHParams creates a new ChangeReleaseExecutionUsingPATCHParams object
// with the default values initialized.
func NewChangeReleaseExecutionUsingPATCHParams() *ChangeReleaseExecutionUsingPATCHParams {
	var ()
	return &ChangeReleaseExecutionUsingPATCHParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewChangeReleaseExecutionUsingPATCHParamsWithTimeout creates a new ChangeReleaseExecutionUsingPATCHParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewChangeReleaseExecutionUsingPATCHParamsWithTimeout(timeout time.Duration) *ChangeReleaseExecutionUsingPATCHParams {
	var ()
	return &ChangeReleaseExecutionUsingPATCHParams{

		timeout: timeout,
	}
}

// NewChangeReleaseExecutionUsingPATCHParamsWithContext creates a new ChangeReleaseExecutionUsingPATCHParams object
// with the default values initialized, and the ability to set a context for a request
func NewChangeReleaseExecutionUsingPATCHParamsWithContext(ctx context.Context) *ChangeReleaseExecutionUsingPATCHParams {
	var ()
	return &ChangeReleaseExecutionUsingPATCHParams{

		Context: ctx,
	}
}

// NewChangeReleaseExecutionUsingPATCHParamsWithHTTPClient creates a new ChangeReleaseExecutionUsingPATCHParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewChangeReleaseExecutionUsingPATCHParamsWithHTTPClient(client *http.Client) *ChangeReleaseExecutionUsingPATCHParams {
	var ()
	return &ChangeReleaseExecutionUsingPATCHParams{
		HTTPClient: client,
	}
}

/*ChangeReleaseExecutionUsingPATCHParams contains all the parameters to send to the API endpoint
for the change release execution using p a t c h operation typically these are written to a http.Request
*/
type ChangeReleaseExecutionUsingPATCHParams struct {

	/*Force
	  isForce

	*/
	Force *bool
	/*ReleaseExecution
	  releaseExecution

	*/
	ReleaseExecution *models.ReleaseExecutionDto
	/*ReleaseID
	  releaseId

	*/
	ReleaseID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the change release execution using p a t c h params
func (o *ChangeReleaseExecutionUsingPATCHParams) WithTimeout(timeout time.Duration) *ChangeReleaseExecutionUsingPATCHParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the change release execution using p a t c h params
func (o *ChangeReleaseExecutionUsingPATCHParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the change release execution using p a t c h params
func (o *ChangeReleaseExecutionUsingPATCHParams) WithContext(ctx context.Context) *ChangeReleaseExecutionUsingPATCHParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the change release execution using p a t c h params
func (o *ChangeReleaseExecutionUsingPATCHParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the change release execution using p a t c h params
func (o *ChangeReleaseExecutionUsingPATCHParams) WithHTTPClient(client *http.Client) *ChangeReleaseExecutionUsingPATCHParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the change release execution using p a t c h params
func (o *ChangeReleaseExecutionUsingPATCHParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithForce adds the force to the change release execution using p a t c h params
func (o *ChangeReleaseExecutionUsingPATCHParams) WithForce(force *bool) *ChangeReleaseExecutionUsingPATCHParams {
	o.SetForce(force)
	return o
}

// SetForce adds the force to the change release execution using p a t c h params
func (o *ChangeReleaseExecutionUsingPATCHParams) SetForce(force *bool) {
	o.Force = force
}

// WithReleaseExecution adds the releaseExecution to the change release execution using p a t c h params
func (o *ChangeReleaseExecutionUsingPATCHParams) WithReleaseExecution(releaseExecution *models.ReleaseExecutionDto) *ChangeReleaseExecutionUsingPATCHParams {
	o.SetReleaseExecution(releaseExecution)
	return o
}

// SetReleaseExecution adds the releaseExecution to the change release execution using p a t c h params
func (o *ChangeReleaseExecutionUsingPATCHParams) SetReleaseExecution(releaseExecution *models.ReleaseExecutionDto) {
	o.ReleaseExecution = releaseExecution
}

// WithReleaseID adds the releaseID to the change release execution using p a t c h params
func (o *ChangeReleaseExecutionUsingPATCHParams) WithReleaseID(releaseID int64) *ChangeReleaseExecutionUsingPATCHParams {
	o.SetReleaseID(releaseID)
	return o
}

// SetReleaseID adds the releaseId to the change release execution using p a t c h params
func (o *ChangeReleaseExecutionUsingPATCHParams) SetReleaseID(releaseID int64) {
	o.ReleaseID = releaseID
}

// WriteToRequest writes these params to a swagger request
func (o *ChangeReleaseExecutionUsingPATCHParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Force != nil {

		// query param force
		var qrForce bool
		if o.Force != nil {
			qrForce = *o.Force
		}
		qForce := swag.FormatBool(qrForce)
		if qForce != "" {
			if err := r.SetQueryParam("force", qForce); err != nil {
				return err
			}
		}

	}

	if o.ReleaseExecution != nil {
		if err := r.SetBodyParam(o.ReleaseExecution); err != nil {
			return err
		}
	}

	// path param releaseId
	if err := r.SetPathParam("releaseId", swag.FormatInt64(o.ReleaseID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}