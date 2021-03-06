// Code generated by go-swagger; DO NOT EDIT.

package system_preferences

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

// NewGetSystemPreferencesDtoUsingGETParams creates a new GetSystemPreferencesDtoUsingGETParams object
// with the default values initialized.
func NewGetSystemPreferencesDtoUsingGETParams() *GetSystemPreferencesDtoUsingGETParams {
	var ()
	return &GetSystemPreferencesDtoUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSystemPreferencesDtoUsingGETParamsWithTimeout creates a new GetSystemPreferencesDtoUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSystemPreferencesDtoUsingGETParamsWithTimeout(timeout time.Duration) *GetSystemPreferencesDtoUsingGETParams {
	var ()
	return &GetSystemPreferencesDtoUsingGETParams{

		timeout: timeout,
	}
}

// NewGetSystemPreferencesDtoUsingGETParamsWithContext creates a new GetSystemPreferencesDtoUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSystemPreferencesDtoUsingGETParamsWithContext(ctx context.Context) *GetSystemPreferencesDtoUsingGETParams {
	var ()
	return &GetSystemPreferencesDtoUsingGETParams{

		Context: ctx,
	}
}

// NewGetSystemPreferencesDtoUsingGETParamsWithHTTPClient creates a new GetSystemPreferencesDtoUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSystemPreferencesDtoUsingGETParamsWithHTTPClient(client *http.Client) *GetSystemPreferencesDtoUsingGETParams {
	var ()
	return &GetSystemPreferencesDtoUsingGETParams{
		HTTPClient: client,
	}
}

/*GetSystemPreferencesDtoUsingGETParams contains all the parameters to send to the API endpoint
for the get system preferences dto using g e t operation typically these are written to a http.Request
*/
type GetSystemPreferencesDtoUsingGETParams struct {

	/*SystemPreference
	  systemPreference

	*/
	SystemPreference []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get system preferences dto using g e t params
func (o *GetSystemPreferencesDtoUsingGETParams) WithTimeout(timeout time.Duration) *GetSystemPreferencesDtoUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get system preferences dto using g e t params
func (o *GetSystemPreferencesDtoUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get system preferences dto using g e t params
func (o *GetSystemPreferencesDtoUsingGETParams) WithContext(ctx context.Context) *GetSystemPreferencesDtoUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get system preferences dto using g e t params
func (o *GetSystemPreferencesDtoUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get system preferences dto using g e t params
func (o *GetSystemPreferencesDtoUsingGETParams) WithHTTPClient(client *http.Client) *GetSystemPreferencesDtoUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get system preferences dto using g e t params
func (o *GetSystemPreferencesDtoUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSystemPreference adds the systemPreference to the get system preferences dto using g e t params
func (o *GetSystemPreferencesDtoUsingGETParams) WithSystemPreference(systemPreference []string) *GetSystemPreferencesDtoUsingGETParams {
	o.SetSystemPreference(systemPreference)
	return o
}

// SetSystemPreference adds the systemPreference to the get system preferences dto using g e t params
func (o *GetSystemPreferencesDtoUsingGETParams) SetSystemPreference(systemPreference []string) {
	o.SystemPreference = systemPreference
}

// WriteToRequest writes these params to a swagger request
func (o *GetSystemPreferencesDtoUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesSystemPreference := o.SystemPreference

	joinedSystemPreference := swag.JoinByFormat(valuesSystemPreference, "multi")
	// query array param system_preference
	if err := r.SetQueryParam("system_preference", joinedSystemPreference...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
