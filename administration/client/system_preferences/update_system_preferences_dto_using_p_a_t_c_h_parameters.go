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

	strfmt "github.com/go-openapi/strfmt"

	"github.com/e88z4/go_cdd/administration/models"
)

// NewUpdateSystemPreferencesDtoUsingPATCHParams creates a new UpdateSystemPreferencesDtoUsingPATCHParams object
// with the default values initialized.
func NewUpdateSystemPreferencesDtoUsingPATCHParams() *UpdateSystemPreferencesDtoUsingPATCHParams {
	var ()
	return &UpdateSystemPreferencesDtoUsingPATCHParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateSystemPreferencesDtoUsingPATCHParamsWithTimeout creates a new UpdateSystemPreferencesDtoUsingPATCHParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateSystemPreferencesDtoUsingPATCHParamsWithTimeout(timeout time.Duration) *UpdateSystemPreferencesDtoUsingPATCHParams {
	var ()
	return &UpdateSystemPreferencesDtoUsingPATCHParams{

		timeout: timeout,
	}
}

// NewUpdateSystemPreferencesDtoUsingPATCHParamsWithContext creates a new UpdateSystemPreferencesDtoUsingPATCHParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateSystemPreferencesDtoUsingPATCHParamsWithContext(ctx context.Context) *UpdateSystemPreferencesDtoUsingPATCHParams {
	var ()
	return &UpdateSystemPreferencesDtoUsingPATCHParams{

		Context: ctx,
	}
}

// NewUpdateSystemPreferencesDtoUsingPATCHParamsWithHTTPClient creates a new UpdateSystemPreferencesDtoUsingPATCHParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateSystemPreferencesDtoUsingPATCHParamsWithHTTPClient(client *http.Client) *UpdateSystemPreferencesDtoUsingPATCHParams {
	var ()
	return &UpdateSystemPreferencesDtoUsingPATCHParams{
		HTTPClient: client,
	}
}

/*UpdateSystemPreferencesDtoUsingPATCHParams contains all the parameters to send to the API endpoint
for the update system preferences dto using p a t c h operation typically these are written to a http.Request
*/
type UpdateSystemPreferencesDtoUsingPATCHParams struct {

	/*PreferenceDtos
	  preferenceDtos

	*/
	PreferenceDtos []*models.ShallowSystemPreferenceDto

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update system preferences dto using p a t c h params
func (o *UpdateSystemPreferencesDtoUsingPATCHParams) WithTimeout(timeout time.Duration) *UpdateSystemPreferencesDtoUsingPATCHParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update system preferences dto using p a t c h params
func (o *UpdateSystemPreferencesDtoUsingPATCHParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update system preferences dto using p a t c h params
func (o *UpdateSystemPreferencesDtoUsingPATCHParams) WithContext(ctx context.Context) *UpdateSystemPreferencesDtoUsingPATCHParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update system preferences dto using p a t c h params
func (o *UpdateSystemPreferencesDtoUsingPATCHParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update system preferences dto using p a t c h params
func (o *UpdateSystemPreferencesDtoUsingPATCHParams) WithHTTPClient(client *http.Client) *UpdateSystemPreferencesDtoUsingPATCHParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update system preferences dto using p a t c h params
func (o *UpdateSystemPreferencesDtoUsingPATCHParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPreferenceDtos adds the preferenceDtos to the update system preferences dto using p a t c h params
func (o *UpdateSystemPreferencesDtoUsingPATCHParams) WithPreferenceDtos(preferenceDtos []*models.ShallowSystemPreferenceDto) *UpdateSystemPreferencesDtoUsingPATCHParams {
	o.SetPreferenceDtos(preferenceDtos)
	return o
}

// SetPreferenceDtos adds the preferenceDtos to the update system preferences dto using p a t c h params
func (o *UpdateSystemPreferencesDtoUsingPATCHParams) SetPreferenceDtos(preferenceDtos []*models.ShallowSystemPreferenceDto) {
	o.PreferenceDtos = preferenceDtos
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateSystemPreferencesDtoUsingPATCHParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.PreferenceDtos != nil {
		if err := r.SetBodyParam(o.PreferenceDtos); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}