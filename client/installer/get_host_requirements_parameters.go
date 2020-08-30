// Code generated by go-swagger; DO NOT EDIT.

package installer

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

// NewGetHostRequirementsParams creates a new GetHostRequirementsParams object
// with the default values initialized.
func NewGetHostRequirementsParams() *GetHostRequirementsParams {

	return &GetHostRequirementsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetHostRequirementsParamsWithTimeout creates a new GetHostRequirementsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetHostRequirementsParamsWithTimeout(timeout time.Duration) *GetHostRequirementsParams {

	return &GetHostRequirementsParams{

		timeout: timeout,
	}
}

// NewGetHostRequirementsParamsWithContext creates a new GetHostRequirementsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetHostRequirementsParamsWithContext(ctx context.Context) *GetHostRequirementsParams {

	return &GetHostRequirementsParams{

		Context: ctx,
	}
}

// NewGetHostRequirementsParamsWithHTTPClient creates a new GetHostRequirementsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetHostRequirementsParamsWithHTTPClient(client *http.Client) *GetHostRequirementsParams {

	return &GetHostRequirementsParams{
		HTTPClient: client,
	}
}

/*GetHostRequirementsParams contains all the parameters to send to the API endpoint
for the get host requirements operation typically these are written to a http.Request
*/
type GetHostRequirementsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get host requirements params
func (o *GetHostRequirementsParams) WithTimeout(timeout time.Duration) *GetHostRequirementsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get host requirements params
func (o *GetHostRequirementsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get host requirements params
func (o *GetHostRequirementsParams) WithContext(ctx context.Context) *GetHostRequirementsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get host requirements params
func (o *GetHostRequirementsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get host requirements params
func (o *GetHostRequirementsParams) WithHTTPClient(client *http.Client) *GetHostRequirementsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get host requirements params
func (o *GetHostRequirementsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetHostRequirementsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
