// Code generated by go-swagger; DO NOT EDIT.

package environment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteEnvironmentUsingDELETEReader is a Reader for the DeleteEnvironmentUsingDELETE structure.
type DeleteEnvironmentUsingDELETEReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteEnvironmentUsingDELETEReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteEnvironmentUsingDELETEOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteEnvironmentUsingDELETENoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteEnvironmentUsingDELETEUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteEnvironmentUsingDELETEForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteEnvironmentUsingDELETEOK creates a DeleteEnvironmentUsingDELETEOK with default headers values
func NewDeleteEnvironmentUsingDELETEOK() *DeleteEnvironmentUsingDELETEOK {
	return &DeleteEnvironmentUsingDELETEOK{}
}

/*DeleteEnvironmentUsingDELETEOK handles this case with default header values.

OK
*/
type DeleteEnvironmentUsingDELETEOK struct {
}

func (o *DeleteEnvironmentUsingDELETEOK) Error() string {
	return fmt.Sprintf("[DELETE /applications/{applicationId}/environments/{environmentId}][%d] deleteEnvironmentUsingDELETEOK ", 200)
}

func (o *DeleteEnvironmentUsingDELETEOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteEnvironmentUsingDELETENoContent creates a DeleteEnvironmentUsingDELETENoContent with default headers values
func NewDeleteEnvironmentUsingDELETENoContent() *DeleteEnvironmentUsingDELETENoContent {
	return &DeleteEnvironmentUsingDELETENoContent{}
}

/*DeleteEnvironmentUsingDELETENoContent handles this case with default header values.

No Content
*/
type DeleteEnvironmentUsingDELETENoContent struct {
}

func (o *DeleteEnvironmentUsingDELETENoContent) Error() string {
	return fmt.Sprintf("[DELETE /applications/{applicationId}/environments/{environmentId}][%d] deleteEnvironmentUsingDELETENoContent ", 204)
}

func (o *DeleteEnvironmentUsingDELETENoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteEnvironmentUsingDELETEUnauthorized creates a DeleteEnvironmentUsingDELETEUnauthorized with default headers values
func NewDeleteEnvironmentUsingDELETEUnauthorized() *DeleteEnvironmentUsingDELETEUnauthorized {
	return &DeleteEnvironmentUsingDELETEUnauthorized{}
}

/*DeleteEnvironmentUsingDELETEUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteEnvironmentUsingDELETEUnauthorized struct {
}

func (o *DeleteEnvironmentUsingDELETEUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /applications/{applicationId}/environments/{environmentId}][%d] deleteEnvironmentUsingDELETEUnauthorized ", 401)
}

func (o *DeleteEnvironmentUsingDELETEUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteEnvironmentUsingDELETEForbidden creates a DeleteEnvironmentUsingDELETEForbidden with default headers values
func NewDeleteEnvironmentUsingDELETEForbidden() *DeleteEnvironmentUsingDELETEForbidden {
	return &DeleteEnvironmentUsingDELETEForbidden{}
}

/*DeleteEnvironmentUsingDELETEForbidden handles this case with default header values.

Forbidden
*/
type DeleteEnvironmentUsingDELETEForbidden struct {
}

func (o *DeleteEnvironmentUsingDELETEForbidden) Error() string {
	return fmt.Sprintf("[DELETE /applications/{applicationId}/environments/{environmentId}][%d] deleteEnvironmentUsingDELETEForbidden ", 403)
}

func (o *DeleteEnvironmentUsingDELETEForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
