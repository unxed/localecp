//go:build windows
// +build windows

package localecp

import (
	"testing"
)

func TestInitSystemLocales_Windows(t *testing.T) {
	// Повторно вызываем инициализацию, чтобы покрыть ветки GetACP/GetOEMCP.
	// Ожидаем, что функция отработает без паник и установит декодеры.
	initSystemLocales()

	if ANSIDecoder == nil {
		t.Error("expected ANSIDecoder to be set on Windows")
	}
	if OEMDecoder == nil {
		t.Error("expected OEMDecoder to be set on Windows")
	}
}