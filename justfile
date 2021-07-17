# Run all the tests.
test *ARGS:
	#!/usr/bin/env bash
	cd wasmer
	# Run the tests.
	GODEBUG=cgocheck=2 go test -v {{ARGS}}

examples:
	#!/usr/bin/env bash
	cd examples
	# Run the examples.
	GODEBUG=cgocheck=2 go test -v

e2e:
	#!/usr/bin/env bash
	cd e2e
	./test.sh

# Preview the documentation (needs `godoc`, see `go get -v golang.org/x/tools/cmd/godoc`).
preview-doc ADDRESS="0.0.0.0" PORT="9090":
	@echo "Starting godoc preview..."
	@echo "Documentation preview will be available at: http://{{ADDRESS}}:{{PORT}}/pkg/github.com/wasmerio/wasmer-go/"
	godoc -http={{ADDRESS}}:{{PORT}} -timestamps

# Local Variables:
# mode: makefile
# End:
