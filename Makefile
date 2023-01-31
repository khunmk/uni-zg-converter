build:
	rm -rf cmd/app/static/main.wasm
	GOOS=js GOARCH=wasm go build -o cmd/app/static/main.wasm cmd/wasm/main.go