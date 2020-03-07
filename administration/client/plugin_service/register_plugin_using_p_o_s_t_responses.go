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

// RegisterPluginUsingPOSTReader is a Reader for the RegisterPluginUsingPOST structure.
type RegisterPluginUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RegisterPluginUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRegisterPluginUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewRegisterPluginUsingPOSTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewRegisterPluginUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRegisterPluginUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRegisterPluginUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRegisterPluginUsingPOSTOK creates a RegisterPluginUsingPOSTOK with default headers values
func NewRegisterPluginUsingPOSTOK() *RegisterPluginUsingPOSTOK {
	return &RegisterPluginUsingPOSTOK{}
}

/*RegisterPluginUsingPOSTOK handles this case with default header values.

OK
*/
type RegisterPluginUsingPOSTOK struct {
	Payload *models.PluginDto
}

func (o *RegisterPluginUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /plugins][%d] registerPluginUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *RegisterPluginUsingPOSTOK) GetPayload() *models.PluginDto {
	return o.Payload
}

func (o *RegisterPluginUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PluginDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRegisterPluginUsingPOSTCreated creates a RegisterPluginUsingPOSTCreated with default headers values
func NewRegisterPluginUsingPOSTCreated() *RegisterPluginUsingPOSTCreated {
	return &RegisterPluginUsingPOSTCreated{}
}

/*RegisterPluginUsingPOSTCreated handles this case with default header values.

Created
*/
type RegisterPluginUsingPOSTCreated struct {
}

func (o *RegisterPluginUsingPOSTCreated) Error() string {
	return fmt.Sprintf("[POST /plugins][%d] registerPluginUsingPOSTCreated ", 201)
}

func (o *RegisterPluginUsingPOSTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRegisterPluginUsingPOSTUnauthorized creates a RegisterPluginUsingPOSTUnauthorized with default headers values
func NewRegisterPluginUsingPOSTUnauthorized() *RegisterPluginUsingPOSTUnauthorized {
	return &RegisterPluginUsingPOSTUnauthorized{}
}

/*RegisterPluginUsingPOSTUnauthorized handles this case with default header values.

Unauthorized
*/
type RegisterPluginUsingPOSTUnauthorized struct {
}

func (o *RegisterPluginUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /plugins][%d] registerPluginUsingPOSTUnauthorized ", 401)
}

func (o *RegisterPluginUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRegisterPluginUsingPOSTForbidden creates a RegisterPluginUsingPOSTForbidden with default headers values
func NewRegisterPluginUsingPOSTForbidden() *RegisterPluginUsingPOSTForbidden {
	return &RegisterPluginUsingPOSTForbidden{}
}

/*RegisterPluginUsingPOSTForbidden handles this case with default header values.

Forbidden
*/
type RegisterPluginUsingPOSTForbidden struct {
}

func (o *RegisterPluginUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /plugins][%d] registerPluginUsingPOSTForbidden ", 403)
}

func (o *RegisterPluginUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRegisterPluginUsingPOSTNotFound creates a RegisterPluginUsingPOSTNotFound with default headers values
func NewRegisterPluginUsingPOSTNotFound() *RegisterPluginUsingPOSTNotFound {
	return &RegisterPluginUsingPOSTNotFound{}
}

/*RegisterPluginUsingPOSTNotFound handles this case with default header values.

Not Found
*/
type RegisterPluginUsingPOSTNotFound struct {
}

func (o *RegisterPluginUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /plugins][%d] registerPluginUsingPOSTNotFound ", 404)
}

func (o *RegisterPluginUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
