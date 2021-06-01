<div align="center">
  <a href="https://wasmer.io" target="_blank" rel="noopener noreferrer">
    <img width="300" src="https://raw.githubusercontent.com/wasmerio/wasmer/master/assets/logo.png" alt="Wasmer logo">
  </a>
  
  <h1>Wasmer Go</h1>
  
  <p>
    <a href="https://github.com/wasmerio/wasmer-go/actions?query=workflow%3A%22Build+and+Test%22">
      <img src="https://github.com/wasmerio/wasmer-go/workflows/Build%20and%20Test/badge.svg" alt="Build Status">
    </a>
    <a href="https://github.com/wasmerio/wasmer-go/blob/master/LICENSE">
      <img src="https://img.shields.io/github/license/wasmerio/wasmer-go.svg" alt="License">
    </a>
    <a href="https://pkg.go.dev/github.com/wasmerio/wasmer-go/wasmer">
      <img src="https://img.shields.io/badge/go.dev-package-f06" alt="Go Package">
    </a> 
    <a href="https://pkg.go.dev/github.com/wasmerio/wasmer-go/wasmer?tab=doc">
      <img src="https://img.shields.io/badge/documentation-API-f06" alt="API Documentation">
    </a> 
  </p>

  <h3>
    <a href="https://wasmer.io/">Website</a>
    <span> • </span>
    <a href="https://docs.wasmer.io">Docs</a>
    <span> • </span>
    <a href="https://slack.wasmer.io/">Slack Channel</a>
  </h3>

</div>

<hr/>

A complete and mature WebAssembly runtime for Go based on
[Wasmer](https://github.com/wasmerio/wasmer).

Features

  * **Easy to use**: The `wasmer` API mimics the standard WebAssembly API,
  * **Fast**: `wasmer` executes the WebAssembly modules as fast as
    possible, close to **native speed**,
  * **Safe**: All calls to WebAssembly will be fast, but more
    importantly, completely safe and sandboxed.
    
**Documentation**: [browse the detailed API
documentation](https://pkg.go.dev/github.com/wasmerio/wasmer-go/wasmer)
full of examples.

**Examples** as tutorials: [browse the `examples/`
directory](https://github.com/wasmerio/wasmer-go/tree/master/examples),
it's the best place for a complete introduction!

# Install

To install the library, follow the classical:

```sh
$ go get github.com/wasmerio/wasmer-go/wasmer
```

And you're ready to get fun!

## Supported platforms

This library embeds the Wasmer runtime compiled as shared library
objects, and so uses [`cgo`](https://golang.org/cmd/cgo/) to consume
it. A set of precompiled shared library objects are provided. Thus
this library works (and is tested) on the following platforms:

<table>
  <thead>
    <tr>
      <th>Platform</th>
      <th>Architecture</th>
      <th>Triple</th>
      <th>Status</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td rowspan="2">Linux</td>
      <td><code>amd64</code></td>
      <td><code>x86_64-unknown-linux-gnu</code></td>
      <td>✅</td>
    </tr>
    <tr>
      <td><code>aarch64</code></td>
      <td><code>aarch64-unknown-linux-gnu</code></td>
      <td>✅</td>
    </tr>
    <tr>
      <td rowspan="2">Darwin</td>
      <td><code>amd64</code></td>
      <td><code>x86_64-apple-darwin</code></td>
      <td>✅</td>
    </tr>
    <tr>
      <td><code>aarch64</code></td>
      <td><code>aarch64-apple-darwin</code></td>
      <td>⏳</td>
    </tr>
    <tr>
      <td>Windows</td>
      <td><code>amd64</code></td>
      <td><code>x86_64-pc-windows-msvc</code></td>
      <td>⏳</td>
    </tr>
  </tbody>
</table>

<details>
  <summary>What to do if your platform is missing?</summary>
  
  Up to now, there is no script to automate that process. [We are
  working on it](https://github.com/wasmerio/wasmer-go/issues/191).

  Here are the steps to do that manually:
  
  ```sh
  $ # Build the new Wasmer C API shared object library.
  $ cargo build --release
  $
  $ # Configure cgo.
  $ export CGO_CFLAGS="-I$(pwd)/wasmer/packaged/include/"
  $ export CGO_LDFLAGS="-Wl,-rpath,$(pwd)/target/release/ -L$(pwd)/target/release/ -lwasmer_go"
  $
  $ # Run the tests.
  $ just test -tags custom_wasmer_runtime
  ```
</details>

# Examples

We highly recommend to read the
[`examples/`](https://github.com/wasmerio/wasmer-go/tree/master/examples)
directory, which contains a sequence of examples/tutorials. It's the
best place to learn by reading examples.

But for the most eager of you, there is a quick toy program in
`examples/appendices/simple.go`, written in Rust:

```rust
#[no_mangle]
pub extern "C" fn sum(x: i32, y: i32) -> i32 {
    x + y
}
```

A compiled WebAssembly binary is included in
[`examples/appendices/simple.wasm`]((https://github.com/wasmerio/wasmer-go/blob/master/examples/appendices/simple.wasm)).

Then, we can execute it in Go:

```go
package main

import (
	"fmt"
	"io/ioutil"
	wasmer "github.com/wasmerio/wasmer-go/wasmer"
)

func main() {
    wasmBytes, _ := ioutil.ReadFile("simple.wasm")

    engine := wasmer.NewEngine()
    store := wasmer.NewStore(engine)

    // Compiles the module
    module, _ := wasmer.NewModule(store, wasmBytes)

    // Instantiates the module
    importObject := wasmer.NewImportObject()
    instance, _ := wasmer.NewInstance(module, importObject)

    // Gets the `sum` exported function from the WebAssembly instance.
    sum, _ := instance.Exports.GetFunction("sum")

    // Calls that exported function with Go standard values. The WebAssembly
    // types are inferred and values are casted automatically.
    result, _ := sum(5, 37)

    fmt.Println(result) // 42!
}
```

And then, finally, enjoy by running:

```sh
$ cd examples/appendices/
$ go run simple.go
42
```

# Testing

Run the tests with the following command:

```sh
$ just test
```

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
