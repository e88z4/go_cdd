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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetRoleUsingGETParams creates a new GetRoleUsingGETParams object
// with the default values initialized.
func NewGetRoleUsingGETParams() *GetRoleUsingGETParams {
	var ()
	return &GetRoleUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRoleUsingGETParamsWithTimeout creates a new GetRoleUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRoleUsingGETParamsWithTimeout(timeout time.Duration) *GetRoleUsingGETParams {
	var ()
	return &GetRoleUsingGETParams{

		timeout: timeout,
	}
}

// NewGetRoleUsingGETParamsWithContext creates a new GetRoleUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRoleUsingGETParamsWithContext(ctx context.Context) *GetRoleUsingGETParams {
	var ()
	return &GetRoleUsingGETParams{

		Context: ctx,
	}
}

// NewGetRoleUsingGETParamsWithHTTPClient creates a new GetRoleUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRoleUsingGETParamsWithHTTPClient(client *http.Client) *GetRoleUsingGETParams {
	var ()
	return &GetRoleUsingGETParams{
		HTTPClient: client,
	}
}

/*GetRoleUsingGETParams contains all the parameters to send to the API endpoint
for the get role using g e t operation typically these are written to a http.Request
*/
type GetRoleUsingGETParams struct {

	/*RoleID
	  roleId

	*/
	RoleID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get role using g e t params
func (o *GetRoleUsingGETParams) WithTimeout(timeout time.Duration) *GetRoleUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get role using g e t params
func (o *GetRoleUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get role using g e t params
func (o *GetRoleUsingGETParams) WithContext(ctx context.Context) *GetRoleUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get role using g e t params
func (o *GetRoleUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get role using g e t params
func (o *GetRoleUsingGETParams) WithHTTPClient(client *http.Client) *GetRoleUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get role using g e t params
func (o *GetRoleUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRoleID adds the roleID to the get role using g e t params
func (o *GetRoleUsingGETParams) WithRoleID(roleID int64) *GetRoleUsingGETParams {
	o.SetRoleID(roleID)
	return o
}

// SetRoleID adds the roleId to the get role using g e t params
func (o *GetRoleUsingGETParams) SetRoleID(roleID int64) {
	o.RoleID = roleID
}

// WriteToRequest writes these params to a swagger request
func (o *GetRoleUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param roleId
	if err := r.SetPathParam("roleId", swag.FormatInt64(o.RoleID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
