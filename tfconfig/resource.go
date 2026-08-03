// Copyright IBM Corp. 2018, 2025
// Copyright 2026 Northwood Labs, LLC <licensing@northwood-labs.com>
// SPDX-License-Identifier: MPL-2.0

package tfconfig

import (
	"fmt"
	"strconv"
	"strings"
)

const (InvalidResourceMode ResourceMode = 0
ManagedResourceMode ResourceMode = 'M'
DataResourceMode ResourceMode = 'D')

type (
	// Resource represents a single "resource" or "data" block within a module.
	Resource struct {
		Mode ResourceMode `json:"mode"`
		Type string       `json:"type"`
		Name string       `json:"name"`

		Provider ProviderRef `json:"provider"`

		Pos SourcePos `json:"pos"`
	}

	// ResourceMode represents the "mode" of a resource, which is used to
	// distinguish between managed resources ("resource" blocks in config) and
	// data resources ("data" blocks in config).
	ResourceMode rune
)

// MapKey returns a string that can be used to uniquely identify the receiver
// in a map[string]*Resource.
func (r *Resource) MapKey() string {
	switch r.Mode {
	case ManagedResourceMode:
		return fmt.Sprintf("%s.%s", r.Type, r.Name)
	case DataResourceMode:
		return fmt.Sprintf("data.%s.%s", r.Type, r.Name)
	default:
		// should never happen
		return fmt.Sprintf("[invalid_mode!].%s.%s", r.Type, r.Name)
	}
}

func (m ResourceMode) String() string {
	switch m {
	case ManagedResourceMode:
		return "managed"
	case DataResourceMode:
		return "data"
	default:
		return ""
	}
}

// MarshalJSON implements encoding/json.Marshaler.
func (m ResourceMode) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Quote(m.String())), nil
}

func resourceTypeDefaultProviderName(typeName string) string {
	if underPos := strings.IndexByte(typeName, '_'); underPos != -1 {
		return typeName[:underPos]
	}

	return typeName
}
