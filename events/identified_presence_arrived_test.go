package events

import (
	"testing"
	"time"

	athenav1 "github.com/ixxet/ashton-proto/gen/go/ashton/athena/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestMarshalAndParseIdentifiedPresenceArrivedRoundTrip(t *testing.T) {
	event := IdentifiedPresenceArrivedEvent{
		ID:        "mock-in-001",
		Timestamp: time.Date(2026, 4, 1, 12, 30, 0, 0, time.UTC),
		Data: &athenav1.IdentifiedPresenceArrived{
			FacilityId:           "ashtonbee",
			ZoneId:               "weight-room",
			ExternalIdentityHash: "tag_tracer2_001",
			Source:               athenav1.PresenceSource_PRESENCE_SOURCE_MOCK,
			RecordedAt:           timestamppb.New(time.Date(2026, 4, 1, 12, 30, 0, 0, time.UTC)),
		},
	}

	payload, err := MarshalIdentifiedPresenceArrived(event)
	if err != nil {
		t.Fatalf("MarshalIdentifiedPresenceArrived() error = %v", err)
	}

	if string(payload) != string(ValidIdentifiedPresenceArrivedFixture()) {
		t.Fatalf("MarshalIdentifiedPresenceArrived() payload = %s, want fixture %s", payload, ValidIdentifiedPresenceArrivedFixture())
	}

	parsed, err := ParseIdentifiedPresenceArrived(payload)
	if err != nil {
		t.Fatalf("ParseIdentifiedPresenceArrived() error = %v", err)
	}

	if parsed.ID != event.ID {
		t.Fatalf("parsed.ID = %q, want %q", parsed.ID, event.ID)
	}
	if !parsed.Timestamp.Equal(event.Timestamp) {
		t.Fatalf("parsed.Timestamp = %s, want %s", parsed.Timestamp, event.Timestamp)
	}
	if parsed.Data.GetSource() != event.Data.GetSource() {
		t.Fatalf("parsed.Data.GetSource() = %q, want %q", parsed.Data.GetSource().String(), event.Data.GetSource().String())
	}
}

func TestValidateIdentifiedPresenceArrivedFixtures(t *testing.T) {
	tests := []struct {
		name    string
		payload []byte
		wantErr bool
	}{
		{
			name:    "valid",
			payload: ValidIdentifiedPresenceArrivedFixture(),
		},
		{
			name:    "invalid source",
			payload: InvalidSourceIdentifiedPresenceArrivedFixture(),
			wantErr: true,
		},
		{
			name:    "invalid recorded_at",
			payload: InvalidRecordedAtIdentifiedPresenceArrivedFixture(),
			wantErr: true,
		},
		{
			name:    "missing facility_id",
			payload: MissingFacilityIDIdentifiedPresenceArrivedFixture(),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateIdentifiedPresenceArrivedJSON(tt.payload)
			if tt.wantErr && err == nil {
				t.Fatal("ValidateIdentifiedPresenceArrivedJSON() error = nil, want error")
			}
			if !tt.wantErr && err != nil {
				t.Fatalf("ValidateIdentifiedPresenceArrivedJSON() error = %v", err)
			}
		})
	}
}

func TestPresenceSourceMappingRejectsUnsupportedValues(t *testing.T) {
	if _, err := PresenceSourceName(athenav1.PresenceSource_PRESENCE_SOURCE_UNSPECIFIED); err == nil {
		t.Fatal("PresenceSourceName() error = nil, want unsupported source error")
	}

	if _, err := ParsePresenceSourceName("infrared"); err == nil {
		t.Fatal("ParsePresenceSourceName() error = nil, want unsupported source error")
	}
}
