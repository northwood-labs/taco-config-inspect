// Copyright IBM Corp. 2018, 2025
// Copyright 2026 Northwood Labs, LLC <licensing@northwood-labs.com>
// SPDX-License-Identifier: MPL-2.0

package tfconfig

// Output represents a single output from a Terraform module.
type Output struct {
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	Deprecated  string    `json:"deprecated,omitempty"`
	Type        string    `json:"type,omitempty"`
	Pos         SourcePos `json:"pos"`
	Sensitive   bool      `json:"sensitive,omitempty"`
}
