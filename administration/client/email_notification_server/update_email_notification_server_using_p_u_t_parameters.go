// Code generated by go-swagger; DO NOT EDIT.

package email_notification_server

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

	"github.com/e88z4/go_cdd/administration/models"
)

// NewUpdateEmailNotificationServerUsingPUTParams creates a new UpdateEmailNotificationServerUsingPUTParams object
// with the default values initialized.
func NewUpdateEmailNotificationServerUsingPUTParams() *UpdateEmailNotificationServerUsingPUTParams {
	var ()
	return &UpdateEmailNotificationServerUsingPUTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateEmailNotificationServerUsingPUTParamsWithTimeout creates a new UpdateEmailNotificationServerUsingPUTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateEmailNotificationServerUsingPUTParamsWithTimeout(timeout time.Duration) *UpdateEmailNotificationServerUsingPUTParams {
	var ()
	return &UpdateEmailNotificationServerUsingPUTParams{

		timeout: timeout,
	}
}

// NewUpdateEmailNotificationServerUsingPUTParamsWithContext creates a new UpdateEmailNotificationServerUsingPUTParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateEmailNotificationServerUsingPUTParamsWithContext(ctx context.Context) *UpdateEmailNotificationServerUsingPUTParams {
	var ()
	return &UpdateEmailNotificationServerUsingPUTParams{

		Context: ctx,
	}
}

// NewUpdateEmailNotificationServerUsingPUTParamsWithHTTPClient creates a new UpdateEmailNotificationServerUsingPUTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateEmailNotificationServerUsingPUTParamsWithHTTPClient(client *http.Client) *UpdateEmailNotificationServerUsingPUTParams {
	var ()
	return &UpdateEmailNotificationServerUsingPUTParams{
		HTTPClient: client,
	}
}

/*UpdateEmailNotificationServerUsingPUTParams contains all the parameters to send to the API endpoint
for the update email notification server using p u t operation typically these are written to a http.Request
*/
type UpdateEmailNotificationServerUsingPUTParams struct {

	/*EmailNotificationServerDto
	  emailNotificationServerDto

	*/
	EmailNotificationServerDto *models.EmailNotificationServerDto
	/*EmailNotificationServerID
	  emailNotificationServerId

	*/
	EmailNotificationServerID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update email notification server using p u t params
func (o *UpdateEmailNotificationServerUsingPUTParams) WithTimeout(timeout time.Duration) *UpdateEmailNotificationServerUsingPUTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update email notification server using p u t params
func (o *UpdateEmailNotificationServerUsingPUTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update email notification server using p u t params
func (o *UpdateEmailNotificationServerUsingPUTParams) WithContext(ctx context.Context) *UpdateEmailNotificationServerUsingPUTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update email notification server using p u t params
func (o *UpdateEmailNotificationServerUsingPUTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update email notification server using p u t params
func (o *UpdateEmailNotificationServerUsingPUTParams) WithHTTPClient(client *http.Client) *UpdateEmailNotificationServerUsingPUTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update email notification server using p u t params
func (o *UpdateEmailNotificationServerUsingPUTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEmailNotificationServerDto adds the emailNotificationServerDto to the update email notification server using p u t params
func (o *UpdateEmailNotificationServerUsingPUTParams) WithEmailNotificationServerDto(emailNotificationServerDto *models.EmailNotificationServerDto) *UpdateEmailNotificationServerUsingPUTParams {
	o.SetEmailNotificationServerDto(emailNotificationServerDto)
	return o
}

// SetEmailNotificationServerDto adds the emailNotificationServerDto to the update email notification server using p u t params
func (o *UpdateEmailNotificationServerUsingPUTParams) SetEmailNotificationServerDto(emailNotificationServerDto *models.EmailNotificationServerDto) {
	o.EmailNotificationServerDto = emailNotificationServerDto
}

// WithEmailNotificationServerID adds the emailNotificationServerID to the update email notification server using p u t params
func (o *UpdateEmailNotificationServerUsingPUTParams) WithEmailNotificationServerID(emailNotificationServerID int64) *UpdateEmailNotificationServerUsingPUTParams {
	o.SetEmailNotificationServerID(emailNotificationServerID)
	return o
}

// SetEmailNotificationServerID adds the emailNotificationServerId to the update email notification server using p u t params
func (o *UpdateEmailNotificationServerUsingPUTParams) SetEmailNotificationServerID(emailNotificationServerID int64) {
	o.EmailNotificationServerID = emailNotificationServerID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateEmailNotificationServerUsingPUTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.EmailNotificationServerDto != nil {
		if err := r.SetBodyParam(o.EmailNotificationServerDto); err != nil {
			return err
		}
	}

	// path param emailNotificationServerId
	if err := r.SetPathParam("emailNotificationServerId", swag.FormatInt64(o.EmailNotificationServerID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}