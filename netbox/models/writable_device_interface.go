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
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WritableDeviceInterface writable device interface
// swagger:model WritableDeviceInterface
type WritableDeviceInterface struct {

	// cable
	Cable *NestedCable `json:"cable,omitempty"`

	// Connected endpoint
	//
	//
	//         Return the appropriate serializer for the type of connected object.
	//
	// Read Only: true
	ConnectedEndpoint map[string]string `json:"connected_endpoint,omitempty"`

	// Connected endpoint type
	// Read Only: true
	ConnectedEndpointType string `json:"connected_endpoint_type,omitempty"`

	// Connection status
	// Enum: [false true]
	ConnectionStatus bool `json:"connection_status,omitempty"`

	// Count ipaddresses
	// Read Only: true
	CountIpaddresses string `json:"count_ipaddresses,omitempty"`

	// Description
	// Max Length: 100
	Description string `json:"description,omitempty"`

	// Device
	// Required: true
	Device *int64 `json:"device"`

	// Enabled
	Enabled bool `json:"enabled,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Parent LAG
	Lag *int64 `json:"lag,omitempty"`

	// MAC Address
	MacAddress *string `json:"mac_address,omitempty"`

	// OOB Management
	//
	// This interface is used only for out-of-band management
	MgmtOnly bool `json:"mgmt_only,omitempty"`

	// Mode
	// Enum: [access tagged tagged-all]
	Mode string `json:"mode,omitempty"`

	// MTU
	// Maximum: 65536
	// Minimum: 1
	Mtu *int64 `json:"mtu,omitempty"`

	// Name
	// Required: true
	// Max Length: 64
	// Min Length: 1
	Name *string `json:"name"`

	// tagged vlans
	// Unique: true
	TaggedVlans []int64 `json:"tagged_vlans"`

	// tags
	Tags []string `json:"tags"`

	// Type
	// Required: true
	// Enum: [virtual lag 100base-tx 1000base-t 2.5gbase-t 5gbase-t 10gbase-t 10gbase-cx4 1000base-x-gbic 1000base-x-sfp 10gbase-x-sfpp 10gbase-x-xfp 10gbase-x-xenpak 10gbase-x-x2 25gbase-x-sfp28 40gbase-x-qsfpp 50gbase-x-sfp28 100gbase-x-cfp 100gbase-x-cfp2 200gbase-x-cfp2 100gbase-x-cfp4 100gbase-x-cpak 100gbase-x-qsfp28 200gbase-x-qsfp56 400gbase-x-qsfpdd 400gbase-x-osfp ieee802.11a ieee802.11g ieee802.11n ieee802.11ac ieee802.11ad ieee802.11ax gsm cdma lte sonet-oc3 sonet-oc12 sonet-oc48 sonet-oc192 sonet-oc768 sonet-oc1920 sonet-oc3840 1gfc-sfp 2gfc-sfp 4gfc-sfp 8gfc-sfpp 16gfc-sfpp 32gfc-sfp28 128gfc-sfp28 inifiband-sdr inifiband-ddr inifiband-qdr inifiband-fdr10 inifiband-fdr inifiband-edr inifiband-hdr inifiband-ndr inifiband-xdr t1 e1 t3 e3 cisco-stackwise cisco-stackwise-plus cisco-flexstack cisco-flexstack-plus juniper-vcp extreme-summitstack extreme-summitstack-128 extreme-summitstack-256 extreme-summitstack-512 other]
	Type *string `json:"type"`

	// Untagged VLAN
	UntaggedVlan *int64 `json:"untagged_vlan,omitempty"`
}

// Validate validates this writable device interface
func (m *WritableDeviceInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnectionStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMtu(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaggedVlans(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritableDeviceInterface) validateCable(formats strfmt.Registry) error {

	if swag.IsZero(m.Cable) { // not required
		return nil
	}

	if m.Cable != nil {
		if err := m.Cable.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cable")
			}
			return err
		}
	}

	return nil
}

var writableDeviceInterfaceTypeConnectionStatusPropEnum []interface{}

func init() {
	var res []bool
	if err := json.Unmarshal([]byte(`[false,true]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableDeviceInterfaceTypeConnectionStatusPropEnum = append(writableDeviceInterfaceTypeConnectionStatusPropEnum, v)
	}
}

// prop value enum
func (m *WritableDeviceInterface) validateConnectionStatusEnum(path, location string, value bool) error {
	if err := validate.Enum(path, location, value, writableDeviceInterfaceTypeConnectionStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *WritableDeviceInterface) validateConnectionStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.ConnectionStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateConnectionStatusEnum("connection_status", "body", m.ConnectionStatus); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceInterface) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", string(m.Description), 100); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceInterface) validateDevice(formats strfmt.Registry) error {

	if err := validate.Required("device", "body", m.Device); err != nil {
		return err
	}

	return nil
}

var writableDeviceInterfaceTypeModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["access","tagged","tagged-all"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableDeviceInterfaceTypeModePropEnum = append(writableDeviceInterfaceTypeModePropEnum, v)
	}
}

const (

	// WritableDeviceInterfaceModeAccess captures enum value "access"
	WritableDeviceInterfaceModeAccess string = "access"

	// WritableDeviceInterfaceModeTagged captures enum value "tagged"
	WritableDeviceInterfaceModeTagged string = "tagged"

	// WritableDeviceInterfaceModeTaggedAll captures enum value "tagged-all"
	WritableDeviceInterfaceModeTaggedAll string = "tagged-all"
)

// prop value enum
func (m *WritableDeviceInterface) validateModeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, writableDeviceInterfaceTypeModePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *WritableDeviceInterface) validateMode(formats strfmt.Registry) error {

	if swag.IsZero(m.Mode) { // not required
		return nil
	}

	// value enum
	if err := m.validateModeEnum("mode", "body", m.Mode); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceInterface) validateMtu(formats strfmt.Registry) error {

	if swag.IsZero(m.Mtu) { // not required
		return nil
	}

	if err := validate.MinimumInt("mtu", "body", int64(*m.Mtu), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("mtu", "body", int64(*m.Mtu), 65536, false); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceInterface) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 64); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceInterface) validateTaggedVlans(formats strfmt.Registry) error {

	if swag.IsZero(m.TaggedVlans) { // not required
		return nil
	}

	if err := validate.UniqueItems("tagged_vlans", "body", m.TaggedVlans); err != nil {
		return err
	}

	return nil
}

func (m *WritableDeviceInterface) validateTags(formats strfmt.Registry) error {

	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {

		if err := validate.MinLength("tags"+"."+strconv.Itoa(i), "body", string(m.Tags[i]), 1); err != nil {
			return err
		}

	}

	return nil
}

var writableDeviceInterfaceTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["virtual","lag","100base-tx","1000base-t","2.5gbase-t","5gbase-t","10gbase-t","10gbase-cx4","1000base-x-gbic","1000base-x-sfp","10gbase-x-sfpp","10gbase-x-xfp","10gbase-x-xenpak","10gbase-x-x2","25gbase-x-sfp28","40gbase-x-qsfpp","50gbase-x-sfp28","100gbase-x-cfp","100gbase-x-cfp2","200gbase-x-cfp2","100gbase-x-cfp4","100gbase-x-cpak","100gbase-x-qsfp28","200gbase-x-qsfp56","400gbase-x-qsfpdd","400gbase-x-osfp","ieee802.11a","ieee802.11g","ieee802.11n","ieee802.11ac","ieee802.11ad","ieee802.11ax","gsm","cdma","lte","sonet-oc3","sonet-oc12","sonet-oc48","sonet-oc192","sonet-oc768","sonet-oc1920","sonet-oc3840","1gfc-sfp","2gfc-sfp","4gfc-sfp","8gfc-sfpp","16gfc-sfpp","32gfc-sfp28","128gfc-sfp28","inifiband-sdr","inifiband-ddr","inifiband-qdr","inifiband-fdr10","inifiband-fdr","inifiband-edr","inifiband-hdr","inifiband-ndr","inifiband-xdr","t1","e1","t3","e3","cisco-stackwise","cisco-stackwise-plus","cisco-flexstack","cisco-flexstack-plus","juniper-vcp","extreme-summitstack","extreme-summitstack-128","extreme-summitstack-256","extreme-summitstack-512","other"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableDeviceInterfaceTypeTypePropEnum = append(writableDeviceInterfaceTypeTypePropEnum, v)
	}
}

