package events

import (
	"bytes"
	_ "embed"
)

var (
	//go:embed testdata/athena.identified_presence.arrived.valid.json
	validIdentifiedPresenceArrivedFixture []byte
	//go:embed testdata/athena.identified_presence.arrived.invalid_source.json
	invalidSourceIdentifiedPresenceArrivedFixture []byte
	//go:embed testdata/athena.identified_presence.arrived.invalid_recorded_at.json
	invalidRecordedAtIdentifiedPresenceArrivedFixture []byte
	//go:embed testdata/athena.identified_presence.arrived.missing_facility_id.json
	missingFacilityIdentifiedPresenceArrivedFixture []byte
)

func ValidIdentifiedPresenceArrivedFixture() []byte {
	return bytes.TrimSpace(bytes.Clone(validIdentifiedPresenceArrivedFixture))
}

func InvalidSourceIdentifiedPresenceArrivedFixture() []byte {
	return bytes.TrimSpace(bytes.Clone(invalidSourceIdentifiedPresenceArrivedFixture))
}

func InvalidRecordedAtIdentifiedPresenceArrivedFixture() []byte {
	return bytes.TrimSpace(bytes.Clone(invalidRecordedAtIdentifiedPresenceArrivedFixture))
}

func MissingFacilityIDIdentifiedPresenceArrivedFixture() []byte {
	return bytes.TrimSpace(bytes.Clone(missingFacilityIdentifiedPresenceArrivedFixture))
}
