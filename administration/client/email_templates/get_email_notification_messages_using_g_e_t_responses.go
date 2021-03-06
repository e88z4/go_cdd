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

// GetEmailNotificationMessagesUsingGETReader is a Reader for the GetEmailNotificationMessagesUsingGET structure.
type GetEmailNotificationMessagesUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEmailNotificationMessagesUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEmailNotificationMessagesUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetEmailNotificationMessagesUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetEmailNotificationMessagesUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetEmailNotificationMessagesUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetEmailNotificationMessagesUsingGETOK creates a GetEmailNotificationMessagesUsingGETOK with default headers values
func NewGetEmailNotificationMessagesUsingGETOK() *GetEmailNotificationMessagesUsingGETOK {
	return &GetEmailNotificationMessagesUsingGETOK{}
}

/*GetEmailNotificationMessagesUsingGETOK handles this case with default header values.

OK
*/
type GetEmailNotificationMessagesUsingGETOK struct {
	Payload *models.ListHolderDtoEmailNotificationDto
}

func (o *GetEmailNotificationMessagesUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /email-notification-messages/{languageTag}][%d] getEmailNotificationMessagesUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetEmailNotificationMessagesUsingGETOK) GetPayload() *models.ListHolderDtoEmailNotificationDto {
	return o.Payload
}

func (o *GetEmailNotificationMessagesUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListHolderDtoEmailNotificationDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEmailNotificationMessagesUsingGETUnauthorized creates a GetEmailNotificationMessagesUsingGETUnauthorized with default headers values
func NewGetEmailNotificationMessagesUsingGETUnauthorized() *GetEmailNotificationMessagesUsingGETUnauthorized {
	return &GetEmailNotificationMessagesUsingGETUnauthorized{}
}

/*GetEmailNotificationMessagesUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type GetEmailNotificationMessagesUsingGETUnauthorized struct {
}

func (o *GetEmailNotificationMessagesUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /email-notification-messages/{languageTag}][%d] getEmailNotificationMessagesUsingGETUnauthorized ", 401)
}

func (o *GetEmailNotificationMessagesUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEmailNotificationMessagesUsingGETForbidden creates a GetEmailNotificationMessagesUsingGETForbidden with default headers values
func NewGetEmailNotificationMessagesUsingGETForbidden() *GetEmailNotificationMessagesUsingGETForbidden {
	return &GetEmailNotificationMessagesUsingGETForbidden{}
}

/*GetEmailNotificationMessagesUsingGETForbidden handles this case with default header values.

Forbidden
*/
type GetEmailNotificationMessagesUsingGETForbidden struct {
}

func (o *GetEmailNotificationMessagesUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /email-notification-messages/{languageTag}][%d] getEmailNotificationMessagesUsingGETForbidden ", 403)
}

func (o *GetEmailNotificationMessagesUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEmailNotificationMessagesUsingGETNotFound creates a GetEmailNotificationMessagesUsingGETNotFound with default headers values
func NewGetEmailNotificationMessagesUsingGETNotFound() *GetEmailNotificationMessagesUsingGETNotFound {
	return &GetEmailNotificationMessagesUsingGETNotFound{}
}

/*GetEmailNotificationMessagesUsingGETNotFound handles this case with default header values.

Not Found
*/
type GetEmailNotificationMessagesUsingGETNotFound struct {
}

func (o *GetEmailNotificationMessagesUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /email-notification-messages/{languageTag}][%d] getEmailNotificationMessagesUsingGETNotFound ", 404)
}

func (o *GetEmailNotificationMessagesUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
