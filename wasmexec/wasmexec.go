// Package wasmexec is a collection of host functions to support
// WebAssembly modules compiled by the TinyGo or the Go compilers.
package wasmexec

// #include <stdlib.h>
//
// // For TinyGo.
// extern int32_t env_io_get_stdout(void *context);
// extern int32_t env_resource_write(void *context, int32_t a, int32_t b, int32_t c);
//
// // For Go.
// extern void go_runtime_wasmExit(void *context, int32_t pointer);
// extern void go_runtime_wasmWrite(void *context, int32_t pointer);
// extern void go_runtime_nanotime(void *context, int32_t pointer);
// extern void go_runtime_walltime(void *context, int32_t pointer);
// extern void go_runtime_scheduleTimeoutEvent(void *context, int32_t pointer);
// extern void go_runtime_clearTimeoutEvent(void *context, int32_t pointer);
// extern void go_runtime_getRandomData(void *context, int32_t pointer);
// extern void go_syscall_js_stringVal(void *context, int32_t pointer);
// extern void go_syscall_js_valueGet(void *context, int32_t pointer);
// extern void go_syscall_js_valueSet(void *context, int32_t pointer);
// extern void go_syscall_js_valueIndex(void *context, int32_t pointer);
// extern void go_syscall_js_valueSetIndex(void *context, int32_t pointer);
// extern void go_syscall_js_valueCall(void *context, int32_t pointer);
// extern void go_syscall_js_valueInvoke(void *context, int32_t pointer);
// extern void go_syscall_js_valueNew(void *context, int32_t pointer);
// extern void go_syscall_js_valueLength(void *context, int32_t pointer);
// extern void go_syscall_js_valuePrepareString(void *context, int32_t pointer);
// extern void go_syscall_js_valueLoadString(void *context, int32_t pointer);
// extern void go_syscall_js_valueInstanceOf(void *context, int32_t pointer);
import "C"

import (
	"unsafe"
)

////
// For TinyGo.
////

//export env_io_get_stdout
func env_io_get_stdout(context unsafe.Pointer) int32 {
	panic("not implemented")
}

//export env_resource_write
func env_resource_write(context unsafe.Pointer, a int32, b int32, c int32) int32 {
	panic("not implemented")
}

////
// For Go.
////

//export go_runtime_wasmExit
func go_runtime_wasmExit(context unsafe.Pointer, pointer int32) {
	panic("not implemented")
}

//export go_runtime_wasmWrite
func go_runtime_wasmWrite(context unsafe.Pointer, pointer int32) {
	panic("not implemented")
}

//export go_runtime_nanotime
func go_runtime_nanotime(context unsafe.Pointer, pointer int32) {
	panic("not implemented")
}

//export go_runtime_walltime
func go_runtime_walltime(context unsafe.Pointer, pointer int32) {
	panic("not implemented")
}

//export go_runtime_scheduleTimeoutEvent
func go_runtime_scheduleTimeoutEvent(context unsafe.Pointer, pointer int32) {
	panic("not implemented")
}

//export go_runtime_clearTimeoutEvent
func go_runtime_clearTimeoutEvent(context unsafe.Pointer, pointer int32) {
	panic("not implemented")
}

//export go_runtime_getRandomData
func go_runtime_getRandomData(context unsafe.Pointer, pointer int32) {
	panic("not implemented")
}

//export go_syscall_js_stringVal
func go_syscall_js_stringVal(context unsafe.Pointer, pointer int32) {
	panic("not implemented")
}

//export go_syscall_js_valueGet
func go_syscall_js_valueGet(context unsafe.Pointer, pointer int32) {
	panic("not implemented")
}

//export go_syscall_js_valueSet
func go_syscall_js_valueSet(context unsafe.Pointer, pointer int32) {
	panic("not implemented")
}

//export go_syscall_js_valueIndex
func go_syscall_js_valueIndex(context unsafe.Pointer, pointer int32) {
	panic("not implemented")
}

//export go_syscall_js_valueSetIndex
func go_syscall_js_valueSetIndex(context unsafe.Pointer, pointer int32) {
	panic("not implemented")
}

//export go_syscall_js_valueCall
func go_syscall_js_valueCall(context unsafe.Pointer, pointer int32) {
	panic("not implemented")
}

//export go_syscall_js_valueInvoke
func go_syscall_js_valueInvoke(context unsafe.Pointer, pointer int32) {
	panic("not implemented")
}

//export go_syscall_js_valueNew
func go_syscall_js_valueNew(context unsafe.Pointer, pointer int32) {
	panic("not implemented")
}

