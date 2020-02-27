// Code generated by go-swagger; DO NOT EDIT.

package payment_plan_controller

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

// NewRedirectToPurchasePlanUsingGETParams creates a new RedirectToPurchasePlanUsingGETParams object
// with the default values initialized.
func NewRedirectToPurchasePlanUsingGETParams() *RedirectToPurchasePlanUsingGETParams {
	var ()
	return &RedirectToPurchasePlanUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRedirectToPurchasePlanUsingGETParamsWithTimeout creates a new RedirectToPurchasePlanUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRedirectToPurchasePlanUsingGETParamsWithTimeout(timeout time.Duration) *RedirectToPurchasePlanUsingGETParams {
	var ()
	return &RedirectToPurchasePlanUsingGETParams{

		timeout: timeout,
	}
}

// NewRedirectToPurchasePlanUsingGETParamsWithContext creates a new RedirectToPurchasePlanUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewRedirectToPurchasePlanUsingGETParamsWithContext(ctx context.Context) *RedirectToPurchasePlanUsingGETParams {
	var ()
	return &RedirectToPurchasePlanUsingGETParams{

		Context: ctx,
	}
}

// NewRedirectToPurchasePlanUsingGETParamsWithHTTPClient creates a new RedirectToPurchasePlanUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRedirectToPurchasePlanUsingGETParamsWithHTTPClient(client *http.Client) *RedirectToPurchasePlanUsingGETParams {
	var ()
	return &RedirectToPurchasePlanUsingGETParams{
		HTTPClient: client,
	}
}

/*RedirectToPurchasePlanUsingGETParams contains all the parameters to send to the API endpoint
for the redirect to purchase plan using g e t operation typically these are written to a http.Request
*/
type RedirectToPurchasePlanUsingGETParams struct {

	/*Plan
	  paymentPlan

	*/
	Plan *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the redirect to purchase plan using g e t params
func (o *RedirectToPurchasePlanUsingGETParams) WithTimeout(timeout time.Duration) *RedirectToPurchasePlanUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the redirect to purchase plan using g e t params
func (o *RedirectToPurchasePlanUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the redirect to purchase plan using g e t params
func (o *RedirectToPurchasePlanUsingGETParams) WithContext(ctx context.Context) *RedirectToPurchasePlanUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the redirect to purchase plan using g e t params
func (o *RedirectToPurchasePlanUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the redirect to purchase plan using g e t params
func (o *RedirectToPurchasePlanUsingGETParams) WithHTTPClient(client *http.Client) *RedirectToPurchasePlanUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the redirect to purchase plan using g e t params
func (o *RedirectToPurchasePlanUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPlan adds the plan to the redirect to purchase plan using g e t params
func (o *RedirectToPurchasePlanUsingGETParams) WithPlan(plan *string) *RedirectToPurchasePlanUsingGETParams {
	o.SetPlan(plan)
	return o
}

// SetPlan adds the plan to the redirect to purchase plan using g e t params
func (o *RedirectToPurchasePlanUsingGETParams) SetPlan(plan *string) {
	o.Plan = plan
}

// WriteToRequest writes these params to a swagger request
func (o *RedirectToPurchasePlanUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Plan != nil {

		// query param plan
		var qrPlan string
		if o.Plan != nil {
			qrPlan = *o.Plan
		}
		qPlan := qrPlan
		if qPlan != "" {
			if err := r.SetQueryParam("plan", qPlan); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
