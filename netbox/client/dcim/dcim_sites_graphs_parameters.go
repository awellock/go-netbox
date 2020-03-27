// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package dcim

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
	"github.com/go-openapi/swag"
)

// NewDcimSitesGraphsParams creates a new DcimSitesGraphsParams object
// with the default values initialized.
func NewDcimSitesGraphsParams() *DcimSitesGraphsParams {
	var ()
	return &DcimSitesGraphsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimSitesGraphsParamsWithTimeout creates a new DcimSitesGraphsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimSitesGraphsParamsWithTimeout(timeout time.Duration) *DcimSitesGraphsParams {
	var ()
	return &DcimSitesGraphsParams{

		timeout: timeout,
	}
}

// NewDcimSitesGraphsParamsWithContext creates a new DcimSitesGraphsParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimSitesGraphsParamsWithContext(ctx context.Context) *DcimSitesGraphsParams {
	var ()
	return &DcimSitesGraphsParams{

		Context: ctx,
	}
}

// NewDcimSitesGraphsParamsWithHTTPClient creates a new DcimSitesGraphsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimSitesGraphsParamsWithHTTPClient(client *http.Client) *DcimSitesGraphsParams {
	var ()
	return &DcimSitesGraphsParams{
		HTTPClient: client,
	}
}

/*DcimSitesGraphsParams contains all the parameters to send to the API endpoint
for the dcim sites graphs operation typically these are written to a http.Request
*/
type DcimSitesGraphsParams struct {

	/*ID
	  A unique integer value identifying this site.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim sites graphs params
func (o *DcimSitesGraphsParams) WithTimeout(timeout time.Duration) *DcimSitesGraphsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim sites graphs params
func (o *DcimSitesGraphsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim sites graphs params
func (o *DcimSitesGraphsParams) WithContext(ctx context.Context) *DcimSitesGraphsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim sites graphs params
func (o *DcimSitesGraphsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim sites graphs params
func (o *DcimSitesGraphsParams) WithHTTPClient(client *http.Client) *DcimSitesGraphsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim sites graphs params
func (o *DcimSitesGraphsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim sites graphs params
func (o *DcimSitesGraphsParams) WithID(id int64) *DcimSitesGraphsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim sites graphs params
func (o *DcimSitesGraphsParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimSitesGraphsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
