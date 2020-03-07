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

// GetEmailNotificationServerUsingGETReader is a Reader for the GetEmailNotificationServerUsingGET structure.
type GetEmailNotificationServerUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEmailNotificationServerUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEmailNotificationServerUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetEmailNotificationServerUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetEmailNotificationServerUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetEmailNotificationServerUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetEmailNotificationServerUsingGETOK creates a GetEmailNotificationServerUsingGETOK with default headers values
func NewGetEmailNotificationServerUsingGETOK() *GetEmailNotificationServerUsingGETOK {
	return &GetEmailNotificationServerUsingGETOK{}
}

/*GetEmailNotificationServerUsingGETOK handles this case with default header values.

OK
*/
type GetEmailNotificationServerUsingGETOK struct {
	Payload *models.EmailNotificationServerDto
}

func (o *GetEmailNotificationServerUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /email-notification-servers/{emailNotificationServerId}][%d] getEmailNotificationServerUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetEmailNotificationServerUsingGETOK) GetPayload() *models.EmailNotificationServerDto {
	return o.Payload
}

func (o *GetEmailNotificationServerUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EmailNotificationServerDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEmailNotificationServerUsingGETUnauthorized creates a GetEmailNotificationServerUsingGETUnauthorized with default headers values
func NewGetEmailNotificationServerUsingGETUnauthorized() *GetEmailNotificationServerUsingGETUnauthorized {
	return &GetEmailNotificationServerUsingGETUnauthorized{}
}

/*GetEmailNotificationServerUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type GetEmailNotificationServerUsingGETUnauthorized struct {
}

func (o *GetEmailNotificationServerUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /email-notification-servers/{emailNotificationServerId}][%d] getEmailNotificationServerUsingGETUnauthorized ", 401)
}

func (o *GetEmailNotificationServerUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEmailNotificationServerUsingGETForbidden creates a GetEmailNotificationServerUsingGETForbidden with default headers values
func NewGetEmailNotificationServerUsingGETForbidden() *GetEmailNotificationServerUsingGETForbidden {
	return &GetEmailNotificationServerUsingGETForbidden{}
}

/*GetEmailNotificationServerUsingGETForbidden handles this case with default header values.

Forbidden
*/
type GetEmailNotificationServerUsingGETForbidden struct {
}

func (o *GetEmailNotificationServerUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /email-notification-servers/{emailNotificationServerId}][%d] getEmailNotificationServerUsingGETForbidden ", 403)
}

func (o *GetEmailNotificationServerUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEmailNotificationServerUsingGETNotFound creates a GetEmailNotificationServerUsingGETNotFound with default headers values
func NewGetEmailNotificationServerUsingGETNotFound() *GetEmailNotificationServerUsingGETNotFound {
	return &GetEmailNotificationServerUsingGETNotFound{}
}

/*GetEmailNotificationServerUsingGETNotFound handles this case with default header values.

Not Found
*/
type GetEmailNotificationServerUsingGETNotFound struct {
}

func (o *GetEmailNotificationServerUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /email-notification-servers/{emailNotificationServerId}][%d] getEmailNotificationServerUsingGETNotFound ", 404)
}

func (o *GetEmailNotificationServerUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
