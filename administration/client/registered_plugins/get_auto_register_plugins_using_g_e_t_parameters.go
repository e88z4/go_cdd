// Code generated by go-swagger; DO NOT EDIT.

package registered_plugins

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
)

// NewGetAutoRegisterPluginsUsingGETParams creates a new GetAutoRegisterPluginsUsingGETParams object
// with the default values initialized.
func NewGetAutoRegisterPluginsUsingGETParams() *GetAutoRegisterPluginsUsingGETParams {

	return &GetAutoRegisterPluginsUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAutoRegisterPluginsUsingGETParamsWithTimeout creates a new GetAutoRegisterPluginsUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAutoRegisterPluginsUsingGETParamsWithTimeout(timeout time.Duration) *GetAutoRegisterPluginsUsingGETParams {

	return &GetAutoRegisterPluginsUsingGETParams{

		timeout: timeout,
	}
}

// NewGetAutoRegisterPluginsUsingGETParamsWithContext creates a new GetAutoRegisterPluginsUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAutoRegisterPluginsUsingGETParamsWithContext(ctx context.Context) *GetAutoRegisterPluginsUsingGETParams {

	return &GetAutoRegisterPluginsUsingGETParams{

		Context: ctx,
	}
}

// NewGetAutoRegisterPluginsUsingGETParamsWithHTTPClient creates a new GetAutoRegisterPluginsUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAutoRegisterPluginsUsingGETParamsWithHTTPClient(client *http.Client) *GetAutoRegisterPluginsUsingGETParams {

	return &GetAutoRegisterPluginsUsingGETParams{
		HTTPClient: client,
	}
}

/*GetAutoRegisterPluginsUsingGETParams contains all the parameters to send to the API endpoint
for the get auto register plugins using g e t operation typically these are written to a http.Request
*/
type GetAutoRegisterPluginsUsingGETParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get auto register plugins using g e t params
func (o *GetAutoRegisterPluginsUsingGETParams) WithTimeout(timeout time.Duration) *GetAutoRegisterPluginsUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get auto register plugins using g e t params
func (o *GetAutoRegisterPluginsUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get auto register plugins using g e t params
func (o *GetAutoRegisterPluginsUsingGETParams) WithContext(ctx context.Context) *GetAutoRegisterPluginsUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get auto register plugins using g e t params
func (o *GetAutoRegisterPluginsUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get auto register plugins using g e t params
func (o *GetAutoRegisterPluginsUsingGETParams) WithHTTPClient(client *http.Client) *GetAutoRegisterPluginsUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get auto register plugins using g e t params
func (o *GetAutoRegisterPluginsUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAutoRegisterPluginsUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
