// Code generated by go-swagger; DO NOT EDIT.

package test_suite

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/e88z4/go_cdd/tar/models"
)

// DeleteTestSuitesUsingDELETEReader is a Reader for the DeleteTestSuitesUsingDELETE structure.
type DeleteTestSuitesUsingDELETEReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteTestSuitesUsingDELETEReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteTestSuitesUsingDELETEOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteTestSuitesUsingDELETENoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteTestSuitesUsingDELETEUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteTestSuitesUsingDELETEForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteTestSuitesUsingDELETEOK creates a DeleteTestSuitesUsingDELETEOK with default headers values
func NewDeleteTestSuitesUsingDELETEOK() *DeleteTestSuitesUsingDELETEOK {
	return &DeleteTestSuitesUsingDELETEOK{}
}

/*DeleteTestSuitesUsingDELETEOK handles this case with default header values.

OK
*/
type DeleteTestSuitesUsingDELETEOK struct {
	Payload *models.DeletedTestsDto
}

func (o *DeleteTestSuitesUsingDELETEOK) Error() string {
	return fmt.Sprintf("[DELETE /applications/{applicationId}/application-versions/{applicationVersionId}/test-suites][%d] deleteTestSuitesUsingDELETEOK  %+v", 200, o.Payload)
}

func (o *DeleteTestSuitesUsingDELETEOK) GetPayload() *models.DeletedTestsDto {
	return o.Payload
}

func (o *DeleteTestSuitesUsingDELETEOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeletedTestsDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTestSuitesUsingDELETENoContent creates a DeleteTestSuitesUsingDELETENoContent with default headers values
func NewDeleteTestSuitesUsingDELETENoContent() *DeleteTestSuitesUsingDELETENoContent {
	return &DeleteTestSuitesUsingDELETENoContent{}
}

/*DeleteTestSuitesUsingDELETENoContent handles this case with default header values.

No Content
*/
type DeleteTestSuitesUsingDELETENoContent struct {
}

func (o *DeleteTestSuitesUsingDELETENoContent) Error() string {
	return fmt.Sprintf("[DELETE /applications/{applicationId}/application-versions/{applicationVersionId}/test-suites][%d] deleteTestSuitesUsingDELETENoContent ", 204)
}

func (o *DeleteTestSuitesUsingDELETENoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteTestSuitesUsingDELETEUnauthorized creates a DeleteTestSuitesUsingDELETEUnauthorized with default headers values
func NewDeleteTestSuitesUsingDELETEUnauthorized() *DeleteTestSuitesUsingDELETEUnauthorized {
	return &DeleteTestSuitesUsingDELETEUnauthorized{}
}

/*DeleteTestSuitesUsingDELETEUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteTestSuitesUsingDELETEUnauthorized struct {
}

func (o *DeleteTestSuitesUsingDELETEUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /applications/{applicationId}/application-versions/{applicationVersionId}/test-suites][%d] deleteTestSuitesUsingDELETEUnauthorized ", 401)
}

func (o *DeleteTestSuitesUsingDELETEUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteTestSuitesUsingDELETEForbidden creates a DeleteTestSuitesUsingDELETEForbidden with default headers values
func NewDeleteTestSuitesUsingDELETEForbidden() *DeleteTestSuitesUsingDELETEForbidden {
	return &DeleteTestSuitesUsingDELETEForbidden{}
}

/*DeleteTestSuitesUsingDELETEForbidden handles this case with default header values.

Forbidden
*/
type DeleteTestSuitesUsingDELETEForbidden struct {
}

func (o *DeleteTestSuitesUsingDELETEForbidden) Error() string {
	return fmt.Sprintf("[DELETE /applications/{applicationId}/application-versions/{applicationVersionId}/test-suites][%d] deleteTestSuitesUsingDELETEForbidden ", 403)
}

func (o *DeleteTestSuitesUsingDELETEForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}