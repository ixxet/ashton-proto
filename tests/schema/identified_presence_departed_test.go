package schema_test

import (
	"encoding/json"
	"path/filepath"
	"runtime"
	"testing"

	protoevents "github.com/ixxet/ashton-proto/events"
	jsonschema "github.com/santhosh-tekuri/jsonschema/v6"
)

func TestIdentifiedPresenceDepartedSchemaAcceptsValidEnvelope(t *testing.T) {
	schema := loadDepartedSchema(t)

	validatePayload(t, schema, protoevents.ValidIdentifiedPresenceDepartedFixture())
}

func TestIdentifiedPresenceDepartedSchemaRejectsMissingFacilityID(t *testing.T) {
	schema := loadDepartedSchema(t)

	assertValidationError(t, schema, protoevents.MissingFacilityIDIdentifiedPresenceDepartedFixture())
}

func TestIdentifiedPresenceDepartedSchemaRejectsMissingRecordedAt(t *testing.T) {
	schema := loadDepartedSchema(t)

	event := validDepartedEvent(t)
	delete(event["data"].(map[string]any), "recorded_at")
	assertValidationError(t, schema, marshalEvent(t, event))
}

func TestIdentifiedPresenceDepartedSchemaRejectsMissingExternalIdentityHash(t *testing.T) {
	schema := loadDepartedSchema(t)

	event := validDepartedEvent(t)
	delete(event["data"].(map[string]any), "external_identity_hash")
	assertValidationError(t, schema, marshalEvent(t, event))
}

func loadDepartedSchema(t *testing.T) *jsonschema.Schema {
	t.Helper()

	_, currentFile, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("runtime.Caller() = false")
	}

	eventsDir := filepath.Join(filepath.Dir(currentFile), "..", "..", "events")
	schemaPath := filepath.Join(eventsDir, "athena.identified_presence.departed.schema.json")
	envelopePath := filepath.Join(eventsDir, "envelope.schema.json")
	compiler := jsonschema.NewCompiler()
	if err := addResource(compiler, "https://github.com/ixxet/ashton-proto/events/athena.identified_presence.departed.schema.json", schemaPath); err != nil {
		t.Fatalf("addResource(schema) error = %v", err)
	}
	if err := addResource(compiler, "https://github.com/ixxet/ashton-proto/events/envelope.schema.json", envelopePath); err != nil {
		t.Fatalf("addResource(envelope) error = %v", err)
	}

	schema, err := compiler.Compile("https://github.com/ixxet/ashton-proto/events/athena.identified_presence.departed.schema.json")
	if err != nil {
		t.Fatalf("compiler.Compile() error = %v", err)
	}

	return schema
}

func validDepartedEvent(t *testing.T) map[string]any {
	t.Helper()

	var event map[string]any
	if err := json.Unmarshal(protoevents.ValidIdentifiedPresenceDepartedFixture(), &event); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}

	return event
}
