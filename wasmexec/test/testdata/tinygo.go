// Compiled with: `tinygo build -o tinygo.wasm -target=wasm tinygo.go`

package main

func main() {}

//go:export sum
func sum(x int, y int) int {
	return x + y
}
