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

// GetTenantUsingGETReader is a Reader for the GetTenantUsingGET structure.
type GetTenantUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTenantUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTenantUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetTenantUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTenantUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTenantUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetTenantUsingGETOK creates a GetTenantUsingGETOK with default headers values
func NewGetTenantUsingGETOK() *GetTenantUsingGETOK {
	return &GetTenantUsingGETOK{}
}

/*GetTenantUsingGETOK handles this case with default header values.

OK
*/
type GetTenantUsingGETOK struct {
	Payload *models.TenantDto
}

func (o *GetTenantUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /tenants/{tenantId}][%d] getTenantUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetTenantUsingGETOK) GetPayload() *models.TenantDto {
	return o.Payload
}

func (o *GetTenantUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TenantDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTenantUsingGETUnauthorized creates a GetTenantUsingGETUnauthorized with default headers values
func NewGetTenantUsingGETUnauthorized() *GetTenantUsingGETUnauthorized {
	return &GetTenantUsingGETUnauthorized{}
}

/*GetTenantUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type GetTenantUsingGETUnauthorized struct {
}

func (o *GetTenantUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /tenants/{tenantId}][%d] getTenantUsingGETUnauthorized ", 401)
}

func (o *GetTenantUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTenantUsingGETForbidden creates a GetTenantUsingGETForbidden with default headers values
func NewGetTenantUsingGETForbidden() *GetTenantUsingGETForbidden {
	return &GetTenantUsingGETForbidden{}
}

/*GetTenantUsingGETForbidden handles this case with default header values.

Forbidden
*/
type GetTenantUsingGETForbidden struct {
}

func (o *GetTenantUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /tenants/{tenantId}][%d] getTenantUsingGETForbidden ", 403)
}

func (o *GetTenantUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTenantUsingGETNotFound creates a GetTenantUsingGETNotFound with default headers values
func NewGetTenantUsingGETNotFound() *GetTenantUsingGETNotFound {
	return &GetTenantUsingGETNotFound{}
}

/*GetTenantUsingGETNotFound handles this case with default header values.

Not Found
*/
type GetTenantUsingGETNotFound struct {
}

func (o *GetTenantUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /tenants/{tenantId}][%d] getTenantUsingGETNotFound ", 404)
}

func (o *GetTenantUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}