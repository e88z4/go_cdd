// Code generated by go-swagger; DO NOT EDIT.

package freeze_period

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

// NewGetFreezePeriodsUsingGETParams creates a new GetFreezePeriodsUsingGETParams object
// with the default values initialized.
func NewGetFreezePeriodsUsingGETParams() *GetFreezePeriodsUsingGETParams {
	var ()
	return &GetFreezePeriodsUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetFreezePeriodsUsingGETParamsWithTimeout creates a new GetFreezePeriodsUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetFreezePeriodsUsingGETParamsWithTimeout(timeout time.Duration) *GetFreezePeriodsUsingGETParams {
	var ()
	return &GetFreezePeriodsUsingGETParams{

		timeout: timeout,
	}
}

// NewGetFreezePeriodsUsingGETParamsWithContext creates a new GetFreezePeriodsUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetFreezePeriodsUsingGETParamsWithContext(ctx context.Context) *GetFreezePeriodsUsingGETParams {
	var ()
	return &GetFreezePeriodsUsingGETParams{

		Context: ctx,
	}
}

// NewGetFreezePeriodsUsingGETParamsWithHTTPClient creates a new GetFreezePeriodsUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetFreezePeriodsUsingGETParamsWithHTTPClient(client *http.Client) *GetFreezePeriodsUsingGETParams {
	var ()
	return &GetFreezePeriodsUsingGETParams{
		HTTPClient: client,
	}
}

/*GetFreezePeriodsUsingGETParams contains all the parameters to send to the API endpoint
for the get freeze periods using g e t operation typically these are written to a http.Request
*/
type GetFreezePeriodsUsingGETParams struct {

	/*Filter
	  freeTextFilter

	*/
	Filter *string
	/*PageNumber
	  pageNumber

	*/
	PageNumber *int32
	/*PageSize
	  pageSize

	*/
	PageSize *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get freeze periods using g e t params
func (o *GetFreezePeriodsUsingGETParams) WithTimeout(timeout time.Duration) *GetFreezePeriodsUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get freeze periods using g e t params
func (o *GetFreezePeriodsUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get freeze periods using g e t params
func (o *GetFreezePeriodsUsingGETParams) WithContext(ctx context.Context) *GetFreezePeriodsUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get freeze periods using g e t params
func (o *GetFreezePeriodsUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get freeze periods using g e t params
func (o *GetFreezePeriodsUsingGETParams) WithHTTPClient(client *http.Client) *GetFreezePeriodsUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get freeze periods using g e t params
func (o *GetFreezePeriodsUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilter adds the filter to the get freeze periods using g e t params
func (o *GetFreezePeriodsUsingGETParams) WithFilter(filter *string) *GetFreezePeriodsUsingGETParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the get freeze periods using g e t params
func (o *GetFreezePeriodsUsingGETParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithPageNumber adds the pageNumber to the get freeze periods using g e t params
func (o *GetFreezePeriodsUsingGETParams) WithPageNumber(pageNumber *int32) *GetFreezePeriodsUsingGETParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get freeze periods using g e t params
func (o *GetFreezePeriodsUsingGETParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get freeze periods using g e t params
func (o *GetFreezePeriodsUsingGETParams) WithPageSize(pageSize *int32) *GetFreezePeriodsUsingGETParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get freeze periods using g e t params
func (o *GetFreezePeriodsUsingGETParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WriteToRequest writes these params to a swagger request
func (o *GetFreezePeriodsUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Filter != nil {

		// query param filter
		var qrFilter string
		if o.Filter != nil {
			qrFilter = *o.Filter
		}
		qFilter := qrFilter
		if qFilter != "" {
			if err := r.SetQueryParam("filter", qFilter); err != nil {
				return err
			}
		}

	}

	if o.PageNumber != nil {

		// query param page_number
		var qrPageNumber int32
		if o.PageNumber != nil {
			qrPageNumber = *o.PageNumber
		}
		qPageNumber := swag.FormatInt32(qrPageNumber)
		if qPageNumber != "" {
			if err := r.SetQueryParam("page_number", qPageNumber); err != nil {
				return err
			}
		}

	}

	if o.PageSize != nil {

		// query param page_size
		var qrPageSize int32
		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt32(qrPageSize)
		if qPageSize != "" {
			if err := r.SetQueryParam("page_size", qPageSize); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
