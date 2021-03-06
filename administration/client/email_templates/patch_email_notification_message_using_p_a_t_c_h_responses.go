// Code generated by go-swagger; DO NOT EDIT.

package email_templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/e88z4/go_cdd/administration/models"
)

// PatchEmailNotificationMessageUsingPATCHReader is a Reader for the PatchEmailNotificationMessageUsingPATCH structure.
type PatchEmailNotificationMessageUsingPATCHReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchEmailNotificationMessageUsingPATCHReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchEmailNotificationMessageUsingPATCHOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewPatchEmailNotificationMessageUsingPATCHNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchEmailNotificationMessageUsingPATCHUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchEmailNotificationMessageUsingPATCHForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchEmailNotificationMessageUsingPATCHOK creates a PatchEmailNotificationMessageUsingPATCHOK with default headers values
func NewPatchEmailNotificationMessageUsingPATCHOK() *PatchEmailNotificationMessageUsingPATCHOK {
	return &PatchEmailNotificationMessageUsingPATCHOK{}
}

/*PatchEmailNotificationMessageUsingPATCHOK handles this case with default header values.

OK
*/
type PatchEmailNotificationMessageUsingPATCHOK struct {
	Payload *models.EmailNotificationMessageDto
}

func (o *PatchEmailNotificationMessageUsingPATCHOK) Error() string {
	return fmt.Sprintf("[PATCH /email-notification-messages/{emailNotificationMessageName}/{languageTag}][%d] patchEmailNotificationMessageUsingPATCHOK  %+v", 200, o.Payload)
}

func (o *PatchEmailNotificationMessageUsingPATCHOK) GetPayload() *models.EmailNotificationMessageDto {
	return o.Payload
}

func (o *PatchEmailNotificationMessageUsingPATCHOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EmailNotificationMessageDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchEmailNotificationMessageUsingPATCHNoContent creates a PatchEmailNotificationMessageUsingPATCHNoContent with default headers values
func NewPatchEmailNotificationMessageUsingPATCHNoContent() *PatchEmailNotificationMessageUsingPATCHNoContent {
	return &PatchEmailNotificationMessageUsingPATCHNoContent{}
}

/*PatchEmailNotificationMessageUsingPATCHNoContent handles this case with default header values.

No Content
*/
type PatchEmailNotificationMessageUsingPATCHNoContent struct {
}

func (o *PatchEmailNotificationMessageUsingPATCHNoContent) Error() string {
	return fmt.Sprintf("[PATCH /email-notification-messages/{emailNotificationMessageName}/{languageTag}][%d] patchEmailNotificationMessageUsingPATCHNoContent ", 204)
}

func (o *PatchEmailNotificationMessageUsingPATCHNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchEmailNotificationMessageUsingPATCHUnauthorized creates a PatchEmailNotificationMessageUsingPATCHUnauthorized with default headers values
func NewPatchEmailNotificationMessageUsingPATCHUnauthorized() *PatchEmailNotificationMessageUsingPATCHUnauthorized {
	return &PatchEmailNotificationMessageUsingPATCHUnauthorized{}
}

/*PatchEmailNotificationMessageUsingPATCHUnauthorized handles this case with default header values.

Unauthorized
*/
type PatchEmailNotificationMessageUsingPATCHUnauthorized struct {
}

func (o *PatchEmailNotificationMessageUsingPATCHUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /email-notification-messages/{emailNotificationMessageName}/{languageTag}][%d] patchEmailNotificationMessageUsingPATCHUnauthorized ", 401)
}

func (o *PatchEmailNotificationMessageUsingPATCHUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchEmailNotificationMessageUsingPATCHForbidden creates a PatchEmailNotificationMessageUsingPATCHForbidden with default headers values
func NewPatchEmailNotificationMessageUsingPATCHForbidden() *PatchEmailNotificationMessageUsingPATCHForbidden {
	return &PatchEmailNotificationMessageUsingPATCHForbidden{}
}

/*PatchEmailNotificationMessageUsingPATCHForbidden handles this case with default header values.

Forbidden
*/
type PatchEmailNotificationMessageUsingPATCHForbidden struct {
}

func (o *PatchEmailNotificationMessageUsingPATCHForbidden) Error() string {
	return fmt.Sprintf("[PATCH /email-notification-messages/{emailNotificationMessageName}/{languageTag}][%d] patchEmailNotificationMessageUsingPATCHForbidden ", 403)
}

func (o *PatchEmailNotificationMessageUsingPATCHForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
