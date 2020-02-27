// Code generated by go-swagger; DO NOT EDIT.

package tenant

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/e88z4/go_cdd/administration/models"
)

// PatchTenantUsingPATCH1Reader is a Reader for the PatchTenantUsingPATCH1 structure.
type PatchTenantUsingPATCH1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchTenantUsingPATCH1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchTenantUsingPATCH1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewPatchTenantUsingPATCH1NoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchTenantUsingPATCH1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchTenantUsingPATCH1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchTenantUsingPATCH1OK creates a PatchTenantUsingPATCH1OK with default headers values
func NewPatchTenantUsingPATCH1OK() *PatchTenantUsingPATCH1OK {
	return &PatchTenantUsingPATCH1OK{}
}

/*PatchTenantUsingPATCH1OK handles this case with default header values.

OK
*/
type PatchTenantUsingPATCH1OK struct {
	Payload *models.TenantDto
}

func (o *PatchTenantUsingPATCH1OK) Error() string {
	return fmt.Sprintf("[PATCH /tenants/{tenantId}][%d] patchTenantUsingPATCH1OK  %+v", 200, o.Payload)
}

func (o *PatchTenantUsingPATCH1OK) GetPayload() *models.TenantDto {
	return o.Payload
}

func (o *PatchTenantUsingPATCH1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TenantDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchTenantUsingPATCH1NoContent creates a PatchTenantUsingPATCH1NoContent with default headers values
func NewPatchTenantUsingPATCH1NoContent() *PatchTenantUsingPATCH1NoContent {
	return &PatchTenantUsingPATCH1NoContent{}
}

/*PatchTenantUsingPATCH1NoContent handles this case with default header values.

No Content
*/
type PatchTenantUsingPATCH1NoContent struct {
}

func (o *PatchTenantUsingPATCH1NoContent) Error() string {
	return fmt.Sprintf("[PATCH /tenants/{tenantId}][%d] patchTenantUsingPATCH1NoContent ", 204)
}

func (o *PatchTenantUsingPATCH1NoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchTenantUsingPATCH1Unauthorized creates a PatchTenantUsingPATCH1Unauthorized with default headers values
func NewPatchTenantUsingPATCH1Unauthorized() *PatchTenantUsingPATCH1Unauthorized {
	return &PatchTenantUsingPATCH1Unauthorized{}
}

/*PatchTenantUsingPATCH1Unauthorized handles this case with default header values.

Unauthorized
*/
type PatchTenantUsingPATCH1Unauthorized struct {
}

func (o *PatchTenantUsingPATCH1Unauthorized) Error() string {
	return fmt.Sprintf("[PATCH /tenants/{tenantId}][%d] patchTenantUsingPATCH1Unauthorized ", 401)
}

func (o *PatchTenantUsingPATCH1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchTenantUsingPATCH1Forbidden creates a PatchTenantUsingPATCH1Forbidden with default headers values
func NewPatchTenantUsingPATCH1Forbidden() *PatchTenantUsingPATCH1Forbidden {
	return &PatchTenantUsingPATCH1Forbidden{}
}

/*PatchTenantUsingPATCH1Forbidden handles this case with default header values.

Forbidden
*/
type PatchTenantUsingPATCH1Forbidden struct {
}

func (o *PatchTenantUsingPATCH1Forbidden) Error() string {
	return fmt.Sprintf("[PATCH /tenants/{tenantId}][%d] patchTenantUsingPATCH1Forbidden ", 403)
}

func (o *PatchTenantUsingPATCH1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
