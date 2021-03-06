// Code generated by go-swagger; DO NOT EDIT.

package roles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/e88z4/go_cdd/administration/models"
)

// CreateRoleUsingPOSTReader is a Reader for the CreateRoleUsingPOST structure.
type CreateRoleUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRoleUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateRoleUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewCreateRoleUsingPOSTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateRoleUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateRoleUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateRoleUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateRoleUsingPOSTOK creates a CreateRoleUsingPOSTOK with default headers values
func NewCreateRoleUsingPOSTOK() *CreateRoleUsingPOSTOK {
	return &CreateRoleUsingPOSTOK{}
}

/*CreateRoleUsingPOSTOK handles this case with default header values.

OK
*/
type CreateRoleUsingPOSTOK struct {
	Payload *models.RoleDto
}

func (o *CreateRoleUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /roles][%d] createRoleUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *CreateRoleUsingPOSTOK) GetPayload() *models.RoleDto {
	return o.Payload
}

func (o *CreateRoleUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RoleDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRoleUsingPOSTCreated creates a CreateRoleUsingPOSTCreated with default headers values
func NewCreateRoleUsingPOSTCreated() *CreateRoleUsingPOSTCreated {
	return &CreateRoleUsingPOSTCreated{}
}

/*CreateRoleUsingPOSTCreated handles this case with default header values.

Created
*/
type CreateRoleUsingPOSTCreated struct {
}

func (o *CreateRoleUsingPOSTCreated) Error() string {
	return fmt.Sprintf("[POST /roles][%d] createRoleUsingPOSTCreated ", 201)
}

func (o *CreateRoleUsingPOSTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRoleUsingPOSTUnauthorized creates a CreateRoleUsingPOSTUnauthorized with default headers values
func NewCreateRoleUsingPOSTUnauthorized() *CreateRoleUsingPOSTUnauthorized {
	return &CreateRoleUsingPOSTUnauthorized{}
}

/*CreateRoleUsingPOSTUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateRoleUsingPOSTUnauthorized struct {
}

func (o *CreateRoleUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /roles][%d] createRoleUsingPOSTUnauthorized ", 401)
}

func (o *CreateRoleUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRoleUsingPOSTForbidden creates a CreateRoleUsingPOSTForbidden with default headers values
func NewCreateRoleUsingPOSTForbidden() *CreateRoleUsingPOSTForbidden {
	return &CreateRoleUsingPOSTForbidden{}
}

/*CreateRoleUsingPOSTForbidden handles this case with default header values.

Forbidden
*/
type CreateRoleUsingPOSTForbidden struct {
}

func (o *CreateRoleUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /roles][%d] createRoleUsingPOSTForbidden ", 403)
}

func (o *CreateRoleUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRoleUsingPOSTNotFound creates a CreateRoleUsingPOSTNotFound with default headers values
func NewCreateRoleUsingPOSTNotFound() *CreateRoleUsingPOSTNotFound {
	return &CreateRoleUsingPOSTNotFound{}
}

/*CreateRoleUsingPOSTNotFound handles this case with default header values.

Not Found
*/
type CreateRoleUsingPOSTNotFound struct {
}

func (o *CreateRoleUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /roles][%d] createRoleUsingPOSTNotFound ", 404)
}

func (o *CreateRoleUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
