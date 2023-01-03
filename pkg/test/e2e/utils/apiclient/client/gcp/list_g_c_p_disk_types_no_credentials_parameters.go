// Code generated by go-swagger; DO NOT EDIT.

package gcp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewListGCPDiskTypesNoCredentialsParams creates a new ListGCPDiskTypesNoCredentialsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListGCPDiskTypesNoCredentialsParams() *ListGCPDiskTypesNoCredentialsParams {
	return &ListGCPDiskTypesNoCredentialsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListGCPDiskTypesNoCredentialsParamsWithTimeout creates a new ListGCPDiskTypesNoCredentialsParams object
// with the ability to set a timeout on a request.
func NewListGCPDiskTypesNoCredentialsParamsWithTimeout(timeout time.Duration) *ListGCPDiskTypesNoCredentialsParams {
	return &ListGCPDiskTypesNoCredentialsParams{
		timeout: timeout,
	}
}

// NewListGCPDiskTypesNoCredentialsParamsWithContext creates a new ListGCPDiskTypesNoCredentialsParams object
// with the ability to set a context for a request.
func NewListGCPDiskTypesNoCredentialsParamsWithContext(ctx context.Context) *ListGCPDiskTypesNoCredentialsParams {
	return &ListGCPDiskTypesNoCredentialsParams{
		Context: ctx,
	}
}

// NewListGCPDiskTypesNoCredentialsParamsWithHTTPClient creates a new ListGCPDiskTypesNoCredentialsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListGCPDiskTypesNoCredentialsParamsWithHTTPClient(client *http.Client) *ListGCPDiskTypesNoCredentialsParams {
	return &ListGCPDiskTypesNoCredentialsParams{
		HTTPClient: client,
	}
}

/* ListGCPDiskTypesNoCredentialsParams contains all the parameters to send to the API endpoint
   for the list g c p disk types no credentials operation.

   Typically these are written to a http.Request.
*/
type ListGCPDiskTypesNoCredentialsParams struct {

	// Zone.
	Zone *string

	// ClusterID.
	ClusterID string

	// Dc.
	DC string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list g c p disk types no credentials params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListGCPDiskTypesNoCredentialsParams) WithDefaults() *ListGCPDiskTypesNoCredentialsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list g c p disk types no credentials params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListGCPDiskTypesNoCredentialsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list g c p disk types no credentials params
func (o *ListGCPDiskTypesNoCredentialsParams) WithTimeout(timeout time.Duration) *ListGCPDiskTypesNoCredentialsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list g c p disk types no credentials params
func (o *ListGCPDiskTypesNoCredentialsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list g c p disk types no credentials params
func (o *ListGCPDiskTypesNoCredentialsParams) WithContext(ctx context.Context) *ListGCPDiskTypesNoCredentialsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list g c p disk types no credentials params
func (o *ListGCPDiskTypesNoCredentialsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list g c p disk types no credentials params
func (o *ListGCPDiskTypesNoCredentialsParams) WithHTTPClient(client *http.Client) *ListGCPDiskTypesNoCredentialsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list g c p disk types no credentials params
func (o *ListGCPDiskTypesNoCredentialsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithZone adds the zone to the list g c p disk types no credentials params
func (o *ListGCPDiskTypesNoCredentialsParams) WithZone(zone *string) *ListGCPDiskTypesNoCredentialsParams {
	o.SetZone(zone)
	return o
}

// SetZone adds the zone to the list g c p disk types no credentials params
func (o *ListGCPDiskTypesNoCredentialsParams) SetZone(zone *string) {
	o.Zone = zone
}

// WithClusterID adds the clusterID to the list g c p disk types no credentials params
func (o *ListGCPDiskTypesNoCredentialsParams) WithClusterID(clusterID string) *ListGCPDiskTypesNoCredentialsParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the list g c p disk types no credentials params
func (o *ListGCPDiskTypesNoCredentialsParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithDC adds the dc to the list g c p disk types no credentials params
func (o *ListGCPDiskTypesNoCredentialsParams) WithDC(dc string) *ListGCPDiskTypesNoCredentialsParams {
	o.SetDC(dc)
	return o
}

// SetDC adds the dc to the list g c p disk types no credentials params
func (o *ListGCPDiskTypesNoCredentialsParams) SetDC(dc string) {
	o.DC = dc
}

// WithProjectID adds the projectID to the list g c p disk types no credentials params
func (o *ListGCPDiskTypesNoCredentialsParams) WithProjectID(projectID string) *ListGCPDiskTypesNoCredentialsParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the list g c p disk types no credentials params
func (o *ListGCPDiskTypesNoCredentialsParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ListGCPDiskTypesNoCredentialsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Zone != nil {

		// header param Zone
		if err := r.SetHeaderParam("Zone", *o.Zone); err != nil {
			return err
		}
	}

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	// path param dc
	if err := r.SetPathParam("dc", o.DC); err != nil {
		return err
	}

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}