// Code generated by go-swagger; DO NOT EDIT.

package roles

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

// NewCreateRoleUsingPOSTParams creates a new CreateRoleUsingPOSTParams object
// with the default values initialized.
func NewCreateRoleUsingPOSTParams() *CreateRoleUsingPOSTParams {
	var ()
	return &CreateRoleUsingPOSTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateRoleUsingPOSTParamsWithTimeout creates a new CreateRoleUsingPOSTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateRoleUsingPOSTParamsWithTimeout(timeout time.Duration) *CreateRoleUsingPOSTParams {
	var ()
	return &CreateRoleUsingPOSTParams{

		timeout: timeout,
	}
}

// NewCreateRoleUsingPOSTParamsWithContext creates a new CreateRoleUsingPOSTParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateRoleUsingPOSTParamsWithContext(ctx context.Context) *CreateRoleUsingPOSTParams {
	var ()
	return &CreateRoleUsingPOSTParams{

		Context: ctx,
	}
}

// NewCreateRoleUsingPOSTParamsWithHTTPClient creates a new CreateRoleUsingPOSTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateRoleUsingPOSTParamsWithHTTPClient(client *http.Client) *CreateRoleUsingPOSTParams {
	var ()
	return &CreateRoleUsingPOSTParams{
		HTTPClient: client,
	}
}

/*CreateRoleUsingPOSTParams contains all the parameters to send to the API endpoint
for the create role using p o s t operation typically these are written to a http.Request
*/
type CreateRoleUsingPOSTParams struct {

	/*RoleDto
	  roleDto

	*/
	RoleDto *models.RoleDto

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create role using p o s t params
func (o *CreateRoleUsingPOSTParams) WithTimeout(timeout time.Duration) *CreateRoleUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create role using p o s t params
func (o *CreateRoleUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create role using p o s t params
func (o *CreateRoleUsingPOSTParams) WithContext(ctx context.Context) *CreateRoleUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create role using p o s t params
func (o *CreateRoleUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create role using p o s t params
func (o *CreateRoleUsingPOSTParams) WithHTTPClient(client *http.Client) *CreateRoleUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create role using p o s t params
func (o *CreateRoleUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRoleDto adds the roleDto to the create role using p o s t params
func (o *CreateRoleUsingPOSTParams) WithRoleDto(roleDto *models.RoleDto) *CreateRoleUsingPOSTParams {
	o.SetRoleDto(roleDto)
	return o
}

// SetRoleDto adds the roleDto to the create role using p o s t params
func (o *CreateRoleUsingPOSTParams) SetRoleDto(roleDto *models.RoleDto) {
	o.RoleDto = roleDto
}

// WriteToRequest writes these params to a swagger request
func (o *CreateRoleUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.RoleDto != nil {
		if err := r.SetBodyParam(o.RoleDto); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
