//go:build windows
// +build windows

package localecp

import (
	"fmt"
	"syscall"
	"unsafe"

	"golang.org/x/text/encoding/htmlindex"
)

const (
	cpUTF8 = 65001

	// LCTYPEs for GetLocaleInfoEx. On a machine where the user has turned
	// on "Use Unicode UTF-8 for worldwide language support", GetACP and
	// GetOEMCP both answer 65001; these two answer the legacy codepages the
	// same locale would have had without that setting (winnls.h).
	localeIUseUTF8LegacyACP   = 0x0666
	localeIUseUTF8LegacyOEMCP = 0x0999
	localeReturnNumber        = 0x20000000
)

var (
	kernel32        = syscall.NewLazyDLL("kernel32.dll")
	procGetACP      = kernel32.NewProc("GetACP")
	procGetOEMCP    = kernel32.NewProc("GetOEMCP")
	procLocaleInfoW = kernel32.NewProc("GetLocaleInfoEx")
)

// legacyLocaleCodepage asks the user's default locale for the codepage it
// would have used before UTF-8 became the system codepage. 0 when the OS
// does not know (pre-1903 Windows) or the call fails.
func legacyLocaleCodepage(lctype uintptr) int {
	if err := procLocaleInfoW.Find(); err != nil {
		return 0
	}
	var value uint32
	// LOCALE_NAME_USER_DEFAULT is a NULL name; with LOCALE_RETURN_NUMBER the
	// buffer receives a DWORD and cchData counts WCHARs, so 2.
	r, _, _ := procLocaleInfoW.Call(0, lctype|localeReturnNumber, uintptr(unsafe.Pointer(&value)), 2)
	if r == 0 {
		return 0
	}
	return int(value)
}

// systemCodepages returns the ANSI and OEM codepage numbers for this
// machine the way a user thinks of them: the legacy codepages of the locale.
// On a UTF-8 system both raw calls say 65001, which is true for the process
// but useless as an "ANSI" or "OEM" identity -- Far shows 1252/850 on such a
// machine, and so does this, via the same locale query (real_ansi/real_oem
// in Far's codepage.cpp).
func systemCodepages() (ansi, oem int) {
	if acp, _, _ := procGetACP.Call(); acp != 0 {
		ansi = int(acp)
	}
	if oemcp, _, _ := procGetOEMCP.Call(); oemcp != 0 {
		oem = int(oemcp)
	}
	if ansi == cpUTF8 {
		if legacy := legacyLocaleCodepage(localeIUseUTF8LegacyACP); legacy != 0 {
			ansi = legacy
		}
	}
	if oem == cpUTF8 {
		if legacy := legacyLocaleCodepage(localeIUseUTF8LegacyOEMCP); legacy != 0 {
			oem = legacy
		}
	}
	return ansi, oem
}

func initSystemLocales() {
	ansi, oem := systemCodepages()
	ANSICodepage, OEMCodepage = ansi, oem

	if ansi != 0 {
		if enc, err := htmlindex.Get(fmt.Sprintf("windows-%d", ansi)); err == nil {
			ANSIDecoder = enc.NewDecoder()
			ANSIEncoder = enc.NewEncoder()
			ANSIEncoding = enc
		}
	}
	if oem != 0 {
		if enc, err := htmlindex.Get(fmt.Sprintf("cp%d", oem)); err == nil {
			OEMDecoder = enc.NewDecoder()
			SystemDecoder = OEMDecoder
			OEMEncoding = enc
		}
	}
}
