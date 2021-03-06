// Code generated by go-swagger; DO NOT EDIT.

package environment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new environment API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for environment API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateEnvironmentAtApplicationUsingPOST(params *CreateEnvironmentAtApplicationUsingPOSTParams) (*CreateEnvironmentAtApplicationUsingPOSTOK, *CreateEnvironmentAtApplicationUsingPOSTCreated, error)

	DeleteEnvironmentUsingDELETE(params *DeleteEnvironmentUsingDELETEParams) (*DeleteEnvironmentUsingDELETEOK, *DeleteEnvironmentUsingDELETENoContent, error)

	GetAllEnvironmentsUsingGET(params *GetAllEnvironmentsUsingGETParams) (*GetAllEnvironmentsUsingGETOK, error)

	GetCalendarEnvironmentEntitiesUsingGET(params *GetCalendarEnvironmentEntitiesUsingGETParams) (*GetCalendarEnvironmentEntitiesUsingGETOK, error)

	GetEnvironmentAtApplicationUsingGET(params *GetEnvironmentAtApplicationUsingGETParams) (*GetEnvironmentAtApplicationUsingGETOK, error)

	GetEnvironmentUsingGET(params *GetEnvironmentUsingGETParams) (*GetEnvironmentUsingGETOK, error)

	GetEnvironmentsAtApplicationUsingGET(params *GetEnvironmentsAtApplicationUsingGETParams) (*GetEnvironmentsAtApplicationUsingGETOK, error)

	PatchEnvironmentAtApplicationUsingPATCH(params *PatchEnvironmentAtApplicationUsingPATCHParams) (*PatchEnvironmentAtApplicationUsingPATCHOK, *PatchEnvironmentAtApplicationUsingPATCHNoContent, error)

	PatchEnvironmentUsingPATCH(params *PatchEnvironmentUsingPATCHParams) (*PatchEnvironmentUsingPATCHOK, *PatchEnvironmentUsingPATCHNoContent, error)

	UpdateEnvironmentAtApplicationUsingPUT(params *UpdateEnvironmentAtApplicationUsingPUTParams) (*UpdateEnvironmentAtApplicationUsingPUTOK, *UpdateEnvironmentAtApplicationUsingPUTCreated, error)

	UpdateEnvironmentUsingPUT(params *UpdateEnvironmentUsingPUTParams) (*UpdateEnvironmentUsingPUTOK, *UpdateEnvironmentUsingPUTCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateEnvironmentAtApplicationUsingPOST creates an environment
*/
func (a *Client) CreateEnvironmentAtApplicationUsingPOST(params *CreateEnvironmentAtApplicationUsingPOSTParams) (*CreateEnvironmentAtApplicationUsingPOSTOK, *CreateEnvironmentAtApplicationUsingPOSTCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateEnvironmentAtApplicationUsingPOSTParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createEnvironmentAtApplicationUsingPOST",
		Method:             "POST",
		PathPattern:        "/applications/{applicationId}/environments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateEnvironmentAtApplicationUsingPOSTReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CreateEnvironmentAtApplicationUsingPOSTOK:
		return value, nil, nil
	case *CreateEnvironmentAtApplicationUsingPOSTCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for environment: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteEnvironmentUsingDELETE deletes an environment
*/
func (a *Client) DeleteEnvironmentUsingDELETE(params *DeleteEnvironmentUsingDELETEParams) (*DeleteEnvironmentUsingDELETEOK, *DeleteEnvironmentUsingDELETENoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteEnvironmentUsingDELETEParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteEnvironmentUsingDELETE",
		Method:             "DELETE",
		PathPattern:        "/applications/{applicationId}/environments/{environmentId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteEnvironmentUsingDELETEReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *DeleteEnvironmentUsingDELETEOK:
		return value, nil, nil
	case *DeleteEnvironmentUsingDELETENoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for environment: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAllEnvironmentsUsingGET gets all environments
*/
func (a *Client) GetAllEnvironmentsUsingGET(params *GetAllEnvironmentsUsingGETParams) (*GetAllEnvironmentsUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllEnvironmentsUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAllEnvironmentsUsingGET",
		Method:             "GET",
		PathPattern:        "/environments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAllEnvironmentsUsingGETReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAllEnvironmentsUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAllEnvironmentsUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetCalendarEnvironmentEntitiesUsingGET gets calendar environment entities phases maintenance windows and freeze periods
*/
func (a *Client) GetCalendarEnvironmentEntitiesUsingGET(params *GetCalendarEnvironmentEntitiesUsingGETParams) (*GetCalendarEnvironmentEntitiesUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCalendarEnvironmentEntitiesUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCalendarEnvironmentEntitiesUsingGET",
		Method:             "GET",
		PathPattern:        "/environments/planner",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetCalendarEnvironmentEntitiesUsingGETReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCalendarEnvironmentEntitiesUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCalendarEnvironmentEntitiesUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetEnvironmentAtApplicationUsingGET gets an environment at application
*/
func (a *Client) GetEnvironmentAtApplicationUsingGET(params *GetEnvironmentAtApplicationUsingGETParams) (*GetEnvironmentAtApplicationUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEnvironmentAtApplicationUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getEnvironmentAtApplicationUsingGET",
		Method:             "GET",
		PathPattern:        "/applications/{applicationId}/environments/{environmentId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetEnvironmentAtApplicationUsingGETReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetEnvironmentAtApplicationUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getEnvironmentAtApplicationUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetEnvironmentUsingGET gets an environment
*/
func (a *Client) GetEnvironmentUsingGET(params *GetEnvironmentUsingGETParams) (*GetEnvironmentUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEnvironmentUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getEnvironmentUsingGET",
		Method:             "GET",
		PathPattern:        "/environments/{environmentId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetEnvironmentUsingGETReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetEnvironmentUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getEnvironmentUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetEnvironmentsAtApplicationUsingGET gets application environments
*/
func (a *Client) GetEnvironmentsAtApplicationUsingGET(params *GetEnvironmentsAtApplicationUsingGETParams) (*GetEnvironmentsAtApplicationUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEnvironmentsAtApplicationUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getEnvironmentsAtApplicationUsingGET",
		Method:             "GET",
		PathPattern:        "/applications/{applicationId}/environments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetEnvironmentsAtApplicationUsingGETReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetEnvironmentsAtApplicationUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getEnvironmentsAtApplicationUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PatchEnvironmentAtApplicationUsingPATCH patches an environment at application
*/
func (a *Client) PatchEnvironmentAtApplicationUsingPATCH(params *PatchEnvironmentAtApplicationUsingPATCHParams) (*PatchEnvironmentAtApplicationUsingPATCHOK, *PatchEnvironmentAtApplicationUsingPATCHNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchEnvironmentAtApplicationUsingPATCHParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "patchEnvironmentAtApplicationUsingPATCH",
		Method:             "PATCH",
		PathPattern:        "/applications/{applicationId}/environments/{environmentId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PatchEnvironmentAtApplicationUsingPATCHReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *PatchEnvironmentAtApplicationUsingPATCHOK:
		return value, nil, nil
	case *PatchEnvironmentAtApplicationUsingPATCHNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for environment: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PatchEnvironmentUsingPATCH patches an environment
*/
func (a *Client) PatchEnvironmentUsingPATCH(params *PatchEnvironmentUsingPATCHParams) (*PatchEnvironmentUsingPATCHOK, *PatchEnvironmentUsingPATCHNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchEnvironmentUsingPATCHParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "patchEnvironmentUsingPATCH",
		Method:             "PATCH",
		PathPattern:        "/environments/{environmentId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PatchEnvironmentUsingPATCHReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *PatchEnvironmentUsingPATCHOK:
		return value, nil, nil
	case *PatchEnvironmentUsingPATCHNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for environment: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateEnvironmentAtApplicationUsingPUT updates an environment at application
*/
func (a *Client) UpdateEnvironmentAtApplicationUsingPUT(params *UpdateEnvironmentAtApplicationUsingPUTParams) (*UpdateEnvironmentAtApplicationUsingPUTOK, *UpdateEnvironmentAtApplicationUsingPUTCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateEnvironmentAtApplicationUsingPUTParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateEnvironmentAtApplicationUsingPUT",
		Method:             "PUT",
		PathPattern:        "/applications/{applicationId}/environments/{environmentId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateEnvironmentAtApplicationUsingPUTReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *UpdateEnvironmentAtApplicationUsingPUTOK:
		return value, nil, nil
	case *UpdateEnvironmentAtApplicationUsingPUTCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for environment: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateEnvironmentUsingPUT updates an environment
*/
func (a *Client) UpdateEnvironmentUsingPUT(params *UpdateEnvironmentUsingPUTParams) (*UpdateEnvironmentUsingPUTOK, *UpdateEnvironmentUsingPUTCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateEnvironmentUsingPUTParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateEnvironmentUsingPUT",
		Method:             "PUT",
		PathPattern:        "/environments/{environmentId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateEnvironmentUsingPUTReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *UpdateEnvironmentUsingPUTOK:
		return value, nil, nil
	case *UpdateEnvironmentUsingPUTCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for environment: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
