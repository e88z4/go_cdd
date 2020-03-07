// Code generated by go-swagger; DO NOT EDIT.

package saml_domain

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

// NewDeleteSamlDomainUsingDELETEParams creates a new DeleteSamlDomainUsingDELETEParams object
// with the default values initialized.
func NewDeleteSamlDomainUsingDELETEParams() *DeleteSamlDomainUsingDELETEParams {
	var ()
	return &DeleteSamlDomainUsingDELETEParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSamlDomainUsingDELETEParamsWithTimeout creates a new DeleteSamlDomainUsingDELETEParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteSamlDomainUsingDELETEParamsWithTimeout(timeout time.Duration) *DeleteSamlDomainUsingDELETEParams {
	var ()
	return &DeleteSamlDomainUsingDELETEParams{

		timeout: timeout,
	}
}

// NewDeleteSamlDomainUsingDELETEParamsWithContext creates a new DeleteSamlDomainUsingDELETEParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteSamlDomainUsingDELETEParamsWithContext(ctx context.Context) *DeleteSamlDomainUsingDELETEParams {
	var ()
	return &DeleteSamlDomainUsingDELETEParams{

		Context: ctx,
	}
}

// NewDeleteSamlDomainUsingDELETEParamsWithHTTPClient creates a new DeleteSamlDomainUsingDELETEParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteSamlDomainUsingDELETEParamsWithHTTPClient(client *http.Client) *DeleteSamlDomainUsingDELETEParams {
	var ()
	return &DeleteSamlDomainUsingDELETEParams{
		HTTPClient: client,
	}
}

/*DeleteSamlDomainUsingDELETEParams contains all the parameters to send to the API endpoint
for the delete saml domain using d e l e t e operation typically these are written to a http.Request
*/
type DeleteSamlDomainUsingDELETEParams struct {

	/*SamlDomainID
	  samlDomainId

	*/
	SamlDomainID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete saml domain using d e l e t e params
func (o *DeleteSamlDomainUsingDELETEParams) WithTimeout(timeout time.Duration) *DeleteSamlDomainUsingDELETEParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete saml domain using d e l e t e params
func (o *DeleteSamlDomainUsingDELETEParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete saml domain using d e l e t e params
func (o *DeleteSamlDomainUsingDELETEParams) WithContext(ctx context.Context) *DeleteSamlDomainUsingDELETEParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete saml domain using d e l e t e params
func (o *DeleteSamlDomainUsingDELETEParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete saml domain using d e l e t e params
func (o *DeleteSamlDomainUsingDELETEParams) WithHTTPClient(client *http.Client) *DeleteSamlDomainUsingDELETEParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete saml domain using d e l e t e params
func (o *DeleteSamlDomainUsingDELETEParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSamlDomainID adds the samlDomainID to the delete saml domain using d e l e t e params
func (o *DeleteSamlDomainUsingDELETEParams) WithSamlDomainID(samlDomainID int64) *DeleteSamlDomainUsingDELETEParams {
	o.SetSamlDomainID(samlDomainID)
	return o
}

// SetSamlDomainID adds the samlDomainId to the delete saml domain using d e l e t e params
func (o *DeleteSamlDomainUsingDELETEParams) SetSamlDomainID(samlDomainID int64) {
	o.SamlDomainID = samlDomainID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSamlDomainUsingDELETEParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param samlDomainId
	if err := r.SetPathParam("samlDomainId", swag.FormatInt64(o.SamlDomainID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}