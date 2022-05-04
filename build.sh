#!/bin/sh
GOOS=js GOARCH=wasm go build -o public/main.wasm cmd/backend/main.go