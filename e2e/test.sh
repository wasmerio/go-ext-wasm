#!/bin/bash
rm -r vendor || echo ''
mkdir -p vendor/github.com/wasmerio/wasmer-go
cp -r ../wasmer vendor/github.com/wasmerio/wasmer-go
GODEBUG=cgocheck=2 go build -o app
ldd app || echo ''
otool -L app || echo ''
./app
