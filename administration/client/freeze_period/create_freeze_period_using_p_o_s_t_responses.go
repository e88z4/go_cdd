// Code generated by go-swagger; DO NOT EDIT.

package freeze_period

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/e88z4/go_cdd/administration/models"
)

// CreateFreezePeriodUsingPOSTReader is a Reader for the CreateFreezePeriodUsingPOST structure.
type CreateFreezePeriodUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateFreezePeriodUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateFreezePeriodUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewCreateFreezePeriodUsingPOSTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateFreezePeriodUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateFreezePeriodUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateFreezePeriodUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateFreezePeriodUsingPOSTOK creates a CreateFreezePeriodUsingPOSTOK with default headers values
func NewCreateFreezePeriodUsingPOSTOK() *CreateFreezePeriodUsingPOSTOK {
	return &CreateFreezePeriodUsingPOSTOK{}
}

/*CreateFreezePeriodUsingPOSTOK handles this case with default header values.

OK
*/
type CreateFreezePeriodUsingPOSTOK struct {
	Payload *models.FreezePeriodDto
}

func (o *CreateFreezePeriodUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /freeze-periods][%d] createFreezePeriodUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *CreateFreezePeriodUsingPOSTOK) GetPayload() *models.FreezePeriodDto {
	return o.Payload
}

func (o *CreateFreezePeriodUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FreezePeriodDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateFreezePeriodUsingPOSTCreated creates a CreateFreezePeriodUsingPOSTCreated with default headers values
func NewCreateFreezePeriodUsingPOSTCreated() *CreateFreezePeriodUsingPOSTCreated {
	return &CreateFreezePeriodUsingPOSTCreated{}
}

/*CreateFreezePeriodUsingPOSTCreated handles this case with default header values.

Created
*/
type CreateFreezePeriodUsingPOSTCreated struct {
}

func (o *CreateFreezePeriodUsingPOSTCreated) Error() string {
	return fmt.Sprintf("[POST /freeze-periods][%d] createFreezePeriodUsingPOSTCreated ", 201)
}

func (o *CreateFreezePeriodUsingPOSTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateFreezePeriodUsingPOSTUnauthorized creates a CreateFreezePeriodUsingPOSTUnauthorized with default headers values
func NewCreateFreezePeriodUsingPOSTUnauthorized() *CreateFreezePeriodUsingPOSTUnauthorized {
	return &CreateFreezePeriodUsingPOSTUnauthorized{}
}

/*CreateFreezePeriodUsingPOSTUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateFreezePeriodUsingPOSTUnauthorized struct {
}

func (o *CreateFreezePeriodUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /freeze-periods][%d] createFreezePeriodUsingPOSTUnauthorized ", 401)
}

func (o *CreateFreezePeriodUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateFreezePeriodUsingPOSTForbidden creates a CreateFreezePeriodUsingPOSTForbidden with default headers values
func NewCreateFreezePeriodUsingPOSTForbidden() *CreateFreezePeriodUsingPOSTForbidden {
	return &CreateFreezePeriodUsingPOSTForbidden{}
}

/*CreateFreezePeriodUsingPOSTForbidden handles this case with default header values.

Forbidden
*/
type CreateFreezePeriodUsingPOSTForbidden struct {
}

func (o *CreateFreezePeriodUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /freeze-periods][%d] createFreezePeriodUsingPOSTForbidden ", 403)
}

func (o *CreateFreezePeriodUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateFreezePeriodUsingPOSTNotFound creates a CreateFreezePeriodUsingPOSTNotFound with default headers values
func NewCreateFreezePeriodUsingPOSTNotFound() *CreateFreezePeriodUsingPOSTNotFound {
	return &CreateFreezePeriodUsingPOSTNotFound{}
}

/*CreateFreezePeriodUsingPOSTNotFound handles this case with default header values.

Not Found
*/
type CreateFreezePeriodUsingPOSTNotFound struct {
}

func (o *CreateFreezePeriodUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /freeze-periods][%d] createFreezePeriodUsingPOSTNotFound ", 404)
}

func (o *CreateFreezePeriodUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
