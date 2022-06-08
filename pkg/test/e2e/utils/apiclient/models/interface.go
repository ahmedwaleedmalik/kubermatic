// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Interface interface
//
// swagger:model Interface
type Interface struct {

	// BootOrder is an integer value > 0, used to determine ordering of boot devices.
	// Lower values take precedence.
	// Each interface or disk that has a boot order must have a unique value.
	// Interfaces without a boot order are not tried.
	// +optional
	BootOrder uint64 `json:"bootOrder,omitempty"`

	// mac address
	// Example: de:ad:00:00:be:af or DE-AD-00-00-BE-AF.
	MacAddress string `json:"macAddress,omitempty"`

	// Interface model.
	// One of: e1000, e1000e, ne2k_pci, pcnet, rtl8139, virtio.
	// Defaults to virtio.
	// TODO:(ihar) switch to enums once opengen-api supports them. See: https://github.com/kubernetes/kube-openapi/issues/51
	Model string `json:"model,omitempty"`

	// Logical name of the interface as well as a reference to the associated networks.
	// Must match the Name of a Network.
	Name string `json:"name,omitempty"`

	// pci address
	// Example: 0000:81:01.10
	PciAddress string `json:"pciAddress,omitempty"`

	// List of ports to be forwarded to the virtual machine.
	Ports []*Port `json:"ports"`

	// If specified, the virtual network interface address and its tag will be provided to the guest via config drive
	// +optional
	Tag string `json:"tag,omitempty"`

	// bridge
	Bridge InterfaceBridge `json:"bridge,omitempty"`

	// dhcp options
	DhcpOptions *DHCPOptions `json:"dhcpOptions,omitempty"`

	// macvtap
	Macvtap InterfaceMacvtap `json:"macvtap,omitempty"`

	// masquerade
	Masquerade InterfaceMasquerade `json:"masquerade,omitempty"`

	// slirp
	Slirp InterfaceSlirp `json:"slirp,omitempty"`

	// sriov
	Sriov InterfaceSRIOV `json:"sriov,omitempty"`
}

// Validate validates this interface
func (m *Interface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePorts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDhcpOptions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Interface) validatePorts(formats strfmt.Registry) error {
	if swag.IsZero(m.Ports) { // not required
		return nil
	}

	for i := 0; i < len(m.Ports); i++ {
		if swag.IsZero(m.Ports[i]) { // not required
			continue
		}

		if m.Ports[i] != nil {
			if err := m.Ports[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ports" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Interface) validateDhcpOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.DhcpOptions) { // not required
		return nil
	}

	if m.DhcpOptions != nil {
		if err := m.DhcpOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dhcpOptions")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this interface based on the context it is used
func (m *Interface) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePorts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDhcpOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Interface) contextValidatePorts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Ports); i++ {

		if m.Ports[i] != nil {
			if err := m.Ports[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ports" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Interface) contextValidateDhcpOptions(ctx context.Context, formats strfmt.Registry) error {

	if m.DhcpOptions != nil {
		if err := m.DhcpOptions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dhcpOptions")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Interface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Interface) UnmarshalBinary(b []byte) error {
	var res Interface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}