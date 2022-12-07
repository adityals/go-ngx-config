GO_ROOT=$(shell go env GOROOT)

.PHONY: build
build:
	GOOS=linux GOARCH=amd64 go build -o ./bin/go-ngx-config ./cmd/go-ngx-config/*.go

.PHONY: build-wasm
build-wasm:
	GOOS=js GOARCH=wasm go build -o ./bin/go-ngx-config-parser.wasm ./cmd/go-ngx-config-wasm/*.go

.PHONY: clean-wasm-dev
clean-wasm-dev:
	rm ./examples/wasm/wasm_exec.js || true
	rm ./examples/wasm/go-ngx-config-parser.wasm || true

.PHONY: serve-wasm-dev
serve-wasm-dev: clean-wasm-dev build-wasm
	bash -c "cp ${GO_ROOT}/misc/wasm/wasm_exec.js ./examples/wasm/"
	@cp ./bin/go-ngx-config-parser.wasm  ./examples/wasm/
	pnpm serve --filter=wasm-example

.PHONY: init-web
init-web:
	pnpm i

.PHONY: prepare-web
prepare-web: build-wasm
	bash -c "cp ${GO_ROOT}/misc/wasm/wasm_exec.js ./web/public/"
	@cp ./bin/go-ngx-config-parser.wasm ./web/public

.PHONY: serve-web
serve-dev-web: prepare-web
	pnpm dev --filter=web
