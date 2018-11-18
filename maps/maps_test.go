package maps

import (
	"log"
	"testing"
)

func TestParsingMaps(t *testing.T) {
	m, err := New("./testfiles/maps")

	if err != nil {
		t.Fatal("Got error", err)
	}

	if m == nil {
		t.Fatal("maps is missing")
	}
	log.Println("maps", m)

	if len(m) != 19 {
		t.Fatal("Expected 19 entries")
	}
}