const (

	// WritableDeviceInterfaceTypeVirtual captures enum value "virtual"
	WritableDeviceInterfaceTypeVirtual string = "virtual"

	// WritableDeviceInterfaceTypeLag captures enum value "lag"
	WritableDeviceInterfaceTypeLag string = "lag"

	// WritableDeviceInterfaceTypeNr100baseTx captures enum value "100base-tx"
	WritableDeviceInterfaceTypeNr100baseTx string = "100base-tx"

	// WritableDeviceInterfaceTypeNr1000baset captures enum value "1000base-t"
	WritableDeviceInterfaceTypeNr1000baset string = "1000base-t"

	// WritableDeviceInterfaceTypeNr25gbaset captures enum value "2.5gbase-t"
	WritableDeviceInterfaceTypeNr25gbaset string = "2.5gbase-t"

	// WritableDeviceInterfaceTypeNr5gbaset captures enum value "5gbase-t"
	WritableDeviceInterfaceTypeNr5gbaset string = "5gbase-t"

	// WritableDeviceInterfaceTypeNr10gbaset captures enum value "10gbase-t"
	WritableDeviceInterfaceTypeNr10gbaset string = "10gbase-t"

	// WritableDeviceInterfaceTypeNr10gbaseCx4 captures enum value "10gbase-cx4"
	WritableDeviceInterfaceTypeNr10gbaseCx4 string = "10gbase-cx4"

	// WritableDeviceInterfaceTypeNr1000basexGbic captures enum value "1000base-x-gbic"
	WritableDeviceInterfaceTypeNr1000basexGbic string = "1000base-x-gbic"

	// WritableDeviceInterfaceTypeNr1000basexSfp captures enum value "1000base-x-sfp"
	WritableDeviceInterfaceTypeNr1000basexSfp string = "1000base-x-sfp"

	// WritableDeviceInterfaceTypeNr10gbasexSfpp captures enum value "10gbase-x-sfpp"
	WritableDeviceInterfaceTypeNr10gbasexSfpp string = "10gbase-x-sfpp"

	// WritableDeviceInterfaceTypeNr10gbasexXfp captures enum value "10gbase-x-xfp"
	WritableDeviceInterfaceTypeNr10gbasexXfp string = "10gbase-x-xfp"

	// WritableDeviceInterfaceTypeNr10gbasexXenpak captures enum value "10gbase-x-xenpak"
	WritableDeviceInterfaceTypeNr10gbasexXenpak string = "10gbase-x-xenpak"

	// WritableDeviceInterfaceTypeNr10gbasexX2 captures enum value "10gbase-x-x2"
	WritableDeviceInterfaceTypeNr10gbasexX2 string = "10gbase-x-x2"

	// WritableDeviceInterfaceTypeNr25gbasexSfp28 captures enum value "25gbase-x-sfp28"
	WritableDeviceInterfaceTypeNr25gbasexSfp28 string = "25gbase-x-sfp28"

	// WritableDeviceInterfaceTypeNr40gbasexQsfpp captures enum value "40gbase-x-qsfpp"
	WritableDeviceInterfaceTypeNr40gbasexQsfpp string = "40gbase-x-qsfpp"

	// WritableDeviceInterfaceTypeNr50gbasexSfp28 captures enum value "50gbase-x-sfp28"
	WritableDeviceInterfaceTypeNr50gbasexSfp28 string = "50gbase-x-sfp28"

	// WritableDeviceInterfaceTypeNr100gbasexCfp captures enum value "100gbase-x-cfp"
	WritableDeviceInterfaceTypeNr100gbasexCfp string = "100gbase-x-cfp"

	// WritableDeviceInterfaceTypeNr100gbasexCfp2 captures enum value "100gbase-x-cfp2"
	WritableDeviceInterfaceTypeNr100gbasexCfp2 string = "100gbase-x-cfp2"

	// WritableDeviceInterfaceTypeNr200gbasexCfp2 captures enum value "200gbase-x-cfp2"
	WritableDeviceInterfaceTypeNr200gbasexCfp2 string = "200gbase-x-cfp2"

	// WritableDeviceInterfaceTypeNr100gbasexCfp4 captures enum value "100gbase-x-cfp4"
	WritableDeviceInterfaceTypeNr100gbasexCfp4 string = "100gbase-x-cfp4"

	// WritableDeviceInterfaceTypeNr100gbasexCpak captures enum value "100gbase-x-cpak"
	WritableDeviceInterfaceTypeNr100gbasexCpak string = "100gbase-x-cpak"

	// WritableDeviceInterfaceTypeNr100gbasexQsfp28 captures enum value "100gbase-x-qsfp28"
	WritableDeviceInterfaceTypeNr100gbasexQsfp28 string = "100gbase-x-qsfp28"

	// WritableDeviceInterfaceTypeNr200gbasexQsfp56 captures enum value "200gbase-x-qsfp56"
	WritableDeviceInterfaceTypeNr200gbasexQsfp56 string = "200gbase-x-qsfp56"

	// WritableDeviceInterfaceTypeNr400gbasexQsfpdd captures enum value "400gbase-x-qsfpdd"
	WritableDeviceInterfaceTypeNr400gbasexQsfpdd string = "400gbase-x-qsfpdd"

	// WritableDeviceInterfaceTypeNr400gbasexOsfp captures enum value "400gbase-x-osfp"
	WritableDeviceInterfaceTypeNr400gbasexOsfp string = "400gbase-x-osfp"

	// WritableDeviceInterfaceTypeIeee80211a captures enum value "ieee802.11a"
	WritableDeviceInterfaceTypeIeee80211a string = "ieee802.11a"

	// WritableDeviceInterfaceTypeIeee80211g captures enum value "ieee802.11g"
	WritableDeviceInterfaceTypeIeee80211g string = "ieee802.11g"

	// WritableDeviceInterfaceTypeIeee80211n captures enum value "ieee802.11n"
	WritableDeviceInterfaceTypeIeee80211n string = "ieee802.11n"

	// WritableDeviceInterfaceTypeIeee80211ac captures enum value "ieee802.11ac"
	WritableDeviceInterfaceTypeIeee80211ac string = "ieee802.11ac"

	// WritableDeviceInterfaceTypeIeee80211ad captures enum value "ieee802.11ad"
	WritableDeviceInterfaceTypeIeee80211ad string = "ieee802.11ad"

	// WritableDeviceInterfaceTypeIeee80211ax captures enum value "ieee802.11ax"
	WritableDeviceInterfaceTypeIeee80211ax string = "ieee802.11ax"

	// WritableDeviceInterfaceTypeGsm captures enum value "gsm"
	WritableDeviceInterfaceTypeGsm string = "gsm"

	// WritableDeviceInterfaceTypeCdma captures enum value "cdma"
	WritableDeviceInterfaceTypeCdma string = "cdma"

	// WritableDeviceInterfaceTypeLte captures enum value "lte"
	WritableDeviceInterfaceTypeLte string = "lte"

	// WritableDeviceInterfaceTypeSonetOc3 captures enum value "sonet-oc3"
	WritableDeviceInterfaceTypeSonetOc3 string = "sonet-oc3"

	// WritableDeviceInterfaceTypeSonetOc12 captures enum value "sonet-oc12"
	WritableDeviceInterfaceTypeSonetOc12 string = "sonet-oc12"

	// WritableDeviceInterfaceTypeSonetOc48 captures enum value "sonet-oc48"
	WritableDeviceInterfaceTypeSonetOc48 string = "sonet-oc48"

	// WritableDeviceInterfaceTypeSonetOc192 captures enum value "sonet-oc192"
	WritableDeviceInterfaceTypeSonetOc192 string = "sonet-oc192"

	// WritableDeviceInterfaceTypeSonetOc768 captures enum value "sonet-oc768"
	WritableDeviceInterfaceTypeSonetOc768 string = "sonet-oc768"

	// WritableDeviceInterfaceTypeSonetOc1920 captures enum value "sonet-oc1920"
	WritableDeviceInterfaceTypeSonetOc1920 string = "sonet-oc1920"

	// WritableDeviceInterfaceTypeSonetOc3840 captures enum value "sonet-oc3840"
	WritableDeviceInterfaceTypeSonetOc3840 string = "sonet-oc3840"

	// WritableDeviceInterfaceTypeNr1gfcSfp captures enum value "1gfc-sfp"
	WritableDeviceInterfaceTypeNr1gfcSfp string = "1gfc-sfp"

	// WritableDeviceInterfaceTypeNr2gfcSfp captures enum value "2gfc-sfp"
	WritableDeviceInterfaceTypeNr2gfcSfp string = "2gfc-sfp"

	// WritableDeviceInterfaceTypeNr4gfcSfp captures enum value "4gfc-sfp"
	WritableDeviceInterfaceTypeNr4gfcSfp string = "4gfc-sfp"

	// WritableDeviceInterfaceTypeNr8gfcSfpp captures enum value "8gfc-sfpp"
	WritableDeviceInterfaceTypeNr8gfcSfpp string = "8gfc-sfpp"

	// WritableDeviceInterfaceTypeNr16gfcSfpp captures enum value "16gfc-sfpp"
	WritableDeviceInterfaceTypeNr16gfcSfpp string = "16gfc-sfpp"

	// WritableDeviceInterfaceTypeNr32gfcSfp28 captures enum value "32gfc-sfp28"
	WritableDeviceInterfaceTypeNr32gfcSfp28 string = "32gfc-sfp28"

	// WritableDeviceInterfaceTypeNr128gfcSfp28 captures enum value "128gfc-sfp28"
	WritableDeviceInterfaceTypeNr128gfcSfp28 string = "128gfc-sfp28"

	// WritableDeviceInterfaceTypeInifibandSdr captures enum value "inifiband-sdr"
	WritableDeviceInterfaceTypeInifibandSdr string = "inifiband-sdr"

	// WritableDeviceInterfaceTypeInifibandDdr captures enum value "inifiband-ddr"
	WritableDeviceInterfaceTypeInifibandDdr string = "inifiband-ddr"

	// WritableDeviceInterfaceTypeInifibandQdr captures enum value "inifiband-qdr"
	WritableDeviceInterfaceTypeInifibandQdr string = "inifiband-qdr"

	// WritableDeviceInterfaceTypeInifibandFdr10 captures enum value "inifiband-fdr10"
	WritableDeviceInterfaceTypeInifibandFdr10 string = "inifiband-fdr10"

	// WritableDeviceInterfaceTypeInifibandFdr captures enum value "inifiband-fdr"
	WritableDeviceInterfaceTypeInifibandFdr string = "inifiband-fdr"

	// WritableDeviceInterfaceTypeInifibandEdr captures enum value "inifiband-edr"
	WritableDeviceInterfaceTypeInifibandEdr string = "inifiband-edr"

	// WritableDeviceInterfaceTypeInifibandHdr captures enum value "inifiband-hdr"
	WritableDeviceInterfaceTypeInifibandHdr string = "inifiband-hdr"

	// WritableDeviceInterfaceTypeInifibandNdr captures enum value "inifiband-ndr"
	WritableDeviceInterfaceTypeInifibandNdr string = "inifiband-ndr"

	// WritableDeviceInterfaceTypeInifibandXdr captures enum value "inifiband-xdr"
	WritableDeviceInterfaceTypeInifibandXdr string = "inifiband-xdr"

	// WritableDeviceInterfaceTypeT1 captures enum value "t1"
	WritableDeviceInterfaceTypeT1 string = "t1"

	// WritableDeviceInterfaceTypeE1 captures enum value "e1"
	WritableDeviceInterfaceTypeE1 string = "e1"

	// WritableDeviceInterfaceTypeT3 captures enum value "t3"
	WritableDeviceInterfaceTypeT3 string = "t3"

	// WritableDeviceInterfaceTypeE3 captures enum value "e3"
	WritableDeviceInterfaceTypeE3 string = "e3"

	// WritableDeviceInterfaceTypeCiscoStackwise captures enum value "cisco-stackwise"
	WritableDeviceInterfaceTypeCiscoStackwise string = "cisco-stackwise"

	// WritableDeviceInterfaceTypeCiscoStackwisePlus captures enum value "cisco-stackwise-plus"
	WritableDeviceInterfaceTypeCiscoStackwisePlus string = "cisco-stackwise-plus"

	// WritableDeviceInterfaceTypeCiscoFlexstack captures enum value "cisco-flexstack"
	WritableDeviceInterfaceTypeCiscoFlexstack string = "cisco-flexstack"

	// WritableDeviceInterfaceTypeCiscoFlexstackPlus captures enum value "cisco-flexstack-plus"
	WritableDeviceInterfaceTypeCiscoFlexstackPlus string = "cisco-flexstack-plus"

	// WritableDeviceInterfaceTypeJuniperVcp captures enum value "juniper-vcp"
	WritableDeviceInterfaceTypeJuniperVcp string = "juniper-vcp"

	// WritableDeviceInterfaceTypeExtremeSummitstack captures enum value "extreme-summitstack"
	WritableDeviceInterfaceTypeExtremeSummitstack string = "extreme-summitstack"

	// WritableDeviceInterfaceTypeExtremeSummitstack128 captures enum value "extreme-summitstack-128"
	WritableDeviceInterfaceTypeExtremeSummitstack128 string = "extreme-summitstack-128"

	// WritableDeviceInterfaceTypeExtremeSummitstack256 captures enum value "extreme-summitstack-256"
	WritableDeviceInterfaceTypeExtremeSummitstack256 string = "extreme-summitstack-256"

	// WritableDeviceInterfaceTypeExtremeSummitstack512 captures enum value "extreme-summitstack-512"
	WritableDeviceInterfaceTypeExtremeSummitstack512 string = "extreme-summitstack-512"

	// WritableDeviceInterfaceTypeOther captures enum value "other"
	WritableDeviceInterfaceTypeOther string = "other"
)

// prop value enum
func (m *WritableDeviceInterface) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, writableDeviceInterfaceTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *WritableDeviceInterface) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WritableDeviceInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritableDeviceInterface) UnmarshalBinary(b []byte) error {
	var res WritableDeviceInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
