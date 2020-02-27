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

// NewUpdateEnvironmentAtApplicationUsingPUTParams creates a new UpdateEnvironmentAtApplicationUsingPUTParams object
// with the default values initialized.
func NewUpdateEnvironmentAtApplicationUsingPUTParams() *UpdateEnvironmentAtApplicationUsingPUTParams {
	var ()
	return &UpdateEnvironmentAtApplicationUsingPUTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateEnvironmentAtApplicationUsingPUTParamsWithTimeout creates a new UpdateEnvironmentAtApplicationUsingPUTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateEnvironmentAtApplicationUsingPUTParamsWithTimeout(timeout time.Duration) *UpdateEnvironmentAtApplicationUsingPUTParams {
	var ()
	return &UpdateEnvironmentAtApplicationUsingPUTParams{

		timeout: timeout,
	}
}

// NewUpdateEnvironmentAtApplicationUsingPUTParamsWithContext creates a new UpdateEnvironmentAtApplicationUsingPUTParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateEnvironmentAtApplicationUsingPUTParamsWithContext(ctx context.Context) *UpdateEnvironmentAtApplicationUsingPUTParams {
	var ()
	return &UpdateEnvironmentAtApplicationUsingPUTParams{

		Context: ctx,
	}
}

// NewUpdateEnvironmentAtApplicationUsingPUTParamsWithHTTPClient creates a new UpdateEnvironmentAtApplicationUsingPUTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateEnvironmentAtApplicationUsingPUTParamsWithHTTPClient(client *http.Client) *UpdateEnvironmentAtApplicationUsingPUTParams {
	var ()
	return &UpdateEnvironmentAtApplicationUsingPUTParams{
		HTTPClient: client,
	}
}

/*UpdateEnvironmentAtApplicationUsingPUTParams contains all the parameters to send to the API endpoint
for the update environment at application using p u t operation typically these are written to a http.Request
*/
type UpdateEnvironmentAtApplicationUsingPUTParams struct {

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

// WithTimeout adds the timeout to the update environment at application using p u t params
func (o *UpdateEnvironmentAtApplicationUsingPUTParams) WithTimeout(timeout time.Duration) *UpdateEnvironmentAtApplicationUsingPUTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update environment at application using p u t params
func (o *UpdateEnvironmentAtApplicationUsingPUTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update environment at application using p u t params
func (o *UpdateEnvironmentAtApplicationUsingPUTParams) WithContext(ctx context.Context) *UpdateEnvironmentAtApplicationUsingPUTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update environment at application using p u t params
func (o *UpdateEnvironmentAtApplicationUsingPUTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update environment at application using p u t params
func (o *UpdateEnvironmentAtApplicationUsingPUTParams) WithHTTPClient(client *http.Client) *UpdateEnvironmentAtApplicationUsingPUTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update environment at application using p u t params
func (o *UpdateEnvironmentAtApplicationUsingPUTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnvironmentDto adds the environmentDto to the update environment at application using p u t params
func (o *UpdateEnvironmentAtApplicationUsingPUTParams) WithEnvironmentDto(environmentDto *models.EnvironmentDto) *UpdateEnvironmentAtApplicationUsingPUTParams {
	o.SetEnvironmentDto(environmentDto)
	return o
}

// SetEnvironmentDto adds the environmentDto to the update environment at application using p u t params
func (o *UpdateEnvironmentAtApplicationUsingPUTParams) SetEnvironmentDto(environmentDto *models.EnvironmentDto) {
	o.EnvironmentDto = environmentDto
}

// WithEnvironmentID adds the environmentID to the update environment at application using p u t params
func (o *UpdateEnvironmentAtApplicationUsingPUTParams) WithEnvironmentID(environmentID int64) *UpdateEnvironmentAtApplicationUsingPUTParams {
	o.SetEnvironmentID(environmentID)
	return o
}

// SetEnvironmentID adds the environmentId to the update environment at application using p u t params
func (o *UpdateEnvironmentAtApplicationUsingPUTParams) SetEnvironmentID(environmentID int64) {
	o.EnvironmentID = environmentID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateEnvironmentAtApplicationUsingPUTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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