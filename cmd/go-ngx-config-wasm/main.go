package main

import (
	"encoding/json"
	"errors"

	"syscall/js"

	"github.com/adityals/go-ngx-config/internal/parser"
)

// TODO: add panic handler

func main() {
	println("[go-ngx-config-wasm] installed")

	c := make(chan struct{})

	registerCallbacks()

	<-c
}

func registerCallbacks() {
	js.Global().Set("parseConfig", parseConfigWrapper())
}

func parseConfigWrapper() js.Func {
	parseFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 1 {
			return errors.New("invalid no of arguments passed")
		}

		stringArg := args[0].String()
		parser := parser.NewStringParser(stringArg)

		ast := parser.Parse()
		if ast == nil {
			return errors.New("cannot be parsed")
		}

		ast_json, err := json.MarshalIndent(ast, "", "  ")
		if err != nil {
			return err
		}

		return string(ast_json)
	})

	return parseFunc
}
