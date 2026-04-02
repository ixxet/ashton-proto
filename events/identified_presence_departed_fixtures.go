package events

import (
	"bytes"
	_ "embed"
)

var (
	//go:embed testdata/athena.identified_presence.departed.valid.json
	validIdentifiedPresenceDepartedFixture []byte
	//go:embed testdata/athena.identified_presence.departed.invalid_source.json
	invalidSourceIdentifiedPresenceDepartedFixture []byte
	//go:embed testdata/athena.identified_presence.departed.invalid_recorded_at.json
	invalidRecordedAtIdentifiedPresenceDepartedFixture []byte
	//go:embed testdata/athena.identified_presence.departed.missing_facility_id.json
	missingFacilityIdentifiedPresenceDepartedFixture []byte
)

func ValidIdentifiedPresenceDepartedFixture() []byte {
	return bytes.TrimSpace(bytes.Clone(validIdentifiedPresenceDepartedFixture))
}

func InvalidSourceIdentifiedPresenceDepartedFixture() []byte {
	return bytes.TrimSpace(bytes.Clone(invalidSourceIdentifiedPresenceDepartedFixture))
}

func InvalidRecordedAtIdentifiedPresenceDepartedFixture() []byte {
	return bytes.TrimSpace(bytes.Clone(invalidRecordedAtIdentifiedPresenceDepartedFixture))
}

func MissingFacilityIDIdentifiedPresenceDepartedFixture() []byte {
	return bytes.TrimSpace(bytes.Clone(missingFacilityIdentifiedPresenceDepartedFixture))
}
