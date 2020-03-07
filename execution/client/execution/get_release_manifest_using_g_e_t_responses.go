// Code generated by go-swagger; DO NOT EDIT.

package execution

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/e88z4/go_cdd/execution/models"
)

// GetReleaseManifestUsingGETReader is a Reader for the GetReleaseManifestUsingGET structure.
type GetReleaseManifestUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReleaseManifestUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetReleaseManifestUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetReleaseManifestUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetReleaseManifestUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetReleaseManifestUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetReleaseManifestUsingGETOK creates a GetReleaseManifestUsingGETOK with default headers values
func NewGetReleaseManifestUsingGETOK() *GetReleaseManifestUsingGETOK {
	return &GetReleaseManifestUsingGETOK{}
}

/*GetReleaseManifestUsingGETOK handles this case with default header values.

OK
*/
type GetReleaseManifestUsingGETOK struct {
	Payload *models.IdentifiableDto
}

func (o *GetReleaseManifestUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /releases-execution/{releaseId}/manifests/{manifestId}][%d] getReleaseManifestUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetReleaseManifestUsingGETOK) GetPayload() *models.IdentifiableDto {
	return o.Payload
}

func (o *GetReleaseManifestUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IdentifiableDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReleaseManifestUsingGETUnauthorized creates a GetReleaseManifestUsingGETUnauthorized with default headers values
func NewGetReleaseManifestUsingGETUnauthorized() *GetReleaseManifestUsingGETUnauthorized {
	return &GetReleaseManifestUsingGETUnauthorized{}
}

/*GetReleaseManifestUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type GetReleaseManifestUsingGETUnauthorized struct {
}

func (o *GetReleaseManifestUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /releases-execution/{releaseId}/manifests/{manifestId}][%d] getReleaseManifestUsingGETUnauthorized ", 401)
}

func (o *GetReleaseManifestUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetReleaseManifestUsingGETForbidden creates a GetReleaseManifestUsingGETForbidden with default headers values
func NewGetReleaseManifestUsingGETForbidden() *GetReleaseManifestUsingGETForbidden {
	return &GetReleaseManifestUsingGETForbidden{}
}

/*GetReleaseManifestUsingGETForbidden handles this case with default header values.

Forbidden
*/
type GetReleaseManifestUsingGETForbidden struct {
}

func (o *GetReleaseManifestUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /releases-execution/{releaseId}/manifests/{manifestId}][%d] getReleaseManifestUsingGETForbidden ", 403)
}

func (o *GetReleaseManifestUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetReleaseManifestUsingGETNotFound creates a GetReleaseManifestUsingGETNotFound with default headers values
func NewGetReleaseManifestUsingGETNotFound() *GetReleaseManifestUsingGETNotFound {
	return &GetReleaseManifestUsingGETNotFound{}
}

/*GetReleaseManifestUsingGETNotFound handles this case with default header values.

Not Found
*/
type GetReleaseManifestUsingGETNotFound struct {
}

func (o *GetReleaseManifestUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /releases-execution/{releaseId}/manifests/{manifestId}][%d] getReleaseManifestUsingGETNotFound ", 404)
}

func (o *GetReleaseManifestUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
