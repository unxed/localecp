//go:build extralite

package localecp

import "testing"

// Under the extralite tag the East Asian names are not resolved; the
// single-byte table still works. The other tests of this package assume the
// full build and are not run with the tag.
func TestExtraLiteNoMultiByte(t *testing.T) {
	for _, name := range []string{"CP932", "CP949", "shift_jis", "euc-kr", "gbk"} {
		if getEncodingByName(name) != nil {
			t.Errorf("%s resolved under extralite", name)
		}
	}
	if getEncodingByName("IBM866") == nil {
		t.Error("IBM866 must still resolve under extralite")
	}
}
