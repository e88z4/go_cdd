// Code generated by go-swagger; DO NOT EDIT.

package plugin_service

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

// NewGetAllPluginsUsingGETParams creates a new GetAllPluginsUsingGETParams object
// with the default values initialized.
func NewGetAllPluginsUsingGETParams() *GetAllPluginsUsingGETParams {

	return &GetAllPluginsUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllPluginsUsingGETParamsWithTimeout creates a new GetAllPluginsUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAllPluginsUsingGETParamsWithTimeout(timeout time.Duration) *GetAllPluginsUsingGETParams {

	return &GetAllPluginsUsingGETParams{

		timeout: timeout,
	}
}

// NewGetAllPluginsUsingGETParamsWithContext creates a new GetAllPluginsUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAllPluginsUsingGETParamsWithContext(ctx context.Context) *GetAllPluginsUsingGETParams {

	return &GetAllPluginsUsingGETParams{

		Context: ctx,
	}
}

// NewGetAllPluginsUsingGETParamsWithHTTPClient creates a new GetAllPluginsUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAllPluginsUsingGETParamsWithHTTPClient(client *http.Client) *GetAllPluginsUsingGETParams {

	return &GetAllPluginsUsingGETParams{
		HTTPClient: client,
	}
}

/*GetAllPluginsUsingGETParams contains all the parameters to send to the API endpoint
for the get all plugins using g e t operation typically these are written to a http.Request
*/
type GetAllPluginsUsingGETParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get all plugins using g e t params
func (o *GetAllPluginsUsingGETParams) WithTimeout(timeout time.Duration) *GetAllPluginsUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get all plugins using g e t params
func (o *GetAllPluginsUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get all plugins using g e t params
func (o *GetAllPluginsUsingGETParams) WithContext(ctx context.Context) *GetAllPluginsUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get all plugins using g e t params
func (o *GetAllPluginsUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get all plugins using g e t params
func (o *GetAllPluginsUsingGETParams) WithHTTPClient(client *http.Client) *GetAllPluginsUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get all plugins using g e t params
func (o *GetAllPluginsUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllPluginsUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
