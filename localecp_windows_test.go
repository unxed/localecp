//go:build windows
// +build windows

package localecp

import "testing"

func TestInitSystemLocales_Windows(t *testing.T) {
	initSystemLocales()

	if ANSIDecoder == nil || ANSIEncoding == nil {
		t.Error("expected ANSI decoder and encoding to be set on Windows")
	}
	if OEMDecoder == nil || OEMEncoding == nil {
		t.Error("expected OEM decoder and encoding to be set on Windows")
	}
	if ANSICodepage == 0 || OEMCodepage == 0 {
		t.Errorf("expected codepage numbers, got ANSI=%d OEM=%d", ANSICodepage, OEMCodepage)
	}
	// On a machine set to a UTF-8 system codepage the raw calls say 65001;
	// the numbers exposed here must be the locale's legacy codepages, so
	// that "ANSI" and "OEM" still name a real single-byte codepage.
	if ANSICodepage == cpUTF8 || OEMCodepage == cpUTF8 {
		t.Errorf("codepage numbers must not be UTF-8: ANSI=%d OEM=%d", ANSICodepage, OEMCodepage)
	}
}

func TestSystemCodepages_NeverUTF8WhenLocaleKnows(t *testing.T) {
	ansi, oem := systemCodepages()
	if ansi == 0 || oem == 0 {
		t.Fatalf("systemCodepages() = %d, %d", ansi, oem)
	}
	// If the OS reports UTF-8 and also knows the legacy codepage, the
	// legacy one wins. On an older OS without the LCTYPE the raw value
	// stays, which the test tolerates only when the locale query fails.
	if ansi == cpUTF8 && legacyLocaleCodepage(localeIUseUTF8LegacyACP) != 0 {
		t.Errorf("ANSI stayed 65001 although the locale knows its legacy codepage")
	}
	if oem == cpUTF8 && legacyLocaleCodepage(localeIUseUTF8LegacyOEMCP) != 0 {
		t.Errorf("OEM stayed 65001 although the locale knows its legacy codepage")
	}
}
