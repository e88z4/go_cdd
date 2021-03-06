// Code generated by go-swagger; DO NOT EDIT.

package tenant

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/e88z4/go_cdd/administration/models"
)

// GetAllTenantsUsingGETReader is a Reader for the GetAllTenantsUsingGET structure.
type GetAllTenantsUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllTenantsUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllTenantsUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetAllTenantsUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAllTenantsUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAllTenantsUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAllTenantsUsingGETOK creates a GetAllTenantsUsingGETOK with default headers values
func NewGetAllTenantsUsingGETOK() *GetAllTenantsUsingGETOK {
	return &GetAllTenantsUsingGETOK{}
}

/*GetAllTenantsUsingGETOK handles this case with default header values.

OK
*/
type GetAllTenantsUsingGETOK struct {
	Payload *models.ListHolderDtoTenantDto
}

func (o *GetAllTenantsUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /tenants][%d] getAllTenantsUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetAllTenantsUsingGETOK) GetPayload() *models.ListHolderDtoTenantDto {
	return o.Payload
}

func (o *GetAllTenantsUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListHolderDtoTenantDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllTenantsUsingGETUnauthorized creates a GetAllTenantsUsingGETUnauthorized with default headers values
func NewGetAllTenantsUsingGETUnauthorized() *GetAllTenantsUsingGETUnauthorized {
	return &GetAllTenantsUsingGETUnauthorized{}
}

/*GetAllTenantsUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type GetAllTenantsUsingGETUnauthorized struct {
}

func (o *GetAllTenantsUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /tenants][%d] getAllTenantsUsingGETUnauthorized ", 401)
}

func (o *GetAllTenantsUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllTenantsUsingGETForbidden creates a GetAllTenantsUsingGETForbidden with default headers values
func NewGetAllTenantsUsingGETForbidden() *GetAllTenantsUsingGETForbidden {
	return &GetAllTenantsUsingGETForbidden{}
}

/*GetAllTenantsUsingGETForbidden handles this case with default header values.

Forbidden
*/
type GetAllTenantsUsingGETForbidden struct {
}

func (o *GetAllTenantsUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /tenants][%d] getAllTenantsUsingGETForbidden ", 403)
}

func (o *GetAllTenantsUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllTenantsUsingGETNotFound creates a GetAllTenantsUsingGETNotFound with default headers values
func NewGetAllTenantsUsingGETNotFound() *GetAllTenantsUsingGETNotFound {
	return &GetAllTenantsUsingGETNotFound{}
}

/*GetAllTenantsUsingGETNotFound handles this case with default header values.

Not Found
*/
type GetAllTenantsUsingGETNotFound struct {
}

func (o *GetAllTenantsUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /tenants][%d] getAllTenantsUsingGETNotFound ", 404)
}

func (o *GetAllTenantsUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
