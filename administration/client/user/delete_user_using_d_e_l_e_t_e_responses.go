// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteUserUsingDELETEReader is a Reader for the DeleteUserUsingDELETE structure.
type DeleteUserUsingDELETEReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUserUsingDELETEReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteUserUsingDELETEOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteUserUsingDELETENoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteUserUsingDELETEUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteUserUsingDELETEForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteUserUsingDELETEOK creates a DeleteUserUsingDELETEOK with default headers values
func NewDeleteUserUsingDELETEOK() *DeleteUserUsingDELETEOK {
	return &DeleteUserUsingDELETEOK{}
}

/*DeleteUserUsingDELETEOK handles this case with default header values.

OK
*/
type DeleteUserUsingDELETEOK struct {
}

func (o *DeleteUserUsingDELETEOK) Error() string {
	return fmt.Sprintf("[DELETE /users/{userId}][%d] deleteUserUsingDELETEOK ", 200)
}

func (o *DeleteUserUsingDELETEOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUserUsingDELETENoContent creates a DeleteUserUsingDELETENoContent with default headers values
func NewDeleteUserUsingDELETENoContent() *DeleteUserUsingDELETENoContent {
	return &DeleteUserUsingDELETENoContent{}
}

/*DeleteUserUsingDELETENoContent handles this case with default header values.

No Content
*/
type DeleteUserUsingDELETENoContent struct {
}

func (o *DeleteUserUsingDELETENoContent) Error() string {
	return fmt.Sprintf("[DELETE /users/{userId}][%d] deleteUserUsingDELETENoContent ", 204)
}

func (o *DeleteUserUsingDELETENoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUserUsingDELETEUnauthorized creates a DeleteUserUsingDELETEUnauthorized with default headers values
func NewDeleteUserUsingDELETEUnauthorized() *DeleteUserUsingDELETEUnauthorized {
	return &DeleteUserUsingDELETEUnauthorized{}
}

/*DeleteUserUsingDELETEUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteUserUsingDELETEUnauthorized struct {
}

func (o *DeleteUserUsingDELETEUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /users/{userId}][%d] deleteUserUsingDELETEUnauthorized ", 401)
}

func (o *DeleteUserUsingDELETEUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUserUsingDELETEForbidden creates a DeleteUserUsingDELETEForbidden with default headers values
func NewDeleteUserUsingDELETEForbidden() *DeleteUserUsingDELETEForbidden {
	return &DeleteUserUsingDELETEForbidden{}
}

/*DeleteUserUsingDELETEForbidden handles this case with default header values.

Forbidden
*/
type DeleteUserUsingDELETEForbidden struct {
}

func (o *DeleteUserUsingDELETEForbidden) Error() string {
	return fmt.Sprintf("[DELETE /users/{userId}][%d] deleteUserUsingDELETEForbidden ", 403)
}

func (o *DeleteUserUsingDELETEForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
