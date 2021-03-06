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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetMaintenanceWindowUsingGETParams creates a new GetMaintenanceWindowUsingGETParams object
// with the default values initialized.
func NewGetMaintenanceWindowUsingGETParams() *GetMaintenanceWindowUsingGETParams {
	var ()
	return &GetMaintenanceWindowUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMaintenanceWindowUsingGETParamsWithTimeout creates a new GetMaintenanceWindowUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMaintenanceWindowUsingGETParamsWithTimeout(timeout time.Duration) *GetMaintenanceWindowUsingGETParams {
	var ()
	return &GetMaintenanceWindowUsingGETParams{

		timeout: timeout,
	}
}

// NewGetMaintenanceWindowUsingGETParamsWithContext creates a new GetMaintenanceWindowUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMaintenanceWindowUsingGETParamsWithContext(ctx context.Context) *GetMaintenanceWindowUsingGETParams {
	var ()
	return &GetMaintenanceWindowUsingGETParams{

		Context: ctx,
	}
}

// NewGetMaintenanceWindowUsingGETParamsWithHTTPClient creates a new GetMaintenanceWindowUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetMaintenanceWindowUsingGETParamsWithHTTPClient(client *http.Client) *GetMaintenanceWindowUsingGETParams {
	var ()
	return &GetMaintenanceWindowUsingGETParams{
		HTTPClient: client,
	}
}

/*GetMaintenanceWindowUsingGETParams contains all the parameters to send to the API endpoint
for the get maintenance window using g e t operation typically these are written to a http.Request
*/
type GetMaintenanceWindowUsingGETParams struct {

	/*MaintenanceWindowID
	  maintenanceWindowId

	*/
	MaintenanceWindowID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get maintenance window using g e t params
func (o *GetMaintenanceWindowUsingGETParams) WithTimeout(timeout time.Duration) *GetMaintenanceWindowUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get maintenance window using g e t params
func (o *GetMaintenanceWindowUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get maintenance window using g e t params
func (o *GetMaintenanceWindowUsingGETParams) WithContext(ctx context.Context) *GetMaintenanceWindowUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get maintenance window using g e t params
func (o *GetMaintenanceWindowUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get maintenance window using g e t params
func (o *GetMaintenanceWindowUsingGETParams) WithHTTPClient(client *http.Client) *GetMaintenanceWindowUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get maintenance window using g e t params
func (o *GetMaintenanceWindowUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMaintenanceWindowID adds the maintenanceWindowID to the get maintenance window using g e t params
func (o *GetMaintenanceWindowUsingGETParams) WithMaintenanceWindowID(maintenanceWindowID int64) *GetMaintenanceWindowUsingGETParams {
	o.SetMaintenanceWindowID(maintenanceWindowID)
	return o
}

// SetMaintenanceWindowID adds the maintenanceWindowId to the get maintenance window using g e t params
func (o *GetMaintenanceWindowUsingGETParams) SetMaintenanceWindowID(maintenanceWindowID int64) {
	o.MaintenanceWindowID = maintenanceWindowID
}

// WriteToRequest writes these params to a swagger request
func (o *GetMaintenanceWindowUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param maintenanceWindowId
	if err := r.SetPathParam("maintenanceWindowId", swag.FormatInt64(o.MaintenanceWindowID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
