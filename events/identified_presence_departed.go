package events

import (
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

const SubjectIdentifiedPresenceDeparted = "athena.identified_presence.departed"

var (
	//go:embed athena.identified_presence.departed.schema.json
	identifiedPresenceDepartedSchemaJSON []byte

	identifiedPresenceDepartedSchemaOnce sync.Once
	compiledIdentifiedPresenceDepartedSchema    *jsonschema.Schema
	compiledIdentifiedPresenceDepartedSchemaErr error
)

type IdentifiedPresenceDepartedEvent struct {
	ID            string
	CorrelationID string
	Timestamp     time.Time
	Data          *athenav1.IdentifiedPresenceDeparted
}

type identifiedPresenceDepartedEnvelope struct {
	ID            string                              `json:"id"`
	Source        string                              `json:"source"`
	Type          string                              `json:"type"`
	Timestamp     string                              `json:"timestamp"`
	CorrelationID string                              `json:"correlation_id,omitempty"`
	Data          identifiedPresenceDepartedEventData `json:"data"`
}

type identifiedPresenceDepartedEventData struct {
	FacilityID           string `json:"facility_id"`
	ZoneID               string `json:"zone_id,omitempty"`
	ExternalIdentityHash string `json:"external_identity_hash"`
	Source               string `json:"source"`
	RecordedAt           string `json:"recorded_at"`
}

func MarshalIdentifiedPresenceDeparted(event IdentifiedPresenceDepartedEvent) ([]byte, error) {
	if strings.TrimSpace(event.ID) == "" {
		return nil, fmt.Errorf("identified presence departed event missing id")
	}
	if event.Timestamp.IsZero() {
		return nil, fmt.Errorf("identified presence departed event missing timestamp")
	}
	if event.Data == nil {
		return nil, fmt.Errorf("identified presence departed event missing data")
	}

	recordedAt := event.Data.GetRecordedAt()
	if recordedAt == nil {
		return nil, fmt.Errorf("identified presence departed event missing recorded_at")
	}
	if err := recordedAt.CheckValid(); err != nil {
		return nil, fmt.Errorf("identified presence departed event recorded_at: %w", err)
	}

	source, err := PresenceSourceName(event.Data.GetSource())
	if err != nil {
		return nil, fmt.Errorf("identified presence departed event source: %w", err)
	}

	envelope := identifiedPresenceDepartedEnvelope{
		ID:            event.ID,
		Source:        ServiceAthena,
		Type:          SubjectIdentifiedPresenceDeparted,
		Timestamp:     event.Timestamp.UTC().Format(time.RFC3339Nano),
		CorrelationID: event.CorrelationID,
		Data: identifiedPresenceDepartedEventData{
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

	if err := ValidateIdentifiedPresenceDepartedJSON(payload); err != nil {
		return nil, err
	}

	return payload, nil
}

func ParseIdentifiedPresenceDeparted(payload []byte) (IdentifiedPresenceDepartedEvent, error) {
	if err := ValidateIdentifiedPresenceDepartedJSON(payload); err != nil {
		return IdentifiedPresenceDepartedEvent{}, err
	}

	var envelope identifiedPresenceDepartedEnvelope
	if err := json.Unmarshal(payload, &envelope); err != nil {
		return IdentifiedPresenceDepartedEvent{}, err
	}

	timestamp, err := time.Parse(time.RFC3339Nano, envelope.Timestamp)
	if err != nil {
		return IdentifiedPresenceDepartedEvent{}, fmt.Errorf("identified presence departed event timestamp: %w", err)
	}

	recordedAt, err := time.Parse(time.RFC3339Nano, envelope.Data.RecordedAt)
	if err != nil {
		return IdentifiedPresenceDepartedEvent{}, fmt.Errorf("identified presence departed event recorded_at: %w", err)
	}

	source, err := ParsePresenceSourceName(envelope.Data.Source)
	if err != nil {
		return IdentifiedPresenceDepartedEvent{}, fmt.Errorf("identified presence departed event source: %w", err)
	}

	return IdentifiedPresenceDepartedEvent{
		ID:            envelope.ID,
		CorrelationID: envelope.CorrelationID,
		Timestamp:     timestamp,
		Data: &athenav1.IdentifiedPresenceDeparted{
			FacilityId:           envelope.Data.FacilityID,
			ZoneId:               envelope.Data.ZoneID,
			ExternalIdentityHash: envelope.Data.ExternalIdentityHash,
			Source:               source,
			RecordedAt:           timestamppb.New(recordedAt.UTC()),
		},
	}, nil
}

func ValidateIdentifiedPresenceDepartedJSON(payload []byte) error {
	schema, err := identifiedPresenceDepartedSchema()
	if err != nil {
		return err
	}

	var value any
	if err := json.Unmarshal(payload, &value); err != nil {
		return fmt.Errorf("invalid identified presence departed payload: %w", err)
	}

	if err := schema.Validate(value); err != nil {
		return fmt.Errorf("identified presence departed schema validation: %w", err)
	}

	var envelope identifiedPresenceDepartedEnvelope
	if err := json.Unmarshal(payload, &envelope); err != nil {
		return fmt.Errorf("invalid identified presence departed payload: %w", err)
	}

	if _, err := time.Parse(time.RFC3339Nano, envelope.Timestamp); err != nil {
		return fmt.Errorf("identified presence departed event timestamp: %w", err)
	}
	if _, err := time.Parse(time.RFC3339Nano, envelope.Data.RecordedAt); err != nil {
		return fmt.Errorf("identified presence departed event recorded_at: %w", err)
	}
	if _, err := ParsePresenceSourceName(envelope.Data.Source); err != nil {
		return fmt.Errorf("identified presence departed event source: %w", err)
	}

	return nil
}

func identifiedPresenceDepartedSchema() (*jsonschema.Schema, error) {
	identifiedPresenceDepartedSchemaOnce.Do(func() {
		compiler := jsonschema.NewCompiler()
		if err := addEmbeddedSchemaResource(compiler, schemaURL("athena.identified_presence.departed.schema.json"), identifiedPresenceDepartedSchemaJSON); err != nil {
			compiledIdentifiedPresenceDepartedSchemaErr = err
			return
		}
		if err := addEmbeddedSchemaResource(compiler, schemaURL("envelope.schema.json"), envelopeSchemaJSON); err != nil {
			compiledIdentifiedPresenceDepartedSchemaErr = err
			return
		}

		compiledIdentifiedPresenceDepartedSchema, compiledIdentifiedPresenceDepartedSchemaErr = compiler.Compile(schemaURL("athena.identified_presence.departed.schema.json"))
	})

	return compiledIdentifiedPresenceDepartedSchema, compiledIdentifiedPresenceDepartedSchemaErr
}
