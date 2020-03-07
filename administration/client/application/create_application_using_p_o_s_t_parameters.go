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

	strfmt "github.com/go-openapi/strfmt"

	"github.com/e88z4/go_cdd/administration/models"
)

// NewCreateApplicationUsingPOSTParams creates a new CreateApplicationUsingPOSTParams object
// with the default values initialized.
func NewCreateApplicationUsingPOSTParams() *CreateApplicationUsingPOSTParams {
	var ()
	return &CreateApplicationUsingPOSTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateApplicationUsingPOSTParamsWithTimeout creates a new CreateApplicationUsingPOSTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateApplicationUsingPOSTParamsWithTimeout(timeout time.Duration) *CreateApplicationUsingPOSTParams {
	var ()
	return &CreateApplicationUsingPOSTParams{

		timeout: timeout,
	}
}

// NewCreateApplicationUsingPOSTParamsWithContext creates a new CreateApplicationUsingPOSTParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateApplicationUsingPOSTParamsWithContext(ctx context.Context) *CreateApplicationUsingPOSTParams {
	var ()
	return &CreateApplicationUsingPOSTParams{

		Context: ctx,
	}
}

// NewCreateApplicationUsingPOSTParamsWithHTTPClient creates a new CreateApplicationUsingPOSTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateApplicationUsingPOSTParamsWithHTTPClient(client *http.Client) *CreateApplicationUsingPOSTParams {
	var ()
	return &CreateApplicationUsingPOSTParams{
		HTTPClient: client,
	}
}

/*CreateApplicationUsingPOSTParams contains all the parameters to send to the API endpoint
for the create application using p o s t operation typically these are written to a http.Request
*/
type CreateApplicationUsingPOSTParams struct {

	/*ApplicationDto
	  applicationDto

	*/
	ApplicationDto *models.ApplicationDto

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create application using p o s t params
func (o *CreateApplicationUsingPOSTParams) WithTimeout(timeout time.Duration) *CreateApplicationUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create application using p o s t params
func (o *CreateApplicationUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create application using p o s t params
func (o *CreateApplicationUsingPOSTParams) WithContext(ctx context.Context) *CreateApplicationUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create application using p o s t params
func (o *CreateApplicationUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create application using p o s t params
func (o *CreateApplicationUsingPOSTParams) WithHTTPClient(client *http.Client) *CreateApplicationUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create application using p o s t params
func (o *CreateApplicationUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApplicationDto adds the applicationDto to the create application using p o s t params
func (o *CreateApplicationUsingPOSTParams) WithApplicationDto(applicationDto *models.ApplicationDto) *CreateApplicationUsingPOSTParams {
	o.SetApplicationDto(applicationDto)
	return o
}

// SetApplicationDto adds the applicationDto to the create application using p o s t params
func (o *CreateApplicationUsingPOSTParams) SetApplicationDto(applicationDto *models.ApplicationDto) {
	o.ApplicationDto = applicationDto
}

// WriteToRequest writes these params to a swagger request
func (o *CreateApplicationUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ApplicationDto != nil {
		if err := r.SetBodyParam(o.ApplicationDto); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}