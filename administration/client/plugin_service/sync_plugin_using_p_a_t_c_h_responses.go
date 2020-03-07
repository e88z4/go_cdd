// Code generated by go-swagger; DO NOT EDIT.

package plugin_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/e88z4/go_cdd/administration/models"
)

// SyncPluginUsingPATCHReader is a Reader for the SyncPluginUsingPATCH structure.
type SyncPluginUsingPATCHReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SyncPluginUsingPATCHReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSyncPluginUsingPATCHOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewSyncPluginUsingPATCHNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewSyncPluginUsingPATCHUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSyncPluginUsingPATCHForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSyncPluginUsingPATCHOK creates a SyncPluginUsingPATCHOK with default headers values
func NewSyncPluginUsingPATCHOK() *SyncPluginUsingPATCHOK {
	return &SyncPluginUsingPATCHOK{}
}

/*SyncPluginUsingPATCHOK handles this case with default header values.

OK
*/
type SyncPluginUsingPATCHOK struct {
	Payload *models.PluginDto
}

func (o *SyncPluginUsingPATCHOK) Error() string {
	return fmt.Sprintf("[PATCH /plugins/{pluginId}/sync][%d] syncPluginUsingPATCHOK  %+v", 200, o.Payload)
}

func (o *SyncPluginUsingPATCHOK) GetPayload() *models.PluginDto {
	return o.Payload
}

func (o *SyncPluginUsingPATCHOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PluginDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSyncPluginUsingPATCHNoContent creates a SyncPluginUsingPATCHNoContent with default headers values
func NewSyncPluginUsingPATCHNoContent() *SyncPluginUsingPATCHNoContent {
	return &SyncPluginUsingPATCHNoContent{}
}

/*SyncPluginUsingPATCHNoContent handles this case with default header values.

No Content
*/
type SyncPluginUsingPATCHNoContent struct {
}

func (o *SyncPluginUsingPATCHNoContent) Error() string {
	return fmt.Sprintf("[PATCH /plugins/{pluginId}/sync][%d] syncPluginUsingPATCHNoContent ", 204)
}

func (o *SyncPluginUsingPATCHNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSyncPluginUsingPATCHUnauthorized creates a SyncPluginUsingPATCHUnauthorized with default headers values
func NewSyncPluginUsingPATCHUnauthorized() *SyncPluginUsingPATCHUnauthorized {
	return &SyncPluginUsingPATCHUnauthorized{}
}

/*SyncPluginUsingPATCHUnauthorized handles this case with default header values.

Unauthorized
*/
type SyncPluginUsingPATCHUnauthorized struct {
}

func (o *SyncPluginUsingPATCHUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /plugins/{pluginId}/sync][%d] syncPluginUsingPATCHUnauthorized ", 401)
}

func (o *SyncPluginUsingPATCHUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSyncPluginUsingPATCHForbidden creates a SyncPluginUsingPATCHForbidden with default headers values
func NewSyncPluginUsingPATCHForbidden() *SyncPluginUsingPATCHForbidden {
	return &SyncPluginUsingPATCHForbidden{}
}

/*SyncPluginUsingPATCHForbidden handles this case with default header values.

Forbidden
*/
type SyncPluginUsingPATCHForbidden struct {
}

func (o *SyncPluginUsingPATCHForbidden) Error() string {
	return fmt.Sprintf("[PATCH /plugins/{pluginId}/sync][%d] syncPluginUsingPATCHForbidden ", 403)
}

func (o *SyncPluginUsingPATCHForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
