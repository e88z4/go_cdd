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

// NewCreateAPIAccessTokenUsingPOST1Params creates a new CreateAPIAccessTokenUsingPOST1Params object
// with the default values initialized.
func NewCreateAPIAccessTokenUsingPOST1Params() *CreateAPIAccessTokenUsingPOST1Params {
	var ()
	return &CreateAPIAccessTokenUsingPOST1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateAPIAccessTokenUsingPOST1ParamsWithTimeout creates a new CreateAPIAccessTokenUsingPOST1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateAPIAccessTokenUsingPOST1ParamsWithTimeout(timeout time.Duration) *CreateAPIAccessTokenUsingPOST1Params {
	var ()
	return &CreateAPIAccessTokenUsingPOST1Params{

		timeout: timeout,
	}
}

// NewCreateAPIAccessTokenUsingPOST1ParamsWithContext creates a new CreateAPIAccessTokenUsingPOST1Params object
// with the default values initialized, and the ability to set a context for a request
func NewCreateAPIAccessTokenUsingPOST1ParamsWithContext(ctx context.Context) *CreateAPIAccessTokenUsingPOST1Params {
	var ()
	return &CreateAPIAccessTokenUsingPOST1Params{

		Context: ctx,
	}
}

// NewCreateAPIAccessTokenUsingPOST1ParamsWithHTTPClient creates a new CreateAPIAccessTokenUsingPOST1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateAPIAccessTokenUsingPOST1ParamsWithHTTPClient(client *http.Client) *CreateAPIAccessTokenUsingPOST1Params {
	var ()
	return &CreateAPIAccessTokenUsingPOST1Params{
		HTTPClient: client,
	}
}

/*CreateAPIAccessTokenUsingPOST1Params contains all the parameters to send to the API endpoint
for the create Api access token using p o s t 1 operation typically these are written to a http.Request
*/
type CreateAPIAccessTokenUsingPOST1Params struct {

	/*UserID
	  userId

	*/
	UserID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create Api access token using p o s t 1 params
func (o *CreateAPIAccessTokenUsingPOST1Params) WithTimeout(timeout time.Duration) *CreateAPIAccessTokenUsingPOST1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create Api access token using p o s t 1 params
func (o *CreateAPIAccessTokenUsingPOST1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create Api access token using p o s t 1 params
func (o *CreateAPIAccessTokenUsingPOST1Params) WithContext(ctx context.Context) *CreateAPIAccessTokenUsingPOST1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create Api access token using p o s t 1 params
func (o *CreateAPIAccessTokenUsingPOST1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create Api access token using p o s t 1 params
func (o *CreateAPIAccessTokenUsingPOST1Params) WithHTTPClient(client *http.Client) *CreateAPIAccessTokenUsingPOST1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create Api access token using p o s t 1 params
func (o *CreateAPIAccessTokenUsingPOST1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserID adds the userID to the create Api access token using p o s t 1 params
func (o *CreateAPIAccessTokenUsingPOST1Params) WithUserID(userID int64) *CreateAPIAccessTokenUsingPOST1Params {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the create Api access token using p o s t 1 params
func (o *CreateAPIAccessTokenUsingPOST1Params) SetUserID(userID int64) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateAPIAccessTokenUsingPOST1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
