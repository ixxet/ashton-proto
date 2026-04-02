package events

import (
	"testing"
	"time"

	athenav1 "github.com/ixxet/ashton-proto/gen/go/ashton/athena/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestMarshalAndParseIdentifiedPresenceDepartedRoundTrip(t *testing.T) {
	event := IdentifiedPresenceDepartedEvent{
		ID:        "mock-out-001",
		Timestamp: time.Date(2026, 4, 1, 12, 45, 0, 0, time.UTC),
		Data: &athenav1.IdentifiedPresenceDeparted{
			FacilityId:           "ashtonbee",
			ZoneId:               "weight-room",
			ExternalIdentityHash: "tag_tracer5_001",
			Source:               athenav1.PresenceSource_PRESENCE_SOURCE_MOCK,
			RecordedAt:           timestamppb.New(time.Date(2026, 4, 1, 12, 45, 0, 0, time.UTC)),
		},
	}

	payload, err := MarshalIdentifiedPresenceDeparted(event)
	if err != nil {
		t.Fatalf("MarshalIdentifiedPresenceDeparted() error = %v", err)
	}

	if string(payload) != string(ValidIdentifiedPresenceDepartedFixture()) {
		t.Fatalf("MarshalIdentifiedPresenceDeparted() payload = %s, want fixture %s", payload, ValidIdentifiedPresenceDepartedFixture())
	}

	parsed, err := ParseIdentifiedPresenceDeparted(payload)
	if err != nil {
		t.Fatalf("ParseIdentifiedPresenceDeparted() error = %v", err)
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

func TestValidateIdentifiedPresenceDepartedFixtures(t *testing.T) {
	tests := []struct {
		name    string
		payload []byte
		wantErr bool
	}{
		{
			name:    "valid",
			payload: ValidIdentifiedPresenceDepartedFixture(),
		},
		{
			name:    "invalid source",
			payload: InvalidSourceIdentifiedPresenceDepartedFixture(),
			wantErr: true,
		},
		{
			name:    "invalid recorded_at",
			payload: InvalidRecordedAtIdentifiedPresenceDepartedFixture(),
			wantErr: true,
		},
		{
			name:    "missing facility_id",
			payload: MissingFacilityIDIdentifiedPresenceDepartedFixture(),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateIdentifiedPresenceDepartedJSON(tt.payload)
			if tt.wantErr && err == nil {
				t.Fatal("ValidateIdentifiedPresenceDepartedJSON() error = nil, want error")
			}
			if !tt.wantErr && err != nil {
				t.Fatalf("ValidateIdentifiedPresenceDepartedJSON() error = %v", err)
			}
		})
	}
}
