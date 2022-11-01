package main

import (
	"encoding/json"
	"errors"

	"syscall/js"

	"github.com/adityals/go-ngx-config/internal/ast"
	"github.com/adityals/go-ngx-config/internal/matcher"
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
	js.Global().Set("testLocation", testLocation())
}

func parseConfig(confString string) (*ast.Config, error) {
	parser := parser.NewStringParser(confString)

	ast := parser.Parse()
	if ast == nil {
		return nil, errors.New("cannot be parsed")
	}

	return ast, nil
}

func parseConfigWrapper() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 1 {
			return errors.New("invalid arguments")
		}

		ngxConfStr := args[0].String()

		handler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			resolve := args[0]
			reject := args[1]

			go func() {
				ast, err := parseConfig(ngxConfStr)
				if err != nil {
					errorConstructor := js.Global().Get("Error")
					errorObject := errorConstructor.New(err.Error())
					reject.Invoke(errorObject)
					return
				}

				ast_json, err := json.MarshalIndent(ast, "", "  ")
				if err != nil {
					errorConstructor := js.Global().Get("Error")
					errorObject := errorConstructor.New(err.Error())
					reject.Invoke(errorObject)
					return
				}

				resolve.Invoke(string(ast_json))
			}()

			return nil
		})

		promiseConstructor := js.Global().Get("Promise")
		return promiseConstructor.New(handler)
	})
}

func testLocation() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 2 {
			return errors.New("arguments is invalid")
		}

		ngxConfStr := args[0].String()
		targetPath := args[1].String()

		handler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			resolve := args[0]
			reject := args[1]

			go func() {
				ast, err := parseConfig(ngxConfStr)

				if err != nil {
					errorConstructor := js.Global().Get("Error")
					errorObject := errorConstructor.New(err.Error())
					reject.Invoke(errorObject)
					return
				}

				match, err := matcher.NewLocationMatcher(ast, targetPath)
				if err != nil {
					errorConstructor := js.Global().Get("Error")
					errorObject := errorConstructor.New(err.Error())
					reject.Invoke(errorObject)
					return
				}

				match_json, err := json.MarshalIndent(match, "", "  ")
				if err != nil {
					// Handle errors here too
					errorConstructor := js.Global().Get("Error")
					errorObject := errorConstructor.New(err.Error())
					reject.Invoke(errorObject)
					return
				}

				resolve.Invoke(string(match_json))
			}()

			return nil
		})

		promiseConstructor := js.Global().Get("Promise")
		return promiseConstructor.New(handler)
	})

}
