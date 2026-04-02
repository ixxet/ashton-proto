package events

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"time"

	athenav1 "github.com/ixxet/ashton-proto/gen/go/ashton/athena/v1"
	jsonschema "github.com/santhosh-tekuri/jsonschema/v6"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	ServiceAthena                    = "athena"
	SubjectIdentifiedPresenceArrived = "athena.identified_presence.arrived"
)

var (
	//go:embed athena.identified_presence.arrived.schema.json
	identifiedPresenceArrivedSchemaJSON []byte
	//go:embed envelope.schema.json
	envelopeSchemaJSON []byte

	identifiedPresenceArrivedSchemaOnce sync.Once
	compiledIdentifiedPresenceSchema    *jsonschema.Schema
	compiledIdentifiedPresenceSchemaErr error
)

type IdentifiedPresenceArrivedEvent struct {
	ID            string
	CorrelationID string
	Timestamp     time.Time
	Data          *athenav1.IdentifiedPresenceArrived
}

type identifiedPresenceArrivedEnvelope struct {
	ID            string                             `json:"id"`
	Source        string                             `json:"source"`
	Type          string                             `json:"type"`
	Timestamp     string                             `json:"timestamp"`
	CorrelationID string                             `json:"correlation_id,omitempty"`
	Data          identifiedPresenceArrivedEventData `json:"data"`
}

type identifiedPresenceArrivedEventData struct {
	FacilityID           string `json:"facility_id"`
	ZoneID               string `json:"zone_id,omitempty"`
	ExternalIdentityHash string `json:"external_identity_hash"`
	Source               string `json:"source"`
	RecordedAt           string `json:"recorded_at"`
}

var presenceSourceNames = map[athenav1.PresenceSource]string{
	athenav1.PresenceSource_PRESENCE_SOURCE_MOCK:     "mock",
	athenav1.PresenceSource_PRESENCE_SOURCE_RFID:     "rfid",
	athenav1.PresenceSource_PRESENCE_SOURCE_TOF:      "tof",
	athenav1.PresenceSource_PRESENCE_SOURCE_DATABASE: "database",
	athenav1.PresenceSource_PRESENCE_SOURCE_CSV:      "csv",
}

var presenceSourceValues = map[string]athenav1.PresenceSource{
	"mock":     athenav1.PresenceSource_PRESENCE_SOURCE_MOCK,
	"rfid":     athenav1.PresenceSource_PRESENCE_SOURCE_RFID,
	"tof":      athenav1.PresenceSource_PRESENCE_SOURCE_TOF,
	"database": athenav1.PresenceSource_PRESENCE_SOURCE_DATABASE,
	"csv":      athenav1.PresenceSource_PRESENCE_SOURCE_CSV,
}

func MarshalIdentifiedPresenceArrived(event IdentifiedPresenceArrivedEvent) ([]byte, error) {
	if strings.TrimSpace(event.ID) == "" {
		return nil, fmt.Errorf("identified presence arrived event missing id")
	}
	if event.Timestamp.IsZero() {
		return nil, fmt.Errorf("identified presence arrived event missing timestamp")
	}
	if event.Data == nil {
		return nil, fmt.Errorf("identified presence arrived event missing data")
	}

	recordedAt := event.Data.GetRecordedAt()
	if recordedAt == nil {
		return nil, fmt.Errorf("identified presence arrived event missing recorded_at")
	}
	if err := recordedAt.CheckValid(); err != nil {
		return nil, fmt.Errorf("identified presence arrived event recorded_at: %w", err)
	}

	source, err := PresenceSourceName(event.Data.GetSource())
	if err != nil {
		return nil, fmt.Errorf("identified presence arrived event source: %w", err)
	}

	envelope := identifiedPresenceArrivedEnvelope{
		ID:            event.ID,
		Source:        ServiceAthena,
		Type:          SubjectIdentifiedPresenceArrived,
		Timestamp:     event.Timestamp.UTC().Format(time.RFC3339Nano),
		CorrelationID: event.CorrelationID,
		Data: identifiedPresenceArrivedEventData{
			FacilityID:           event.Data.GetFacilityId(),
			ZoneID:               event.Data.GetZoneId(),
			ExternalIdentityHash: event.Data.GetExternalIdentityHash(),
			Source:               source,
			RecordedAt:           recordedAt.AsTime().UTC().Format(time.RFC3339Nano),
		},
	}

	payload, err := json.Marshal(envelope)
	if err != nil {
		return nil, err
	}

	if err := ValidateIdentifiedPresenceArrivedJSON(payload); err != nil {
		return nil, err
	}

	return payload, nil
}

