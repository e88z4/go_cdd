// Code generated by go-swagger; DO NOT EDIT.

package payment_plan_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/e88z4/go_cdd/administration/models"
)

// GetCurrentPaymentPlanUsingGETReader is a Reader for the GetCurrentPaymentPlanUsingGET structure.
type GetCurrentPaymentPlanUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCurrentPaymentPlanUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCurrentPaymentPlanUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetCurrentPaymentPlanUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCurrentPaymentPlanUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetCurrentPaymentPlanUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCurrentPaymentPlanUsingGETOK creates a GetCurrentPaymentPlanUsingGETOK with default headers values
func NewGetCurrentPaymentPlanUsingGETOK() *GetCurrentPaymentPlanUsingGETOK {
	return &GetCurrentPaymentPlanUsingGETOK{}
}

/*GetCurrentPaymentPlanUsingGETOK handles this case with default header values.

OK
*/
type GetCurrentPaymentPlanUsingGETOK struct {
	Payload *models.PaymentPlanDto
}

func (o *GetCurrentPaymentPlanUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /payment-plans/current][%d] getCurrentPaymentPlanUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetCurrentPaymentPlanUsingGETOK) GetPayload() *models.PaymentPlanDto {
	return o.Payload
}

func (o *GetCurrentPaymentPlanUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PaymentPlanDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCurrentPaymentPlanUsingGETUnauthorized creates a GetCurrentPaymentPlanUsingGETUnauthorized with default headers values
func NewGetCurrentPaymentPlanUsingGETUnauthorized() *GetCurrentPaymentPlanUsingGETUnauthorized {
	return &GetCurrentPaymentPlanUsingGETUnauthorized{}
}

/*GetCurrentPaymentPlanUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type GetCurrentPaymentPlanUsingGETUnauthorized struct {
}

func (o *GetCurrentPaymentPlanUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /payment-plans/current][%d] getCurrentPaymentPlanUsingGETUnauthorized ", 401)
}

func (o *GetCurrentPaymentPlanUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCurrentPaymentPlanUsingGETForbidden creates a GetCurrentPaymentPlanUsingGETForbidden with default headers values
func NewGetCurrentPaymentPlanUsingGETForbidden() *GetCurrentPaymentPlanUsingGETForbidden {
	return &GetCurrentPaymentPlanUsingGETForbidden{}
}

/*GetCurrentPaymentPlanUsingGETForbidden handles this case with default header values.

Forbidden
*/
type GetCurrentPaymentPlanUsingGETForbidden struct {
}

func (o *GetCurrentPaymentPlanUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /payment-plans/current][%d] getCurrentPaymentPlanUsingGETForbidden ", 403)
}

func (o *GetCurrentPaymentPlanUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCurrentPaymentPlanUsingGETNotFound creates a GetCurrentPaymentPlanUsingGETNotFound with default headers values
func NewGetCurrentPaymentPlanUsingGETNotFound() *GetCurrentPaymentPlanUsingGETNotFound {
	return &GetCurrentPaymentPlanUsingGETNotFound{}
}

/*GetCurrentPaymentPlanUsingGETNotFound handles this case with default header values.

Not Found
*/
type GetCurrentPaymentPlanUsingGETNotFound struct {
}

func (o *GetCurrentPaymentPlanUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /payment-plans/current][%d] getCurrentPaymentPlanUsingGETNotFound ", 404)
}

func (o *GetCurrentPaymentPlanUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
