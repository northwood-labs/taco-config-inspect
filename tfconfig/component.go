// Copyright IBM Corp. 2018, 2025
// Copyright 2026 Northwood Labs, LLC <licensing@northwood-labs.com>
// SPDX-License-Identifier: MPL-2.0

package tfconfig

// Component represents a single component from a Terraform stack config.
type Component struct {
	Name   string    `json:"name"`
	Source string    `json:"source"`
	Pos    SourcePos `json:"pos"`
}
