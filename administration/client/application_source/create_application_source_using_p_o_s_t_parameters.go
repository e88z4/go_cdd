// Code generated by go-swagger; DO NOT EDIT.

package application_source

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

// NewCreateApplicationSourceUsingPOSTParams creates a new CreateApplicationSourceUsingPOSTParams object
// with the default values initialized.
func NewCreateApplicationSourceUsingPOSTParams() *CreateApplicationSourceUsingPOSTParams {
	var ()
	return &CreateApplicationSourceUsingPOSTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateApplicationSourceUsingPOSTParamsWithTimeout creates a new CreateApplicationSourceUsingPOSTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateApplicationSourceUsingPOSTParamsWithTimeout(timeout time.Duration) *CreateApplicationSourceUsingPOSTParams {
	var ()
	return &CreateApplicationSourceUsingPOSTParams{

		timeout: timeout,
	}
}

// NewCreateApplicationSourceUsingPOSTParamsWithContext creates a new CreateApplicationSourceUsingPOSTParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateApplicationSourceUsingPOSTParamsWithContext(ctx context.Context) *CreateApplicationSourceUsingPOSTParams {
	var ()
	return &CreateApplicationSourceUsingPOSTParams{

		Context: ctx,
	}
}

// NewCreateApplicationSourceUsingPOSTParamsWithHTTPClient creates a new CreateApplicationSourceUsingPOSTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateApplicationSourceUsingPOSTParamsWithHTTPClient(client *http.Client) *CreateApplicationSourceUsingPOSTParams {
	var ()
	return &CreateApplicationSourceUsingPOSTParams{
		HTTPClient: client,
	}
}

/*CreateApplicationSourceUsingPOSTParams contains all the parameters to send to the API endpoint
for the create application source using p o s t operation typically these are written to a http.Request
*/
type CreateApplicationSourceUsingPOSTParams struct {

	/*ApplicationSourceDto
	  applicationSourceDto

	*/
	ApplicationSourceDto *models.ApplicationSourceDto

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create application source using p o s t params
func (o *CreateApplicationSourceUsingPOSTParams) WithTimeout(timeout time.Duration) *CreateApplicationSourceUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create application source using p o s t params
func (o *CreateApplicationSourceUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create application source using p o s t params
func (o *CreateApplicationSourceUsingPOSTParams) WithContext(ctx context.Context) *CreateApplicationSourceUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create application source using p o s t params
func (o *CreateApplicationSourceUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create application source using p o s t params
func (o *CreateApplicationSourceUsingPOSTParams) WithHTTPClient(client *http.Client) *CreateApplicationSourceUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create application source using p o s t params
func (o *CreateApplicationSourceUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApplicationSourceDto adds the applicationSourceDto to the create application source using p o s t params
func (o *CreateApplicationSourceUsingPOSTParams) WithApplicationSourceDto(applicationSourceDto *models.ApplicationSourceDto) *CreateApplicationSourceUsingPOSTParams {
	o.SetApplicationSourceDto(applicationSourceDto)
	return o
}

// SetApplicationSourceDto adds the applicationSourceDto to the create application source using p o s t params
func (o *CreateApplicationSourceUsingPOSTParams) SetApplicationSourceDto(applicationSourceDto *models.ApplicationSourceDto) {
	o.ApplicationSourceDto = applicationSourceDto
}

// WriteToRequest writes these params to a swagger request
func (o *CreateApplicationSourceUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ApplicationSourceDto != nil {
		if err := r.SetBodyParam(o.ApplicationSourceDto); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
