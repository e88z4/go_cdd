// Code generated by go-swagger; DO NOT EDIT.

package email_notification_server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/e88z4/go_cdd/administration/models"
)

// PatchEmailNotificationServerUsingPATCHReader is a Reader for the PatchEmailNotificationServerUsingPATCH structure.
type PatchEmailNotificationServerUsingPATCHReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchEmailNotificationServerUsingPATCHReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchEmailNotificationServerUsingPATCHOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewPatchEmailNotificationServerUsingPATCHNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchEmailNotificationServerUsingPATCHUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchEmailNotificationServerUsingPATCHForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchEmailNotificationServerUsingPATCHOK creates a PatchEmailNotificationServerUsingPATCHOK with default headers values
func NewPatchEmailNotificationServerUsingPATCHOK() *PatchEmailNotificationServerUsingPATCHOK {
	return &PatchEmailNotificationServerUsingPATCHOK{}
}

/*PatchEmailNotificationServerUsingPATCHOK handles this case with default header values.

OK
*/
type PatchEmailNotificationServerUsingPATCHOK struct {
	Payload *models.EmailNotificationServerDto
}

func (o *PatchEmailNotificationServerUsingPATCHOK) Error() string {
	return fmt.Sprintf("[PATCH /email-notification-servers/{emailNotificationServerId}][%d] patchEmailNotificationServerUsingPATCHOK  %+v", 200, o.Payload)
}

func (o *PatchEmailNotificationServerUsingPATCHOK) GetPayload() *models.EmailNotificationServerDto {
	return o.Payload
}

func (o *PatchEmailNotificationServerUsingPATCHOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EmailNotificationServerDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchEmailNotificationServerUsingPATCHNoContent creates a PatchEmailNotificationServerUsingPATCHNoContent with default headers values
func NewPatchEmailNotificationServerUsingPATCHNoContent() *PatchEmailNotificationServerUsingPATCHNoContent {
	return &PatchEmailNotificationServerUsingPATCHNoContent{}
}

/*PatchEmailNotificationServerUsingPATCHNoContent handles this case with default header values.

No Content
*/
type PatchEmailNotificationServerUsingPATCHNoContent struct {
}

func (o *PatchEmailNotificationServerUsingPATCHNoContent) Error() string {
	return fmt.Sprintf("[PATCH /email-notification-servers/{emailNotificationServerId}][%d] patchEmailNotificationServerUsingPATCHNoContent ", 204)
}

func (o *PatchEmailNotificationServerUsingPATCHNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchEmailNotificationServerUsingPATCHUnauthorized creates a PatchEmailNotificationServerUsingPATCHUnauthorized with default headers values
func NewPatchEmailNotificationServerUsingPATCHUnauthorized() *PatchEmailNotificationServerUsingPATCHUnauthorized {
	return &PatchEmailNotificationServerUsingPATCHUnauthorized{}
}

/*PatchEmailNotificationServerUsingPATCHUnauthorized handles this case with default header values.

Unauthorized
*/
type PatchEmailNotificationServerUsingPATCHUnauthorized struct {
}

func (o *PatchEmailNotificationServerUsingPATCHUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /email-notification-servers/{emailNotificationServerId}][%d] patchEmailNotificationServerUsingPATCHUnauthorized ", 401)
}

func (o *PatchEmailNotificationServerUsingPATCHUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchEmailNotificationServerUsingPATCHForbidden creates a PatchEmailNotificationServerUsingPATCHForbidden with default headers values
func NewPatchEmailNotificationServerUsingPATCHForbidden() *PatchEmailNotificationServerUsingPATCHForbidden {
	return &PatchEmailNotificationServerUsingPATCHForbidden{}
}

/*PatchEmailNotificationServerUsingPATCHForbidden handles this case with default header values.

Forbidden
*/
type PatchEmailNotificationServerUsingPATCHForbidden struct {
}

func (o *PatchEmailNotificationServerUsingPATCHForbidden) Error() string {
	return fmt.Sprintf("[PATCH /email-notification-servers/{emailNotificationServerId}][%d] patchEmailNotificationServerUsingPATCHForbidden ", 403)
}

func (o *PatchEmailNotificationServerUsingPATCHForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
