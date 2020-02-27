// Code generated by go-swagger; DO NOT EDIT.

package user_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/e88z4/go_cdd/administration/models"
)

// CreateUserGroupsUsingPOSTReader is a Reader for the CreateUserGroupsUsingPOST structure.
type CreateUserGroupsUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateUserGroupsUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateUserGroupsUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewCreateUserGroupsUsingPOSTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateUserGroupsUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateUserGroupsUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateUserGroupsUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateUserGroupsUsingPOSTOK creates a CreateUserGroupsUsingPOSTOK with default headers values
func NewCreateUserGroupsUsingPOSTOK() *CreateUserGroupsUsingPOSTOK {
	return &CreateUserGroupsUsingPOSTOK{}
}

/*CreateUserGroupsUsingPOSTOK handles this case with default header values.

OK
*/
type CreateUserGroupsUsingPOSTOK struct {
	Payload *models.ListHolderDtoUserGroupDto
}

func (o *CreateUserGroupsUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /user-groups][%d] createUserGroupsUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *CreateUserGroupsUsingPOSTOK) GetPayload() *models.ListHolderDtoUserGroupDto {
	return o.Payload
}

func (o *CreateUserGroupsUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListHolderDtoUserGroupDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserGroupsUsingPOSTCreated creates a CreateUserGroupsUsingPOSTCreated with default headers values
func NewCreateUserGroupsUsingPOSTCreated() *CreateUserGroupsUsingPOSTCreated {
	return &CreateUserGroupsUsingPOSTCreated{}
}

/*CreateUserGroupsUsingPOSTCreated handles this case with default header values.

Created
*/
type CreateUserGroupsUsingPOSTCreated struct {
}

func (o *CreateUserGroupsUsingPOSTCreated) Error() string {
	return fmt.Sprintf("[POST /user-groups][%d] createUserGroupsUsingPOSTCreated ", 201)
}

func (o *CreateUserGroupsUsingPOSTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateUserGroupsUsingPOSTUnauthorized creates a CreateUserGroupsUsingPOSTUnauthorized with default headers values
func NewCreateUserGroupsUsingPOSTUnauthorized() *CreateUserGroupsUsingPOSTUnauthorized {
	return &CreateUserGroupsUsingPOSTUnauthorized{}
}

/*CreateUserGroupsUsingPOSTUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateUserGroupsUsingPOSTUnauthorized struct {
}

func (o *CreateUserGroupsUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /user-groups][%d] createUserGroupsUsingPOSTUnauthorized ", 401)
}

func (o *CreateUserGroupsUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateUserGroupsUsingPOSTForbidden creates a CreateUserGroupsUsingPOSTForbidden with default headers values
func NewCreateUserGroupsUsingPOSTForbidden() *CreateUserGroupsUsingPOSTForbidden {
	return &CreateUserGroupsUsingPOSTForbidden{}
}

/*CreateUserGroupsUsingPOSTForbidden handles this case with default header values.

Forbidden
*/
type CreateUserGroupsUsingPOSTForbidden struct {
}

func (o *CreateUserGroupsUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /user-groups][%d] createUserGroupsUsingPOSTForbidden ", 403)
}

func (o *CreateUserGroupsUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateUserGroupsUsingPOSTNotFound creates a CreateUserGroupsUsingPOSTNotFound with default headers values
func NewCreateUserGroupsUsingPOSTNotFound() *CreateUserGroupsUsingPOSTNotFound {
	return &CreateUserGroupsUsingPOSTNotFound{}
}

/*CreateUserGroupsUsingPOSTNotFound handles this case with default header values.

Not Found
*/
type CreateUserGroupsUsingPOSTNotFound struct {
}

func (o *CreateUserGroupsUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /user-groups][%d] createUserGroupsUsingPOSTNotFound ", 404)
}

func (o *CreateUserGroupsUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
