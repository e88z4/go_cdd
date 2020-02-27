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

// UpdateTagsAndIsFavoriteForTestSuiteUsingPATCHReader is a Reader for the UpdateTagsAndIsFavoriteForTestSuiteUsingPATCH structure.
type UpdateTagsAndIsFavoriteForTestSuiteUsingPATCHReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateTagsAndIsFavoriteForTestSuiteUsingPATCHReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateTagsAndIsFavoriteForTestSuiteUsingPATCHOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewUpdateTagsAndIsFavoriteForTestSuiteUsingPATCHNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateTagsAndIsFavoriteForTestSuiteUsingPATCHUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateTagsAndIsFavoriteForTestSuiteUsingPATCHForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateTagsAndIsFavoriteForTestSuiteUsingPATCHOK creates a UpdateTagsAndIsFavoriteForTestSuiteUsingPATCHOK with default headers values
func NewUpdateTagsAndIsFavoriteForTestSuiteUsingPATCHOK() *UpdateTagsAndIsFavoriteForTestSuiteUsingPATCHOK {
	return &UpdateTagsAndIsFavoriteForTestSuiteUsingPATCHOK{}
}

/*UpdateTagsAndIsFavoriteForTestSuiteUsingPATCHOK handles this case with default header values.

OK
*/
type UpdateTagsAndIsFavoriteForTestSuiteUsingPATCHOK struct {
	Payload *models.TestSuiteDto
}

func (o *UpdateTagsAndIsFavoriteForTestSuiteUsingPATCHOK) Error() string {
	return fmt.Sprintf("[PATCH /applications/{applicationId}/application-versions/{applicationVersionId}/test-suites/{testSuiteHexId}][%d] updateTagsAndIsFavoriteForTestSuiteUsingPATCHOK  %+v", 200, o.Payload)
}

func (o *UpdateTagsAndIsFavoriteForTestSuiteUsingPATCHOK) GetPayload() *models.TestSuiteDto {
	return o.Payload
}

func (o *UpdateTagsAndIsFavoriteForTestSuiteUsingPATCHOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TestSuiteDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTagsAndIsFavoriteForTestSuiteUsingPATCHNoContent creates a UpdateTagsAndIsFavoriteForTestSuiteUsingPATCHNoContent with default headers values
func NewUpdateTagsAndIsFavoriteForTestSuiteUsingPATCHNoContent() *UpdateTagsAndIsFavoriteForTestSuiteUsingPATCHNoContent {
	return &UpdateTagsAndIsFavoriteForTestSuiteUsingPATCHNoContent{}
}

/*UpdateTagsAndIsFavoriteForTestSuiteUsingPATCHNoContent handles this case with default header values.

No Content
*/
type UpdateTagsAndIsFavoriteForTestSuiteUsingPATCHNoContent struct {
}

func (o *UpdateTagsAndIsFavoriteForTestSuiteUsingPATCHNoContent) Error() string {
	return fmt.Sprintf("[PATCH /applications/{applicationId}/application-versions/{applicationVersionId}/test-suites/{testSuiteHexId}][%d] updateTagsAndIsFavoriteForTestSuiteUsingPATCHNoContent ", 204)
}

func (o *UpdateTagsAndIsFavoriteForTestSuiteUsingPATCHNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateTagsAndIsFavoriteForTestSuiteUsingPATCHUnauthorized creates a UpdateTagsAndIsFavoriteForTestSuiteUsingPATCHUnauthorized with default headers values
func NewUpdateTagsAndIsFavoriteForTestSuiteUsingPATCHUnauthorized() *UpdateTagsAndIsFavoriteForTestSuiteUsingPATCHUnauthorized {
	return &UpdateTagsAndIsFavoriteForTestSuiteUsingPATCHUnauthorized{}
}

/*UpdateTagsAndIsFavoriteForTestSuiteUsingPATCHUnauthorized handles this case with default header values.

Unauthorized
*/
type UpdateTagsAndIsFavoriteForTestSuiteUsingPATCHUnauthorized struct {
}

func (o *UpdateTagsAndIsFavoriteForTestSuiteUsingPATCHUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /applications/{applicationId}/application-versions/{applicationVersionId}/test-suites/{testSuiteHexId}][%d] updateTagsAndIsFavoriteForTestSuiteUsingPATCHUnauthorized ", 401)
}

func (o *UpdateTagsAndIsFavoriteForTestSuiteUsingPATCHUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateTagsAndIsFavoriteForTestSuiteUsingPATCHForbidden creates a UpdateTagsAndIsFavoriteForTestSuiteUsingPATCHForbidden with default headers values
func NewUpdateTagsAndIsFavoriteForTestSuiteUsingPATCHForbidden() *UpdateTagsAndIsFavoriteForTestSuiteUsingPATCHForbidden {
	return &UpdateTagsAndIsFavoriteForTestSuiteUsingPATCHForbidden{}
}

/*UpdateTagsAndIsFavoriteForTestSuiteUsingPATCHForbidden handles this case with default header values.

Forbidden
*/
type UpdateTagsAndIsFavoriteForTestSuiteUsingPATCHForbidden struct {
}

func (o *UpdateTagsAndIsFavoriteForTestSuiteUsingPATCHForbidden) Error() string {
	return fmt.Sprintf("[PATCH /applications/{applicationId}/application-versions/{applicationVersionId}/test-suites/{testSuiteHexId}][%d] updateTagsAndIsFavoriteForTestSuiteUsingPATCHForbidden ", 403)
}

func (o *UpdateTagsAndIsFavoriteForTestSuiteUsingPATCHForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}