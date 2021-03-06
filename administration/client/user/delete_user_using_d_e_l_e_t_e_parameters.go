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
)

// NewDeleteUserUsingDELETEParams creates a new DeleteUserUsingDELETEParams object
// with the default values initialized.
func NewDeleteUserUsingDELETEParams() *DeleteUserUsingDELETEParams {
	var ()
	return &DeleteUserUsingDELETEParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteUserUsingDELETEParamsWithTimeout creates a new DeleteUserUsingDELETEParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteUserUsingDELETEParamsWithTimeout(timeout time.Duration) *DeleteUserUsingDELETEParams {
	var ()
	return &DeleteUserUsingDELETEParams{

		timeout: timeout,
	}
}

// NewDeleteUserUsingDELETEParamsWithContext creates a new DeleteUserUsingDELETEParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteUserUsingDELETEParamsWithContext(ctx context.Context) *DeleteUserUsingDELETEParams {
	var ()
	return &DeleteUserUsingDELETEParams{

		Context: ctx,
	}
}

// NewDeleteUserUsingDELETEParamsWithHTTPClient creates a new DeleteUserUsingDELETEParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteUserUsingDELETEParamsWithHTTPClient(client *http.Client) *DeleteUserUsingDELETEParams {
	var ()
	return &DeleteUserUsingDELETEParams{
		HTTPClient: client,
	}
}

/*DeleteUserUsingDELETEParams contains all the parameters to send to the API endpoint
for the delete user using d e l e t e operation typically these are written to a http.Request
*/
type DeleteUserUsingDELETEParams struct {

	/*UserID
	  userId

	*/
	UserID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete user using d e l e t e params
func (o *DeleteUserUsingDELETEParams) WithTimeout(timeout time.Duration) *DeleteUserUsingDELETEParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete user using d e l e t e params
func (o *DeleteUserUsingDELETEParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete user using d e l e t e params
func (o *DeleteUserUsingDELETEParams) WithContext(ctx context.Context) *DeleteUserUsingDELETEParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete user using d e l e t e params
func (o *DeleteUserUsingDELETEParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete user using d e l e t e params
func (o *DeleteUserUsingDELETEParams) WithHTTPClient(client *http.Client) *DeleteUserUsingDELETEParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete user using d e l e t e params
func (o *DeleteUserUsingDELETEParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserID adds the userID to the delete user using d e l e t e params
func (o *DeleteUserUsingDELETEParams) WithUserID(userID int64) *DeleteUserUsingDELETEParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the delete user using d e l e t e params
func (o *DeleteUserUsingDELETEParams) SetUserID(userID int64) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteUserUsingDELETEParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param userId
	if err := r.SetPathParam("userId", swag.FormatInt64(o.UserID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
