# Build the runtime shared library for this specific system.
build-runtime:
	#!/usr/bin/env bash
	set -euo pipefail

	# Build the shared library.
	cargo build --release

	# Find the shared library extension.
	case "{{os()}}" in
		"macos")
			shared_library_path=$( ls -t target/release/deps/libwasmer_runtime_c_api*.dylib | head -n 1 )
			shared_library=libwasmer_runtime_c_api.dylib
			;;
		"windows")
			shared_library_path=$( ls -t target/release/deps/wasmer_runtime_c_api*.dll | head -n 1 )
			shared_library=wasmer_runtime_c_api.dll
			;;
		*)
			shared_library_path=$( ls -t target/release/deps/libwasmer_runtime_c_api*.so | head -n 1 )
			shared_library=libwasmer_runtime_c_api.so
	esac

	# Link `wasmer/*wasmer_runtime_c_api.*`.
	rm -f wasmer/${shared_library}
	ln -s "../${shared_library_path}" wasmer/${shared_library}

# Build the `wasmer` library.
build go-build-args='-v':
	cd wasmer && go build {{go-build-args}} .

# Build the `go-wasmer` bin.
build-bin go-build-args='-v':
	cd go-wasmer && go build {{go-build-args}} -o ../target/go/go-wasmer .

# Compile the Rust part for your specific system.
rust:
	cargo build --release

# Generate cgo debug objects.
debug-cgo:
	cd wasmer/ && \
		go tool cgo bridge.go && \
		echo "Browse wasmer/_obj/"

# Run all the tests.
test: test-wasmer

# Run all the tests of the `wasmer` package.
test-wasmer:
	#!/usr/bin/env bash
	cd wasmer
	# Run the tests.
	GODEBUG=cgocheck=2 go test -test.v $( find test -type f \( -name "*_test.go" \! -name "example_*.go" \! -name "benchmark*.go" \) ) test/imports.go
	# Run the short examples.
	go test -test.v example_test.go
	# Run the long examples.
	go test -test.v $(find . -type f \( -name "example_*_test.go" \! -name "_example_import_test.go" \) )

# Run all the tests of the `wasmexec` package.
test-wasmexec:
	#!/usr/bin/env bash
	cd wasmexec
	GODEBUG=cgocheck=2 go test -test.v $( find test -type f \( -name "*_test.go" \) )

# Run benchmarks. Subjects can be `wasmer`, `wagon` or `life`. Filter is a regex to select the benchmarks.
bench subject='wasmer' filter='.*':
	cd benchmarks/{{subject}} && go test -bench '{{filter}}' benchmarks_test.go

# Local Variables:
# mode: makefile
# End:
