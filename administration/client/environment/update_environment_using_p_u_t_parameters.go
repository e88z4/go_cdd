// Code generated by go-swagger; DO NOT EDIT.

package environment

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

// NewUpdateEnvironmentUsingPUTParams creates a new UpdateEnvironmentUsingPUTParams object
// with the default values initialized.
func NewUpdateEnvironmentUsingPUTParams() *UpdateEnvironmentUsingPUTParams {
	var ()
	return &UpdateEnvironmentUsingPUTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateEnvironmentUsingPUTParamsWithTimeout creates a new UpdateEnvironmentUsingPUTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateEnvironmentUsingPUTParamsWithTimeout(timeout time.Duration) *UpdateEnvironmentUsingPUTParams {
	var ()
	return &UpdateEnvironmentUsingPUTParams{

		timeout: timeout,
	}
}

// NewUpdateEnvironmentUsingPUTParamsWithContext creates a new UpdateEnvironmentUsingPUTParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateEnvironmentUsingPUTParamsWithContext(ctx context.Context) *UpdateEnvironmentUsingPUTParams {
	var ()
	return &UpdateEnvironmentUsingPUTParams{

		Context: ctx,
	}
}

// NewUpdateEnvironmentUsingPUTParamsWithHTTPClient creates a new UpdateEnvironmentUsingPUTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateEnvironmentUsingPUTParamsWithHTTPClient(client *http.Client) *UpdateEnvironmentUsingPUTParams {
	var ()
	return &UpdateEnvironmentUsingPUTParams{
		HTTPClient: client,
	}
}

/*UpdateEnvironmentUsingPUTParams contains all the parameters to send to the API endpoint
for the update environment using p u t operation typically these are written to a http.Request
*/
type UpdateEnvironmentUsingPUTParams struct {

	/*EnvironmentDto
	  environmentDto

	*/
	EnvironmentDto *models.EnvironmentDto
	/*EnvironmentID
	  environmentId

	*/
	EnvironmentID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update environment using p u t params
func (o *UpdateEnvironmentUsingPUTParams) WithTimeout(timeout time.Duration) *UpdateEnvironmentUsingPUTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update environment using p u t params
func (o *UpdateEnvironmentUsingPUTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update environment using p u t params
func (o *UpdateEnvironmentUsingPUTParams) WithContext(ctx context.Context) *UpdateEnvironmentUsingPUTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update environment using p u t params
func (o *UpdateEnvironmentUsingPUTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update environment using p u t params
func (o *UpdateEnvironmentUsingPUTParams) WithHTTPClient(client *http.Client) *UpdateEnvironmentUsingPUTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update environment using p u t params
func (o *UpdateEnvironmentUsingPUTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnvironmentDto adds the environmentDto to the update environment using p u t params
func (o *UpdateEnvironmentUsingPUTParams) WithEnvironmentDto(environmentDto *models.EnvironmentDto) *UpdateEnvironmentUsingPUTParams {
	o.SetEnvironmentDto(environmentDto)
	return o
}

// SetEnvironmentDto adds the environmentDto to the update environment using p u t params
func (o *UpdateEnvironmentUsingPUTParams) SetEnvironmentDto(environmentDto *models.EnvironmentDto) {
	o.EnvironmentDto = environmentDto
}

// WithEnvironmentID adds the environmentID to the update environment using p u t params
func (o *UpdateEnvironmentUsingPUTParams) WithEnvironmentID(environmentID int64) *UpdateEnvironmentUsingPUTParams {
	o.SetEnvironmentID(environmentID)
	return o
}

// SetEnvironmentID adds the environmentId to the update environment using p u t params
func (o *UpdateEnvironmentUsingPUTParams) SetEnvironmentID(environmentID int64) {
	o.EnvironmentID = environmentID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateEnvironmentUsingPUTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.EnvironmentDto != nil {
		if err := r.SetBodyParam(o.EnvironmentDto); err != nil {
			return err
		}
	}

	// path param environmentId
	if err := r.SetPathParam("environmentId", swag.FormatInt64(o.EnvironmentID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
