// Code generated by go-swagger; DO NOT EDIT.

package saml_domain

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/e88z4/go_cdd/administration/models"
)

// NewPatchAllSamlDomainsUsingPATCHParams creates a new PatchAllSamlDomainsUsingPATCHParams object
// with the default values initialized.
func NewPatchAllSamlDomainsUsingPATCHParams() *PatchAllSamlDomainsUsingPATCHParams {
	var ()
	return &PatchAllSamlDomainsUsingPATCHParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchAllSamlDomainsUsingPATCHParamsWithTimeout creates a new PatchAllSamlDomainsUsingPATCHParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchAllSamlDomainsUsingPATCHParamsWithTimeout(timeout time.Duration) *PatchAllSamlDomainsUsingPATCHParams {
	var ()
	return &PatchAllSamlDomainsUsingPATCHParams{

		timeout: timeout,
	}
}

// NewPatchAllSamlDomainsUsingPATCHParamsWithContext creates a new PatchAllSamlDomainsUsingPATCHParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchAllSamlDomainsUsingPATCHParamsWithContext(ctx context.Context) *PatchAllSamlDomainsUsingPATCHParams {
	var ()
	return &PatchAllSamlDomainsUsingPATCHParams{

		Context: ctx,
	}
}

// NewPatchAllSamlDomainsUsingPATCHParamsWithHTTPClient creates a new PatchAllSamlDomainsUsingPATCHParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchAllSamlDomainsUsingPATCHParamsWithHTTPClient(client *http.Client) *PatchAllSamlDomainsUsingPATCHParams {
	var ()
	return &PatchAllSamlDomainsUsingPATCHParams{
		HTTPClient: client,
	}
}

/*PatchAllSamlDomainsUsingPATCHParams contains all the parameters to send to the API endpoint
for the patch all saml domains using p a t c h operation typically these are written to a http.Request
*/
type PatchAllSamlDomainsUsingPATCHParams struct {

	/*DomainName
	  domainName

	*/
	DomainName *string
	/*SamlDomainDto
	  samlDomainDto

	*/
	SamlDomainDto *models.SamlDomainDto
	/*TenantID
	  targetTenantId

	*/
	TenantID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch all saml domains using p a t c h params
func (o *PatchAllSamlDomainsUsingPATCHParams) WithTimeout(timeout time.Duration) *PatchAllSamlDomainsUsingPATCHParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch all saml domains using p a t c h params
func (o *PatchAllSamlDomainsUsingPATCHParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch all saml domains using p a t c h params
func (o *PatchAllSamlDomainsUsingPATCHParams) WithContext(ctx context.Context) *PatchAllSamlDomainsUsingPATCHParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch all saml domains using p a t c h params
func (o *PatchAllSamlDomainsUsingPATCHParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch all saml domains using p a t c h params
func (o *PatchAllSamlDomainsUsingPATCHParams) WithHTTPClient(client *http.Client) *PatchAllSamlDomainsUsingPATCHParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch all saml domains using p a t c h params
func (o *PatchAllSamlDomainsUsingPATCHParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDomainName adds the domainName to the patch all saml domains using p a t c h params
func (o *PatchAllSamlDomainsUsingPATCHParams) WithDomainName(domainName *string) *PatchAllSamlDomainsUsingPATCHParams {
	o.SetDomainName(domainName)
	return o
}

// SetDomainName adds the domainName to the patch all saml domains using p a t c h params
func (o *PatchAllSamlDomainsUsingPATCHParams) SetDomainName(domainName *string) {
	o.DomainName = domainName
}

// WithSamlDomainDto adds the samlDomainDto to the patch all saml domains using p a t c h params
func (o *PatchAllSamlDomainsUsingPATCHParams) WithSamlDomainDto(samlDomainDto *models.SamlDomainDto) *PatchAllSamlDomainsUsingPATCHParams {
	o.SetSamlDomainDto(samlDomainDto)
	return o
}

// SetSamlDomainDto adds the samlDomainDto to the patch all saml domains using p a t c h params
func (o *PatchAllSamlDomainsUsingPATCHParams) SetSamlDomainDto(samlDomainDto *models.SamlDomainDto) {
	o.SamlDomainDto = samlDomainDto
}

// WithTenantID adds the tenantID to the patch all saml domains using p a t c h params
func (o *PatchAllSamlDomainsUsingPATCHParams) WithTenantID(tenantID string) *PatchAllSamlDomainsUsingPATCHParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the patch all saml domains using p a t c h params
func (o *PatchAllSamlDomainsUsingPATCHParams) SetTenantID(tenantID string) {
	o.TenantID = tenantID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchAllSamlDomainsUsingPATCHParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DomainName != nil {

		// query param domain_name
		var qrDomainName string
		if o.DomainName != nil {
			qrDomainName = *o.DomainName
		}
		qDomainName := qrDomainName
		if qDomainName != "" {
			if err := r.SetQueryParam("domain_name", qDomainName); err != nil {
				return err
			}
		}

	}

	if o.SamlDomainDto != nil {
		if err := r.SetBodyParam(o.SamlDomainDto); err != nil {
			return err
		}
	}

	// query param tenant_id
	qrTenantID := o.TenantID
	qTenantID := qrTenantID
	if qTenantID != "" {
		if err := r.SetQueryParam("tenant_id", qTenantID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
