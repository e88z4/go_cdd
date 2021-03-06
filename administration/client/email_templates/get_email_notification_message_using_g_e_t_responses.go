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

// GetEmailNotificationMessageUsingGETReader is a Reader for the GetEmailNotificationMessageUsingGET structure.
type GetEmailNotificationMessageUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEmailNotificationMessageUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEmailNotificationMessageUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetEmailNotificationMessageUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetEmailNotificationMessageUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetEmailNotificationMessageUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetEmailNotificationMessageUsingGETOK creates a GetEmailNotificationMessageUsingGETOK with default headers values
func NewGetEmailNotificationMessageUsingGETOK() *GetEmailNotificationMessageUsingGETOK {
	return &GetEmailNotificationMessageUsingGETOK{}
}

/*GetEmailNotificationMessageUsingGETOK handles this case with default header values.

OK
*/
type GetEmailNotificationMessageUsingGETOK struct {
	Payload *models.EmailNotificationMessageDto
}

func (o *GetEmailNotificationMessageUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /email-notification-messages/{emailNotificationMessageName}/{languageTag}][%d] getEmailNotificationMessageUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetEmailNotificationMessageUsingGETOK) GetPayload() *models.EmailNotificationMessageDto {
	return o.Payload
}

func (o *GetEmailNotificationMessageUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EmailNotificationMessageDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEmailNotificationMessageUsingGETUnauthorized creates a GetEmailNotificationMessageUsingGETUnauthorized with default headers values
func NewGetEmailNotificationMessageUsingGETUnauthorized() *GetEmailNotificationMessageUsingGETUnauthorized {
	return &GetEmailNotificationMessageUsingGETUnauthorized{}
}

/*GetEmailNotificationMessageUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type GetEmailNotificationMessageUsingGETUnauthorized struct {
}

func (o *GetEmailNotificationMessageUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /email-notification-messages/{emailNotificationMessageName}/{languageTag}][%d] getEmailNotificationMessageUsingGETUnauthorized ", 401)
}

func (o *GetEmailNotificationMessageUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEmailNotificationMessageUsingGETForbidden creates a GetEmailNotificationMessageUsingGETForbidden with default headers values
func NewGetEmailNotificationMessageUsingGETForbidden() *GetEmailNotificationMessageUsingGETForbidden {
	return &GetEmailNotificationMessageUsingGETForbidden{}
}

/*GetEmailNotificationMessageUsingGETForbidden handles this case with default header values.

Forbidden
*/
type GetEmailNotificationMessageUsingGETForbidden struct {
}

func (o *GetEmailNotificationMessageUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /email-notification-messages/{emailNotificationMessageName}/{languageTag}][%d] getEmailNotificationMessageUsingGETForbidden ", 403)
}

func (o *GetEmailNotificationMessageUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEmailNotificationMessageUsingGETNotFound creates a GetEmailNotificationMessageUsingGETNotFound with default headers values
func NewGetEmailNotificationMessageUsingGETNotFound() *GetEmailNotificationMessageUsingGETNotFound {
	return &GetEmailNotificationMessageUsingGETNotFound{}
}

/*GetEmailNotificationMessageUsingGETNotFound handles this case with default header values.

Not Found
*/
type GetEmailNotificationMessageUsingGETNotFound struct {
}

func (o *GetEmailNotificationMessageUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /email-notification-messages/{emailNotificationMessageName}/{languageTag}][%d] getEmailNotificationMessageUsingGETNotFound ", 404)
}

func (o *GetEmailNotificationMessageUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
