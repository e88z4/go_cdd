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

// CreateAPIAccessTokenUsingPOST1Reader is a Reader for the CreateAPIAccessTokenUsingPOST1 structure.
type CreateAPIAccessTokenUsingPOST1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAPIAccessTokenUsingPOST1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateAPIAccessTokenUsingPOST1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewCreateAPIAccessTokenUsingPOST1Created()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateAPIAccessTokenUsingPOST1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateAPIAccessTokenUsingPOST1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateAPIAccessTokenUsingPOST1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateAPIAccessTokenUsingPOST1OK creates a CreateAPIAccessTokenUsingPOST1OK with default headers values
func NewCreateAPIAccessTokenUsingPOST1OK() *CreateAPIAccessTokenUsingPOST1OK {
	return &CreateAPIAccessTokenUsingPOST1OK{}
}

/*CreateAPIAccessTokenUsingPOST1OK handles this case with default header values.

OK
*/
type CreateAPIAccessTokenUsingPOST1OK struct {
	Payload *models.APIAccessTokenDto
}

func (o *CreateAPIAccessTokenUsingPOST1OK) Error() string {
	return fmt.Sprintf("[POST /users/{userId}/api-access-tokens][%d] createApiAccessTokenUsingPOST1OK  %+v", 200, o.Payload)
}

func (o *CreateAPIAccessTokenUsingPOST1OK) GetPayload() *models.APIAccessTokenDto {
	return o.Payload
}

func (o *CreateAPIAccessTokenUsingPOST1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIAccessTokenDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAPIAccessTokenUsingPOST1Created creates a CreateAPIAccessTokenUsingPOST1Created with default headers values
func NewCreateAPIAccessTokenUsingPOST1Created() *CreateAPIAccessTokenUsingPOST1Created {
	return &CreateAPIAccessTokenUsingPOST1Created{}
}

/*CreateAPIAccessTokenUsingPOST1Created handles this case with default header values.

Created
*/
type CreateAPIAccessTokenUsingPOST1Created struct {
}

func (o *CreateAPIAccessTokenUsingPOST1Created) Error() string {
	return fmt.Sprintf("[POST /users/{userId}/api-access-tokens][%d] createApiAccessTokenUsingPOST1Created ", 201)
}

func (o *CreateAPIAccessTokenUsingPOST1Created) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateAPIAccessTokenUsingPOST1Unauthorized creates a CreateAPIAccessTokenUsingPOST1Unauthorized with default headers values
func NewCreateAPIAccessTokenUsingPOST1Unauthorized() *CreateAPIAccessTokenUsingPOST1Unauthorized {
	return &CreateAPIAccessTokenUsingPOST1Unauthorized{}
}

/*CreateAPIAccessTokenUsingPOST1Unauthorized handles this case with default header values.

Unauthorized
*/
type CreateAPIAccessTokenUsingPOST1Unauthorized struct {
}

func (o *CreateAPIAccessTokenUsingPOST1Unauthorized) Error() string {
	return fmt.Sprintf("[POST /users/{userId}/api-access-tokens][%d] createApiAccessTokenUsingPOST1Unauthorized ", 401)
}

func (o *CreateAPIAccessTokenUsingPOST1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateAPIAccessTokenUsingPOST1Forbidden creates a CreateAPIAccessTokenUsingPOST1Forbidden with default headers values
func NewCreateAPIAccessTokenUsingPOST1Forbidden() *CreateAPIAccessTokenUsingPOST1Forbidden {
	return &CreateAPIAccessTokenUsingPOST1Forbidden{}
}

/*CreateAPIAccessTokenUsingPOST1Forbidden handles this case with default header values.

Forbidden
*/
type CreateAPIAccessTokenUsingPOST1Forbidden struct {
}

func (o *CreateAPIAccessTokenUsingPOST1Forbidden) Error() string {
	return fmt.Sprintf("[POST /users/{userId}/api-access-tokens][%d] createApiAccessTokenUsingPOST1Forbidden ", 403)
}

func (o *CreateAPIAccessTokenUsingPOST1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateAPIAccessTokenUsingPOST1NotFound creates a CreateAPIAccessTokenUsingPOST1NotFound with default headers values
func NewCreateAPIAccessTokenUsingPOST1NotFound() *CreateAPIAccessTokenUsingPOST1NotFound {
	return &CreateAPIAccessTokenUsingPOST1NotFound{}
}

/*CreateAPIAccessTokenUsingPOST1NotFound handles this case with default header values.

Not Found
*/
type CreateAPIAccessTokenUsingPOST1NotFound struct {
}

func (o *CreateAPIAccessTokenUsingPOST1NotFound) Error() string {
	return fmt.Sprintf("[POST /users/{userId}/api-access-tokens][%d] createApiAccessTokenUsingPOST1NotFound ", 404)
}

func (o *CreateAPIAccessTokenUsingPOST1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}