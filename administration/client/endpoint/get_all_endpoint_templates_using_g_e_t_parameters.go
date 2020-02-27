// Code generated by go-swagger; DO NOT EDIT.

package endpoint

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

// NewGetAllEndpointTemplatesUsingGETParams creates a new GetAllEndpointTemplatesUsingGETParams object
// with the default values initialized.
func NewGetAllEndpointTemplatesUsingGETParams() *GetAllEndpointTemplatesUsingGETParams {

	return &GetAllEndpointTemplatesUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllEndpointTemplatesUsingGETParamsWithTimeout creates a new GetAllEndpointTemplatesUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAllEndpointTemplatesUsingGETParamsWithTimeout(timeout time.Duration) *GetAllEndpointTemplatesUsingGETParams {

	return &GetAllEndpointTemplatesUsingGETParams{

		timeout: timeout,
	}
}

// NewGetAllEndpointTemplatesUsingGETParamsWithContext creates a new GetAllEndpointTemplatesUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAllEndpointTemplatesUsingGETParamsWithContext(ctx context.Context) *GetAllEndpointTemplatesUsingGETParams {

	return &GetAllEndpointTemplatesUsingGETParams{

		Context: ctx,
	}
}

// NewGetAllEndpointTemplatesUsingGETParamsWithHTTPClient creates a new GetAllEndpointTemplatesUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAllEndpointTemplatesUsingGETParamsWithHTTPClient(client *http.Client) *GetAllEndpointTemplatesUsingGETParams {

	return &GetAllEndpointTemplatesUsingGETParams{
		HTTPClient: client,
	}
}

/*GetAllEndpointTemplatesUsingGETParams contains all the parameters to send to the API endpoint
for the get all endpoint templates using g e t operation typically these are written to a http.Request
*/
type GetAllEndpointTemplatesUsingGETParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get all endpoint templates using g e t params
func (o *GetAllEndpointTemplatesUsingGETParams) WithTimeout(timeout time.Duration) *GetAllEndpointTemplatesUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get all endpoint templates using g e t params
func (o *GetAllEndpointTemplatesUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get all endpoint templates using g e t params
func (o *GetAllEndpointTemplatesUsingGETParams) WithContext(ctx context.Context) *GetAllEndpointTemplatesUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get all endpoint templates using g e t params
func (o *GetAllEndpointTemplatesUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get all endpoint templates using g e t params
func (o *GetAllEndpointTemplatesUsingGETParams) WithHTTPClient(client *http.Client) *GetAllEndpointTemplatesUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get all endpoint templates using g e t params
func (o *GetAllEndpointTemplatesUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllEndpointTemplatesUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
