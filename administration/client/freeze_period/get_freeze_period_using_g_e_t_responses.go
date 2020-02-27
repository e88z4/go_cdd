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

// GetFreezePeriodUsingGETReader is a Reader for the GetFreezePeriodUsingGET structure.
type GetFreezePeriodUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFreezePeriodUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFreezePeriodUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetFreezePeriodUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetFreezePeriodUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetFreezePeriodUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetFreezePeriodUsingGETOK creates a GetFreezePeriodUsingGETOK with default headers values
func NewGetFreezePeriodUsingGETOK() *GetFreezePeriodUsingGETOK {
	return &GetFreezePeriodUsingGETOK{}
}

/*GetFreezePeriodUsingGETOK handles this case with default header values.

OK
*/
type GetFreezePeriodUsingGETOK struct {
	Payload *models.FreezePeriodDto
}

func (o *GetFreezePeriodUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /freeze-periods/{freezePeriodId}][%d] getFreezePeriodUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetFreezePeriodUsingGETOK) GetPayload() *models.FreezePeriodDto {
	return o.Payload
}

func (o *GetFreezePeriodUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FreezePeriodDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFreezePeriodUsingGETUnauthorized creates a GetFreezePeriodUsingGETUnauthorized with default headers values
func NewGetFreezePeriodUsingGETUnauthorized() *GetFreezePeriodUsingGETUnauthorized {
	return &GetFreezePeriodUsingGETUnauthorized{}
}

/*GetFreezePeriodUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type GetFreezePeriodUsingGETUnauthorized struct {
}

func (o *GetFreezePeriodUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /freeze-periods/{freezePeriodId}][%d] getFreezePeriodUsingGETUnauthorized ", 401)
}

func (o *GetFreezePeriodUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFreezePeriodUsingGETForbidden creates a GetFreezePeriodUsingGETForbidden with default headers values
func NewGetFreezePeriodUsingGETForbidden() *GetFreezePeriodUsingGETForbidden {
	return &GetFreezePeriodUsingGETForbidden{}
}

/*GetFreezePeriodUsingGETForbidden handles this case with default header values.

Forbidden
*/
type GetFreezePeriodUsingGETForbidden struct {
}

func (o *GetFreezePeriodUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /freeze-periods/{freezePeriodId}][%d] getFreezePeriodUsingGETForbidden ", 403)
}

func (o *GetFreezePeriodUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFreezePeriodUsingGETNotFound creates a GetFreezePeriodUsingGETNotFound with default headers values
func NewGetFreezePeriodUsingGETNotFound() *GetFreezePeriodUsingGETNotFound {
	return &GetFreezePeriodUsingGETNotFound{}
}

/*GetFreezePeriodUsingGETNotFound handles this case with default header values.

Not Found
*/
type GetFreezePeriodUsingGETNotFound struct {
}

func (o *GetFreezePeriodUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /freeze-periods/{freezePeriodId}][%d] getFreezePeriodUsingGETNotFound ", 404)
}

func (o *GetFreezePeriodUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}