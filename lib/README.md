# Vendored native libraries

Each `<os>_<arch>/` directory holds the prebuilt static library `libpdf_ffi.a`
for one platform. `link_dist.go` statically links the one matching the host
`GOOS`/`GOARCH`, so `go get` + `go build` needs no external native library.

These archives are committed at every release tag (one per supported platform):

```
lib/darwin_amd64/libpdf_ffi.a
lib/darwin_arm64/libpdf_ffi.a
lib/linux_amd64/libpdf_ffi.a
lib/linux_arm64/libpdf_ffi.a
lib/windows_amd64/libpdf_ffi.a
```

They are built from the rust-pdf core with the production license key and copied
in by the release process.
