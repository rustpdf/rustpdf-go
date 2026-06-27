// Native linkage: statically link the prebuilt libpdf_ffi.a vendored under
// lib/<os>_<arch>/, so `go get` + `go build` works with no external native
// library. The .a for each platform is committed at every release tag.
//
// The trailing system libraries are the static dependencies the Rust staticlib
// pulls in. To override the library (e.g. link a locally built one), set the
// cgo flags yourself, e.g. CGO_LDFLAGS="-L/path -lpdf_ffi ...".
package rustpdf

/*
#cgo darwin,amd64  LDFLAGS: -L${SRCDIR}/lib/darwin_amd64 -lpdf_ffi -liconv -lm
#cgo darwin,arm64  LDFLAGS: -L${SRCDIR}/lib/darwin_arm64 -lpdf_ffi -liconv -lm
#cgo linux,amd64   LDFLAGS: -L${SRCDIR}/lib/linux_amd64 -lpdf_ffi -lpthread -ldl -lm -lrt
#cgo linux,arm64   LDFLAGS: -L${SRCDIR}/lib/linux_arm64 -lpdf_ffi -lpthread -ldl -lm -lrt
#cgo windows,amd64 LDFLAGS: -L${SRCDIR}/lib/windows_amd64 -lpdf_ffi -lws2_32 -luserenv -lbcrypt -lntdll -ladvapi32 -lkernel32
*/
import "C"
