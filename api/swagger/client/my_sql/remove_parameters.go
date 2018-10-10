// Code generated by go-swagger; DO NOT EDIT.

package my_sql

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewRemoveParams creates a new RemoveParams object
// with the default values initialized.
func NewRemoveParams() *RemoveParams {
	var ()
	return &RemoveParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRemoveParamsWithTimeout creates a new RemoveParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRemoveParamsWithTimeout(timeout time.Duration) *RemoveParams {
	var ()
	return &RemoveParams{

		timeout: timeout,
	}
}

// NewRemoveParamsWithContext creates a new RemoveParams object
// with the default values initialized, and the ability to set a context for a request
func NewRemoveParamsWithContext(ctx context.Context) *RemoveParams {
	var ()
	return &RemoveParams{

		Context: ctx,
	}
}

// NewRemoveParamsWithHTTPClient creates a new RemoveParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRemoveParamsWithHTTPClient(client *http.Client) *RemoveParams {
	var ()
	return &RemoveParams{
		HTTPClient: client,
	}
}

/*RemoveParams contains all the parameters to send to the API endpoint
for the remove operation typically these are written to a http.Request
*/
type RemoveParams struct {

	/*ID*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the remove params
func (o *RemoveParams) WithTimeout(timeout time.Duration) *RemoveParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the remove params
func (o *RemoveParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the remove params
func (o *RemoveParams) WithContext(ctx context.Context) *RemoveParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the remove params
func (o *RemoveParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the remove params
func (o *RemoveParams) WithHTTPClient(client *http.Client) *RemoveParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the remove params
func (o *RemoveParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the remove params
func (o *RemoveParams) WithID(id int64) *RemoveParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the remove params
func (o *RemoveParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *RemoveParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
