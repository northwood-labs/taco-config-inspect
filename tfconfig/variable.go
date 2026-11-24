// Copyright IBM Corp. 2018, 2025
// Copyright 2026 Northwood Labs, LLC <licensing@northwood-labs.com>
// SPDX-License-Identifier: MPL-2.0

package tfconfig

// Variable represents a single variable from a Terraform module.
type Variable struct {
	Default     any       `json:"default"`
	Name        string    `json:"name"`
	Type        string    `json:"type,omitempty"`
	Description string    `json:"description,omitempty"`
	Deprecated  string    `json:"deprecated,omitempty"`
	Pos         SourcePos `json:"pos"`
	Required    bool      `json:"required"`
	Sensitive   bool      `json:"sensitive,omitempty"`
}
