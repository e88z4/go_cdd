// Code generated by go-swagger; DO NOT EDIT.

package log_files

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetLogsUsingGETReader is a Reader for the GetLogsUsingGET structure.
type GetLogsUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLogsUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLogsUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetLogsUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetLogsUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetLogsUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLogsUsingGETOK creates a GetLogsUsingGETOK with default headers values
func NewGetLogsUsingGETOK() *GetLogsUsingGETOK {
	return &GetLogsUsingGETOK{}
}

/*GetLogsUsingGETOK handles this case with default header values.

OK
*/
type GetLogsUsingGETOK struct {
}

func (o *GetLogsUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /product/settings/home-folder/files][%d] getLogsUsingGETOK ", 200)
}

func (o *GetLogsUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetLogsUsingGETUnauthorized creates a GetLogsUsingGETUnauthorized with default headers values
func NewGetLogsUsingGETUnauthorized() *GetLogsUsingGETUnauthorized {
	return &GetLogsUsingGETUnauthorized{}
}

/*GetLogsUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type GetLogsUsingGETUnauthorized struct {
}

func (o *GetLogsUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /product/settings/home-folder/files][%d] getLogsUsingGETUnauthorized ", 401)
}

func (o *GetLogsUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetLogsUsingGETForbidden creates a GetLogsUsingGETForbidden with default headers values
func NewGetLogsUsingGETForbidden() *GetLogsUsingGETForbidden {
	return &GetLogsUsingGETForbidden{}
}

/*GetLogsUsingGETForbidden handles this case with default header values.

Forbidden
*/
type GetLogsUsingGETForbidden struct {
}

func (o *GetLogsUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /product/settings/home-folder/files][%d] getLogsUsingGETForbidden ", 403)
}

func (o *GetLogsUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetLogsUsingGETNotFound creates a GetLogsUsingGETNotFound with default headers values
func NewGetLogsUsingGETNotFound() *GetLogsUsingGETNotFound {
	return &GetLogsUsingGETNotFound{}
}

/*GetLogsUsingGETNotFound handles this case with default header values.

Not Found
*/
type GetLogsUsingGETNotFound struct {
}

func (o *GetLogsUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /product/settings/home-folder/files][%d] getLogsUsingGETNotFound ", 404)
}

func (o *GetLogsUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
