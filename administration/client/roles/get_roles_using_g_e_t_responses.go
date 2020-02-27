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

// GetRolesUsingGETReader is a Reader for the GetRolesUsingGET structure.
type GetRolesUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRolesUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRolesUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetRolesUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRolesUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRolesUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRolesUsingGETOK creates a GetRolesUsingGETOK with default headers values
func NewGetRolesUsingGETOK() *GetRolesUsingGETOK {
	return &GetRolesUsingGETOK{}
}

/*GetRolesUsingGETOK handles this case with default header values.

OK
*/
type GetRolesUsingGETOK struct {
	Payload *models.ListHolderDtoRoleDto
}

func (o *GetRolesUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /roles][%d] getRolesUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetRolesUsingGETOK) GetPayload() *models.ListHolderDtoRoleDto {
	return o.Payload
}

func (o *GetRolesUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListHolderDtoRoleDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRolesUsingGETUnauthorized creates a GetRolesUsingGETUnauthorized with default headers values
func NewGetRolesUsingGETUnauthorized() *GetRolesUsingGETUnauthorized {
	return &GetRolesUsingGETUnauthorized{}
}

/*GetRolesUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type GetRolesUsingGETUnauthorized struct {
}

func (o *GetRolesUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /roles][%d] getRolesUsingGETUnauthorized ", 401)
}

func (o *GetRolesUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRolesUsingGETForbidden creates a GetRolesUsingGETForbidden with default headers values
func NewGetRolesUsingGETForbidden() *GetRolesUsingGETForbidden {
	return &GetRolesUsingGETForbidden{}
}

/*GetRolesUsingGETForbidden handles this case with default header values.

Forbidden
*/
type GetRolesUsingGETForbidden struct {
}

func (o *GetRolesUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /roles][%d] getRolesUsingGETForbidden ", 403)
}

func (o *GetRolesUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRolesUsingGETNotFound creates a GetRolesUsingGETNotFound with default headers values
func NewGetRolesUsingGETNotFound() *GetRolesUsingGETNotFound {
	return &GetRolesUsingGETNotFound{}
}

/*GetRolesUsingGETNotFound handles this case with default header values.

Not Found
*/
type GetRolesUsingGETNotFound struct {
}

func (o *GetRolesUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /roles][%d] getRolesUsingGETNotFound ", 404)
}

func (o *GetRolesUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}