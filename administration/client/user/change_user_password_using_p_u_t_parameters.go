// Code generated by go-swagger; DO NOT EDIT.

package user

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

// NewChangeUserPasswordUsingPUTParams creates a new ChangeUserPasswordUsingPUTParams object
// with the default values initialized.
func NewChangeUserPasswordUsingPUTParams() *ChangeUserPasswordUsingPUTParams {
	var ()
	return &ChangeUserPasswordUsingPUTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewChangeUserPasswordUsingPUTParamsWithTimeout creates a new ChangeUserPasswordUsingPUTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewChangeUserPasswordUsingPUTParamsWithTimeout(timeout time.Duration) *ChangeUserPasswordUsingPUTParams {
	var ()
	return &ChangeUserPasswordUsingPUTParams{

		timeout: timeout,
	}
}

// NewChangeUserPasswordUsingPUTParamsWithContext creates a new ChangeUserPasswordUsingPUTParams object
// with the default values initialized, and the ability to set a context for a request
func NewChangeUserPasswordUsingPUTParamsWithContext(ctx context.Context) *ChangeUserPasswordUsingPUTParams {
	var ()
	return &ChangeUserPasswordUsingPUTParams{

		Context: ctx,
	}
}

// NewChangeUserPasswordUsingPUTParamsWithHTTPClient creates a new ChangeUserPasswordUsingPUTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewChangeUserPasswordUsingPUTParamsWithHTTPClient(client *http.Client) *ChangeUserPasswordUsingPUTParams {
	var ()
	return &ChangeUserPasswordUsingPUTParams{
		HTTPClient: client,
	}
}

/*ChangeUserPasswordUsingPUTParams contains all the parameters to send to the API endpoint
for the change user password using p u t operation typically these are written to a http.Request
*/
type ChangeUserPasswordUsingPUTParams struct {

	/*ChangeUserPasswordDto
	  changeUserPasswordDto

	*/
	ChangeUserPasswordDto *models.ChangeUserPasswordDto
	/*UserID
	  userId

	*/
	UserID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the change user password using p u t params
func (o *ChangeUserPasswordUsingPUTParams) WithTimeout(timeout time.Duration) *ChangeUserPasswordUsingPUTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the change user password using p u t params
func (o *ChangeUserPasswordUsingPUTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the change user password using p u t params
func (o *ChangeUserPasswordUsingPUTParams) WithContext(ctx context.Context) *ChangeUserPasswordUsingPUTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the change user password using p u t params
func (o *ChangeUserPasswordUsingPUTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the change user password using p u t params
func (o *ChangeUserPasswordUsingPUTParams) WithHTTPClient(client *http.Client) *ChangeUserPasswordUsingPUTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the change user password using p u t params
func (o *ChangeUserPasswordUsingPUTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChangeUserPasswordDto adds the changeUserPasswordDto to the change user password using p u t params
func (o *ChangeUserPasswordUsingPUTParams) WithChangeUserPasswordDto(changeUserPasswordDto *models.ChangeUserPasswordDto) *ChangeUserPasswordUsingPUTParams {
	o.SetChangeUserPasswordDto(changeUserPasswordDto)
	return o
}

// SetChangeUserPasswordDto adds the changeUserPasswordDto to the change user password using p u t params
func (o *ChangeUserPasswordUsingPUTParams) SetChangeUserPasswordDto(changeUserPasswordDto *models.ChangeUserPasswordDto) {
	o.ChangeUserPasswordDto = changeUserPasswordDto
}

// WithUserID adds the userID to the change user password using p u t params
func (o *ChangeUserPasswordUsingPUTParams) WithUserID(userID int64) *ChangeUserPasswordUsingPUTParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the change user password using p u t params
func (o *ChangeUserPasswordUsingPUTParams) SetUserID(userID int64) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *ChangeUserPasswordUsingPUTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ChangeUserPasswordDto != nil {
		if err := r.SetBodyParam(o.ChangeUserPasswordDto); err != nil {
			return err
		}
	}

	// path param userId
	if err := r.SetPathParam("userId", swag.FormatInt64(o.UserID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
