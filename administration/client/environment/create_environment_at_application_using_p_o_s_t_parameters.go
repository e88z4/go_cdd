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

// NewCreateEnvironmentAtApplicationUsingPOSTParams creates a new CreateEnvironmentAtApplicationUsingPOSTParams object
// with the default values initialized.
func NewCreateEnvironmentAtApplicationUsingPOSTParams() *CreateEnvironmentAtApplicationUsingPOSTParams {
	var ()
	return &CreateEnvironmentAtApplicationUsingPOSTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateEnvironmentAtApplicationUsingPOSTParamsWithTimeout creates a new CreateEnvironmentAtApplicationUsingPOSTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateEnvironmentAtApplicationUsingPOSTParamsWithTimeout(timeout time.Duration) *CreateEnvironmentAtApplicationUsingPOSTParams {
	var ()
	return &CreateEnvironmentAtApplicationUsingPOSTParams{

		timeout: timeout,
	}
}

// NewCreateEnvironmentAtApplicationUsingPOSTParamsWithContext creates a new CreateEnvironmentAtApplicationUsingPOSTParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateEnvironmentAtApplicationUsingPOSTParamsWithContext(ctx context.Context) *CreateEnvironmentAtApplicationUsingPOSTParams {
	var ()
	return &CreateEnvironmentAtApplicationUsingPOSTParams{

		Context: ctx,
	}
}

// NewCreateEnvironmentAtApplicationUsingPOSTParamsWithHTTPClient creates a new CreateEnvironmentAtApplicationUsingPOSTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateEnvironmentAtApplicationUsingPOSTParamsWithHTTPClient(client *http.Client) *CreateEnvironmentAtApplicationUsingPOSTParams {
	var ()
	return &CreateEnvironmentAtApplicationUsingPOSTParams{
		HTTPClient: client,
	}
}

/*CreateEnvironmentAtApplicationUsingPOSTParams contains all the parameters to send to the API endpoint
for the create environment at application using p o s t operation typically these are written to a http.Request
*/
type CreateEnvironmentAtApplicationUsingPOSTParams struct {

	/*ApplicationID
	  applicationId

	*/
	ApplicationID int64
	/*EnvironmentDto
	  environmentDto

	*/
	EnvironmentDto *models.EnvironmentDto

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create environment at application using p o s t params
func (o *CreateEnvironmentAtApplicationUsingPOSTParams) WithTimeout(timeout time.Duration) *CreateEnvironmentAtApplicationUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create environment at application using p o s t params
func (o *CreateEnvironmentAtApplicationUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create environment at application using p o s t params
func (o *CreateEnvironmentAtApplicationUsingPOSTParams) WithContext(ctx context.Context) *CreateEnvironmentAtApplicationUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create environment at application using p o s t params
func (o *CreateEnvironmentAtApplicationUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create environment at application using p o s t params
func (o *CreateEnvironmentAtApplicationUsingPOSTParams) WithHTTPClient(client *http.Client) *CreateEnvironmentAtApplicationUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create environment at application using p o s t params
func (o *CreateEnvironmentAtApplicationUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApplicationID adds the applicationID to the create environment at application using p o s t params
func (o *CreateEnvironmentAtApplicationUsingPOSTParams) WithApplicationID(applicationID int64) *CreateEnvironmentAtApplicationUsingPOSTParams {
	o.SetApplicationID(applicationID)
	return o
}

// SetApplicationID adds the applicationId to the create environment at application using p o s t params
func (o *CreateEnvironmentAtApplicationUsingPOSTParams) SetApplicationID(applicationID int64) {
	o.ApplicationID = applicationID
}

// WithEnvironmentDto adds the environmentDto to the create environment at application using p o s t params
func (o *CreateEnvironmentAtApplicationUsingPOSTParams) WithEnvironmentDto(environmentDto *models.EnvironmentDto) *CreateEnvironmentAtApplicationUsingPOSTParams {
	o.SetEnvironmentDto(environmentDto)
	return o
}

// SetEnvironmentDto adds the environmentDto to the create environment at application using p o s t params
func (o *CreateEnvironmentAtApplicationUsingPOSTParams) SetEnvironmentDto(environmentDto *models.EnvironmentDto) {
	o.EnvironmentDto = environmentDto
}

// WriteToRequest writes these params to a swagger request
func (o *CreateEnvironmentAtApplicationUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param applicationId
	if err := r.SetPathParam("applicationId", swag.FormatInt64(o.ApplicationID)); err != nil {
		return err
	}

	if o.EnvironmentDto != nil {
		if err := r.SetBodyParam(o.EnvironmentDto); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}