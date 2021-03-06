// Code generated by go-swagger; DO NOT EDIT.

package directory_server_controller

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

// NewUpdateDirectoryServerUsingPUTParams creates a new UpdateDirectoryServerUsingPUTParams object
// with the default values initialized.
func NewUpdateDirectoryServerUsingPUTParams() *UpdateDirectoryServerUsingPUTParams {
	var ()
	return &UpdateDirectoryServerUsingPUTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateDirectoryServerUsingPUTParamsWithTimeout creates a new UpdateDirectoryServerUsingPUTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateDirectoryServerUsingPUTParamsWithTimeout(timeout time.Duration) *UpdateDirectoryServerUsingPUTParams {
	var ()
	return &UpdateDirectoryServerUsingPUTParams{

		timeout: timeout,
	}
}

// NewUpdateDirectoryServerUsingPUTParamsWithContext creates a new UpdateDirectoryServerUsingPUTParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateDirectoryServerUsingPUTParamsWithContext(ctx context.Context) *UpdateDirectoryServerUsingPUTParams {
	var ()
	return &UpdateDirectoryServerUsingPUTParams{

		Context: ctx,
	}
}

// NewUpdateDirectoryServerUsingPUTParamsWithHTTPClient creates a new UpdateDirectoryServerUsingPUTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateDirectoryServerUsingPUTParamsWithHTTPClient(client *http.Client) *UpdateDirectoryServerUsingPUTParams {
	var ()
	return &UpdateDirectoryServerUsingPUTParams{
		HTTPClient: client,
	}
}

/*UpdateDirectoryServerUsingPUTParams contains all the parameters to send to the API endpoint
for the update directory server using p u t operation typically these are written to a http.Request
*/
type UpdateDirectoryServerUsingPUTParams struct {

	/*DirectoryServerDto
	  directoryServerDto

	*/
	DirectoryServerDto *models.DirectoryServerDto
	/*DirectoryServerID
	  directoryServerId

	*/
	DirectoryServerID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update directory server using p u t params
func (o *UpdateDirectoryServerUsingPUTParams) WithTimeout(timeout time.Duration) *UpdateDirectoryServerUsingPUTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update directory server using p u t params
func (o *UpdateDirectoryServerUsingPUTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update directory server using p u t params
func (o *UpdateDirectoryServerUsingPUTParams) WithContext(ctx context.Context) *UpdateDirectoryServerUsingPUTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update directory server using p u t params
func (o *UpdateDirectoryServerUsingPUTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update directory server using p u t params
func (o *UpdateDirectoryServerUsingPUTParams) WithHTTPClient(client *http.Client) *UpdateDirectoryServerUsingPUTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update directory server using p u t params
func (o *UpdateDirectoryServerUsingPUTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDirectoryServerDto adds the directoryServerDto to the update directory server using p u t params
func (o *UpdateDirectoryServerUsingPUTParams) WithDirectoryServerDto(directoryServerDto *models.DirectoryServerDto) *UpdateDirectoryServerUsingPUTParams {
	o.SetDirectoryServerDto(directoryServerDto)
	return o
}

// SetDirectoryServerDto adds the directoryServerDto to the update directory server using p u t params
func (o *UpdateDirectoryServerUsingPUTParams) SetDirectoryServerDto(directoryServerDto *models.DirectoryServerDto) {
	o.DirectoryServerDto = directoryServerDto
}

// WithDirectoryServerID adds the directoryServerID to the update directory server using p u t params
func (o *UpdateDirectoryServerUsingPUTParams) WithDirectoryServerID(directoryServerID int64) *UpdateDirectoryServerUsingPUTParams {
	o.SetDirectoryServerID(directoryServerID)
	return o
}

// SetDirectoryServerID adds the directoryServerId to the update directory server using p u t params
func (o *UpdateDirectoryServerUsingPUTParams) SetDirectoryServerID(directoryServerID int64) {
	o.DirectoryServerID = directoryServerID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateDirectoryServerUsingPUTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DirectoryServerDto != nil {
		if err := r.SetBodyParam(o.DirectoryServerDto); err != nil {
			return err
		}
	}

	// path param directoryServerId
	if err := r.SetPathParam("directoryServerId", swag.FormatInt64(o.DirectoryServerID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
