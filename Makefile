.PHONY: build
build:
	go build -o ./bin/go-ngx-config ./cmd/go-ngx-config/*.go

.PHONY: build-wasm
build-wasm:
	GOOS=js GOARCH=wasm go build -o ./bin/go-ngx-config-parser.wasm ./cmd/go-ngx-config-wasm/main.go

.PHONY: clean-wasm-dev
clean-wasm-dev:
	rm ./examples/wasm/wasm_exec.js || true
	rm ./examples/wasm/go-ngx-config-parser.wasm || true

.PHONY: serve-wasm-dev
serve-wasm-dev: clean-wasm-dev build-wasm
	bash -c "cp /usr/local/Cellar/go@1.14/1.14.15/libexec/misc/wasm/wasm_exec.js ./examples/wasm/"
	@cp ./bin/go-ngx-config-parser.wasm  ./examples/wasm/
	pnpm serve --filter=wasm-example