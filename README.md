# rustpdf-go

Idiomatic Go binding for the **rust-pdf** core over its C ABI (`libpdf_ffi`), via
**cgo**. It covers the whole product surface: vector graphics, embedded/subsetted
fonts and text, wrapping paragraphs, images, **PDF/A** (levels 1b–3a),
**tagged/accessible** output, embedded-file attachments, **AcroForm** fields,
manipulation (merge/split/rotate/optimize/incremental update), **text
extraction**, **encryption** (RC4 / AES-128 / AES-256) and **digital signatures**
(PKCS#7 / PAdES) — plus **feature licensing**.

This repository is a **distribution mirror**: it carries the Go sources plus a
prebuilt static `libpdf_ffi.a` for each platform, so it installs and builds with
no native toolchain beyond a C compiler. The source of truth is the rust-pdf
core.

## Install

```sh
go get github.com/rustpdf/rustpdf-go@latest
```

The module is **self-contained**: `pdf.h` is vendored and a prebuilt static
`libpdf_ffi.a` for your platform ships under `lib/<os>_<arch>/`, so the build
statically links it — nothing external to install. cgo (a C toolchain +
`CGO_ENABLED=1`, the default) is the only requirement.

Supported platforms: `darwin/arm64`, `darwin/amd64`, `linux/amd64`,
`linux/arm64`, `windows/amd64`.

## Quick start

```go
package main

import (
	"fmt"
	"os"

	rustpdf "github.com/rustpdf/rustpdf-go"
)

func main() {
	// A token in RUSTPDF_LICENSE is auto-activated; or call ActivateLicense.
	// _ = rustpdf.ActivateLicense(token)

	d, _ := rustpdf.New()
	defer d.Close()
	_ = d.AddPage()
	_ = d.SetFillRGB(0.1, 0.2, 0.8)
	_ = d.Rect(72, 72, 200, 120)
	_ = d.Fill()
	data, _ := d.ToBytes()
	_ = os.WriteFile("out.pdf", data, 0o644)
	fmt.Println("wrote", len(data), "bytes; rustpdf", rustpdf.Version())
}
```

Corporate features (PDF/A, signing, encryption, accessibility) require a license;
without one they return an `*Error`. Set `RUSTPDF_LICENSE` (the token) or
`RUSTPDF_LICENSE_FILE` (a path) in the environment and it is auto-activated — no
code change needed.

## API

* `rustpdf.go` — package-level funcs (`Version`, `ActivateLicense`,
  `ExtractText`, `Sign`, `Timestamp`, `AddDss`), enums, the `Error` type, helpers;
* `document.go` — the `Document` authoring type;
* `editable.go` — the `EditableDoc` manipulation type;
* `link_dist.go` — the per-platform cgo linker flags.

## License

Proprietary — see [LICENSE](LICENSE). Use of the binding and the bundled
`libpdf_ffi` is governed by the rust-pdf product license.
