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

// PatchEmailNotificationTemplateUsingPATCHReader is a Reader for the PatchEmailNotificationTemplateUsingPATCH structure.
type PatchEmailNotificationTemplateUsingPATCHReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchEmailNotificationTemplateUsingPATCHReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchEmailNotificationTemplateUsingPATCHOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewPatchEmailNotificationTemplateUsingPATCHNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchEmailNotificationTemplateUsingPATCHUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchEmailNotificationTemplateUsingPATCHForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchEmailNotificationTemplateUsingPATCHOK creates a PatchEmailNotificationTemplateUsingPATCHOK with default headers values
func NewPatchEmailNotificationTemplateUsingPATCHOK() *PatchEmailNotificationTemplateUsingPATCHOK {
	return &PatchEmailNotificationTemplateUsingPATCHOK{}
}

/*PatchEmailNotificationTemplateUsingPATCHOK handles this case with default header values.

OK
*/
type PatchEmailNotificationTemplateUsingPATCHOK struct {
	Payload *models.EmailNotificationTemplateDto
}

func (o *PatchEmailNotificationTemplateUsingPATCHOK) Error() string {
	return fmt.Sprintf("[PATCH /email-notification-templates/{emailNotificationTemplateName}/{languageTag}][%d] patchEmailNotificationTemplateUsingPATCHOK  %+v", 200, o.Payload)
}

func (o *PatchEmailNotificationTemplateUsingPATCHOK) GetPayload() *models.EmailNotificationTemplateDto {
	return o.Payload
}

func (o *PatchEmailNotificationTemplateUsingPATCHOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EmailNotificationTemplateDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchEmailNotificationTemplateUsingPATCHNoContent creates a PatchEmailNotificationTemplateUsingPATCHNoContent with default headers values
func NewPatchEmailNotificationTemplateUsingPATCHNoContent() *PatchEmailNotificationTemplateUsingPATCHNoContent {
	return &PatchEmailNotificationTemplateUsingPATCHNoContent{}
}

/*PatchEmailNotificationTemplateUsingPATCHNoContent handles this case with default header values.

No Content
*/
type PatchEmailNotificationTemplateUsingPATCHNoContent struct {
}

func (o *PatchEmailNotificationTemplateUsingPATCHNoContent) Error() string {
	return fmt.Sprintf("[PATCH /email-notification-templates/{emailNotificationTemplateName}/{languageTag}][%d] patchEmailNotificationTemplateUsingPATCHNoContent ", 204)
}

func (o *PatchEmailNotificationTemplateUsingPATCHNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchEmailNotificationTemplateUsingPATCHUnauthorized creates a PatchEmailNotificationTemplateUsingPATCHUnauthorized with default headers values
func NewPatchEmailNotificationTemplateUsingPATCHUnauthorized() *PatchEmailNotificationTemplateUsingPATCHUnauthorized {
	return &PatchEmailNotificationTemplateUsingPATCHUnauthorized{}
}

/*PatchEmailNotificationTemplateUsingPATCHUnauthorized handles this case with default header values.

Unauthorized
*/
type PatchEmailNotificationTemplateUsingPATCHUnauthorized struct {
}

func (o *PatchEmailNotificationTemplateUsingPATCHUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /email-notification-templates/{emailNotificationTemplateName}/{languageTag}][%d] patchEmailNotificationTemplateUsingPATCHUnauthorized ", 401)
}

func (o *PatchEmailNotificationTemplateUsingPATCHUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchEmailNotificationTemplateUsingPATCHForbidden creates a PatchEmailNotificationTemplateUsingPATCHForbidden with default headers values
func NewPatchEmailNotificationTemplateUsingPATCHForbidden() *PatchEmailNotificationTemplateUsingPATCHForbidden {
	return &PatchEmailNotificationTemplateUsingPATCHForbidden{}
}

/*PatchEmailNotificationTemplateUsingPATCHForbidden handles this case with default header values.

Forbidden
*/
type PatchEmailNotificationTemplateUsingPATCHForbidden struct {
}

func (o *PatchEmailNotificationTemplateUsingPATCHForbidden) Error() string {
	return fmt.Sprintf("[PATCH /email-notification-templates/{emailNotificationTemplateName}/{languageTag}][%d] patchEmailNotificationTemplateUsingPATCHForbidden ", 403)
}

func (o *PatchEmailNotificationTemplateUsingPATCHForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
