<p align="center">
  <a href="https://wasmer.io" target="_blank" rel="noopener">
    <img width="300" src="https://raw.githubusercontent.com/wasmerio/wasmer/master/logo.png" alt="Wasmer logo">
  </a>
</p>

<p align="center">
  <a href="https://spectrum.chat/wasmer">
    <img src="https://withspectrum.github.io/badge/badge.svg" alt="Join the Wasmer Community" valign="middle"></a>
  <a href="https://godoc.org/github.com/wasmerio/go-ext-wasm/wasmer">
    <img src="https://godoc.org/github.com/wasmerio/go-ext-wasm?status.svg" alt="Read our API documentation" valign="middle"></a>
  <a href="https://goreportcard.com/report/github.com/wasmerio/go-ext-wasm">
    <img src="https://goreportcard.com/badge/github.com/wasmerio/go-ext-wasm" alt="Go Report Card" valign="middle" /></a>
  <a href="https://github.com/wasmerio/wasmer/blob/master/LICENSE">
    <img src="https://img.shields.io/github/license/wasmerio/wasmer.svg" alt="License" valign="middle"></a>
</p>

Wasmer is a Go library for executing WebAssembly binaries.

# Install

To install the library, follow the classical:

```sh
# Enable cgo
$ export CGO_ENABLED=1; export CC=gcc;

$ go get github.com/wasmerio/go-ext-wasm/wasmer
```

It will work on many macOS and Linux distributions. It will not work
on Windows yet, we are working on it.

If the pre-compiled shared libraries are not compatible with your system, try installing manually with the following command:

```sh
$ just build-runtime
$ just build
$ go install github.com/wasmerio/go-ext-wasm/wasmer
```

# Documentation

[The documentation can be read online on godoc.org][documentation]. It
contains function descriptions, short examples, long examples
etc. Everything one need to start using Wasmer with Go!

Also, there is this article written for the announcement that
introduces the project: [Announcing the fastest WebAssembly runtime
for Go: wasmer][medium].

[documentation]: https://godoc.org/github.com/wasmerio/go-ext-wasm/wasmer
[medium]: https://medium.com/wasmer/announcing-the-fastest-webassembly-runtime-for-go-wasmer-19832d77c050

# Examples

## Basic example: Exported function

There is a toy program in `wasmer/test/testdata/examples/simple.rs`,
written in Rust (or any other language that compiles to WebAssembly):

```rust
#[no_mangle]
pub extern fn sum(x: i32, y: i32) -> i32 {
    x + y
}
```