//export go_syscall_js_valueLength
func go_syscall_js_valueLength(context unsafe.Pointer, pointer int32) {
	panic("not implemented")
}

//export go_syscall_js_valuePrepareString
func go_syscall_js_valuePrepareString(context unsafe.Pointer, pointer int32) {
	panic("not implemented")
}

//export go_syscall_js_valueLoadString
func go_syscall_js_valueLoadString(context unsafe.Pointer, pointer int32) {
	panic("not implemented")
}

//export go_syscall_js_valueInstanceOf
func go_syscall_js_valueInstanceOf(context unsafe.Pointer, pointer int32) {
	panic("not implemented")
}

// HostFunction is a pair (go, c) host function implementations.
type HostFunction struct {
	// Go implementation.
	Go interface{}

	// C implementation.
	C unsafe.Pointer
}

// HostFunctionNames is a type alias representing a collection of host
// function with respective names.
type HostFunctionNames = map[string]HostFunction

// HostFunctionNamespaces is a type alias representing a collection of
// host function with respective namespaces.
type HostFunctionNamespaces = map[string]HostFunctionNames

// hostFunctions is a private variable holding all the host functions.
var hostFunctions = HostFunctionNamespaces{
	"env": {
		"io_get_stdout": HostFunction{
			Go: env_io_get_stdout,
			C:  C.env_io_get_stdout,
		},
		"resource_write": HostFunction{
			Go: env_resource_write,
			C:  C.env_resource_write,
		},
	},
	"go": {
		"runtime.wasmExit": HostFunction{
			Go: go_runtime_wasmExit,
			C:  C.go_runtime_wasmExit,
		},
		"runtime.wasmWrite": HostFunction{
			Go: go_runtime_wasmWrite,
			C:  C.go_runtime_wasmWrite,
		},
		"runtime.nanotime": HostFunction{
			Go: go_runtime_nanotime,
			C:  C.go_runtime_nanotime,
		},
		"runtime.walltime": HostFunction{
			Go: go_runtime_walltime,
			C:  C.go_runtime_walltime,
		},
		"runtime.scheduleTimeoutEvent": HostFunction{
			Go: go_runtime_scheduleTimeoutEvent,
			C:  C.go_runtime_scheduleTimeoutEvent,
		},
		"runtime.clearTimeoutEvent": HostFunction{
			Go: go_runtime_clearTimeoutEvent,
			C:  C.go_runtime_clearTimeoutEvent,
		},
		"runtime.getRandomData": HostFunction{
			Go: go_runtime_getRandomData,
			C:  C.go_runtime_getRandomData,
		},
		"syscall/js.stringVal": HostFunction{
			Go: go_syscall_js_stringVal,
			C:  C.go_syscall_js_stringVal,
		},
		"syscall/js.valueGet": HostFunction{
			Go: go_syscall_js_valueGet,
			C:  C.go_syscall_js_valueGet,
		},
		"syscall/js.valueSet": HostFunction{
			Go: go_syscall_js_valueSet,
			C:  C.go_syscall_js_valueSet,
		},
		"syscall/js.valueIndex": HostFunction{
			Go: go_syscall_js_valueIndex,
			C:  C.go_syscall_js_valueIndex,
		},
		"syscall/js.valueSetIndex": HostFunction{
			Go: go_syscall_js_valueSetIndex,
			C:  C.go_syscall_js_valueSetIndex,
		},
		"syscall/js.valueCall": HostFunction{
			Go: go_syscall_js_valueCall,
			C:  C.go_syscall_js_valueCall,
		},
		"syscall/js.valueInvoke": HostFunction{
			Go: go_syscall_js_valueInvoke,
			C:  C.go_syscall_js_valueInvoke,
		},
		"syscall/js.valueNew": HostFunction{
			Go: go_syscall_js_valueNew,
			C:  C.go_syscall_js_valueNew,
		},
		"syscall/js.valueLength": HostFunction{
			Go: go_syscall_js_valueLength,
			C:  C.go_syscall_js_valueLength,
		},
		"syscall/js.valuePrepareString": HostFunction{
			Go: go_syscall_js_valuePrepareString,
			C:  C.go_syscall_js_valuePrepareString,
		},
		"syscall/js.valueLoadString": HostFunction{
			Go: go_syscall_js_valueLoadString,
			C:  C.go_syscall_js_valueLoadString,
		},
		"syscall/js.valueInstanceOf": HostFunction{
			Go: go_syscall_js_valueInstanceOf,
			C:  C.go_syscall_js_valueInstanceOf,
		},
	},
}

// HostFunctions returns all the host functions for TinyGo and Go.
func HostFunctions() HostFunctionNamespaces {
	return hostFunctions
}
