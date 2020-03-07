// Code generated by go-swagger; DO NOT EDIT.

package directory_server_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new directory server controller API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for directory server controller API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateDirectoryServerUsingPOST(params *CreateDirectoryServerUsingPOSTParams) (*CreateDirectoryServerUsingPOSTOK, *CreateDirectoryServerUsingPOSTCreated, error)

	DeleteDirectoryServerUsingDELETE(params *DeleteDirectoryServerUsingDELETEParams) (*DeleteDirectoryServerUsingDELETEOK, *DeleteDirectoryServerUsingDELETENoContent, error)

	GetDirectoryServerUsingGET(params *GetDirectoryServerUsingGETParams) (*GetDirectoryServerUsingGETOK, error)

	GetDirectoryServersUsingGET(params *GetDirectoryServersUsingGETParams) (*GetDirectoryServersUsingGETOK, error)

	GetPotentialGroupsUsingGET(params *GetPotentialGroupsUsingGETParams) (*GetPotentialGroupsUsingGETOK, error)

	GetPotentialUsersUsingGET(params *GetPotentialUsersUsingGETParams) (*GetPotentialUsersUsingGETOK, error)

	PatchDirectoryServerUsingPATCH(params *PatchDirectoryServerUsingPATCHParams) (*PatchDirectoryServerUsingPATCHOK, *PatchDirectoryServerUsingPATCHNoContent, error)

	TestDirectoryServerConnectivityUsingPOST(params *TestDirectoryServerConnectivityUsingPOSTParams) (*TestDirectoryServerConnectivityUsingPOSTOK, *TestDirectoryServerConnectivityUsingPOSTCreated, error)

	UpdateDirectoryServerUsingPUT(params *UpdateDirectoryServerUsingPUTParams) (*UpdateDirectoryServerUsingPUTOK, *UpdateDirectoryServerUsingPUTCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateDirectoryServerUsingPOST creates a new directory server definition
*/
func (a *Client) CreateDirectoryServerUsingPOST(params *CreateDirectoryServerUsingPOSTParams) (*CreateDirectoryServerUsingPOSTOK, *CreateDirectoryServerUsingPOSTCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateDirectoryServerUsingPOSTParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createDirectoryServerUsingPOST",
		Method:             "POST",
		PathPattern:        "/directory-servers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateDirectoryServerUsingPOSTReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CreateDirectoryServerUsingPOSTOK:
		return value, nil, nil
	case *CreateDirectoryServerUsingPOSTCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for directory_server_controller: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteDirectoryServerUsingDELETE deletes a directory server definition
*/
func (a *Client) DeleteDirectoryServerUsingDELETE(params *DeleteDirectoryServerUsingDELETEParams) (*DeleteDirectoryServerUsingDELETEOK, *DeleteDirectoryServerUsingDELETENoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteDirectoryServerUsingDELETEParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteDirectoryServerUsingDELETE",
		Method:             "DELETE",
		PathPattern:        "/directory-servers/{directoryServerId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteDirectoryServerUsingDELETEReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *DeleteDirectoryServerUsingDELETEOK:
		return value, nil, nil
	case *DeleteDirectoryServerUsingDELETENoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for directory_server_controller: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetDirectoryServerUsingGET retrieves a directory server definition
*/
func (a *Client) GetDirectoryServerUsingGET(params *GetDirectoryServerUsingGETParams) (*GetDirectoryServerUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDirectoryServerUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDirectoryServerUsingGET",
		Method:             "GET",
		PathPattern:        "/directory-servers/{directoryServerId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetDirectoryServerUsingGETReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDirectoryServerUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDirectoryServerUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetDirectoryServersUsingGET retrieves all directory server definitions
*/
func (a *Client) GetDirectoryServersUsingGET(params *GetDirectoryServersUsingGETParams) (*GetDirectoryServersUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDirectoryServersUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDirectoryServersUsingGET",
		Method:             "GET",
		PathPattern:        "/directory-servers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetDirectoryServersUsingGETReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDirectoryServersUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDirectoryServersUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetPotentialGroupsUsingGET searches for groups to be imported inside a given directory server
*/
func (a *Client) GetPotentialGroupsUsingGET(params *GetPotentialGroupsUsingGETParams) (*GetPotentialGroupsUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPotentialGroupsUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPotentialGroupsUsingGET",
		Method:             "GET",
		PathPattern:        "/directory-user-groups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPotentialGroupsUsingGETReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPotentialGroupsUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getPotentialGroupsUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetPotentialUsersUsingGET searches for users to be imported inside a given directory server
*/
func (a *Client) GetPotentialUsersUsingGET(params *GetPotentialUsersUsingGETParams) (*GetPotentialUsersUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPotentialUsersUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPotentialUsersUsingGET",
		Method:             "GET",
		PathPattern:        "/directory-users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPotentialUsersUsingGETReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPotentialUsersUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getPotentialUsersUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PatchDirectoryServerUsingPATCH patches a directory server definition
*/
func (a *Client) PatchDirectoryServerUsingPATCH(params *PatchDirectoryServerUsingPATCHParams) (*PatchDirectoryServerUsingPATCHOK, *PatchDirectoryServerUsingPATCHNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchDirectoryServerUsingPATCHParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "patchDirectoryServerUsingPATCH",
		Method:             "PATCH",
		PathPattern:        "/directory-servers/{directoryServerId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PatchDirectoryServerUsingPATCHReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *PatchDirectoryServerUsingPATCHOK:
		return value, nil, nil
	case *PatchDirectoryServerUsingPATCHNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for directory_server_controller: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TestDirectoryServerConnectivityUsingPOST connectivities test for directory server
*/
func (a *Client) TestDirectoryServerConnectivityUsingPOST(params *TestDirectoryServerConnectivityUsingPOSTParams) (*TestDirectoryServerConnectivityUsingPOSTOK, *TestDirectoryServerConnectivityUsingPOSTCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTestDirectoryServerConnectivityUsingPOSTParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "testDirectoryServerConnectivityUsingPOST",
		Method:             "POST",
		PathPattern:        "/directory-servers/connectivity-tests",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &TestDirectoryServerConnectivityUsingPOSTReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *TestDirectoryServerConnectivityUsingPOSTOK:
		return value, nil, nil
	case *TestDirectoryServerConnectivityUsingPOSTCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for directory_server_controller: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateDirectoryServerUsingPUT updates a directory server definition
*/
func (a *Client) UpdateDirectoryServerUsingPUT(params *UpdateDirectoryServerUsingPUTParams) (*UpdateDirectoryServerUsingPUTOK, *UpdateDirectoryServerUsingPUTCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateDirectoryServerUsingPUTParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateDirectoryServerUsingPUT",
		Method:             "PUT",
		PathPattern:        "/directory-servers/{directoryServerId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateDirectoryServerUsingPUTReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *UpdateDirectoryServerUsingPUTOK:
		return value, nil, nil
	case *UpdateDirectoryServerUsingPUTCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for directory_server_controller: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
