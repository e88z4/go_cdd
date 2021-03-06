// Code generated by go-swagger; DO NOT EDIT.

package api_access_token

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteAPIAccessTokenUsingDELETEReader is a Reader for the DeleteAPIAccessTokenUsingDELETE structure.
type DeleteAPIAccessTokenUsingDELETEReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAPIAccessTokenUsingDELETEReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteAPIAccessTokenUsingDELETEOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteAPIAccessTokenUsingDELETENoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteAPIAccessTokenUsingDELETEUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteAPIAccessTokenUsingDELETEForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteAPIAccessTokenUsingDELETEOK creates a DeleteAPIAccessTokenUsingDELETEOK with default headers values
func NewDeleteAPIAccessTokenUsingDELETEOK() *DeleteAPIAccessTokenUsingDELETEOK {
	return &DeleteAPIAccessTokenUsingDELETEOK{}
}

/*DeleteAPIAccessTokenUsingDELETEOK handles this case with default header values.

OK
*/
type DeleteAPIAccessTokenUsingDELETEOK struct {
}

func (o *DeleteAPIAccessTokenUsingDELETEOK) Error() string {
	return fmt.Sprintf("[DELETE /api-access-tokens/{apiAccessTokenId}][%d] deleteApiAccessTokenUsingDELETEOK ", 200)
}

func (o *DeleteAPIAccessTokenUsingDELETEOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAPIAccessTokenUsingDELETENoContent creates a DeleteAPIAccessTokenUsingDELETENoContent with default headers values
func NewDeleteAPIAccessTokenUsingDELETENoContent() *DeleteAPIAccessTokenUsingDELETENoContent {
	return &DeleteAPIAccessTokenUsingDELETENoContent{}
}

/*DeleteAPIAccessTokenUsingDELETENoContent handles this case with default header values.

No Content
*/
type DeleteAPIAccessTokenUsingDELETENoContent struct {
}

func (o *DeleteAPIAccessTokenUsingDELETENoContent) Error() string {
	return fmt.Sprintf("[DELETE /api-access-tokens/{apiAccessTokenId}][%d] deleteApiAccessTokenUsingDELETENoContent ", 204)
}

func (o *DeleteAPIAccessTokenUsingDELETENoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAPIAccessTokenUsingDELETEUnauthorized creates a DeleteAPIAccessTokenUsingDELETEUnauthorized with default headers values
func NewDeleteAPIAccessTokenUsingDELETEUnauthorized() *DeleteAPIAccessTokenUsingDELETEUnauthorized {
	return &DeleteAPIAccessTokenUsingDELETEUnauthorized{}
}

/*DeleteAPIAccessTokenUsingDELETEUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteAPIAccessTokenUsingDELETEUnauthorized struct {
}

func (o *DeleteAPIAccessTokenUsingDELETEUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api-access-tokens/{apiAccessTokenId}][%d] deleteApiAccessTokenUsingDELETEUnauthorized ", 401)
}

func (o *DeleteAPIAccessTokenUsingDELETEUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAPIAccessTokenUsingDELETEForbidden creates a DeleteAPIAccessTokenUsingDELETEForbidden with default headers values
func NewDeleteAPIAccessTokenUsingDELETEForbidden() *DeleteAPIAccessTokenUsingDELETEForbidden {
	return &DeleteAPIAccessTokenUsingDELETEForbidden{}
}

/*DeleteAPIAccessTokenUsingDELETEForbidden handles this case with default header values.

Forbidden
*/
type DeleteAPIAccessTokenUsingDELETEForbidden struct {
}

func (o *DeleteAPIAccessTokenUsingDELETEForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api-access-tokens/{apiAccessTokenId}][%d] deleteApiAccessTokenUsingDELETEForbidden ", 403)
}

func (o *DeleteAPIAccessTokenUsingDELETEForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
