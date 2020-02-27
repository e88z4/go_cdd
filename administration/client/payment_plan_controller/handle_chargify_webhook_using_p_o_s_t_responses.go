// Code generated by go-swagger; DO NOT EDIT.

package payment_plan_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// HandleChargifyWebhookUsingPOSTReader is a Reader for the HandleChargifyWebhookUsingPOST structure.
type HandleChargifyWebhookUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HandleChargifyWebhookUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHandleChargifyWebhookUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewHandleChargifyWebhookUsingPOSTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewHandleChargifyWebhookUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewHandleChargifyWebhookUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewHandleChargifyWebhookUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewHandleChargifyWebhookUsingPOSTOK creates a HandleChargifyWebhookUsingPOSTOK with default headers values
func NewHandleChargifyWebhookUsingPOSTOK() *HandleChargifyWebhookUsingPOSTOK {
	return &HandleChargifyWebhookUsingPOSTOK{}
}

/*HandleChargifyWebhookUsingPOSTOK handles this case with default header values.

OK
*/
type HandleChargifyWebhookUsingPOSTOK struct {
}

func (o *HandleChargifyWebhookUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /payment-plans][%d] handleChargifyWebhookUsingPOSTOK ", 200)
}

func (o *HandleChargifyWebhookUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewHandleChargifyWebhookUsingPOSTCreated creates a HandleChargifyWebhookUsingPOSTCreated with default headers values
func NewHandleChargifyWebhookUsingPOSTCreated() *HandleChargifyWebhookUsingPOSTCreated {
	return &HandleChargifyWebhookUsingPOSTCreated{}
}

/*HandleChargifyWebhookUsingPOSTCreated handles this case with default header values.

Created
*/
type HandleChargifyWebhookUsingPOSTCreated struct {
}

func (o *HandleChargifyWebhookUsingPOSTCreated) Error() string {
	return fmt.Sprintf("[POST /payment-plans][%d] handleChargifyWebhookUsingPOSTCreated ", 201)
}

func (o *HandleChargifyWebhookUsingPOSTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewHandleChargifyWebhookUsingPOSTUnauthorized creates a HandleChargifyWebhookUsingPOSTUnauthorized with default headers values
func NewHandleChargifyWebhookUsingPOSTUnauthorized() *HandleChargifyWebhookUsingPOSTUnauthorized {
	return &HandleChargifyWebhookUsingPOSTUnauthorized{}
}

/*HandleChargifyWebhookUsingPOSTUnauthorized handles this case with default header values.

Unauthorized
*/
type HandleChargifyWebhookUsingPOSTUnauthorized struct {
}

func (o *HandleChargifyWebhookUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /payment-plans][%d] handleChargifyWebhookUsingPOSTUnauthorized ", 401)
}

func (o *HandleChargifyWebhookUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewHandleChargifyWebhookUsingPOSTForbidden creates a HandleChargifyWebhookUsingPOSTForbidden with default headers values
func NewHandleChargifyWebhookUsingPOSTForbidden() *HandleChargifyWebhookUsingPOSTForbidden {
	return &HandleChargifyWebhookUsingPOSTForbidden{}
}

/*HandleChargifyWebhookUsingPOSTForbidden handles this case with default header values.

Forbidden
*/
type HandleChargifyWebhookUsingPOSTForbidden struct {
}

func (o *HandleChargifyWebhookUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /payment-plans][%d] handleChargifyWebhookUsingPOSTForbidden ", 403)
}

func (o *HandleChargifyWebhookUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewHandleChargifyWebhookUsingPOSTNotFound creates a HandleChargifyWebhookUsingPOSTNotFound with default headers values
func NewHandleChargifyWebhookUsingPOSTNotFound() *HandleChargifyWebhookUsingPOSTNotFound {
	return &HandleChargifyWebhookUsingPOSTNotFound{}
}

/*HandleChargifyWebhookUsingPOSTNotFound handles this case with default header values.

Not Found
*/
type HandleChargifyWebhookUsingPOSTNotFound struct {
}

func (o *HandleChargifyWebhookUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /payment-plans][%d] handleChargifyWebhookUsingPOSTNotFound ", 404)
}

func (o *HandleChargifyWebhookUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}