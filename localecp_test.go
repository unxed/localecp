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
}