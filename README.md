# localecp

`localecp` (Locale Codepage) is a zero-dependency Go micro-library that deduces active legacy system codepages (OEM and ANSI) based on the host operating system's locale.

It provides a clean, independent way of managing character set translation without any bindings to a specific archive container format.

## Features

* **Platform-Aware:** Detects system locales on Unix/macOS (via `LANG`/`LC_ALL` variables) and Windows (using `GetACP` and `GetOEMCP` API calls).
* **Pre-configured Decoders:** Exposes active `OEMDecoder`, `ANSIDecoder`, and `SystemDecoder` ready to use with standard text processing.
* **No Heavy Dependencies:** Relies solely on `golang.org/x/text`.

## Installation

```bash
go get github.com/unxed/localecp
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/unxed/localecp"
)

func main() {
	rawBytes := []byte{0x8f, 0xe0, 0xa8, 0xa2, 0xa5, 0xe2} // "Привет" in CP866

	// Decode using the automatically initialized OEM decoder
	decoded, _ := localecp.OEMDecoder.Bytes(rawBytes)
	fmt.Println(string(decoded))
}
```