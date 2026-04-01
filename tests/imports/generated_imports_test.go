package imports_test

import (
	"testing"

	athenav1 "github.com/ixxet/ashton-proto/gen/go/ashton/athena/v1"
	commonv1 "github.com/ixxet/ashton-proto/gen/go/ashton/common/v1"
)

func TestGeneratedPackagesCompile(t *testing.T) {
	if (&commonv1.CheckRequest{}) == nil {
		t.Fatal("expected common health request type to be importable")
	}

	response := &athenav1.GetCurrentOccupancyResponse{
		Occupancy: &athenav1.OccupancyState{},
	}
	if response.Occupancy == nil {
		t.Fatal("expected athena occupancy response type to be importable")
	}

	arrival := &athenav1.IdentifiedPresenceArrived{}
	if arrival == nil {
		t.Fatal("expected identified presence arrival type to be importable")
	}
}
