// Code generated by go-swagger; DO NOT EDIT.

package api_access_token

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/e88z4/go_cdd/administration/models"
)

// GetAPIAccessTokensUsingGET1Reader is a Reader for the GetAPIAccessTokensUsingGET1 structure.
type GetAPIAccessTokensUsingGET1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIAccessTokensUsingGET1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIAccessTokensUsingGET1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetAPIAccessTokensUsingGET1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAPIAccessTokensUsingGET1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAPIAccessTokensUsingGET1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAPIAccessTokensUsingGET1OK creates a GetAPIAccessTokensUsingGET1OK with default headers values
func NewGetAPIAccessTokensUsingGET1OK() *GetAPIAccessTokensUsingGET1OK {
	return &GetAPIAccessTokensUsingGET1OK{}
}

/*GetAPIAccessTokensUsingGET1OK handles this case with default header values.

OK
*/
type GetAPIAccessTokensUsingGET1OK struct {
	Payload *models.ListHolderDto
}

func (o *GetAPIAccessTokensUsingGET1OK) Error() string {
	return fmt.Sprintf("[GET /api-access-tokens][%d] getApiAccessTokensUsingGET1OK  %+v", 200, o.Payload)
}

func (o *GetAPIAccessTokensUsingGET1OK) GetPayload() *models.ListHolderDto {
	return o.Payload
}

func (o *GetAPIAccessTokensUsingGET1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListHolderDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIAccessTokensUsingGET1Unauthorized creates a GetAPIAccessTokensUsingGET1Unauthorized with default headers values
func NewGetAPIAccessTokensUsingGET1Unauthorized() *GetAPIAccessTokensUsingGET1Unauthorized {
	return &GetAPIAccessTokensUsingGET1Unauthorized{}
}

/*GetAPIAccessTokensUsingGET1Unauthorized handles this case with default header values.

Unauthorized
*/
type GetAPIAccessTokensUsingGET1Unauthorized struct {
}

func (o *GetAPIAccessTokensUsingGET1Unauthorized) Error() string {
	return fmt.Sprintf("[GET /api-access-tokens][%d] getApiAccessTokensUsingGET1Unauthorized ", 401)
}

func (o *GetAPIAccessTokensUsingGET1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAPIAccessTokensUsingGET1Forbidden creates a GetAPIAccessTokensUsingGET1Forbidden with default headers values
func NewGetAPIAccessTokensUsingGET1Forbidden() *GetAPIAccessTokensUsingGET1Forbidden {
	return &GetAPIAccessTokensUsingGET1Forbidden{}
}

/*GetAPIAccessTokensUsingGET1Forbidden handles this case with default header values.

Forbidden
*/
type GetAPIAccessTokensUsingGET1Forbidden struct {
}

func (o *GetAPIAccessTokensUsingGET1Forbidden) Error() string {
	return fmt.Sprintf("[GET /api-access-tokens][%d] getApiAccessTokensUsingGET1Forbidden ", 403)
}

func (o *GetAPIAccessTokensUsingGET1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAPIAccessTokensUsingGET1NotFound creates a GetAPIAccessTokensUsingGET1NotFound with default headers values
func NewGetAPIAccessTokensUsingGET1NotFound() *GetAPIAccessTokensUsingGET1NotFound {
	return &GetAPIAccessTokensUsingGET1NotFound{}
}

/*GetAPIAccessTokensUsingGET1NotFound handles this case with default header values.

Not Found
*/
type GetAPIAccessTokensUsingGET1NotFound struct {
}

func (o *GetAPIAccessTokensUsingGET1NotFound) Error() string {
	return fmt.Sprintf("[GET /api-access-tokens][%d] getApiAccessTokensUsingGET1NotFound ", 404)
}

func (o *GetAPIAccessTokensUsingGET1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
