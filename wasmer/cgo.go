// +build !custom_wasmer_runtime

package wasmer

// #cgo CFLAGS: -I${SRCDIR}/packaged/include
// #cgo LDFLAGS: -lwasmer
//
// #cgo linux,amd64 LDFLAGS: -L${SRCDIR}/packaged/lib/linux-amd64 -lm -ldl
// //#cgo linux,arm64 LDFLAGS: -Wl,-rpath,${SRCDIR}/packaged/lib/linux-aarch64 -L${SRCDIR}/packaged/lib/linux-aarch64
// #cgo darwin,amd64 LDFLAGS: -Wl,-rpath,${SRCDIR}/packaged/lib/darwin-amd64 -L${SRCDIR}/packaged/lib/darwin-amd64
// #cgo darwin,arm64 LDFLAGS: -Wl,-rpath,${SRCDIR}/packaged/lib/darwin-aarch64 -L${SRCDIR}/packaged/lib/darwin-aarch64
//
// #include <wasmer.h>
import "C"

// See https://github.com/golang/go/issues/26366.
import (
	_ "github.com/wasmerio/wasmer-go/wasmer/packaged/include"
	_ "github.com/wasmerio/wasmer-go/wasmer/packaged/lib"
)
