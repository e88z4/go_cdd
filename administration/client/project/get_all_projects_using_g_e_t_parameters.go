// Code generated by go-swagger; DO NOT EDIT.

package project

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

// NewGetAllProjectsUsingGETParams creates a new GetAllProjectsUsingGETParams object
// with the default values initialized.
func NewGetAllProjectsUsingGETParams() *GetAllProjectsUsingGETParams {
	var ()
	return &GetAllProjectsUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllProjectsUsingGETParamsWithTimeout creates a new GetAllProjectsUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAllProjectsUsingGETParamsWithTimeout(timeout time.Duration) *GetAllProjectsUsingGETParams {
	var ()
	return &GetAllProjectsUsingGETParams{

		timeout: timeout,
	}
}

// NewGetAllProjectsUsingGETParamsWithContext creates a new GetAllProjectsUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAllProjectsUsingGETParamsWithContext(ctx context.Context) *GetAllProjectsUsingGETParams {
	var ()
	return &GetAllProjectsUsingGETParams{

		Context: ctx,
	}
}

// NewGetAllProjectsUsingGETParamsWithHTTPClient creates a new GetAllProjectsUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAllProjectsUsingGETParamsWithHTTPClient(client *http.Client) *GetAllProjectsUsingGETParams {
	var ()
	return &GetAllProjectsUsingGETParams{
		HTTPClient: client,
	}
}

/*GetAllProjectsUsingGETParams contains all the parameters to send to the API endpoint
for the get all projects using g e t operation typically these are written to a http.Request
*/
type GetAllProjectsUsingGETParams struct {

	/*CurrentUser
	  currentUser

	*/
	CurrentUser *bool
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

// WithTimeout adds the timeout to the get all projects using g e t params
func (o *GetAllProjectsUsingGETParams) WithTimeout(timeout time.Duration) *GetAllProjectsUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get all projects using g e t params
func (o *GetAllProjectsUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get all projects using g e t params
func (o *GetAllProjectsUsingGETParams) WithContext(ctx context.Context) *GetAllProjectsUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get all projects using g e t params
func (o *GetAllProjectsUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get all projects using g e t params
func (o *GetAllProjectsUsingGETParams) WithHTTPClient(client *http.Client) *GetAllProjectsUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get all projects using g e t params
func (o *GetAllProjectsUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCurrentUser adds the currentUser to the get all projects using g e t params
func (o *GetAllProjectsUsingGETParams) WithCurrentUser(currentUser *bool) *GetAllProjectsUsingGETParams {
	o.SetCurrentUser(currentUser)
	return o
}

// SetCurrentUser adds the currentUser to the get all projects using g e t params
func (o *GetAllProjectsUsingGETParams) SetCurrentUser(currentUser *bool) {
	o.CurrentUser = currentUser
}

// WithFilter adds the filter to the get all projects using g e t params
func (o *GetAllProjectsUsingGETParams) WithFilter(filter *string) *GetAllProjectsUsingGETParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the get all projects using g e t params
func (o *GetAllProjectsUsingGETParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithPageNumber adds the pageNumber to the get all projects using g e t params
func (o *GetAllProjectsUsingGETParams) WithPageNumber(pageNumber *int32) *GetAllProjectsUsingGETParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get all projects using g e t params
func (o *GetAllProjectsUsingGETParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get all projects using g e t params
func (o *GetAllProjectsUsingGETParams) WithPageSize(pageSize *int32) *GetAllProjectsUsingGETParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get all projects using g e t params
func (o *GetAllProjectsUsingGETParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllProjectsUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CurrentUser != nil {

		// query param current_user
		var qrCurrentUser bool
		if o.CurrentUser != nil {
			qrCurrentUser = *o.CurrentUser
		}
		qCurrentUser := swag.FormatBool(qrCurrentUser)
		if qCurrentUser != "" {
			if err := r.SetQueryParam("current_user", qCurrentUser); err != nil {
				return err
			}
		}

	}

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
