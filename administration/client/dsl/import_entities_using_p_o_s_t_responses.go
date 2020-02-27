// Code generated by go-swagger; DO NOT EDIT.

package dsl

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/e88z4/go_cdd/administration/models"
)

// ImportEntitiesUsingPOSTReader is a Reader for the ImportEntitiesUsingPOST structure.
type ImportEntitiesUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImportEntitiesUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewImportEntitiesUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewImportEntitiesUsingPOSTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewImportEntitiesUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewImportEntitiesUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewImportEntitiesUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewImportEntitiesUsingPOSTOK creates a ImportEntitiesUsingPOSTOK with default headers values
func NewImportEntitiesUsingPOSTOK() *ImportEntitiesUsingPOSTOK {
	return &ImportEntitiesUsingPOSTOK{}
}

/*ImportEntitiesUsingPOSTOK handles this case with default header values.

OK
*/
type ImportEntitiesUsingPOSTOK struct {
	Payload *models.DslResponseDto
}

func (o *ImportEntitiesUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /dsl-manifest/attachment][%d] importEntitiesUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *ImportEntitiesUsingPOSTOK) GetPayload() *models.DslResponseDto {
	return o.Payload
}

func (o *ImportEntitiesUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DslResponseDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImportEntitiesUsingPOSTCreated creates a ImportEntitiesUsingPOSTCreated with default headers values
func NewImportEntitiesUsingPOSTCreated() *ImportEntitiesUsingPOSTCreated {
	return &ImportEntitiesUsingPOSTCreated{}
}

/*ImportEntitiesUsingPOSTCreated handles this case with default header values.

Created
*/
type ImportEntitiesUsingPOSTCreated struct {
}

func (o *ImportEntitiesUsingPOSTCreated) Error() string {
	return fmt.Sprintf("[POST /dsl-manifest/attachment][%d] importEntitiesUsingPOSTCreated ", 201)
}

func (o *ImportEntitiesUsingPOSTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewImportEntitiesUsingPOSTUnauthorized creates a ImportEntitiesUsingPOSTUnauthorized with default headers values
func NewImportEntitiesUsingPOSTUnauthorized() *ImportEntitiesUsingPOSTUnauthorized {
	return &ImportEntitiesUsingPOSTUnauthorized{}
}

/*ImportEntitiesUsingPOSTUnauthorized handles this case with default header values.

Unauthorized
*/
type ImportEntitiesUsingPOSTUnauthorized struct {
}

func (o *ImportEntitiesUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /dsl-manifest/attachment][%d] importEntitiesUsingPOSTUnauthorized ", 401)
}

func (o *ImportEntitiesUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewImportEntitiesUsingPOSTForbidden creates a ImportEntitiesUsingPOSTForbidden with default headers values
func NewImportEntitiesUsingPOSTForbidden() *ImportEntitiesUsingPOSTForbidden {
	return &ImportEntitiesUsingPOSTForbidden{}
}

/*ImportEntitiesUsingPOSTForbidden handles this case with default header values.

Forbidden
*/
type ImportEntitiesUsingPOSTForbidden struct {
}

func (o *ImportEntitiesUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /dsl-manifest/attachment][%d] importEntitiesUsingPOSTForbidden ", 403)
}

func (o *ImportEntitiesUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewImportEntitiesUsingPOSTNotFound creates a ImportEntitiesUsingPOSTNotFound with default headers values
func NewImportEntitiesUsingPOSTNotFound() *ImportEntitiesUsingPOSTNotFound {
	return &ImportEntitiesUsingPOSTNotFound{}
}

/*ImportEntitiesUsingPOSTNotFound handles this case with default header values.

Not Found
*/
type ImportEntitiesUsingPOSTNotFound struct {
}

func (o *ImportEntitiesUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /dsl-manifest/attachment][%d] importEntitiesUsingPOSTNotFound ", 404)
}

func (o *ImportEntitiesUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}