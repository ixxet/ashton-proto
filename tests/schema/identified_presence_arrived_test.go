package schema_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	jsonschema "github.com/santhosh-tekuri/jsonschema/v6"
)

func TestIdentifiedPresenceArrivedSchemaAcceptsValidEnvelope(t *testing.T) {
	schema := loadSchema(t)

	event := map[string]any{
		"id":        "mock-in-001",
		"source":    "athena",
		"type":      "athena.identified_presence.arrived",
		"timestamp": "2026-04-01T12:30:00Z",
		"data": map[string]any{
			"facility_id":            "ashtonbee",
			"zone_id":                "weight-room",
			"external_identity_hash": "tag_tracer2_001",
			"source":                 "mock",
			"recorded_at":            "2026-04-01T12:30:00Z",
		},
	}

	validateEvent(t, schema, event)
}

func TestIdentifiedPresenceArrivedSchemaRejectsMissingFacilityID(t *testing.T) {
	schema := loadSchema(t)

	event := validEvent()
	delete(event["data"].(map[string]any), "facility_id")

	assertValidationError(t, schema, event)
}

func TestIdentifiedPresenceArrivedSchemaRejectsMissingRecordedAt(t *testing.T) {
	schema := loadSchema(t)

	event := validEvent()
	delete(event["data"].(map[string]any), "recorded_at")

	assertValidationError(t, schema, event)
}

func TestIdentifiedPresenceArrivedSchemaRejectsMissingExternalIdentityHash(t *testing.T) {
	schema := loadSchema(t)

	event := validEvent()
	delete(event["data"].(map[string]any), "external_identity_hash")

	assertValidationError(t, schema, event)
}

func loadSchema(t *testing.T) *jsonschema.Schema {
	t.Helper()

	_, currentFile, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("runtime.Caller() = false")
	}

	eventsDir := filepath.Join(filepath.Dir(currentFile), "..", "..", "events")
	schemaPath := filepath.Join(eventsDir, "athena.identified_presence.arrived.schema.json")
	envelopePath := filepath.Join(eventsDir, "envelope.schema.json")
	compiler := jsonschema.NewCompiler()
	if err := addResource(compiler, "https://github.com/ixxet/ashton-proto/events/athena.identified_presence.arrived.schema.json", schemaPath); err != nil {
		t.Fatalf("addResource(schema) error = %v", err)
	}
	if err := addResource(compiler, "https://github.com/ixxet/ashton-proto/events/envelope.schema.json", envelopePath); err != nil {
		t.Fatalf("addResource(envelope) error = %v", err)
	}

	schema, err := compiler.Compile("https://github.com/ixxet/ashton-proto/events/athena.identified_presence.arrived.schema.json")
	if err != nil {
		t.Fatalf("compiler.Compile() error = %v", err)
	}

	return schema
}

func addResource(compiler *jsonschema.Compiler, url string, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	doc, err := jsonschema.UnmarshalJSON(file)
	if err != nil {
		return err
	}

	return compiler.AddResource(url, doc)
}

func validateEvent(t *testing.T, schema *jsonschema.Schema, event map[string]any) {
	t.Helper()

	payload, err := json.Marshal(event)
	if err != nil {
		t.Fatalf("json.Marshal() error = %v", err)
	}

	var value any
	if err := json.Unmarshal(payload, &value); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}

	if err := schema.Validate(value); err != nil {
		t.Fatalf("schema.Validate() error = %v", err)
	}
}

func assertValidationError(t *testing.T, schema *jsonschema.Schema, event map[string]any) {
	t.Helper()

	payload, err := json.Marshal(event)
	if err != nil {
		t.Fatalf("json.Marshal() error = %v", err)
	}

	var value any
	if err := json.Unmarshal(payload, &value); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}

	if err := schema.Validate(value); err == nil {
		t.Fatal("schema.Validate() error = nil, want validation failure")
	}
}

func validEvent() map[string]any {
	return map[string]any{
		"id":        "mock-in-001",
		"source":    "athena",
		"type":      "athena.identified_presence.arrived",
		"timestamp": "2026-04-01T12:30:00Z",
		"data": map[string]any{
			"facility_id":            "ashtonbee",
			"zone_id":                "weight-room",
			"external_identity_hash": "tag_tracer2_001",
			"source":                 "mock",
			"recorded_at":            "2026-04-01T12:30:00Z",
		},
	}
}
