// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GKEUpgradeSettings GKEUpgradeSettings These upgrade settings control the level of
// parallelism and the level of disruption caused by an upgrade.
// maxUnavailable controls the number of nodes that can be
// simultaneously unavailable. maxSurge controls the number of
// additional nodes that can be added to the node pool temporarily for
// the time of the upgrade to increase the number of available nodes.
// (maxUnavailable + maxSurge) determines the level of parallelism (how
// many nodes are being upgraded at the same time). Note: upgrades
// inevitably introduce some disruption since workloads need to be moved
// from old nodes to new, upgraded ones. Even if maxUnavailable=0, this
// holds true. (Disruption stays within the limits of
// PodDisruptionBudget, if it is configured.) Consider a hypothetical
// node pool with 5 nodes having maxSurge=2, maxUnavailable=1. This
// means the upgrade process upgrades 3 nodes simultaneously. It creates
// 2 additional (upgraded) nodes, then it brings down 3 old (not yet
// upgraded) nodes at the same time. This ensures that there are always
// at least 4 nodes available.
//
// swagger:model GKEUpgradeSettings
type GKEUpgradeSettings struct {

	// MaxSurge: The maximum number of nodes that can be created beyond the
	// current size of the node pool during the upgrade process.
	MaxSurge int64 `json:"maxSurge,omitempty"`

	// MaxUnavailable: The maximum number of nodes that can be
	// simultaneously unavailable during the upgrade process. A node is
	// considered available if its status is Ready.
	MaxUnavailable int64 `json:"maxUnavailable,omitempty"`
}

// Validate validates this g k e upgrade settings
func (m *GKEUpgradeSettings) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this g k e upgrade settings based on context it is used
func (m *GKEUpgradeSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GKEUpgradeSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GKEUpgradeSettings) UnmarshalBinary(b []byte) error {
	var res GKEUpgradeSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}