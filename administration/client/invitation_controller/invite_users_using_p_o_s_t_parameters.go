// Code generated by go-swagger; DO NOT EDIT.

package invitation_controller

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

// NewInviteUsersUsingPOSTParams creates a new InviteUsersUsingPOSTParams object
// with the default values initialized.
func NewInviteUsersUsingPOSTParams() *InviteUsersUsingPOSTParams {
	var ()
	return &InviteUsersUsingPOSTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewInviteUsersUsingPOSTParamsWithTimeout creates a new InviteUsersUsingPOSTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewInviteUsersUsingPOSTParamsWithTimeout(timeout time.Duration) *InviteUsersUsingPOSTParams {
	var ()
	return &InviteUsersUsingPOSTParams{

		timeout: timeout,
	}
}

// NewInviteUsersUsingPOSTParamsWithContext creates a new InviteUsersUsingPOSTParams object
// with the default values initialized, and the ability to set a context for a request
func NewInviteUsersUsingPOSTParamsWithContext(ctx context.Context) *InviteUsersUsingPOSTParams {
	var ()
	return &InviteUsersUsingPOSTParams{

		Context: ctx,
	}
}

// NewInviteUsersUsingPOSTParamsWithHTTPClient creates a new InviteUsersUsingPOSTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewInviteUsersUsingPOSTParamsWithHTTPClient(client *http.Client) *InviteUsersUsingPOSTParams {
	var ()
	return &InviteUsersUsingPOSTParams{
		HTTPClient: client,
	}
}

/*InviteUsersUsingPOSTParams contains all the parameters to send to the API endpoint
for the invite users using p o s t operation typically these are written to a http.Request
*/
type InviteUsersUsingPOSTParams struct {

	/*InvitationDto
	  invitationDto

	*/
	InvitationDto *models.InvitationDto

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the invite users using p o s t params
func (o *InviteUsersUsingPOSTParams) WithTimeout(timeout time.Duration) *InviteUsersUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the invite users using p o s t params
func (o *InviteUsersUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the invite users using p o s t params
func (o *InviteUsersUsingPOSTParams) WithContext(ctx context.Context) *InviteUsersUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the invite users using p o s t params
func (o *InviteUsersUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the invite users using p o s t params
func (o *InviteUsersUsingPOSTParams) WithHTTPClient(client *http.Client) *InviteUsersUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the invite users using p o s t params
func (o *InviteUsersUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInvitationDto adds the invitationDto to the invite users using p o s t params
func (o *InviteUsersUsingPOSTParams) WithInvitationDto(invitationDto *models.InvitationDto) *InviteUsersUsingPOSTParams {
	o.SetInvitationDto(invitationDto)
	return o
}

// SetInvitationDto adds the invitationDto to the invite users using p o s t params
func (o *InviteUsersUsingPOSTParams) SetInvitationDto(invitationDto *models.InvitationDto) {
	o.InvitationDto = invitationDto
}

// WriteToRequest writes these params to a swagger request
func (o *InviteUsersUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.InvitationDto != nil {
		if err := r.SetBodyParam(o.InvitationDto); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}