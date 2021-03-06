// Code generated by go-swagger; DO NOT EDIT.

package freeze_period

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

	"github.com/e88z4/go_cdd/administration/models"
)

// NewUpdateFreezePeriodUsingPUTParams creates a new UpdateFreezePeriodUsingPUTParams object
// with the default values initialized.
func NewUpdateFreezePeriodUsingPUTParams() *UpdateFreezePeriodUsingPUTParams {
	var ()
	return &UpdateFreezePeriodUsingPUTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateFreezePeriodUsingPUTParamsWithTimeout creates a new UpdateFreezePeriodUsingPUTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateFreezePeriodUsingPUTParamsWithTimeout(timeout time.Duration) *UpdateFreezePeriodUsingPUTParams {
	var ()
	return &UpdateFreezePeriodUsingPUTParams{

		timeout: timeout,
	}
}

// NewUpdateFreezePeriodUsingPUTParamsWithContext creates a new UpdateFreezePeriodUsingPUTParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateFreezePeriodUsingPUTParamsWithContext(ctx context.Context) *UpdateFreezePeriodUsingPUTParams {
	var ()
	return &UpdateFreezePeriodUsingPUTParams{

		Context: ctx,
	}
}

// NewUpdateFreezePeriodUsingPUTParamsWithHTTPClient creates a new UpdateFreezePeriodUsingPUTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateFreezePeriodUsingPUTParamsWithHTTPClient(client *http.Client) *UpdateFreezePeriodUsingPUTParams {
	var ()
	return &UpdateFreezePeriodUsingPUTParams{
		HTTPClient: client,
	}
}

/*UpdateFreezePeriodUsingPUTParams contains all the parameters to send to the API endpoint
for the update freeze period using p u t operation typically these are written to a http.Request
*/
type UpdateFreezePeriodUsingPUTParams struct {

	/*FreezePeriodDto
	  freezePeriodDto

	*/
	FreezePeriodDto *models.FreezePeriodDto
	/*FreezePeriodID
	  freezePeriodId

	*/
	FreezePeriodID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update freeze period using p u t params
func (o *UpdateFreezePeriodUsingPUTParams) WithTimeout(timeout time.Duration) *UpdateFreezePeriodUsingPUTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update freeze period using p u t params
func (o *UpdateFreezePeriodUsingPUTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update freeze period using p u t params
func (o *UpdateFreezePeriodUsingPUTParams) WithContext(ctx context.Context) *UpdateFreezePeriodUsingPUTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update freeze period using p u t params
func (o *UpdateFreezePeriodUsingPUTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update freeze period using p u t params
func (o *UpdateFreezePeriodUsingPUTParams) WithHTTPClient(client *http.Client) *UpdateFreezePeriodUsingPUTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update freeze period using p u t params
func (o *UpdateFreezePeriodUsingPUTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFreezePeriodDto adds the freezePeriodDto to the update freeze period using p u t params
func (o *UpdateFreezePeriodUsingPUTParams) WithFreezePeriodDto(freezePeriodDto *models.FreezePeriodDto) *UpdateFreezePeriodUsingPUTParams {
	o.SetFreezePeriodDto(freezePeriodDto)
	return o
}

// SetFreezePeriodDto adds the freezePeriodDto to the update freeze period using p u t params
func (o *UpdateFreezePeriodUsingPUTParams) SetFreezePeriodDto(freezePeriodDto *models.FreezePeriodDto) {
	o.FreezePeriodDto = freezePeriodDto
}

// WithFreezePeriodID adds the freezePeriodID to the update freeze period using p u t params
func (o *UpdateFreezePeriodUsingPUTParams) WithFreezePeriodID(freezePeriodID int64) *UpdateFreezePeriodUsingPUTParams {
	o.SetFreezePeriodID(freezePeriodID)
	return o
}

// SetFreezePeriodID adds the freezePeriodId to the update freeze period using p u t params
func (o *UpdateFreezePeriodUsingPUTParams) SetFreezePeriodID(freezePeriodID int64) {
	o.FreezePeriodID = freezePeriodID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateFreezePeriodUsingPUTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FreezePeriodDto != nil {
		if err := r.SetBodyParam(o.FreezePeriodDto); err != nil {
			return err
		}
	}

	// path param freezePeriodId
	if err := r.SetPathParam("freezePeriodId", swag.FormatInt64(o.FreezePeriodID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
