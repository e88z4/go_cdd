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

// AddPartiesToProjectUsingPOSTReader is a Reader for the AddPartiesToProjectUsingPOST structure.
type AddPartiesToProjectUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddPartiesToProjectUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddPartiesToProjectUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewAddPartiesToProjectUsingPOSTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAddPartiesToProjectUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAddPartiesToProjectUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAddPartiesToProjectUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddPartiesToProjectUsingPOSTOK creates a AddPartiesToProjectUsingPOSTOK with default headers values
func NewAddPartiesToProjectUsingPOSTOK() *AddPartiesToProjectUsingPOSTOK {
	return &AddPartiesToProjectUsingPOSTOK{}
}

/*AddPartiesToProjectUsingPOSTOK handles this case with default header values.

OK
*/
type AddPartiesToProjectUsingPOSTOK struct {
	Payload *models.ListHolderDtoPartyDto
}

func (o *AddPartiesToProjectUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /projects/{projectId}/parties][%d] addPartiesToProjectUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *AddPartiesToProjectUsingPOSTOK) GetPayload() *models.ListHolderDtoPartyDto {
	return o.Payload
}

func (o *AddPartiesToProjectUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListHolderDtoPartyDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddPartiesToProjectUsingPOSTCreated creates a AddPartiesToProjectUsingPOSTCreated with default headers values
func NewAddPartiesToProjectUsingPOSTCreated() *AddPartiesToProjectUsingPOSTCreated {
	return &AddPartiesToProjectUsingPOSTCreated{}
}

/*AddPartiesToProjectUsingPOSTCreated handles this case with default header values.

Created
*/
type AddPartiesToProjectUsingPOSTCreated struct {
}

func (o *AddPartiesToProjectUsingPOSTCreated) Error() string {
	return fmt.Sprintf("[POST /projects/{projectId}/parties][%d] addPartiesToProjectUsingPOSTCreated ", 201)
}

func (o *AddPartiesToProjectUsingPOSTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddPartiesToProjectUsingPOSTUnauthorized creates a AddPartiesToProjectUsingPOSTUnauthorized with default headers values
func NewAddPartiesToProjectUsingPOSTUnauthorized() *AddPartiesToProjectUsingPOSTUnauthorized {
	return &AddPartiesToProjectUsingPOSTUnauthorized{}
}

/*AddPartiesToProjectUsingPOSTUnauthorized handles this case with default header values.

Unauthorized
*/
type AddPartiesToProjectUsingPOSTUnauthorized struct {
}

func (o *AddPartiesToProjectUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /projects/{projectId}/parties][%d] addPartiesToProjectUsingPOSTUnauthorized ", 401)
}

func (o *AddPartiesToProjectUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddPartiesToProjectUsingPOSTForbidden creates a AddPartiesToProjectUsingPOSTForbidden with default headers values
func NewAddPartiesToProjectUsingPOSTForbidden() *AddPartiesToProjectUsingPOSTForbidden {
	return &AddPartiesToProjectUsingPOSTForbidden{}
}

/*AddPartiesToProjectUsingPOSTForbidden handles this case with default header values.

Forbidden
*/
type AddPartiesToProjectUsingPOSTForbidden struct {
}

func (o *AddPartiesToProjectUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /projects/{projectId}/parties][%d] addPartiesToProjectUsingPOSTForbidden ", 403)
}

func (o *AddPartiesToProjectUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddPartiesToProjectUsingPOSTNotFound creates a AddPartiesToProjectUsingPOSTNotFound with default headers values
func NewAddPartiesToProjectUsingPOSTNotFound() *AddPartiesToProjectUsingPOSTNotFound {
	return &AddPartiesToProjectUsingPOSTNotFound{}
}

/*AddPartiesToProjectUsingPOSTNotFound handles this case with default header values.

Not Found
*/
type AddPartiesToProjectUsingPOSTNotFound struct {
}

func (o *AddPartiesToProjectUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /projects/{projectId}/parties][%d] addPartiesToProjectUsingPOSTNotFound ", 404)
}

func (o *AddPartiesToProjectUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
