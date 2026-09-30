//go:build extralite

package localecp

import "golang.org/x/text/encoding"

// multiByteEncoding knows no further encodings under the extralite tag: the
// WHATWG index links every encoding table of x/text (about 0.6 MB of East
// Asian ones), which a size-constrained build (f4's extra-lite profile) leaves
// out. A locale whose charset is not in the single-byte table keeps the
// default encoding, as for any unknown name.
func multiByteEncoding(string) encoding.Encoding { return nil }
