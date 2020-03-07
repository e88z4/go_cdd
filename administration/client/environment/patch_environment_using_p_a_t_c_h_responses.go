// Code generated by go-swagger; DO NOT EDIT.

package environment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/e88z4/go_cdd/administration/models"
)

// PatchEnvironmentUsingPATCHReader is a Reader for the PatchEnvironmentUsingPATCH structure.
type PatchEnvironmentUsingPATCHReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchEnvironmentUsingPATCHReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchEnvironmentUsingPATCHOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewPatchEnvironmentUsingPATCHNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchEnvironmentUsingPATCHUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchEnvironmentUsingPATCHForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchEnvironmentUsingPATCHOK creates a PatchEnvironmentUsingPATCHOK with default headers values
func NewPatchEnvironmentUsingPATCHOK() *PatchEnvironmentUsingPATCHOK {
	return &PatchEnvironmentUsingPATCHOK{}
}

/*PatchEnvironmentUsingPATCHOK handles this case with default header values.

OK
*/
type PatchEnvironmentUsingPATCHOK struct {
	Payload *models.EnvironmentDto
}

func (o *PatchEnvironmentUsingPATCHOK) Error() string {
	return fmt.Sprintf("[PATCH /environments/{environmentId}][%d] patchEnvironmentUsingPATCHOK  %+v", 200, o.Payload)
}

func (o *PatchEnvironmentUsingPATCHOK) GetPayload() *models.EnvironmentDto {
	return o.Payload
}

func (o *PatchEnvironmentUsingPATCHOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EnvironmentDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchEnvironmentUsingPATCHNoContent creates a PatchEnvironmentUsingPATCHNoContent with default headers values
func NewPatchEnvironmentUsingPATCHNoContent() *PatchEnvironmentUsingPATCHNoContent {
	return &PatchEnvironmentUsingPATCHNoContent{}
}

/*PatchEnvironmentUsingPATCHNoContent handles this case with default header values.

No Content
*/
type PatchEnvironmentUsingPATCHNoContent struct {
}

func (o *PatchEnvironmentUsingPATCHNoContent) Error() string {
	return fmt.Sprintf("[PATCH /environments/{environmentId}][%d] patchEnvironmentUsingPATCHNoContent ", 204)
}

func (o *PatchEnvironmentUsingPATCHNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchEnvironmentUsingPATCHUnauthorized creates a PatchEnvironmentUsingPATCHUnauthorized with default headers values
func NewPatchEnvironmentUsingPATCHUnauthorized() *PatchEnvironmentUsingPATCHUnauthorized {
	return &PatchEnvironmentUsingPATCHUnauthorized{}
}

/*PatchEnvironmentUsingPATCHUnauthorized handles this case with default header values.

Unauthorized
*/
type PatchEnvironmentUsingPATCHUnauthorized struct {
}

func (o *PatchEnvironmentUsingPATCHUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /environments/{environmentId}][%d] patchEnvironmentUsingPATCHUnauthorized ", 401)
}

func (o *PatchEnvironmentUsingPATCHUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchEnvironmentUsingPATCHForbidden creates a PatchEnvironmentUsingPATCHForbidden with default headers values
func NewPatchEnvironmentUsingPATCHForbidden() *PatchEnvironmentUsingPATCHForbidden {
	return &PatchEnvironmentUsingPATCHForbidden{}
}

/*PatchEnvironmentUsingPATCHForbidden handles this case with default header values.

Forbidden
*/
type PatchEnvironmentUsingPATCHForbidden struct {
}

func (o *PatchEnvironmentUsingPATCHForbidden) Error() string {
	return fmt.Sprintf("[PATCH /environments/{environmentId}][%d] patchEnvironmentUsingPATCHForbidden ", 403)
}

func (o *PatchEnvironmentUsingPATCHForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}