// Code generated by go-swagger; DO NOT EDIT.

package maintenance_window

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

// NewGetMaintenanceWindowsUsingGETParams creates a new GetMaintenanceWindowsUsingGETParams object
// with the default values initialized.
func NewGetMaintenanceWindowsUsingGETParams() *GetMaintenanceWindowsUsingGETParams {
	var ()
	return &GetMaintenanceWindowsUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMaintenanceWindowsUsingGETParamsWithTimeout creates a new GetMaintenanceWindowsUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMaintenanceWindowsUsingGETParamsWithTimeout(timeout time.Duration) *GetMaintenanceWindowsUsingGETParams {
	var ()
	return &GetMaintenanceWindowsUsingGETParams{

		timeout: timeout,
	}
}

// NewGetMaintenanceWindowsUsingGETParamsWithContext creates a new GetMaintenanceWindowsUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMaintenanceWindowsUsingGETParamsWithContext(ctx context.Context) *GetMaintenanceWindowsUsingGETParams {
	var ()
	return &GetMaintenanceWindowsUsingGETParams{

		Context: ctx,
	}
}

// NewGetMaintenanceWindowsUsingGETParamsWithHTTPClient creates a new GetMaintenanceWindowsUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetMaintenanceWindowsUsingGETParamsWithHTTPClient(client *http.Client) *GetMaintenanceWindowsUsingGETParams {
	var ()
	return &GetMaintenanceWindowsUsingGETParams{
		HTTPClient: client,
	}
}

/*GetMaintenanceWindowsUsingGETParams contains all the parameters to send to the API endpoint
for the get maintenance windows using g e t operation typically these are written to a http.Request
*/
type GetMaintenanceWindowsUsingGETParams struct {

	/*Filter
	  freeTextFilter

	*/
	Filter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get maintenance windows using g e t params
func (o *GetMaintenanceWindowsUsingGETParams) WithTimeout(timeout time.Duration) *GetMaintenanceWindowsUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get maintenance windows using g e t params
func (o *GetMaintenanceWindowsUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get maintenance windows using g e t params
func (o *GetMaintenanceWindowsUsingGETParams) WithContext(ctx context.Context) *GetMaintenanceWindowsUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get maintenance windows using g e t params
func (o *GetMaintenanceWindowsUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get maintenance windows using g e t params
func (o *GetMaintenanceWindowsUsingGETParams) WithHTTPClient(client *http.Client) *GetMaintenanceWindowsUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get maintenance windows using g e t params
func (o *GetMaintenanceWindowsUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilter adds the filter to the get maintenance windows using g e t params
func (o *GetMaintenanceWindowsUsingGETParams) WithFilter(filter *string) *GetMaintenanceWindowsUsingGETParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the get maintenance windows using g e t params
func (o *GetMaintenanceWindowsUsingGETParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WriteToRequest writes these params to a swagger request
func (o *GetMaintenanceWindowsUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Filter != nil {

		// query param filter
		var qrFilter string
		if o.Filter != nil {
			qrFilter = *o.Filter
		}
		qFilter := qrFilter
		if qFilter != "" {
			if err := r.SetQueryParam("filter", qFilter); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
