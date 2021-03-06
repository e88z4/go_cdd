// Code generated by go-swagger; DO NOT EDIT.

package endpoint

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

// NewUpdateEndpointUsingPUTParams creates a new UpdateEndpointUsingPUTParams object
// with the default values initialized.
func NewUpdateEndpointUsingPUTParams() *UpdateEndpointUsingPUTParams {
	var ()
	return &UpdateEndpointUsingPUTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateEndpointUsingPUTParamsWithTimeout creates a new UpdateEndpointUsingPUTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateEndpointUsingPUTParamsWithTimeout(timeout time.Duration) *UpdateEndpointUsingPUTParams {
	var ()
	return &UpdateEndpointUsingPUTParams{

		timeout: timeout,
	}
}

// NewUpdateEndpointUsingPUTParamsWithContext creates a new UpdateEndpointUsingPUTParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateEndpointUsingPUTParamsWithContext(ctx context.Context) *UpdateEndpointUsingPUTParams {
	var ()
	return &UpdateEndpointUsingPUTParams{

		Context: ctx,
	}
}

// NewUpdateEndpointUsingPUTParamsWithHTTPClient creates a new UpdateEndpointUsingPUTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateEndpointUsingPUTParamsWithHTTPClient(client *http.Client) *UpdateEndpointUsingPUTParams {
	var ()
	return &UpdateEndpointUsingPUTParams{
		HTTPClient: client,
	}
}

/*UpdateEndpointUsingPUTParams contains all the parameters to send to the API endpoint
for the update endpoint using p u t operation typically these are written to a http.Request
*/
type UpdateEndpointUsingPUTParams struct {

	/*EndpointDto
	  endpointDto

	*/
	EndpointDto *models.EndpointDto
	/*EndpointID
	  endpointId

	*/
	EndpointID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update endpoint using p u t params
func (o *UpdateEndpointUsingPUTParams) WithTimeout(timeout time.Duration) *UpdateEndpointUsingPUTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update endpoint using p u t params
func (o *UpdateEndpointUsingPUTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update endpoint using p u t params
func (o *UpdateEndpointUsingPUTParams) WithContext(ctx context.Context) *UpdateEndpointUsingPUTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update endpoint using p u t params
func (o *UpdateEndpointUsingPUTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update endpoint using p u t params
func (o *UpdateEndpointUsingPUTParams) WithHTTPClient(client *http.Client) *UpdateEndpointUsingPUTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update endpoint using p u t params
func (o *UpdateEndpointUsingPUTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndpointDto adds the endpointDto to the update endpoint using p u t params
func (o *UpdateEndpointUsingPUTParams) WithEndpointDto(endpointDto *models.EndpointDto) *UpdateEndpointUsingPUTParams {
	o.SetEndpointDto(endpointDto)
	return o
}

// SetEndpointDto adds the endpointDto to the update endpoint using p u t params
func (o *UpdateEndpointUsingPUTParams) SetEndpointDto(endpointDto *models.EndpointDto) {
	o.EndpointDto = endpointDto
}

// WithEndpointID adds the endpointID to the update endpoint using p u t params
func (o *UpdateEndpointUsingPUTParams) WithEndpointID(endpointID int64) *UpdateEndpointUsingPUTParams {
	o.SetEndpointID(endpointID)
	return o
}

// SetEndpointID adds the endpointId to the update endpoint using p u t params
func (o *UpdateEndpointUsingPUTParams) SetEndpointID(endpointID int64) {
	o.EndpointID = endpointID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateEndpointUsingPUTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.EndpointDto != nil {
		if err := r.SetBodyParam(o.EndpointDto); err != nil {
			return err
		}
	}

	// path param endpointId
	if err := r.SetPathParam("endpointId", swag.FormatInt64(o.EndpointID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
