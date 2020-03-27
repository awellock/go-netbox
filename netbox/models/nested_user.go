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

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NestedUser User
//
// swagger:model NestedUser
type NestedUser struct {

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Username
	//
	// Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only.
	// Required: true
	// Max Length: 150
	// Min Length: 1
	// Pattern: ^[\w.@+-]+$
	Username *string `json:"username"`
}

// Validate validates this nested user
func (m *NestedUser) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NestedUser) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("username", "body", m.Username); err != nil {
		return err
	}

	if err := validate.MinLength("username", "body", string(*m.Username), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("username", "body", string(*m.Username), 150); err != nil {
		return err
	}

	if err := validate.Pattern("username", "body", string(*m.Username), `^[\w.@+-]+$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NestedUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NestedUser) UnmarshalBinary(b []byte) error {
	var res NestedUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
