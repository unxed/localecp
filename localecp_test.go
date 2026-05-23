package localecp

import (
	"bytes"
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

func TestDecoding_ActualBytes(t *testing.T) {
	// Test CP866 decoding
	cp866Enc := getEncodingByName("IBM866")
	if cp866Enc == nil {
		t.Fatal("IBM866 not found")
	}
	dec := cp866Enc.NewDecoder()
	// "Привет" in CP866: 0x8f, 0xe0, 0xa8, 0xa2, 0xa5, 0xe2
	raw := []byte{0x8f, 0xe0, 0xa8, 0xa2, 0xa5, 0xe2}
	decoded, err := dec.Bytes(raw)
	if err != nil {
		t.Fatalf("failed to decode CP866: %v", err)
	}
	if string(decoded) != "Привет" {
		t.Errorf("expected Привет, got %s", string(decoded))
	}

	// Test Windows-1251 decoding
	cp1251Enc := getEncodingByName("WINDOWS-1251")
	if cp1251Enc == nil {
		t.Fatal("WINDOWS-1251 not found")
	}
	dec1251 := cp1251Enc.NewDecoder()
	// "Привет" in Windows-1251: 0xcf, 0xf0, 0xe8, 0xe2, 0xe5, 0xf2
	raw1251 := []byte{0xcf, 0xf0, 0xe8, 0xe2, 0xe5, 0xf2}
	decoded1251, err := dec1251.Bytes(raw1251)
	if err != nil {
		t.Fatalf("failed to decode Windows-1251: %v", err)
	}
	if string(decoded1251) != "Привет" {
		t.Errorf("expected Привет, got %s", string(decoded1251))
	}
}
func TestEncoding_ActualBytes(t *testing.T) {
	// Test CP1251 encoding
	cp1251Enc := getEncodingByName("WINDOWS-1251")
	if cp1251Enc == nil {
		t.Fatal("WINDOWS-1251 not found")
	}
	encoder := cp1251Enc.NewEncoder()
	// "Привет" in Windows-1251: 0xcf, 0xf0, 0xe8, 0xe2, 0xe5, 0xf2
	encoded, err := encoder.Bytes([]byte("Привет"))
	if err != nil {
		t.Fatalf("failed to encode Windows-1251: %v", err)
	}
	expected := []byte{0xcf, 0xf0, 0xe8, 0xe2, 0xe5, 0xf2}
	if !bytes.Equal(encoded, expected) {
		t.Errorf("expected %v, got %v", expected, encoded)
	}
}
