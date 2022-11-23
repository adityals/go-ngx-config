import Head from 'next/head';
import { ChangeEvent, FormEvent, useCallback, useState } from 'react';
import InputConfig from '../components/InputConfig';
import ParseOptions from '../components/ParseOptions';
import { debounce } from '../utilities/common/debounce';

interface ParseOptions {
  skipCtx: boolean;
}

export default function Home() {
  const [parsedConf, setParsedConf] = useState(null);
  const [errParse, setErrParse] = useState('');
  const [parseOpts, setParseOpts] = useState<ParseOptions>({ skipCtx: false });
  const [confVal, setConfVal] = useState('');

  const handleChangeParseOpts = (name: string, e: ChangeEvent<HTMLInputElement>) => {
    const val = e.target.checked;
    switch (name) {
      case 'skip-check-ctx':
        setParseOpts({ skipCtx: val });
        break;
      default:
      // noop
    }
  };

  const parseNginxConfig = debounce(async (e: ChangeEvent<HTMLTextAreaElement>) => {
    const configVal = e.target.value;

    setConfVal(configVal);
    setErrParse('');

    // @ts-ignore - injected from WASM
    const goNgxParseConfig = window?.goNgxParseConfig;
    if (goNgxParseConfig) {
      try {
        const parsed = await goNgxParseConfig(configVal, parseOpts);
        setParsedConf(parsed);
      } catch (err) {
        setErrParse(String(err));
        setParsedConf(null);
      }
    }
  }, 500);

  const handleChangeConfig = useCallback(parseNginxConfig, [parseNginxConfig]);

  const handleCheckTestUrl = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    // @ts-ignore - skip, too lazy to fix
    const targetUrl = e.target['url'].value;

    setErrParse('');

    // @ts-ignore - injected from WASM
    const goNgxTestLocation = window?.goNgxTestLocation;
    if (goNgxTestLocation) {
      try {
        const matchConf = await goNgxTestLocation(confVal, targetUrl, parseOpts);
        setParsedConf(matchConf);
      } catch (err) {
        setErrParse(String(err));
        setParsedConf(null);
      }
    }
  };

  return (
    <div>
      <Head>
        <title>Go Nginx Config | Playground</title>
        <meta name="description" content="Nginx Config Parser and Tester with WASM" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <div>
        <div className="flex">
          <div className="flex-none w-1/12 p-2 bg-slate-800">
            <ParseOptions handleChangeOpts={handleChangeParseOpts} />
          </div>
          <div className="flex-none w-1/2">
            <form>
              <InputConfig name="nginx_config" onChange={handleChangeConfig} />
            </form>
          </div>
          <div className="flex-none h-screen w-5/12">
            <div className="flex flex-col h-screen">
              <div className="bg-slate-400 h-1/6 p-5">
                <form onSubmit={handleCheckTestUrl}>
                  <div>
                    <input
                      type="string"
                      id="url"
                      className="block w-full p-2 text-sm text-gray-900 border border-gray-300 rounded-lg bg-gray-50 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                      placeholder="Test your url (e.g: /my-location)"
                      required
                    />
                  </div>
                  <button
                    type="submit"
                    className="text-white mt-3 right-2.5 bottom-2.5 bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium text-sm px-4 py-2 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
                  >
                    Test
                  </button>
                </form>
              </div>
              <div className="overflow-y-auto h-5/6">
                {parsedConf && <pre className="py-2 pl-3 pr-3 h-full">{parsedConf}</pre>}
                {errParse && <div className="py-2 pl-3 pr-3 h-full text-red-500">{errParse}</div>}
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
