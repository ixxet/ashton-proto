package mcp_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

type manifest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ReadOnly    bool   `json:"read_only"`
	Input       struct {
		Required   []string `json:"required"`
		Properties map[string]struct {
			Type        string `json:"type"`
			Description string `json:"description"`
		} `json:"properties"`
	} `json:"input"`
	Upstream struct {
		Service string            `json:"service"`
		Method  string            `json:"method"`
		Path    string            `json:"path"`
		Query   map[string]string `json:"query"`
	} `json:"upstream"`
}

func TestAthenaCurrentOccupancyManifestIsNarrowAndReadOnly(t *testing.T) {
	t.Parallel()

	manifestPath := filepath.Join("..", "..", "mcp", "athena.get_current_occupancy.json")
	payload, err := os.ReadFile(manifestPath)
	if err != nil {
		t.Fatalf("os.ReadFile(%q) error = %v", manifestPath, err)
	}

	var value manifest
	if err := json.Unmarshal(payload, &value); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}

	if value.Name != "athena.get_current_occupancy" {
		t.Fatalf("manifest name = %q, want %q", value.Name, "athena.get_current_occupancy")
	}
	if value.Description == "" {
		t.Fatal("manifest description = empty, want non-empty")
	}
	if !value.ReadOnly {
		t.Fatal("manifest read_only = false, want true")
	}
	if len(value.Input.Required) != 1 || value.Input.Required[0] != "facility_id" {
		t.Fatalf("manifest required inputs = %#v, want [facility_id]", value.Input.Required)
	}
	property, ok := value.Input.Properties["facility_id"]
	if !ok {
		t.Fatal("manifest input properties missing facility_id")
	}
	if property.Type != "string" {
		t.Fatalf("manifest facility_id type = %q, want %q", property.Type, "string")
	}
	if value.Upstream.Service != "athena" {
		t.Fatalf("manifest upstream service = %q, want %q", value.Upstream.Service, "athena")
	}
	if value.Upstream.Method != "GET" {
		t.Fatalf("manifest upstream method = %q, want %q", value.Upstream.Method, "GET")
	}
	if value.Upstream.Path != "/api/v1/presence/count" {
		t.Fatalf("manifest upstream path = %q, want %q", value.Upstream.Path, "/api/v1/presence/count")
	}
	if value.Upstream.Query["facility"] != "facility_id" {
		t.Fatalf("manifest query map = %#v, want facility -> facility_id", value.Upstream.Query)
	}
}
