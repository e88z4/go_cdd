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

	strfmt "github.com/go-openapi/strfmt"

	"github.com/e88z4/go_cdd/administration/models"
)

// NewTestDirectoryServerConnectivityUsingPOSTParams creates a new TestDirectoryServerConnectivityUsingPOSTParams object
// with the default values initialized.
func NewTestDirectoryServerConnectivityUsingPOSTParams() *TestDirectoryServerConnectivityUsingPOSTParams {
	var ()
	return &TestDirectoryServerConnectivityUsingPOSTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTestDirectoryServerConnectivityUsingPOSTParamsWithTimeout creates a new TestDirectoryServerConnectivityUsingPOSTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTestDirectoryServerConnectivityUsingPOSTParamsWithTimeout(timeout time.Duration) *TestDirectoryServerConnectivityUsingPOSTParams {
	var ()
	return &TestDirectoryServerConnectivityUsingPOSTParams{

		timeout: timeout,
	}
}

// NewTestDirectoryServerConnectivityUsingPOSTParamsWithContext creates a new TestDirectoryServerConnectivityUsingPOSTParams object
// with the default values initialized, and the ability to set a context for a request
func NewTestDirectoryServerConnectivityUsingPOSTParamsWithContext(ctx context.Context) *TestDirectoryServerConnectivityUsingPOSTParams {
	var ()
	return &TestDirectoryServerConnectivityUsingPOSTParams{

		Context: ctx,
	}
}

// NewTestDirectoryServerConnectivityUsingPOSTParamsWithHTTPClient creates a new TestDirectoryServerConnectivityUsingPOSTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTestDirectoryServerConnectivityUsingPOSTParamsWithHTTPClient(client *http.Client) *TestDirectoryServerConnectivityUsingPOSTParams {
	var ()
	return &TestDirectoryServerConnectivityUsingPOSTParams{
		HTTPClient: client,
	}
}

/*TestDirectoryServerConnectivityUsingPOSTParams contains all the parameters to send to the API endpoint
for the test directory server connectivity using p o s t operation typically these are written to a http.Request
*/
type TestDirectoryServerConnectivityUsingPOSTParams struct {

	/*DirectoryServerDto
	  directoryServerDto

	*/
	DirectoryServerDto *models.DirectoryServerDto

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the test directory server connectivity using p o s t params
func (o *TestDirectoryServerConnectivityUsingPOSTParams) WithTimeout(timeout time.Duration) *TestDirectoryServerConnectivityUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the test directory server connectivity using p o s t params
func (o *TestDirectoryServerConnectivityUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the test directory server connectivity using p o s t params
func (o *TestDirectoryServerConnectivityUsingPOSTParams) WithContext(ctx context.Context) *TestDirectoryServerConnectivityUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the test directory server connectivity using p o s t params
func (o *TestDirectoryServerConnectivityUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the test directory server connectivity using p o s t params
func (o *TestDirectoryServerConnectivityUsingPOSTParams) WithHTTPClient(client *http.Client) *TestDirectoryServerConnectivityUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the test directory server connectivity using p o s t params
func (o *TestDirectoryServerConnectivityUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDirectoryServerDto adds the directoryServerDto to the test directory server connectivity using p o s t params
func (o *TestDirectoryServerConnectivityUsingPOSTParams) WithDirectoryServerDto(directoryServerDto *models.DirectoryServerDto) *TestDirectoryServerConnectivityUsingPOSTParams {
	o.SetDirectoryServerDto(directoryServerDto)
	return o
}

// SetDirectoryServerDto adds the directoryServerDto to the test directory server connectivity using p o s t params
func (o *TestDirectoryServerConnectivityUsingPOSTParams) SetDirectoryServerDto(directoryServerDto *models.DirectoryServerDto) {
	o.DirectoryServerDto = directoryServerDto
}

// WriteToRequest writes these params to a swagger request
func (o *TestDirectoryServerConnectivityUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DirectoryServerDto != nil {
		if err := r.SetBodyParam(o.DirectoryServerDto); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}