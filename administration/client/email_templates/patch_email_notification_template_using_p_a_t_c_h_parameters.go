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

	"github.com/e88z4/go_cdd/administration/models"
)

// NewPatchEmailNotificationTemplateUsingPATCHParams creates a new PatchEmailNotificationTemplateUsingPATCHParams object
// with the default values initialized.
func NewPatchEmailNotificationTemplateUsingPATCHParams() *PatchEmailNotificationTemplateUsingPATCHParams {
	var ()
	return &PatchEmailNotificationTemplateUsingPATCHParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchEmailNotificationTemplateUsingPATCHParamsWithTimeout creates a new PatchEmailNotificationTemplateUsingPATCHParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchEmailNotificationTemplateUsingPATCHParamsWithTimeout(timeout time.Duration) *PatchEmailNotificationTemplateUsingPATCHParams {
	var ()
	return &PatchEmailNotificationTemplateUsingPATCHParams{

		timeout: timeout,
	}
}

// NewPatchEmailNotificationTemplateUsingPATCHParamsWithContext creates a new PatchEmailNotificationTemplateUsingPATCHParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchEmailNotificationTemplateUsingPATCHParamsWithContext(ctx context.Context) *PatchEmailNotificationTemplateUsingPATCHParams {
	var ()
	return &PatchEmailNotificationTemplateUsingPATCHParams{

		Context: ctx,
	}
}

// NewPatchEmailNotificationTemplateUsingPATCHParamsWithHTTPClient creates a new PatchEmailNotificationTemplateUsingPATCHParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchEmailNotificationTemplateUsingPATCHParamsWithHTTPClient(client *http.Client) *PatchEmailNotificationTemplateUsingPATCHParams {
	var ()
	return &PatchEmailNotificationTemplateUsingPATCHParams{
		HTTPClient: client,
	}
}

/*PatchEmailNotificationTemplateUsingPATCHParams contains all the parameters to send to the API endpoint
for the patch email notification template using p a t c h operation typically these are written to a http.Request
*/
type PatchEmailNotificationTemplateUsingPATCHParams struct {

	/*EmailNotificationTemplateDto
	  emailNotificationTemplateDto

	*/
	EmailNotificationTemplateDto *models.EmailNotificationTemplateDto
	/*EmailNotificationTemplateName
	  templateName

	*/
	EmailNotificationTemplateName string
	/*LanguageTag
	  language

	*/
	LanguageTag string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch email notification template using p a t c h params
func (o *PatchEmailNotificationTemplateUsingPATCHParams) WithTimeout(timeout time.Duration) *PatchEmailNotificationTemplateUsingPATCHParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch email notification template using p a t c h params
func (o *PatchEmailNotificationTemplateUsingPATCHParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch email notification template using p a t c h params
func (o *PatchEmailNotificationTemplateUsingPATCHParams) WithContext(ctx context.Context) *PatchEmailNotificationTemplateUsingPATCHParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch email notification template using p a t c h params
func (o *PatchEmailNotificationTemplateUsingPATCHParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch email notification template using p a t c h params
func (o *PatchEmailNotificationTemplateUsingPATCHParams) WithHTTPClient(client *http.Client) *PatchEmailNotificationTemplateUsingPATCHParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch email notification template using p a t c h params
func (o *PatchEmailNotificationTemplateUsingPATCHParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEmailNotificationTemplateDto adds the emailNotificationTemplateDto to the patch email notification template using p a t c h params
func (o *PatchEmailNotificationTemplateUsingPATCHParams) WithEmailNotificationTemplateDto(emailNotificationTemplateDto *models.EmailNotificationTemplateDto) *PatchEmailNotificationTemplateUsingPATCHParams {
	o.SetEmailNotificationTemplateDto(emailNotificationTemplateDto)
	return o
}

// SetEmailNotificationTemplateDto adds the emailNotificationTemplateDto to the patch email notification template using p a t c h params
func (o *PatchEmailNotificationTemplateUsingPATCHParams) SetEmailNotificationTemplateDto(emailNotificationTemplateDto *models.EmailNotificationTemplateDto) {
	o.EmailNotificationTemplateDto = emailNotificationTemplateDto
}

// WithEmailNotificationTemplateName adds the emailNotificationTemplateName to the patch email notification template using p a t c h params
func (o *PatchEmailNotificationTemplateUsingPATCHParams) WithEmailNotificationTemplateName(emailNotificationTemplateName string) *PatchEmailNotificationTemplateUsingPATCHParams {
	o.SetEmailNotificationTemplateName(emailNotificationTemplateName)
	return o
}

// SetEmailNotificationTemplateName adds the emailNotificationTemplateName to the patch email notification template using p a t c h params
func (o *PatchEmailNotificationTemplateUsingPATCHParams) SetEmailNotificationTemplateName(emailNotificationTemplateName string) {
	o.EmailNotificationTemplateName = emailNotificationTemplateName
}

// WithLanguageTag adds the languageTag to the patch email notification template using p a t c h params
func (o *PatchEmailNotificationTemplateUsingPATCHParams) WithLanguageTag(languageTag string) *PatchEmailNotificationTemplateUsingPATCHParams {
	o.SetLanguageTag(languageTag)
	return o
}

// SetLanguageTag adds the languageTag to the patch email notification template using p a t c h params
func (o *PatchEmailNotificationTemplateUsingPATCHParams) SetLanguageTag(languageTag string) {
	o.LanguageTag = languageTag
}

// WriteToRequest writes these params to a swagger request
func (o *PatchEmailNotificationTemplateUsingPATCHParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.EmailNotificationTemplateDto != nil {
		if err := r.SetBodyParam(o.EmailNotificationTemplateDto); err != nil {
			return err
		}
	}

	// path param emailNotificationTemplateName
	if err := r.SetPathParam("emailNotificationTemplateName", o.EmailNotificationTemplateName); err != nil {
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
