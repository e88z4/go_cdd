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

// EnableDisableFreezePeriodUsingPATCHReader is a Reader for the EnableDisableFreezePeriodUsingPATCH structure.
type EnableDisableFreezePeriodUsingPATCHReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EnableDisableFreezePeriodUsingPATCHReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEnableDisableFreezePeriodUsingPATCHOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewEnableDisableFreezePeriodUsingPATCHNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewEnableDisableFreezePeriodUsingPATCHUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEnableDisableFreezePeriodUsingPATCHForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewEnableDisableFreezePeriodUsingPATCHOK creates a EnableDisableFreezePeriodUsingPATCHOK with default headers values
func NewEnableDisableFreezePeriodUsingPATCHOK() *EnableDisableFreezePeriodUsingPATCHOK {
	return &EnableDisableFreezePeriodUsingPATCHOK{}
}

/*EnableDisableFreezePeriodUsingPATCHOK handles this case with default header values.

OK
*/
type EnableDisableFreezePeriodUsingPATCHOK struct {
	Payload *models.FreezePeriodDto
}

func (o *EnableDisableFreezePeriodUsingPATCHOK) Error() string {
	return fmt.Sprintf("[PATCH /freeze-periods/{freezePeriodId}][%d] enableDisableFreezePeriodUsingPATCHOK  %+v", 200, o.Payload)
}

func (o *EnableDisableFreezePeriodUsingPATCHOK) GetPayload() *models.FreezePeriodDto {
	return o.Payload
}

func (o *EnableDisableFreezePeriodUsingPATCHOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FreezePeriodDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnableDisableFreezePeriodUsingPATCHNoContent creates a EnableDisableFreezePeriodUsingPATCHNoContent with default headers values
func NewEnableDisableFreezePeriodUsingPATCHNoContent() *EnableDisableFreezePeriodUsingPATCHNoContent {
	return &EnableDisableFreezePeriodUsingPATCHNoContent{}
}

/*EnableDisableFreezePeriodUsingPATCHNoContent handles this case with default header values.

No Content
*/
type EnableDisableFreezePeriodUsingPATCHNoContent struct {
}

func (o *EnableDisableFreezePeriodUsingPATCHNoContent) Error() string {
	return fmt.Sprintf("[PATCH /freeze-periods/{freezePeriodId}][%d] enableDisableFreezePeriodUsingPATCHNoContent ", 204)
}

func (o *EnableDisableFreezePeriodUsingPATCHNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEnableDisableFreezePeriodUsingPATCHUnauthorized creates a EnableDisableFreezePeriodUsingPATCHUnauthorized with default headers values
func NewEnableDisableFreezePeriodUsingPATCHUnauthorized() *EnableDisableFreezePeriodUsingPATCHUnauthorized {
	return &EnableDisableFreezePeriodUsingPATCHUnauthorized{}
}

/*EnableDisableFreezePeriodUsingPATCHUnauthorized handles this case with default header values.

Unauthorized
*/
type EnableDisableFreezePeriodUsingPATCHUnauthorized struct {
}

func (o *EnableDisableFreezePeriodUsingPATCHUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /freeze-periods/{freezePeriodId}][%d] enableDisableFreezePeriodUsingPATCHUnauthorized ", 401)
}

func (o *EnableDisableFreezePeriodUsingPATCHUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEnableDisableFreezePeriodUsingPATCHForbidden creates a EnableDisableFreezePeriodUsingPATCHForbidden with default headers values
func NewEnableDisableFreezePeriodUsingPATCHForbidden() *EnableDisableFreezePeriodUsingPATCHForbidden {
	return &EnableDisableFreezePeriodUsingPATCHForbidden{}
}

/*EnableDisableFreezePeriodUsingPATCHForbidden handles this case with default header values.

Forbidden
*/
type EnableDisableFreezePeriodUsingPATCHForbidden struct {
}

func (o *EnableDisableFreezePeriodUsingPATCHForbidden) Error() string {
	return fmt.Sprintf("[PATCH /freeze-periods/{freezePeriodId}][%d] enableDisableFreezePeriodUsingPATCHForbidden ", 403)
}

func (o *EnableDisableFreezePeriodUsingPATCHForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
