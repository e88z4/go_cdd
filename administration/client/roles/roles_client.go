// Code generated by go-swagger; DO NOT EDIT.

package roles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new roles API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for roles API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateRoleUsingPOST(params *CreateRoleUsingPOSTParams) (*CreateRoleUsingPOSTOK, *CreateRoleUsingPOSTCreated, error)

	DeleteRoleUsingDELETE(params *DeleteRoleUsingDELETEParams) (*DeleteRoleUsingDELETEOK, *DeleteRoleUsingDELETENoContent, error)

	GetRoleUsingGET(params *GetRoleUsingGETParams) (*GetRoleUsingGETOK, error)

	GetRolesUsingGET(params *GetRolesUsingGETParams) (*GetRolesUsingGETOK, error)

	UpdateRoleUsingPUT(params *UpdateRoleUsingPUTParams) (*UpdateRoleUsingPUTOK, *UpdateRoleUsingPUTCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateRoleUsingPOST creates a role
*/
func (a *Client) CreateRoleUsingPOST(params *CreateRoleUsingPOSTParams) (*CreateRoleUsingPOSTOK, *CreateRoleUsingPOSTCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateRoleUsingPOSTParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createRoleUsingPOST",
		Method:             "POST",
		PathPattern:        "/roles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateRoleUsingPOSTReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CreateRoleUsingPOSTOK:
		return value, nil, nil
	case *CreateRoleUsingPOSTCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for roles: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteRoleUsingDELETE deletes a role
*/
func (a *Client) DeleteRoleUsingDELETE(params *DeleteRoleUsingDELETEParams) (*DeleteRoleUsingDELETEOK, *DeleteRoleUsingDELETENoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteRoleUsingDELETEParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteRoleUsingDELETE",
		Method:             "DELETE",
		PathPattern:        "/roles/{roleId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteRoleUsingDELETEReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *DeleteRoleUsingDELETEOK:
		return value, nil, nil
	case *DeleteRoleUsingDELETENoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for roles: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetRoleUsingGET gets a role
*/
func (a *Client) GetRoleUsingGET(params *GetRoleUsingGETParams) (*GetRoleUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRoleUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getRoleUsingGET",
		Method:             "GET",
		PathPattern:        "/roles/{roleId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetRoleUsingGETReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetRoleUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getRoleUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetRolesUsingGET gets roles
*/
func (a *Client) GetRolesUsingGET(params *GetRolesUsingGETParams) (*GetRolesUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRolesUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getRolesUsingGET",
		Method:             "GET",
		PathPattern:        "/roles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetRolesUsingGETReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetRolesUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getRolesUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateRoleUsingPUT updates a role
*/
func (a *Client) UpdateRoleUsingPUT(params *UpdateRoleUsingPUTParams) (*UpdateRoleUsingPUTOK, *UpdateRoleUsingPUTCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateRoleUsingPUTParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateRoleUsingPUT",
		Method:             "PUT",
		PathPattern:        "/roles/{roleId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateRoleUsingPUTReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *UpdateRoleUsingPUTOK:
		return value, nil, nil
	case *UpdateRoleUsingPUTCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for roles: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
