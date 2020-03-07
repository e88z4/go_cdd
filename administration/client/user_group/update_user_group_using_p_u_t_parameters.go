// Code generated by go-swagger; DO NOT EDIT.

package user_group

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

// NewUpdateUserGroupUsingPUTParams creates a new UpdateUserGroupUsingPUTParams object
// with the default values initialized.
func NewUpdateUserGroupUsingPUTParams() *UpdateUserGroupUsingPUTParams {
	var ()
	return &UpdateUserGroupUsingPUTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateUserGroupUsingPUTParamsWithTimeout creates a new UpdateUserGroupUsingPUTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateUserGroupUsingPUTParamsWithTimeout(timeout time.Duration) *UpdateUserGroupUsingPUTParams {
	var ()
	return &UpdateUserGroupUsingPUTParams{

		timeout: timeout,
	}
}

// NewUpdateUserGroupUsingPUTParamsWithContext creates a new UpdateUserGroupUsingPUTParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateUserGroupUsingPUTParamsWithContext(ctx context.Context) *UpdateUserGroupUsingPUTParams {
	var ()
	return &UpdateUserGroupUsingPUTParams{

		Context: ctx,
	}
}

// NewUpdateUserGroupUsingPUTParamsWithHTTPClient creates a new UpdateUserGroupUsingPUTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateUserGroupUsingPUTParamsWithHTTPClient(client *http.Client) *UpdateUserGroupUsingPUTParams {
	var ()
	return &UpdateUserGroupUsingPUTParams{
		HTTPClient: client,
	}
}

/*UpdateUserGroupUsingPUTParams contains all the parameters to send to the API endpoint
for the update user group using p u t operation typically these are written to a http.Request
*/
type UpdateUserGroupUsingPUTParams struct {

	/*UserGroupDto
	  userGroupDto

	*/
	UserGroupDto *models.UserGroupDto
	/*UserGroupID
	  userGroupId

	*/
	UserGroupID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update user group using p u t params
func (o *UpdateUserGroupUsingPUTParams) WithTimeout(timeout time.Duration) *UpdateUserGroupUsingPUTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update user group using p u t params
func (o *UpdateUserGroupUsingPUTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update user group using p u t params
func (o *UpdateUserGroupUsingPUTParams) WithContext(ctx context.Context) *UpdateUserGroupUsingPUTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update user group using p u t params
func (o *UpdateUserGroupUsingPUTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update user group using p u t params
func (o *UpdateUserGroupUsingPUTParams) WithHTTPClient(client *http.Client) *UpdateUserGroupUsingPUTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update user group using p u t params
func (o *UpdateUserGroupUsingPUTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserGroupDto adds the userGroupDto to the update user group using p u t params
func (o *UpdateUserGroupUsingPUTParams) WithUserGroupDto(userGroupDto *models.UserGroupDto) *UpdateUserGroupUsingPUTParams {
	o.SetUserGroupDto(userGroupDto)
	return o
}

// SetUserGroupDto adds the userGroupDto to the update user group using p u t params
func (o *UpdateUserGroupUsingPUTParams) SetUserGroupDto(userGroupDto *models.UserGroupDto) {
	o.UserGroupDto = userGroupDto
}

// WithUserGroupID adds the userGroupID to the update user group using p u t params
func (o *UpdateUserGroupUsingPUTParams) WithUserGroupID(userGroupID int64) *UpdateUserGroupUsingPUTParams {
	o.SetUserGroupID(userGroupID)
	return o
}

// SetUserGroupID adds the userGroupId to the update user group using p u t params
func (o *UpdateUserGroupUsingPUTParams) SetUserGroupID(userGroupID int64) {
	o.UserGroupID = userGroupID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateUserGroupUsingPUTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.UserGroupDto != nil {
		if err := r.SetBodyParam(o.UserGroupDto); err != nil {
			return err
		}
	}

	// path param userGroupId
	if err := r.SetPathParam("userGroupId", swag.FormatInt64(o.UserGroupID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
