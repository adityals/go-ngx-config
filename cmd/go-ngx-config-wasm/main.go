package main

import (
	"syscall/js"
)

// TODO: add panic recover handler, is it possible??
func main() {
	println("[go-ngx-config-wasm] installed and ready to use!")

	c := make(chan struct{})

	registerCallbacks()

	<-c
}

func registerCallbacks() {
	js.Global().Set("goNgxParseConfig", parseConfigWrapper())
	js.Global().Set("goNgxTestLocation", testLocation())
}
