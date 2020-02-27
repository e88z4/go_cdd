// Code generated by go-swagger; DO NOT EDIT.

package product

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

// NewGetPortfolioLicensingAgreementUsingGETParams creates a new GetPortfolioLicensingAgreementUsingGETParams object
// with the default values initialized.
func NewGetPortfolioLicensingAgreementUsingGETParams() *GetPortfolioLicensingAgreementUsingGETParams {

	return &GetPortfolioLicensingAgreementUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPortfolioLicensingAgreementUsingGETParamsWithTimeout creates a new GetPortfolioLicensingAgreementUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPortfolioLicensingAgreementUsingGETParamsWithTimeout(timeout time.Duration) *GetPortfolioLicensingAgreementUsingGETParams {

	return &GetPortfolioLicensingAgreementUsingGETParams{

		timeout: timeout,
	}
}

// NewGetPortfolioLicensingAgreementUsingGETParamsWithContext creates a new GetPortfolioLicensingAgreementUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPortfolioLicensingAgreementUsingGETParamsWithContext(ctx context.Context) *GetPortfolioLicensingAgreementUsingGETParams {

	return &GetPortfolioLicensingAgreementUsingGETParams{

		Context: ctx,
	}
}

// NewGetPortfolioLicensingAgreementUsingGETParamsWithHTTPClient creates a new GetPortfolioLicensingAgreementUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPortfolioLicensingAgreementUsingGETParamsWithHTTPClient(client *http.Client) *GetPortfolioLicensingAgreementUsingGETParams {

	return &GetPortfolioLicensingAgreementUsingGETParams{
		HTTPClient: client,
	}
}

/*GetPortfolioLicensingAgreementUsingGETParams contains all the parameters to send to the API endpoint
for the get portfolio licensing agreement using g e t operation typically these are written to a http.Request
*/
type GetPortfolioLicensingAgreementUsingGETParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get portfolio licensing agreement using g e t params
func (o *GetPortfolioLicensingAgreementUsingGETParams) WithTimeout(timeout time.Duration) *GetPortfolioLicensingAgreementUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get portfolio licensing agreement using g e t params
func (o *GetPortfolioLicensingAgreementUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get portfolio licensing agreement using g e t params
func (o *GetPortfolioLicensingAgreementUsingGETParams) WithContext(ctx context.Context) *GetPortfolioLicensingAgreementUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get portfolio licensing agreement using g e t params
func (o *GetPortfolioLicensingAgreementUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get portfolio licensing agreement using g e t params
func (o *GetPortfolioLicensingAgreementUsingGETParams) WithHTTPClient(client *http.Client) *GetPortfolioLicensingAgreementUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get portfolio licensing agreement using g e t params
func (o *GetPortfolioLicensingAgreementUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetPortfolioLicensingAgreementUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}