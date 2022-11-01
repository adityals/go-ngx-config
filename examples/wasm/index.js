export const wasmBrowserInstantiate = async (wasmModuleUrl, importObject) => {
    let response = undefined;
  
    // Check if the browser supports streaming instantiation
    if (WebAssembly.instantiateStreaming) {
      // Fetch the module, and instantiate it as it is downloading
      response = await WebAssembly.instantiateStreaming(
        fetch(wasmModuleUrl),
        importObject
      );
    } else {
      // Fallback to using fetch to download the entire module
      // And then instantiate the module
      const fetchAndInstantiateTask = async () => {
        const wasmArrayBuffer = await fetch(wasmModuleUrl).then(response =>
          response.arrayBuffer()
        );
        return WebAssembly.instantiate(wasmArrayBuffer, importObject);
      };
  
      response = await fetchAndInstantiateTask();
    }
  
    return response;
};

const go = new Go(); // Defined in wasm_exec.js. Don't forget to add this in your index.html.

const registerWasm = async () => {
    try {
    // Get the importObject from the go instance.
    const importObject = go.importObject;

    // Instantiate our wasm module
    const wasmModule = await wasmBrowserInstantiate("./go-ngx-config-parser.wasm", importObject);

    await go.run(wasmModule.instance)

    // Allow the wasm_exec go instance, bootstrap and execute our wasm module
    // Call the Add function export from wasm, save the result

    // Set the result onto the body
    // document.body.textContent = `Hello World! addResult: ${addResult}`;
    } catch (ex) {
        console.error('[registerWasm]', ex)
    }
};


const actButton = document.getElementById('act-btn')
actButton.addEventListener('click', () => {
  const configVal = document.getElementById('nginx-conf').value
  const jsonAst = parseConfig(configVal)

  const astResultPlaceholder = document.getElementById('ast-result')
  astResultPlaceholder.innerHTML = jsonAst
})



registerWasm();