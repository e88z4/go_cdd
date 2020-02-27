// Code generated by go-swagger; DO NOT EDIT.

package email_templates

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

// NewGetEmailNotificationTemplatesUsingGETParams creates a new GetEmailNotificationTemplatesUsingGETParams object
// with the default values initialized.
func NewGetEmailNotificationTemplatesUsingGETParams() *GetEmailNotificationTemplatesUsingGETParams {
	var ()
	return &GetEmailNotificationTemplatesUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetEmailNotificationTemplatesUsingGETParamsWithTimeout creates a new GetEmailNotificationTemplatesUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetEmailNotificationTemplatesUsingGETParamsWithTimeout(timeout time.Duration) *GetEmailNotificationTemplatesUsingGETParams {
	var ()
	return &GetEmailNotificationTemplatesUsingGETParams{

		timeout: timeout,
	}
}

// NewGetEmailNotificationTemplatesUsingGETParamsWithContext creates a new GetEmailNotificationTemplatesUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetEmailNotificationTemplatesUsingGETParamsWithContext(ctx context.Context) *GetEmailNotificationTemplatesUsingGETParams {
	var ()
	return &GetEmailNotificationTemplatesUsingGETParams{

		Context: ctx,
	}
}

// NewGetEmailNotificationTemplatesUsingGETParamsWithHTTPClient creates a new GetEmailNotificationTemplatesUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetEmailNotificationTemplatesUsingGETParamsWithHTTPClient(client *http.Client) *GetEmailNotificationTemplatesUsingGETParams {
	var ()
	return &GetEmailNotificationTemplatesUsingGETParams{
		HTTPClient: client,
	}
}

/*GetEmailNotificationTemplatesUsingGETParams contains all the parameters to send to the API endpoint
for the get email notification templates using g e t operation typically these are written to a http.Request
*/
type GetEmailNotificationTemplatesUsingGETParams struct {

	/*LanguageTag
	  language

	*/
	LanguageTag string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get email notification templates using g e t params
func (o *GetEmailNotificationTemplatesUsingGETParams) WithTimeout(timeout time.Duration) *GetEmailNotificationTemplatesUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get email notification templates using g e t params
func (o *GetEmailNotificationTemplatesUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get email notification templates using g e t params
func (o *GetEmailNotificationTemplatesUsingGETParams) WithContext(ctx context.Context) *GetEmailNotificationTemplatesUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get email notification templates using g e t params
func (o *GetEmailNotificationTemplatesUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get email notification templates using g e t params
func (o *GetEmailNotificationTemplatesUsingGETParams) WithHTTPClient(client *http.Client) *GetEmailNotificationTemplatesUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get email notification templates using g e t params
func (o *GetEmailNotificationTemplatesUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLanguageTag adds the languageTag to the get email notification templates using g e t params
func (o *GetEmailNotificationTemplatesUsingGETParams) WithLanguageTag(languageTag string) *GetEmailNotificationTemplatesUsingGETParams {
	o.SetLanguageTag(languageTag)
	return o
}

// SetLanguageTag adds the languageTag to the get email notification templates using g e t params
func (o *GetEmailNotificationTemplatesUsingGETParams) SetLanguageTag(languageTag string) {
	o.LanguageTag = languageTag
}

// WriteToRequest writes these params to a swagger request
func (o *GetEmailNotificationTemplatesUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param languageTag
	if err := r.SetPathParam("languageTag", o.LanguageTag); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}