func ParseIdentifiedPresenceArrived(payload []byte) (IdentifiedPresenceArrivedEvent, error) {
	if err := ValidateIdentifiedPresenceArrivedJSON(payload); err != nil {
		return IdentifiedPresenceArrivedEvent{}, err
	}

	var envelope identifiedPresenceArrivedEnvelope
	if err := json.Unmarshal(payload, &envelope); err != nil {
		return IdentifiedPresenceArrivedEvent{}, err
	}

	timestamp, err := time.Parse(time.RFC3339Nano, envelope.Timestamp)
	if err != nil {
		return IdentifiedPresenceArrivedEvent{}, fmt.Errorf("identified presence arrived event timestamp: %w", err)
	}

	recordedAt, err := time.Parse(time.RFC3339Nano, envelope.Data.RecordedAt)
	if err != nil {
		return IdentifiedPresenceArrivedEvent{}, fmt.Errorf("identified presence arrived event recorded_at: %w", err)
	}

	source, err := ParsePresenceSourceName(envelope.Data.Source)
	if err != nil {
		return IdentifiedPresenceArrivedEvent{}, fmt.Errorf("identified presence arrived event source: %w", err)
	}

	return IdentifiedPresenceArrivedEvent{
		ID:            envelope.ID,
		CorrelationID: envelope.CorrelationID,
		Timestamp:     timestamp,
		Data: &athenav1.IdentifiedPresenceArrived{
			FacilityId:           envelope.Data.FacilityID,
			ZoneId:               envelope.Data.ZoneID,
			ExternalIdentityHash: envelope.Data.ExternalIdentityHash,
			Source:               source,
			RecordedAt:           timestamppb.New(recordedAt.UTC()),
		},
	}, nil
}

func ValidateIdentifiedPresenceArrivedJSON(payload []byte) error {
	schema, err := identifiedPresenceArrivedSchema()
	if err != nil {
		return err
	}

	var value any
	if err := json.Unmarshal(payload, &value); err != nil {
		return fmt.Errorf("invalid identified presence arrived payload: %w", err)
	}

	if err := schema.Validate(value); err != nil {
		return fmt.Errorf("identified presence arrived schema validation: %w", err)
	}

	var envelope identifiedPresenceArrivedEnvelope
	if err := json.Unmarshal(payload, &envelope); err != nil {
		return fmt.Errorf("invalid identified presence arrived payload: %w", err)
	}

	if _, err := time.Parse(time.RFC3339Nano, envelope.Timestamp); err != nil {
		return fmt.Errorf("identified presence arrived event timestamp: %w", err)
	}
	if _, err := time.Parse(time.RFC3339Nano, envelope.Data.RecordedAt); err != nil {
		return fmt.Errorf("identified presence arrived event recorded_at: %w", err)
	}
	if _, err := ParsePresenceSourceName(envelope.Data.Source); err != nil {
		return fmt.Errorf("identified presence arrived event source: %w", err)
	}

	return nil
}

func PresenceSourceName(source athenav1.PresenceSource) (string, error) {
	name, ok := presenceSourceNames[source]
	if !ok {
		return "", fmt.Errorf("unsupported presence source %q", source.String())
	}

	return name, nil
}

func ParsePresenceSourceName(name string) (athenav1.PresenceSource, error) {
	source, ok := presenceSourceValues[strings.TrimSpace(name)]
	if !ok {
		return athenav1.PresenceSource_PRESENCE_SOURCE_UNSPECIFIED, fmt.Errorf("unsupported presence source %q", name)
	}

	return source, nil
}

func identifiedPresenceArrivedSchema() (*jsonschema.Schema, error) {
	identifiedPresenceArrivedSchemaOnce.Do(func() {
		compiler := jsonschema.NewCompiler()
		if err := addEmbeddedSchemaResource(compiler, schemaURL("athena.identified_presence.arrived.schema.json"), identifiedPresenceArrivedSchemaJSON); err != nil {
			compiledIdentifiedPresenceSchemaErr = err
			return
		}
		if err := addEmbeddedSchemaResource(compiler, schemaURL("envelope.schema.json"), envelopeSchemaJSON); err != nil {
			compiledIdentifiedPresenceSchemaErr = err
			return
		}

		compiledIdentifiedPresenceSchema, compiledIdentifiedPresenceSchemaErr = compiler.Compile(schemaURL("athena.identified_presence.arrived.schema.json"))
	})

	return compiledIdentifiedPresenceSchema, compiledIdentifiedPresenceSchemaErr
}

func addEmbeddedSchemaResource(compiler *jsonschema.Compiler, url string, schemaJSON []byte) error {
	document, err := jsonschema.UnmarshalJSON(bytes.NewReader(schemaJSON))
	if err != nil {
		return err
	}

	return compiler.AddResource(url, document)
}

func schemaURL(name string) string {
	return "https://github.com/ixxet/ashton-proto/events/" + name
}
