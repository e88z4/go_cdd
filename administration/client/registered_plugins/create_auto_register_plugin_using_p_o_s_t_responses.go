// Code generated by go-swagger; DO NOT EDIT.

package registered_plugins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/e88z4/go_cdd/administration/models"
)

// CreateAutoRegisterPluginUsingPOSTReader is a Reader for the CreateAutoRegisterPluginUsingPOST structure.
type CreateAutoRegisterPluginUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAutoRegisterPluginUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateAutoRegisterPluginUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewCreateAutoRegisterPluginUsingPOSTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateAutoRegisterPluginUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateAutoRegisterPluginUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateAutoRegisterPluginUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateAutoRegisterPluginUsingPOSTOK creates a CreateAutoRegisterPluginUsingPOSTOK with default headers values
func NewCreateAutoRegisterPluginUsingPOSTOK() *CreateAutoRegisterPluginUsingPOSTOK {
	return &CreateAutoRegisterPluginUsingPOSTOK{}
}

/*CreateAutoRegisterPluginUsingPOSTOK handles this case with default header values.

OK
*/
type CreateAutoRegisterPluginUsingPOSTOK struct {
	Payload *models.RegisteredPluginDto
}

func (o *CreateAutoRegisterPluginUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /registered-plugins][%d] createAutoRegisterPluginUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *CreateAutoRegisterPluginUsingPOSTOK) GetPayload() *models.RegisteredPluginDto {
	return o.Payload
}

func (o *CreateAutoRegisterPluginUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RegisteredPluginDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAutoRegisterPluginUsingPOSTCreated creates a CreateAutoRegisterPluginUsingPOSTCreated with default headers values
func NewCreateAutoRegisterPluginUsingPOSTCreated() *CreateAutoRegisterPluginUsingPOSTCreated {
	return &CreateAutoRegisterPluginUsingPOSTCreated{}
}

/*CreateAutoRegisterPluginUsingPOSTCreated handles this case with default header values.

Created
*/
type CreateAutoRegisterPluginUsingPOSTCreated struct {
}

func (o *CreateAutoRegisterPluginUsingPOSTCreated) Error() string {
	return fmt.Sprintf("[POST /registered-plugins][%d] createAutoRegisterPluginUsingPOSTCreated ", 201)
}

func (o *CreateAutoRegisterPluginUsingPOSTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateAutoRegisterPluginUsingPOSTUnauthorized creates a CreateAutoRegisterPluginUsingPOSTUnauthorized with default headers values
func NewCreateAutoRegisterPluginUsingPOSTUnauthorized() *CreateAutoRegisterPluginUsingPOSTUnauthorized {
	return &CreateAutoRegisterPluginUsingPOSTUnauthorized{}
}

/*CreateAutoRegisterPluginUsingPOSTUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateAutoRegisterPluginUsingPOSTUnauthorized struct {
}

func (o *CreateAutoRegisterPluginUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /registered-plugins][%d] createAutoRegisterPluginUsingPOSTUnauthorized ", 401)
}

func (o *CreateAutoRegisterPluginUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateAutoRegisterPluginUsingPOSTForbidden creates a CreateAutoRegisterPluginUsingPOSTForbidden with default headers values
func NewCreateAutoRegisterPluginUsingPOSTForbidden() *CreateAutoRegisterPluginUsingPOSTForbidden {
	return &CreateAutoRegisterPluginUsingPOSTForbidden{}
}

/*CreateAutoRegisterPluginUsingPOSTForbidden handles this case with default header values.

Forbidden
*/
type CreateAutoRegisterPluginUsingPOSTForbidden struct {
}

func (o *CreateAutoRegisterPluginUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /registered-plugins][%d] createAutoRegisterPluginUsingPOSTForbidden ", 403)
}

func (o *CreateAutoRegisterPluginUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateAutoRegisterPluginUsingPOSTNotFound creates a CreateAutoRegisterPluginUsingPOSTNotFound with default headers values
func NewCreateAutoRegisterPluginUsingPOSTNotFound() *CreateAutoRegisterPluginUsingPOSTNotFound {
	return &CreateAutoRegisterPluginUsingPOSTNotFound{}
}

/*CreateAutoRegisterPluginUsingPOSTNotFound handles this case with default header values.

Not Found
*/
type CreateAutoRegisterPluginUsingPOSTNotFound struct {
}

func (o *CreateAutoRegisterPluginUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /registered-plugins][%d] createAutoRegisterPluginUsingPOSTNotFound ", 404)
}

func (o *CreateAutoRegisterPluginUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
