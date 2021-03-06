// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/e88z4/go_cdd/administration/models"
)

// CreateUsersUsingPOSTReader is a Reader for the CreateUsersUsingPOST structure.
type CreateUsersUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateUsersUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateUsersUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewCreateUsersUsingPOSTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateUsersUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateUsersUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateUsersUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateUsersUsingPOSTOK creates a CreateUsersUsingPOSTOK with default headers values
func NewCreateUsersUsingPOSTOK() *CreateUsersUsingPOSTOK {
	return &CreateUsersUsingPOSTOK{}
}

/*CreateUsersUsingPOSTOK handles this case with default header values.

OK
*/
type CreateUsersUsingPOSTOK struct {
	Payload *models.ListHolderDtoUserDto
}

func (o *CreateUsersUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /users][%d] createUsersUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *CreateUsersUsingPOSTOK) GetPayload() *models.ListHolderDtoUserDto {
	return o.Payload
}

func (o *CreateUsersUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListHolderDtoUserDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUsersUsingPOSTCreated creates a CreateUsersUsingPOSTCreated with default headers values
func NewCreateUsersUsingPOSTCreated() *CreateUsersUsingPOSTCreated {
	return &CreateUsersUsingPOSTCreated{}
}

/*CreateUsersUsingPOSTCreated handles this case with default header values.

Created
*/
type CreateUsersUsingPOSTCreated struct {
}

func (o *CreateUsersUsingPOSTCreated) Error() string {
	return fmt.Sprintf("[POST /users][%d] createUsersUsingPOSTCreated ", 201)
}

func (o *CreateUsersUsingPOSTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateUsersUsingPOSTUnauthorized creates a CreateUsersUsingPOSTUnauthorized with default headers values
func NewCreateUsersUsingPOSTUnauthorized() *CreateUsersUsingPOSTUnauthorized {
	return &CreateUsersUsingPOSTUnauthorized{}
}

/*CreateUsersUsingPOSTUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateUsersUsingPOSTUnauthorized struct {
}

func (o *CreateUsersUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /users][%d] createUsersUsingPOSTUnauthorized ", 401)
}

func (o *CreateUsersUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateUsersUsingPOSTForbidden creates a CreateUsersUsingPOSTForbidden with default headers values
func NewCreateUsersUsingPOSTForbidden() *CreateUsersUsingPOSTForbidden {
	return &CreateUsersUsingPOSTForbidden{}
}

/*CreateUsersUsingPOSTForbidden handles this case with default header values.

Forbidden
*/
type CreateUsersUsingPOSTForbidden struct {
}

func (o *CreateUsersUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /users][%d] createUsersUsingPOSTForbidden ", 403)
}

func (o *CreateUsersUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateUsersUsingPOSTNotFound creates a CreateUsersUsingPOSTNotFound with default headers values
func NewCreateUsersUsingPOSTNotFound() *CreateUsersUsingPOSTNotFound {
	return &CreateUsersUsingPOSTNotFound{}
}

/*CreateUsersUsingPOSTNotFound handles this case with default header values.

Not Found
*/
type CreateUsersUsingPOSTNotFound struct {
}

func (o *CreateUsersUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /users][%d] createUsersUsingPOSTNotFound ", 404)
}

func (o *CreateUsersUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
