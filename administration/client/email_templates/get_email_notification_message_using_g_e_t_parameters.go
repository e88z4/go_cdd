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

// NewGetEmailNotificationMessageUsingGETParams creates a new GetEmailNotificationMessageUsingGETParams object
// with the default values initialized.
func NewGetEmailNotificationMessageUsingGETParams() *GetEmailNotificationMessageUsingGETParams {
	var ()
	return &GetEmailNotificationMessageUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetEmailNotificationMessageUsingGETParamsWithTimeout creates a new GetEmailNotificationMessageUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetEmailNotificationMessageUsingGETParamsWithTimeout(timeout time.Duration) *GetEmailNotificationMessageUsingGETParams {
	var ()
	return &GetEmailNotificationMessageUsingGETParams{

		timeout: timeout,
	}
}

// NewGetEmailNotificationMessageUsingGETParamsWithContext creates a new GetEmailNotificationMessageUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetEmailNotificationMessageUsingGETParamsWithContext(ctx context.Context) *GetEmailNotificationMessageUsingGETParams {
	var ()
	return &GetEmailNotificationMessageUsingGETParams{

		Context: ctx,
	}
}

// NewGetEmailNotificationMessageUsingGETParamsWithHTTPClient creates a new GetEmailNotificationMessageUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetEmailNotificationMessageUsingGETParamsWithHTTPClient(client *http.Client) *GetEmailNotificationMessageUsingGETParams {
	var ()
	return &GetEmailNotificationMessageUsingGETParams{
		HTTPClient: client,
	}
}

/*GetEmailNotificationMessageUsingGETParams contains all the parameters to send to the API endpoint
for the get email notification message using g e t operation typically these are written to a http.Request
*/
type GetEmailNotificationMessageUsingGETParams struct {

	/*EmailNotificationMessageName
	  msgName

	*/
	EmailNotificationMessageName string
	/*LanguageTag
	  language

	*/
	LanguageTag string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get email notification message using g e t params
func (o *GetEmailNotificationMessageUsingGETParams) WithTimeout(timeout time.Duration) *GetEmailNotificationMessageUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get email notification message using g e t params
func (o *GetEmailNotificationMessageUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get email notification message using g e t params
func (o *GetEmailNotificationMessageUsingGETParams) WithContext(ctx context.Context) *GetEmailNotificationMessageUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get email notification message using g e t params
func (o *GetEmailNotificationMessageUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get email notification message using g e t params
func (o *GetEmailNotificationMessageUsingGETParams) WithHTTPClient(client *http.Client) *GetEmailNotificationMessageUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get email notification message using g e t params
func (o *GetEmailNotificationMessageUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEmailNotificationMessageName adds the emailNotificationMessageName to the get email notification message using g e t params
func (o *GetEmailNotificationMessageUsingGETParams) WithEmailNotificationMessageName(emailNotificationMessageName string) *GetEmailNotificationMessageUsingGETParams {
	o.SetEmailNotificationMessageName(emailNotificationMessageName)
	return o
}

// SetEmailNotificationMessageName adds the emailNotificationMessageName to the get email notification message using g e t params
func (o *GetEmailNotificationMessageUsingGETParams) SetEmailNotificationMessageName(emailNotificationMessageName string) {
	o.EmailNotificationMessageName = emailNotificationMessageName
}

// WithLanguageTag adds the languageTag to the get email notification message using g e t params
func (o *GetEmailNotificationMessageUsingGETParams) WithLanguageTag(languageTag string) *GetEmailNotificationMessageUsingGETParams {
	o.SetLanguageTag(languageTag)
	return o
}

// SetLanguageTag adds the languageTag to the get email notification message using g e t params
func (o *GetEmailNotificationMessageUsingGETParams) SetLanguageTag(languageTag string) {
	o.LanguageTag = languageTag
}

// WriteToRequest writes these params to a swagger request
func (o *GetEmailNotificationMessageUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param emailNotificationMessageName
	if err := r.SetPathParam("emailNotificationMessageName", o.EmailNotificationMessageName); err != nil {
		return err
	}

	// path param languageTag
	if err := r.SetPathParam("languageTag", o.LanguageTag); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
