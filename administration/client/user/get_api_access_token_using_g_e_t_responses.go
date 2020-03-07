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

// GetAPIAccessTokenUsingGETReader is a Reader for the GetAPIAccessTokenUsingGET structure.
type GetAPIAccessTokenUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIAccessTokenUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIAccessTokenUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetAPIAccessTokenUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAPIAccessTokenUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAPIAccessTokenUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAPIAccessTokenUsingGETOK creates a GetAPIAccessTokenUsingGETOK with default headers values
func NewGetAPIAccessTokenUsingGETOK() *GetAPIAccessTokenUsingGETOK {
	return &GetAPIAccessTokenUsingGETOK{}
}

/*GetAPIAccessTokenUsingGETOK handles this case with default header values.

OK
*/
type GetAPIAccessTokenUsingGETOK struct {
	Payload *models.APIAccessTokenDto
}

func (o *GetAPIAccessTokenUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /users/{userId}/api-access-tokens][%d] getApiAccessTokenUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetAPIAccessTokenUsingGETOK) GetPayload() *models.APIAccessTokenDto {
	return o.Payload
}

func (o *GetAPIAccessTokenUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIAccessTokenDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIAccessTokenUsingGETUnauthorized creates a GetAPIAccessTokenUsingGETUnauthorized with default headers values
func NewGetAPIAccessTokenUsingGETUnauthorized() *GetAPIAccessTokenUsingGETUnauthorized {
	return &GetAPIAccessTokenUsingGETUnauthorized{}
}

/*GetAPIAccessTokenUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type GetAPIAccessTokenUsingGETUnauthorized struct {
}

func (o *GetAPIAccessTokenUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /users/{userId}/api-access-tokens][%d] getApiAccessTokenUsingGETUnauthorized ", 401)
}

func (o *GetAPIAccessTokenUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAPIAccessTokenUsingGETForbidden creates a GetAPIAccessTokenUsingGETForbidden with default headers values
func NewGetAPIAccessTokenUsingGETForbidden() *GetAPIAccessTokenUsingGETForbidden {
	return &GetAPIAccessTokenUsingGETForbidden{}
}

/*GetAPIAccessTokenUsingGETForbidden handles this case with default header values.

Forbidden
*/
type GetAPIAccessTokenUsingGETForbidden struct {
}

func (o *GetAPIAccessTokenUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /users/{userId}/api-access-tokens][%d] getApiAccessTokenUsingGETForbidden ", 403)
}

func (o *GetAPIAccessTokenUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAPIAccessTokenUsingGETNotFound creates a GetAPIAccessTokenUsingGETNotFound with default headers values
func NewGetAPIAccessTokenUsingGETNotFound() *GetAPIAccessTokenUsingGETNotFound {
	return &GetAPIAccessTokenUsingGETNotFound{}
}

/*GetAPIAccessTokenUsingGETNotFound handles this case with default header values.

Not Found
*/
type GetAPIAccessTokenUsingGETNotFound struct {
}

func (o *GetAPIAccessTokenUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /users/{userId}/api-access-tokens][%d] getApiAccessTokenUsingGETNotFound ", 404)
}

func (o *GetAPIAccessTokenUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
