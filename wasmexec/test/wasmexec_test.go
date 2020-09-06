package wasmexectests

import (
	"github.com/stretchr/testify/assert"
	wasm "github.com/wasmerio/go-ext-wasm/wasmer"
	"path"
	"runtime"
	"testing"
)

func GetBytes(file string) []byte {
	_, filename, _, _ := runtime.Caller(0)
	modulePath := path.Join(path.Dir(filename), "testdata", file)

	bytes, _ := wasm.ReadBytes(modulePath)

	return bytes
}

func TestTinyGo(t *testing.T) {
	imports, err := wasm.ImportsForGo()
	assert.NoError(t, err)

	defer imports.Close()

	instance, err := wasm.NewInstanceWithImports(GetBytes("tinygo.wasm"), imports)
	assert.NoError(t, err)

	defer instance.Close()

	sum, exists := instance.Exports["sum"]
	assert.Equal(t, exists, true)

	output, err := sum(1, 2)

	assert.Equal(t, wasm.I32(3), output)
	assert.NoError(t, err)
}
