// Code generated by go-swagger; DO NOT EDIT.

package test_suite_executions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/e88z4/go_cdd/cds/models"
)

// GetTestSuiteByIDUsingGETReader is a Reader for the GetTestSuiteByIDUsingGET structure.
type GetTestSuiteByIDUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTestSuiteByIDUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTestSuiteByIDUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetTestSuiteByIDUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTestSuiteByIDUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTestSuiteByIDUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetTestSuiteByIDUsingGETOK creates a GetTestSuiteByIDUsingGETOK with default headers values
func NewGetTestSuiteByIDUsingGETOK() *GetTestSuiteByIDUsingGETOK {
	return &GetTestSuiteByIDUsingGETOK{}
}

/*GetTestSuiteByIDUsingGETOK handles this case with default header values.

OK
*/
type GetTestSuiteByIDUsingGETOK struct {
	Payload *models.TestSuitesExecutionDto
}

func (o *GetTestSuiteByIDUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /applications/{applicationId}/application-versions/{applicationVersionId}/test-suite-executions/{testSuiteExecutionsHexId}][%d] getTestSuiteByIdUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetTestSuiteByIDUsingGETOK) GetPayload() *models.TestSuitesExecutionDto {
	return o.Payload
}

func (o *GetTestSuiteByIDUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TestSuitesExecutionDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTestSuiteByIDUsingGETUnauthorized creates a GetTestSuiteByIDUsingGETUnauthorized with default headers values
func NewGetTestSuiteByIDUsingGETUnauthorized() *GetTestSuiteByIDUsingGETUnauthorized {
	return &GetTestSuiteByIDUsingGETUnauthorized{}
}

/*GetTestSuiteByIDUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type GetTestSuiteByIDUsingGETUnauthorized struct {
}

func (o *GetTestSuiteByIDUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /applications/{applicationId}/application-versions/{applicationVersionId}/test-suite-executions/{testSuiteExecutionsHexId}][%d] getTestSuiteByIdUsingGETUnauthorized ", 401)
}

func (o *GetTestSuiteByIDUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTestSuiteByIDUsingGETForbidden creates a GetTestSuiteByIDUsingGETForbidden with default headers values
func NewGetTestSuiteByIDUsingGETForbidden() *GetTestSuiteByIDUsingGETForbidden {
	return &GetTestSuiteByIDUsingGETForbidden{}
}

/*GetTestSuiteByIDUsingGETForbidden handles this case with default header values.

Forbidden
*/
type GetTestSuiteByIDUsingGETForbidden struct {
}

func (o *GetTestSuiteByIDUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /applications/{applicationId}/application-versions/{applicationVersionId}/test-suite-executions/{testSuiteExecutionsHexId}][%d] getTestSuiteByIdUsingGETForbidden ", 403)
}

func (o *GetTestSuiteByIDUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTestSuiteByIDUsingGETNotFound creates a GetTestSuiteByIDUsingGETNotFound with default headers values
func NewGetTestSuiteByIDUsingGETNotFound() *GetTestSuiteByIDUsingGETNotFound {
	return &GetTestSuiteByIDUsingGETNotFound{}
}

/*GetTestSuiteByIDUsingGETNotFound handles this case with default header values.

Not Found
*/
type GetTestSuiteByIDUsingGETNotFound struct {
}

func (o *GetTestSuiteByIDUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /applications/{applicationId}/application-versions/{applicationVersionId}/test-suite-executions/{testSuiteExecutionsHexId}][%d] getTestSuiteByIdUsingGETNotFound ", 404)
}

func (o *GetTestSuiteByIDUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
