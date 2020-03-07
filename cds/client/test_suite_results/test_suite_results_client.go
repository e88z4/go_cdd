// Code generated by go-swagger; DO NOT EDIT.

package test_suite_results

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new test suite results API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for test suite results API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetTestSuiteResultUsingGET(params *GetTestSuiteResultUsingGETParams) (*GetTestSuiteResultUsingGETOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetTestSuiteResultUsingGET retrieves test suite result
*/
func (a *Client) GetTestSuiteResultUsingGET(params *GetTestSuiteResultUsingGETParams) (*GetTestSuiteResultUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTestSuiteResultUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getTestSuiteResultUsingGET",
		Method:             "GET",
		PathPattern:        "/applications/{applicationId}/application-versions/{applicationVersionId}/test-suite-results",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetTestSuiteResultUsingGETReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetTestSuiteResultUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getTestSuiteResultUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}