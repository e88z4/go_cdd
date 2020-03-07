// Code generated by go-swagger; DO NOT EDIT.

package system_preferences

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/e88z4/go_cdd/administration/models"
)

// GetSystemPreferencesDtoUsingGETReader is a Reader for the GetSystemPreferencesDtoUsingGET structure.
type GetSystemPreferencesDtoUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSystemPreferencesDtoUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSystemPreferencesDtoUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetSystemPreferencesDtoUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSystemPreferencesDtoUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetSystemPreferencesDtoUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSystemPreferencesDtoUsingGETOK creates a GetSystemPreferencesDtoUsingGETOK with default headers values
func NewGetSystemPreferencesDtoUsingGETOK() *GetSystemPreferencesDtoUsingGETOK {
	return &GetSystemPreferencesDtoUsingGETOK{}
}

/*GetSystemPreferencesDtoUsingGETOK handles this case with default header values.

OK
*/
type GetSystemPreferencesDtoUsingGETOK struct {
	Payload *models.SystemPreferencesDto
}

func (o *GetSystemPreferencesDtoUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /product/settings/system-preferences][%d] getSystemPreferencesDtoUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetSystemPreferencesDtoUsingGETOK) GetPayload() *models.SystemPreferencesDto {
	return o.Payload
}

func (o *GetSystemPreferencesDtoUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SystemPreferencesDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSystemPreferencesDtoUsingGETUnauthorized creates a GetSystemPreferencesDtoUsingGETUnauthorized with default headers values
func NewGetSystemPreferencesDtoUsingGETUnauthorized() *GetSystemPreferencesDtoUsingGETUnauthorized {
	return &GetSystemPreferencesDtoUsingGETUnauthorized{}
}

/*GetSystemPreferencesDtoUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type GetSystemPreferencesDtoUsingGETUnauthorized struct {
}

func (o *GetSystemPreferencesDtoUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /product/settings/system-preferences][%d] getSystemPreferencesDtoUsingGETUnauthorized ", 401)
}

func (o *GetSystemPreferencesDtoUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSystemPreferencesDtoUsingGETForbidden creates a GetSystemPreferencesDtoUsingGETForbidden with default headers values
func NewGetSystemPreferencesDtoUsingGETForbidden() *GetSystemPreferencesDtoUsingGETForbidden {
	return &GetSystemPreferencesDtoUsingGETForbidden{}
}

/*GetSystemPreferencesDtoUsingGETForbidden handles this case with default header values.

Forbidden
*/
type GetSystemPreferencesDtoUsingGETForbidden struct {
}

func (o *GetSystemPreferencesDtoUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /product/settings/system-preferences][%d] getSystemPreferencesDtoUsingGETForbidden ", 403)
}

func (o *GetSystemPreferencesDtoUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSystemPreferencesDtoUsingGETNotFound creates a GetSystemPreferencesDtoUsingGETNotFound with default headers values
func NewGetSystemPreferencesDtoUsingGETNotFound() *GetSystemPreferencesDtoUsingGETNotFound {
	return &GetSystemPreferencesDtoUsingGETNotFound{}
}

/*GetSystemPreferencesDtoUsingGETNotFound handles this case with default header values.

Not Found
*/
type GetSystemPreferencesDtoUsingGETNotFound struct {
}

func (o *GetSystemPreferencesDtoUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /product/settings/system-preferences][%d] getSystemPreferencesDtoUsingGETNotFound ", 404)
}

func (o *GetSystemPreferencesDtoUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
