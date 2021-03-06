// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/e88z4/go_cdd/administration/models"
)

// CreateProjectUsingPOSTReader is a Reader for the CreateProjectUsingPOST structure.
type CreateProjectUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateProjectUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateProjectUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewCreateProjectUsingPOSTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateProjectUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateProjectUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateProjectUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateProjectUsingPOSTOK creates a CreateProjectUsingPOSTOK with default headers values
func NewCreateProjectUsingPOSTOK() *CreateProjectUsingPOSTOK {
	return &CreateProjectUsingPOSTOK{}
}

/*CreateProjectUsingPOSTOK handles this case with default header values.

OK
*/
type CreateProjectUsingPOSTOK struct {
	Payload *models.ProjectDto
}

func (o *CreateProjectUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /projects][%d] createProjectUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *CreateProjectUsingPOSTOK) GetPayload() *models.ProjectDto {
	return o.Payload
}

func (o *CreateProjectUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProjectDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateProjectUsingPOSTCreated creates a CreateProjectUsingPOSTCreated with default headers values
func NewCreateProjectUsingPOSTCreated() *CreateProjectUsingPOSTCreated {
	return &CreateProjectUsingPOSTCreated{}
}

/*CreateProjectUsingPOSTCreated handles this case with default header values.

Created
*/
type CreateProjectUsingPOSTCreated struct {
}

func (o *CreateProjectUsingPOSTCreated) Error() string {
	return fmt.Sprintf("[POST /projects][%d] createProjectUsingPOSTCreated ", 201)
}

func (o *CreateProjectUsingPOSTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateProjectUsingPOSTUnauthorized creates a CreateProjectUsingPOSTUnauthorized with default headers values
func NewCreateProjectUsingPOSTUnauthorized() *CreateProjectUsingPOSTUnauthorized {
	return &CreateProjectUsingPOSTUnauthorized{}
}

/*CreateProjectUsingPOSTUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateProjectUsingPOSTUnauthorized struct {
}

func (o *CreateProjectUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /projects][%d] createProjectUsingPOSTUnauthorized ", 401)
}

func (o *CreateProjectUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateProjectUsingPOSTForbidden creates a CreateProjectUsingPOSTForbidden with default headers values
func NewCreateProjectUsingPOSTForbidden() *CreateProjectUsingPOSTForbidden {
	return &CreateProjectUsingPOSTForbidden{}
}

/*CreateProjectUsingPOSTForbidden handles this case with default header values.

Forbidden
*/
type CreateProjectUsingPOSTForbidden struct {
}

func (o *CreateProjectUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /projects][%d] createProjectUsingPOSTForbidden ", 403)
}

func (o *CreateProjectUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateProjectUsingPOSTNotFound creates a CreateProjectUsingPOSTNotFound with default headers values
func NewCreateProjectUsingPOSTNotFound() *CreateProjectUsingPOSTNotFound {
	return &CreateProjectUsingPOSTNotFound{}
}

/*CreateProjectUsingPOSTNotFound handles this case with default header values.

Not Found
*/
type CreateProjectUsingPOSTNotFound struct {
}

func (o *CreateProjectUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /projects][%d] createProjectUsingPOSTNotFound ", 404)
}

func (o *CreateProjectUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
