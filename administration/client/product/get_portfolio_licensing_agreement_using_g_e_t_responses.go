// Code generated by go-swagger; DO NOT EDIT.

package product

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/e88z4/go_cdd/administration/models"
)

// GetPortfolioLicensingAgreementUsingGETReader is a Reader for the GetPortfolioLicensingAgreementUsingGET structure.
type GetPortfolioLicensingAgreementUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPortfolioLicensingAgreementUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPortfolioLicensingAgreementUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetPortfolioLicensingAgreementUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetPortfolioLicensingAgreementUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetPortfolioLicensingAgreementUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPortfolioLicensingAgreementUsingGETOK creates a GetPortfolioLicensingAgreementUsingGETOK with default headers values
func NewGetPortfolioLicensingAgreementUsingGETOK() *GetPortfolioLicensingAgreementUsingGETOK {
	return &GetPortfolioLicensingAgreementUsingGETOK{}
}

/*GetPortfolioLicensingAgreementUsingGETOK handles this case with default header values.

OK
*/
type GetPortfolioLicensingAgreementUsingGETOK struct {
	Payload *models.PortfolioLicensingAgreementDto
}

func (o *GetPortfolioLicensingAgreementUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /product/license][%d] getPortfolioLicensingAgreementUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetPortfolioLicensingAgreementUsingGETOK) GetPayload() *models.PortfolioLicensingAgreementDto {
	return o.Payload
}

func (o *GetPortfolioLicensingAgreementUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PortfolioLicensingAgreementDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPortfolioLicensingAgreementUsingGETUnauthorized creates a GetPortfolioLicensingAgreementUsingGETUnauthorized with default headers values
func NewGetPortfolioLicensingAgreementUsingGETUnauthorized() *GetPortfolioLicensingAgreementUsingGETUnauthorized {
	return &GetPortfolioLicensingAgreementUsingGETUnauthorized{}
}

/*GetPortfolioLicensingAgreementUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type GetPortfolioLicensingAgreementUsingGETUnauthorized struct {
}

func (o *GetPortfolioLicensingAgreementUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /product/license][%d] getPortfolioLicensingAgreementUsingGETUnauthorized ", 401)
}

func (o *GetPortfolioLicensingAgreementUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPortfolioLicensingAgreementUsingGETForbidden creates a GetPortfolioLicensingAgreementUsingGETForbidden with default headers values
func NewGetPortfolioLicensingAgreementUsingGETForbidden() *GetPortfolioLicensingAgreementUsingGETForbidden {
	return &GetPortfolioLicensingAgreementUsingGETForbidden{}
}

/*GetPortfolioLicensingAgreementUsingGETForbidden handles this case with default header values.

Forbidden
*/
type GetPortfolioLicensingAgreementUsingGETForbidden struct {
}

func (o *GetPortfolioLicensingAgreementUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /product/license][%d] getPortfolioLicensingAgreementUsingGETForbidden ", 403)
}

func (o *GetPortfolioLicensingAgreementUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPortfolioLicensingAgreementUsingGETNotFound creates a GetPortfolioLicensingAgreementUsingGETNotFound with default headers values
func NewGetPortfolioLicensingAgreementUsingGETNotFound() *GetPortfolioLicensingAgreementUsingGETNotFound {
	return &GetPortfolioLicensingAgreementUsingGETNotFound{}
}

/*GetPortfolioLicensingAgreementUsingGETNotFound handles this case with default header values.

Not Found
*/
type GetPortfolioLicensingAgreementUsingGETNotFound struct {
}

func (o *GetPortfolioLicensingAgreementUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /product/license][%d] getPortfolioLicensingAgreementUsingGETNotFound ", 404)
}

func (o *GetPortfolioLicensingAgreementUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
