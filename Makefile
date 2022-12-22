GO_ROOT=$(shell go env GOROOT)
PLATFORMS=darwin/amd64 linux/amd64

.PHONY: build-platforms
build-platforms:
	for p in $(PLATFORMS); do \
		platform_split=($${p//\// }); \
		os=$${platform_split[0]}; \
		arch=$${platform_split[1]}; \
		output_name=go-ngx-config-$$os-$$arch; \
		GOOS=$$os GOARCH=$$arch go build -o bin/$$output_name ./cmd/go-ngx-config/*.go;\
	done

.PHONY: build-dev
build-dev:
	go build -o ./bin/go-ngx-config ./cmd/go-ngx-config/*.go

.PHONY: build-wasm
build-wasm:
	GOOS=js GOARCH=wasm go build -o ./bin/go-ngx-config.wasm ./cmd/go-ngx-config-wasm/*.go

.PHONY: clean-wasm-dev
clean-wasm-dev:
	rm ./examples/wasm/wasm_exec.js || true
	rm ./examples/wasm/go-ngx-config.wasm || true

.PHONY: serve-wasm-dev
serve-wasm-dev: clean-wasm-dev build-wasm
	bash -c "cp ${GO_ROOT}/misc/wasm/wasm_exec.js ./examples/wasm/"
	@cp ./bin/go-ngx-config.wasm  ./examples/wasm/
	pnpm serve --filter=wasm-example

.PHONY: init-web
init-web:
	pnpm i

.PHONY: prepare-web
prepare-web: build-wasm
	bash -c "cp ${GO_ROOT}/misc/wasm/wasm_exec.js ./web/public/"
	@cp ./bin/go-ngx-config.wasm ./web/public

.PHONY: serve-web
serve-dev-web: prepare-web
	pnpm dev --filter=web
