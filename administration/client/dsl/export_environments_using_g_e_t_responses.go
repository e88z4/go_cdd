// Code generated by go-swagger; DO NOT EDIT.

package dsl

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/e88z4/go_cdd/administration/models"
)

// ExportEnvironmentsUsingGETReader is a Reader for the ExportEnvironmentsUsingGET structure.
type ExportEnvironmentsUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExportEnvironmentsUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExportEnvironmentsUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewExportEnvironmentsUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewExportEnvironmentsUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewExportEnvironmentsUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewExportEnvironmentsUsingGETOK creates a ExportEnvironmentsUsingGETOK with default headers values
func NewExportEnvironmentsUsingGETOK() *ExportEnvironmentsUsingGETOK {
	return &ExportEnvironmentsUsingGETOK{}
}

/*ExportEnvironmentsUsingGETOK handles this case with default header values.

OK
*/
type ExportEnvironmentsUsingGETOK struct {
	Payload *models.DslTemplateDto
}

func (o *ExportEnvironmentsUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /environments/dsl-manifest][%d] exportEnvironmentsUsingGETOK  %+v", 200, o.Payload)
}

func (o *ExportEnvironmentsUsingGETOK) GetPayload() *models.DslTemplateDto {
	return o.Payload
}

func (o *ExportEnvironmentsUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DslTemplateDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExportEnvironmentsUsingGETUnauthorized creates a ExportEnvironmentsUsingGETUnauthorized with default headers values
func NewExportEnvironmentsUsingGETUnauthorized() *ExportEnvironmentsUsingGETUnauthorized {
	return &ExportEnvironmentsUsingGETUnauthorized{}
}

/*ExportEnvironmentsUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type ExportEnvironmentsUsingGETUnauthorized struct {
}

func (o *ExportEnvironmentsUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /environments/dsl-manifest][%d] exportEnvironmentsUsingGETUnauthorized ", 401)
}

func (o *ExportEnvironmentsUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExportEnvironmentsUsingGETForbidden creates a ExportEnvironmentsUsingGETForbidden with default headers values
func NewExportEnvironmentsUsingGETForbidden() *ExportEnvironmentsUsingGETForbidden {
	return &ExportEnvironmentsUsingGETForbidden{}
}

/*ExportEnvironmentsUsingGETForbidden handles this case with default header values.

Forbidden
*/
type ExportEnvironmentsUsingGETForbidden struct {
}

func (o *ExportEnvironmentsUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /environments/dsl-manifest][%d] exportEnvironmentsUsingGETForbidden ", 403)
}

func (o *ExportEnvironmentsUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExportEnvironmentsUsingGETNotFound creates a ExportEnvironmentsUsingGETNotFound with default headers values
func NewExportEnvironmentsUsingGETNotFound() *ExportEnvironmentsUsingGETNotFound {
	return &ExportEnvironmentsUsingGETNotFound{}
}

/*ExportEnvironmentsUsingGETNotFound handles this case with default header values.

Not Found
*/
type ExportEnvironmentsUsingGETNotFound struct {
}

func (o *ExportEnvironmentsUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /environments/dsl-manifest][%d] exportEnvironmentsUsingGETNotFound ", 404)
}

func (o *ExportEnvironmentsUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
