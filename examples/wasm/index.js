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


const actAstBtn = document.getElementById('act-ast-btn');
const actLocationTestBtn = document.getElementById('act-test-location');
const astResultPlaceholder = document.getElementById('result');
const configEl = document.getElementById('nginx-conf');
const targetUrlEl = document.getElementById('target-url');


actAstBtn.addEventListener('click', async () => {
  const configVal = configEl.value;
  try {
    const jsonAst = await parseConfig(configVal);
    astResultPlaceholder.innerHTML = jsonAst;
  } catch(err) {
    console.error('[gen-ast] error:', err);
    astResultPlaceholder.innerHTML = err.toString();
  }
})

actLocationTestBtn.addEventListener('click', async () => {
  const configVal = configEl.value;
  const targetUrlVal = targetUrlEl.value;

  try {
    const jsonAst = await testLocation(configVal, targetUrlVal);
    astResultPlaceholder.innerHTML = jsonAst;
  } catch (err) {
    console.error('[location-test] error:', err);
    astResultPlaceholder.innerHTML = err.toString();
  }
})



registerWasm();