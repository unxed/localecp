package localecp

import (
	"testing"
)

func TestDefaultDecoders(t *testing.T) {
	if OEMDecoder == nil {
		t.Error("expected OEMDecoder to be initialized")
	}
	if ANSIDecoder == nil {
		t.Error("expected ANSIDecoder to be initialized")
	}
	if SystemDecoder == nil {
		t.Error("expected SystemDecoder to be initialized")
	}
}

// TestTableIntegrity verifies that every string encoding name defined in the
// translation tables maps to a valid, resolvable text encoding in x/text.
func TestTableIntegrity(t *testing.T) {
	for locale, oemName := range lcToOemTable {
		enc := getEncodingByName(oemName)
		if enc == nil {
			t.Errorf("locale %q: OEM encoding %q is not resolvable", locale, oemName)
		}
	}

	for locale, ansiName := range lcToAnsiTable {
		enc := getEncodingByName(ansiName)
		if enc == nil {
			t.Errorf("locale %q: ANSI encoding %q is not resolvable", locale, ansiName)
		}
	}
}

func TestGetEncodingByName_Fallback(t *testing.T) {
	enc := getEncodingByName("UNKNOWN_CODEPAGE")
	if enc != nil {
		t.Error("expected nil for unknown codepage, got valid encoding")
	}
}