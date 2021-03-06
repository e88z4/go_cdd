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
)

// NewDeleteEnvironmentUsingDELETEParams creates a new DeleteEnvironmentUsingDELETEParams object
// with the default values initialized.
func NewDeleteEnvironmentUsingDELETEParams() *DeleteEnvironmentUsingDELETEParams {
	var ()
	return &DeleteEnvironmentUsingDELETEParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteEnvironmentUsingDELETEParamsWithTimeout creates a new DeleteEnvironmentUsingDELETEParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteEnvironmentUsingDELETEParamsWithTimeout(timeout time.Duration) *DeleteEnvironmentUsingDELETEParams {
	var ()
	return &DeleteEnvironmentUsingDELETEParams{

		timeout: timeout,
	}
}

// NewDeleteEnvironmentUsingDELETEParamsWithContext creates a new DeleteEnvironmentUsingDELETEParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteEnvironmentUsingDELETEParamsWithContext(ctx context.Context) *DeleteEnvironmentUsingDELETEParams {
	var ()
	return &DeleteEnvironmentUsingDELETEParams{

		Context: ctx,
	}
}

// NewDeleteEnvironmentUsingDELETEParamsWithHTTPClient creates a new DeleteEnvironmentUsingDELETEParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteEnvironmentUsingDELETEParamsWithHTTPClient(client *http.Client) *DeleteEnvironmentUsingDELETEParams {
	var ()
	return &DeleteEnvironmentUsingDELETEParams{
		HTTPClient: client,
	}
}

/*DeleteEnvironmentUsingDELETEParams contains all the parameters to send to the API endpoint
for the delete environment using d e l e t e operation typically these are written to a http.Request
*/
type DeleteEnvironmentUsingDELETEParams struct {

	/*ApplicationID
	  applicationId

	*/
	ApplicationID int64
	/*EnvironmentID
	  environmentId

	*/
	EnvironmentID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete environment using d e l e t e params
func (o *DeleteEnvironmentUsingDELETEParams) WithTimeout(timeout time.Duration) *DeleteEnvironmentUsingDELETEParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete environment using d e l e t e params
func (o *DeleteEnvironmentUsingDELETEParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete environment using d e l e t e params
func (o *DeleteEnvironmentUsingDELETEParams) WithContext(ctx context.Context) *DeleteEnvironmentUsingDELETEParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete environment using d e l e t e params
func (o *DeleteEnvironmentUsingDELETEParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete environment using d e l e t e params
func (o *DeleteEnvironmentUsingDELETEParams) WithHTTPClient(client *http.Client) *DeleteEnvironmentUsingDELETEParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete environment using d e l e t e params
func (o *DeleteEnvironmentUsingDELETEParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApplicationID adds the applicationID to the delete environment using d e l e t e params
func (o *DeleteEnvironmentUsingDELETEParams) WithApplicationID(applicationID int64) *DeleteEnvironmentUsingDELETEParams {
	o.SetApplicationID(applicationID)
	return o
}

// SetApplicationID adds the applicationId to the delete environment using d e l e t e params
func (o *DeleteEnvironmentUsingDELETEParams) SetApplicationID(applicationID int64) {
	o.ApplicationID = applicationID
}

// WithEnvironmentID adds the environmentID to the delete environment using d e l e t e params
func (o *DeleteEnvironmentUsingDELETEParams) WithEnvironmentID(environmentID int64) *DeleteEnvironmentUsingDELETEParams {
	o.SetEnvironmentID(environmentID)
	return o
}

// SetEnvironmentID adds the environmentId to the delete environment using d e l e t e params
func (o *DeleteEnvironmentUsingDELETEParams) SetEnvironmentID(environmentID int64) {
	o.EnvironmentID = environmentID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteEnvironmentUsingDELETEParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param applicationId
	if err := r.SetPathParam("applicationId", swag.FormatInt64(o.ApplicationID)); err != nil {
		return err
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
