package main

import (
	"encoding/json"
	"errors"
	"syscall/js"

	"github.com/adityals/go-ngx-config/internal/crossplane"
	"github.com/adityals/go-ngx-config/pkg/matcher"
	"github.com/adityals/go-ngx-config/pkg/parser"
)

type parseOptsConf struct {
	SkipCtx bool
}

func getParseOptsConf(mapArg js.Value) parseOptsConf {
	skipCtx := mapArg.Get("skipCtx")

	return parseOptsConf{
		SkipCtx: skipCtx.Bool(),
	}
}

func parseConfig(confString string, opts parseOptsConf) (*crossplane.Payload, error) {
	// ! we mark single file as true because from web cannot locate another file path and still getting from string
	parsed, err := parser.NewNgxConfStringParser(confString, &crossplane.ParseOptions{
		SingleFile:                true,
		StopParsingOnError:        true,
		SkipDirectiveContextCheck: opts.SkipCtx,
	})
	if err != nil {
		return nil, err
	}

	return parsed, nil
}

func parseConfigWrapper() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 2 {
			return errors.New("invalid arguments")
		}

		ngxConfStr := args[0].String()
		parseOpts := getParseOptsConf(args[1])

		handler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			resolve := args[0]
			reject := args[1]

			go func() {
				ast, err := parseConfig(ngxConfStr, parseOpts)
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
		if len(args) != 3 {
			return errors.New("arguments is invalid")
		}

		ngxConfStr := args[0].String()
		targetPath := args[1].String()
		parseOpts := getParseOptsConf(args[2])

		handler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			resolve := args[0]
			reject := args[1]

			go func() {
				ast, err := parseConfig(ngxConfStr, parseOpts)

				if err != nil {
					errorConstructor := js.Global().Get("Error")
					errorObject := errorConstructor.New(err.Error())
					reject.Invoke(errorObject)
					return
				}

				match, err := matcher.NewLocationMatcherFromPayload(ast, targetPath)
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
