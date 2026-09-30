//go:build !extralite

package localecp

import (
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/htmlindex"
)

// multiByteEncoding resolves a charset name the single-byte table above does
// not know, through the WHATWG index: the East Asian encodings and the rest.
func multiByteEncoding(name string) encoding.Encoding {
	htmlName := name
	switch name {
	case "CP932":
		htmlName = "shift_jis"
	case "CP949":
		htmlName = "euc-kr"
	}
	enc, err := htmlindex.Get(htmlName)
	if err == nil {
		return enc
	}
	return nil
}
