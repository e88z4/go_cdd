// Code generated by go-swagger; DO NOT EDIT.

package registered_plugins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/e88z4/go_cdd/administration/models"
)

// PatchRegisteredPluginsUsingPATCHReader is a Reader for the PatchRegisteredPluginsUsingPATCH structure.
type PatchRegisteredPluginsUsingPATCHReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchRegisteredPluginsUsingPATCHReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchRegisteredPluginsUsingPATCHOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewPatchRegisteredPluginsUsingPATCHNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchRegisteredPluginsUsingPATCHUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchRegisteredPluginsUsingPATCHForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchRegisteredPluginsUsingPATCHOK creates a PatchRegisteredPluginsUsingPATCHOK with default headers values
func NewPatchRegisteredPluginsUsingPATCHOK() *PatchRegisteredPluginsUsingPATCHOK {
	return &PatchRegisteredPluginsUsingPATCHOK{}
}

/*PatchRegisteredPluginsUsingPATCHOK handles this case with default header values.

OK
*/
type PatchRegisteredPluginsUsingPATCHOK struct {
	Payload *models.RegisteredPluginDtoList
}

func (o *PatchRegisteredPluginsUsingPATCHOK) Error() string {
	return fmt.Sprintf("[PATCH /registered-plugins][%d] patchRegisteredPluginsUsingPATCHOK  %+v", 200, o.Payload)
}

func (o *PatchRegisteredPluginsUsingPATCHOK) GetPayload() *models.RegisteredPluginDtoList {
	return o.Payload
}

func (o *PatchRegisteredPluginsUsingPATCHOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RegisteredPluginDtoList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRegisteredPluginsUsingPATCHNoContent creates a PatchRegisteredPluginsUsingPATCHNoContent with default headers values
func NewPatchRegisteredPluginsUsingPATCHNoContent() *PatchRegisteredPluginsUsingPATCHNoContent {
	return &PatchRegisteredPluginsUsingPATCHNoContent{}
}

/*PatchRegisteredPluginsUsingPATCHNoContent handles this case with default header values.

No Content
*/
type PatchRegisteredPluginsUsingPATCHNoContent struct {
}

func (o *PatchRegisteredPluginsUsingPATCHNoContent) Error() string {
	return fmt.Sprintf("[PATCH /registered-plugins][%d] patchRegisteredPluginsUsingPATCHNoContent ", 204)
}

func (o *PatchRegisteredPluginsUsingPATCHNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchRegisteredPluginsUsingPATCHUnauthorized creates a PatchRegisteredPluginsUsingPATCHUnauthorized with default headers values
func NewPatchRegisteredPluginsUsingPATCHUnauthorized() *PatchRegisteredPluginsUsingPATCHUnauthorized {
	return &PatchRegisteredPluginsUsingPATCHUnauthorized{}
}

/*PatchRegisteredPluginsUsingPATCHUnauthorized handles this case with default header values.

Unauthorized
*/
type PatchRegisteredPluginsUsingPATCHUnauthorized struct {
}

func (o *PatchRegisteredPluginsUsingPATCHUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /registered-plugins][%d] patchRegisteredPluginsUsingPATCHUnauthorized ", 401)
}

func (o *PatchRegisteredPluginsUsingPATCHUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchRegisteredPluginsUsingPATCHForbidden creates a PatchRegisteredPluginsUsingPATCHForbidden with default headers values
func NewPatchRegisteredPluginsUsingPATCHForbidden() *PatchRegisteredPluginsUsingPATCHForbidden {
	return &PatchRegisteredPluginsUsingPATCHForbidden{}
}

/*PatchRegisteredPluginsUsingPATCHForbidden handles this case with default header values.

Forbidden
*/
type PatchRegisteredPluginsUsingPATCHForbidden struct {
}

func (o *PatchRegisteredPluginsUsingPATCHForbidden) Error() string {
	return fmt.Sprintf("[PATCH /registered-plugins][%d] patchRegisteredPluginsUsingPATCHForbidden ", 403)
}

func (o *PatchRegisteredPluginsUsingPATCHForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}