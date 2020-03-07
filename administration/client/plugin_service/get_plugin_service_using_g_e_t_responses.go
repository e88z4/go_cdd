// Code generated by go-swagger; DO NOT EDIT.

package plugin_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/e88z4/go_cdd/administration/models"
)

// GetPluginServiceUsingGETReader is a Reader for the GetPluginServiceUsingGET structure.
type GetPluginServiceUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPluginServiceUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPluginServiceUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetPluginServiceUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetPluginServiceUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetPluginServiceUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPluginServiceUsingGETOK creates a GetPluginServiceUsingGETOK with default headers values
func NewGetPluginServiceUsingGETOK() *GetPluginServiceUsingGETOK {
	return &GetPluginServiceUsingGETOK{}
}

/*GetPluginServiceUsingGETOK handles this case with default header values.

OK
*/
type GetPluginServiceUsingGETOK struct {
	Payload *models.PluginServiceDto
}

func (o *GetPluginServiceUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /plugin-services/{pluginServiceId}][%d] getPluginServiceUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetPluginServiceUsingGETOK) GetPayload() *models.PluginServiceDto {
	return o.Payload
}

func (o *GetPluginServiceUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PluginServiceDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPluginServiceUsingGETUnauthorized creates a GetPluginServiceUsingGETUnauthorized with default headers values
func NewGetPluginServiceUsingGETUnauthorized() *GetPluginServiceUsingGETUnauthorized {
	return &GetPluginServiceUsingGETUnauthorized{}
}

/*GetPluginServiceUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type GetPluginServiceUsingGETUnauthorized struct {
}

func (o *GetPluginServiceUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /plugin-services/{pluginServiceId}][%d] getPluginServiceUsingGETUnauthorized ", 401)
}

func (o *GetPluginServiceUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPluginServiceUsingGETForbidden creates a GetPluginServiceUsingGETForbidden with default headers values
func NewGetPluginServiceUsingGETForbidden() *GetPluginServiceUsingGETForbidden {
	return &GetPluginServiceUsingGETForbidden{}
}

/*GetPluginServiceUsingGETForbidden handles this case with default header values.

Forbidden
*/
type GetPluginServiceUsingGETForbidden struct {
}

func (o *GetPluginServiceUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /plugin-services/{pluginServiceId}][%d] getPluginServiceUsingGETForbidden ", 403)
}

func (o *GetPluginServiceUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPluginServiceUsingGETNotFound creates a GetPluginServiceUsingGETNotFound with default headers values
func NewGetPluginServiceUsingGETNotFound() *GetPluginServiceUsingGETNotFound {
	return &GetPluginServiceUsingGETNotFound{}
}

/*GetPluginServiceUsingGETNotFound handles this case with default header values.

Not Found
*/
type GetPluginServiceUsingGETNotFound struct {
}

func (o *GetPluginServiceUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /plugin-services/{pluginServiceId}][%d] getPluginServiceUsingGETNotFound ", 404)
}

func (o *GetPluginServiceUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}