//go:build !windows
// +build !windows

package localecp

import (
	"os"
	"testing"
)

func TestInitSystemLocales_Unix_Env(t *testing.T) {
	// Backup original env
	origLang := os.Getenv("LANG")
	origLcAll := os.Getenv("LC_ALL")
	origLcCtype := os.Getenv("LC_CTYPE")

	defer func() {
		os.Setenv("LANG", origLang)
		os.Setenv("LC_ALL", origLcAll)
		os.Setenv("LC_CTYPE", origLcCtype)
	}()

	// Test French locale
	os.Setenv("LC_ALL", "fr_FR.UTF-8")
	os.Setenv("LC_CTYPE", "")
	os.Setenv("LANG", "")

	initSystemLocales()

	// fr_FR should map to IBM850 (OEM) and WINDOWS-1252 (ANSI)
	// We check if the decoders are initialized (not nil)
	if OEMDecoder == nil {
		t.Error("expected OEMDecoder to be set")
	}
	if ANSIDecoder == nil {
		t.Error("expected ANSIDecoder to be set")
	}
}

func TestInitSystemLocales_Unix_POSIX(t *testing.T) {
	origLang := os.Getenv("LANG")
	origLcAll := os.Getenv("LC_ALL")
	origLcCtype := os.Getenv("LC_CTYPE")

	defer func() {
		os.Setenv("LANG", origLang)
		os.Setenv("LC_ALL", origLcAll)
		os.Setenv("LC_CTYPE", origLcCtype)
	}()

	// Тестируем fallback для C/POSIX локали
	os.Setenv("LC_ALL", "POSIX")
	os.Setenv("LC_CTYPE", "")
	os.Setenv("LANG", "")

	oldOEM := OEMDecoder
	initSystemLocales()

	if OEMDecoder != oldOEM {
		t.Error("OEMDecoder should not change for POSIX locale")
	}
}