After compilation to WebAssembly, the
[`wasmer/test/testdata/examples/simple.wasm`](https://github.com/wasmerio/go-ext-wasm/blob/master/wasmer/test/testdata/examples/simple.wasm)
binary file is generated. ([Download
it](https://github.com/wasmerio/go-ext-wasm/raw/master/wasmer/test/testdata/examples/simple.wasm)).

Then, we can execute it in Go:

```go
package main

import (
	"fmt"
	wasm "github.com/wasmerio/go-ext-wasm/wasmer"
)

func main() {
	// Reads the WebAssembly module as bytes.
	bytes, _ := wasm.ReadBytes("simple.wasm")
	
	// Instantiates the WebAssembly module.
	instance, _ := wasm.NewInstance(bytes)
	defer instance.Close()

	// Gets the `sum` exported function from the WebAssembly instance.
	sum := instance.Exports["sum"]

	// Calls that exported function with Go standard values. The WebAssembly
	// types are inferred and values are casted automatically.
	result, _ := sum(5, 37)

	fmt.Println(result) // 42!
}
```

## Imported function

A WebAssembly module can export functions, this is how to run a
WebAssembly function, like we did in the previous example with
`instance.Exports["sum"](1, 2)`. Nonetheless, a WebAssembly module can
depend on “extern functions”, then called imported functions. For
instance, let's consider the basic following Rust program:

```rust
extern {
    fn sum(x: i32, y: i32) -> i32;
}

#[no_mangle]
pub extern fn add1(x: i32, y: i32) -> i32 {
    unsafe { sum(x, y) + 1 }
}
```

In this case, the `add1` function is a WebAssembly exported function,
whilst the `sum` function is a WebAssembly imported function (the
WebAssembly instance needs to _import_ it to complete the
program). Good news: We can write the implementation of the `sum`
function directly in Go!

First, we need to declare the `sum` function signature in C inside a
Go comment (with the help of [cgo]):

```go
package main

// #include <stdlib.h>
//
// extern int32_t sum(void *context, int32_t x, int32_t y);
import "C"
```

Second, we declare the `sum` function implementation in Go. Notice the
`//export` which is the way cgo uses to map Go code to C code.

```go
//export sum
func sum(context unsafe.Pointer, x int32, y int32) int32 {
	return x + y
}
```

Third, we use `NewImports` to create the WebAssembly imports. In this
code:

* `"sum"` is the imported function name,
* `sum` is the Go function pointer, and
* `C.sum` is the cgo function pointer.

```go
imports, _ := wasm.NewImports().Append("sum", sum, C.sum)
```

Finally, we use `NewInstanceWithImports` to inject the imports:

```go
bytes, _ := wasm.ReadBytes("imported_function.wasm")
instance, _ := wasm.NewInstanceWithImports(bytes, imports)
defer instance.Close()

// Gets and calls the `add1` exported function from the WebAssembly instance.
results, _ := instance.Exports["add1"](1, 2)

fmt.Println(result)
//   add1(1, 2)
// = sum(1 + 2) + 1
// = 1 + 2 + 1
// = 4
// QED
```

[cgo]: https://golang.org/cmd/cgo/

## Read the memory

A WebAssembly instance has a linear memory. Let's see how to read
it. Consider the following Rust program:

```rust
#[no_mangle]
pub extern fn return_hello() -> *const u8 {
    b"Hello, World!\0".as_ptr()
}
```

The `return_hello` function returns a pointer to a string. This string
is stored in the WebAssembly memory. Let's read it.

```go
bytes, _ := wasm.ReadBytes("memory.wasm")
instance, _ := wasm.NewInstance(bytes)
defer instance.Close()

// Calls the `return_hello` exported function.
// This function returns a pointer to a string.
result, _ := instance.Exports["return_hello"]()

// Gets the pointer value as an integer.
pointer := result.ToI32()

// Reads the memory.
memory := instance.Memory.Data()

fmt.Println(string(memory[pointer : pointer+13])) // Hello, World!
```

In this example, we already know the string length, and we use a slice
to read a portion of the memory directly. Notice that the string
terminates by a null byte, which means we could iterate over the
memory starting from `pointer` until a null byte is met; that's a
similar approach.

For a more complete example, see the [Greet Example][greet-example].

[greet-example]: https://godoc.org/github.com/wasmerio/go-ext-wasm/wasmer#example-package--Greet

# Use with bazel

To use this with bazel. Instead of installing it with `go get` (which will not compile with gazelle generated `BUILD.bazel`), add the following to your `WORKSPACE` file:

```
git_repository(
    name = "com_github_wasmerio_go_ext_wasm",
    remote = "https://github.com/wasmerio/go-ext-wasm",
    commit = "90c00ac78c6cc99bf1ed3755f828cc63e5f3920a", // replace with commits of your preference.
    patch_args = ["-p1"],
    patches = ["//external:com_github_wasmerio_go_ext_wasm.patch"],
)
```

and add `external/com_github_wasmerio_go_ext_wasm.patch` file with following:

```
diff --git a/wasmer/BUILD.bazel b/wasmer/BUILD.bazel
new file mode 100644
index 0000000..d3ee99a
--- /dev/null
+++ b/wasmer/BUILD.bazel
@@ -0,0 +1,36 @@
+package(default_visibility = ["//visibility:public"])
+load("@io_bazel_rules_go//go:def.bzl", "go_library")
+load("@rules_cc//cc:defs.bzl", "cc_library")
+
+cc_library(
+    name = "wasm",
+    srcs = glob(["*.h"]) + select({
+        "@io_bazel_rules_go//go/platform:darwin": ["libwasmer_runtime_c_api.dylib"],
+        "@io_bazel_rules_go//go/platform:linux_amd64": ["libwasmer_runtime_c_api.so"],
+        "@io_bazel_rules_go//go/platform:windows_amd64": ["lwasmer_runtime_c_api.dll"],
+    }),
+    hdrs = glob(["*.h"]),
+    visibility = ["//visibility:public"],
+    includes = ["."],
+)
+
+go_library(
+    name = "go_default_library",
+    srcs = [
+        "bridge.go",
+        "error.go",
+        "import.go",
+        "instance.go",
+        "memory.go",
+        "module.go",
+        "value.go",
+        "wasi.go",
+        "wasmer.go",
+        "wasmer.h",
+     ],
+    cgo = True,
+    clinkopts = ["-Wl,-rpath,wasmer -Lwasmer/wasmer -lwasmer_runtime_c_api"],
+    importpath = "github.com/wasmerio/go-ext-wasm/wasmer",
+    visibility = ["//visibility:public"],
+    cdeps = [":wasm"],
+)
```

# Development

The Go library is written in Go and Rust.

To build both parts, run the following commands:

```sh
$ just build-runtime
$ just build
```

To build the Go part, run:

```sh
$ just build
```

(Yes, you need [`just`]).

[`just`]: https://github.com/casey/just/

# Testing

Once the library is build, run the following command:

```sh
$ just test
```

# Benchmarks

We compared Wasmer to [Wagon][wagon] and [Life][life]. The benchmarks
are in `benchmarks/`. The computer that ran these benchmarks is a
MacBook Pro 15" from 2016, 2.9Ghz Core i7 with 16Gb of memory. Here
are the results in a table (the lower the ratio is, the better):

<table>
  <thead>
    <tr>
      <th>Benchmark</th>
      <th>Runtime</th>
      <th align="right">Time (ms)</th>
      <th align="right">Ratio</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td rowspan="3">N-Body</td>
      <td>Wasmer</td>
      <td align="right">42.078</td>
      <td align="right">1×</td>
    </tr>
    <tr>
      <td>Wagon</td>
      <td align="right">1841.950</td>
      <td align="right">44×</td>
    </tr>
    <tr>
      <td>Life</td>
      <td align="right">1976.215</td>
      <td align="right">47×</td>
    </tr>
    <tr>
      <td rowspan="3">Fibonacci (recursive)</td>
      <td>Wasmer</td>
      <td align="right">28.559</td>
      <td align="right">1×</td>
    </tr>
    <tr>
      <td>Wagon</td>
      <td align="right">3238.050</td>
      <td align="right">113×</td>
    </tr>
    <tr>
      <td>Life</td>
      <td align="right">3029.209</td>
      <td align="right">106×</td>
    </tr>
    <tr>
      <td rowspan="3">Pollard rho 128</td>
      <td>Wasmer</td>
      <td align="right">37.478</td>
      <td align="right">1×</td>
    </tr>
    <tr>
      <td>Wagon</td>
      <td align="right">2165.563</td>
      <td align="right">58×</td>
    </tr>
    <tr>
      <td>Life</td>
      <td align="right">2407.752</td>
      <td align="right">64×</td>
    </tr>
  </tbody>
</table>

While both Life and Wagon provide on average the same speed, Wasmer is
on average 72 times faster.

Put on a graph, it looks like this (reminder: the lower, the better):

![Benchmark results](https://cdn-images-1.medium.com/max/1200/1*08ymx9shShohcPCKi1XjlA.png)

[wagon]: https://github.com/go-interpreter/wagon
[life]: https://github.com/perlin-network/life

# What is WebAssembly?

Quoting [the WebAssembly site](https://webassembly.org/):

> WebAssembly (abbreviated Wasm) is a binary instruction format for a
> stack-based virtual machine. Wasm is designed as a portable target
> for compilation of high-level languages like C/C++/Rust, enabling
> deployment on the web for client and server applications.

About speed:

> WebAssembly aims to execute at native speed by taking advantage of
> [common hardware
> capabilities](https://webassembly.org/docs/portability/#assumptions-for-efficient-execution)
> available on a wide range of platforms.

About safety:

> WebAssembly describes a memory-safe, sandboxed [execution
> environment](https://webassembly.org/docs/semantics/#linear-memory) […].

# License

The entire project is under the MIT License. Please read [the
`LICENSE` file][license].


[license]: https://github.com/wasmerio/wasmer/blob/master/LICENSE
