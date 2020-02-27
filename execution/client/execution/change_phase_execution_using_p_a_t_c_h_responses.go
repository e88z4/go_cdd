// Code generated by go-swagger; DO NOT EDIT.

package execution

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/e88z4/go_cdd/execution/models"
)

// ChangePhaseExecutionUsingPATCHReader is a Reader for the ChangePhaseExecutionUsingPATCH structure.
type ChangePhaseExecutionUsingPATCHReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangePhaseExecutionUsingPATCHReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewChangePhaseExecutionUsingPATCHOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewChangePhaseExecutionUsingPATCHNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewChangePhaseExecutionUsingPATCHUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewChangePhaseExecutionUsingPATCHForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewChangePhaseExecutionUsingPATCHOK creates a ChangePhaseExecutionUsingPATCHOK with default headers values
func NewChangePhaseExecutionUsingPATCHOK() *ChangePhaseExecutionUsingPATCHOK {
	return &ChangePhaseExecutionUsingPATCHOK{}
}

/*ChangePhaseExecutionUsingPATCHOK handles this case with default header values.

OK
*/
type ChangePhaseExecutionUsingPATCHOK struct {
	Payload *models.PhaseExecutionDto
}

func (o *ChangePhaseExecutionUsingPATCHOK) Error() string {
	return fmt.Sprintf("[PATCH /releases-execution/{releaseId}/phases-execution/{phaseId}][%d] changePhaseExecutionUsingPATCHOK  %+v", 200, o.Payload)
}

func (o *ChangePhaseExecutionUsingPATCHOK) GetPayload() *models.PhaseExecutionDto {
	return o.Payload
}

func (o *ChangePhaseExecutionUsingPATCHOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PhaseExecutionDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChangePhaseExecutionUsingPATCHNoContent creates a ChangePhaseExecutionUsingPATCHNoContent with default headers values
func NewChangePhaseExecutionUsingPATCHNoContent() *ChangePhaseExecutionUsingPATCHNoContent {
	return &ChangePhaseExecutionUsingPATCHNoContent{}
}

/*ChangePhaseExecutionUsingPATCHNoContent handles this case with default header values.

No Content
*/
type ChangePhaseExecutionUsingPATCHNoContent struct {
}

func (o *ChangePhaseExecutionUsingPATCHNoContent) Error() string {
	return fmt.Sprintf("[PATCH /releases-execution/{releaseId}/phases-execution/{phaseId}][%d] changePhaseExecutionUsingPATCHNoContent ", 204)
}

func (o *ChangePhaseExecutionUsingPATCHNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChangePhaseExecutionUsingPATCHUnauthorized creates a ChangePhaseExecutionUsingPATCHUnauthorized with default headers values
func NewChangePhaseExecutionUsingPATCHUnauthorized() *ChangePhaseExecutionUsingPATCHUnauthorized {
	return &ChangePhaseExecutionUsingPATCHUnauthorized{}
}

/*ChangePhaseExecutionUsingPATCHUnauthorized handles this case with default header values.

Unauthorized
*/
type ChangePhaseExecutionUsingPATCHUnauthorized struct {
}

func (o *ChangePhaseExecutionUsingPATCHUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /releases-execution/{releaseId}/phases-execution/{phaseId}][%d] changePhaseExecutionUsingPATCHUnauthorized ", 401)
}

func (o *ChangePhaseExecutionUsingPATCHUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChangePhaseExecutionUsingPATCHForbidden creates a ChangePhaseExecutionUsingPATCHForbidden with default headers values
func NewChangePhaseExecutionUsingPATCHForbidden() *ChangePhaseExecutionUsingPATCHForbidden {
	return &ChangePhaseExecutionUsingPATCHForbidden{}
}

/*ChangePhaseExecutionUsingPATCHForbidden handles this case with default header values.

Forbidden
*/
type ChangePhaseExecutionUsingPATCHForbidden struct {
}

func (o *ChangePhaseExecutionUsingPATCHForbidden) Error() string {
	return fmt.Sprintf("[PATCH /releases-execution/{releaseId}/phases-execution/{phaseId}][%d] changePhaseExecutionUsingPATCHForbidden ", 403)
}

func (o *ChangePhaseExecutionUsingPATCHForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
