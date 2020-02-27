// Code generated by go-swagger; DO NOT EDIT.

package directory_server_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteDirectoryServerUsingDELETEReader is a Reader for the DeleteDirectoryServerUsingDELETE structure.
type DeleteDirectoryServerUsingDELETEReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDirectoryServerUsingDELETEReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteDirectoryServerUsingDELETEOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteDirectoryServerUsingDELETENoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteDirectoryServerUsingDELETEUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteDirectoryServerUsingDELETEForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteDirectoryServerUsingDELETEOK creates a DeleteDirectoryServerUsingDELETEOK with default headers values
func NewDeleteDirectoryServerUsingDELETEOK() *DeleteDirectoryServerUsingDELETEOK {
	return &DeleteDirectoryServerUsingDELETEOK{}
}

/*DeleteDirectoryServerUsingDELETEOK handles this case with default header values.

OK
*/
type DeleteDirectoryServerUsingDELETEOK struct {
}

func (o *DeleteDirectoryServerUsingDELETEOK) Error() string {
	return fmt.Sprintf("[DELETE /directory-servers/{directoryServerId}][%d] deleteDirectoryServerUsingDELETEOK ", 200)
}

func (o *DeleteDirectoryServerUsingDELETEOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteDirectoryServerUsingDELETENoContent creates a DeleteDirectoryServerUsingDELETENoContent with default headers values
func NewDeleteDirectoryServerUsingDELETENoContent() *DeleteDirectoryServerUsingDELETENoContent {
	return &DeleteDirectoryServerUsingDELETENoContent{}
}

/*DeleteDirectoryServerUsingDELETENoContent handles this case with default header values.

No Content
*/
type DeleteDirectoryServerUsingDELETENoContent struct {
}

func (o *DeleteDirectoryServerUsingDELETENoContent) Error() string {
	return fmt.Sprintf("[DELETE /directory-servers/{directoryServerId}][%d] deleteDirectoryServerUsingDELETENoContent ", 204)
}

func (o *DeleteDirectoryServerUsingDELETENoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteDirectoryServerUsingDELETEUnauthorized creates a DeleteDirectoryServerUsingDELETEUnauthorized with default headers values
func NewDeleteDirectoryServerUsingDELETEUnauthorized() *DeleteDirectoryServerUsingDELETEUnauthorized {
	return &DeleteDirectoryServerUsingDELETEUnauthorized{}
}

/*DeleteDirectoryServerUsingDELETEUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteDirectoryServerUsingDELETEUnauthorized struct {
}

func (o *DeleteDirectoryServerUsingDELETEUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /directory-servers/{directoryServerId}][%d] deleteDirectoryServerUsingDELETEUnauthorized ", 401)
}

func (o *DeleteDirectoryServerUsingDELETEUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteDirectoryServerUsingDELETEForbidden creates a DeleteDirectoryServerUsingDELETEForbidden with default headers values
func NewDeleteDirectoryServerUsingDELETEForbidden() *DeleteDirectoryServerUsingDELETEForbidden {
	return &DeleteDirectoryServerUsingDELETEForbidden{}
}

/*DeleteDirectoryServerUsingDELETEForbidden handles this case with default header values.

Forbidden
*/
type DeleteDirectoryServerUsingDELETEForbidden struct {
}

func (o *DeleteDirectoryServerUsingDELETEForbidden) Error() string {
	return fmt.Sprintf("[DELETE /directory-servers/{directoryServerId}][%d] deleteDirectoryServerUsingDELETEForbidden ", 403)
}

func (o *DeleteDirectoryServerUsingDELETEForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}