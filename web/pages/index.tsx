import Head from 'next/head';
import { useCallback, useEffect, useMemo, useState } from 'react';
import type { ChangeEvent, FormEvent } from 'react';
import InputConfig from '../components/ConfigEditor';
import ParseOptions from '../components/ParseOptions';
import ParserResult from '../components/ParserResult';
import { debounce } from '../utilities/common/debounce';

interface ParseOptions {
  skipCtx: boolean;
}

export default function Home() {
  const [config, setConfig] = useState('');
  const [parseOpts, setParseOpts] = useState<ParseOptions>({ skipCtx: false });

  const [parsedConf, setParsedConf] = useState(null);
  const [submittedTestLocation, setSubmittedTestLocation] = useState(false);

  const [errParse, setErrParse] = useState('');

  const handleChangeParseOpts = useCallback((name: string, e: ChangeEvent<HTMLInputElement>) => {
    const val = e.target.checked;
    switch (name) {
      case 'skip-check-ctx':
        setParseOpts({ skipCtx: val });
        break;
      default:
      // noop
    }
  }, []);

  const debounceParseNgxConf = debounce(async (config: string) => {
    // @ts-ignore - injected from WASM
    const goNgxParseConfig = window?.goNgxParseConfig;
    if (goNgxParseConfig) {
      try {
        const parsed = await goNgxParseConfig(config, parseOpts);
        setParsedConf(parsed);
      } catch (err) {
        setErrParse(String(err));
        setParsedConf(null);
      }
    }
  }, 700);

  const parseNgxConf = useMemo(() => debounceParseNgxConf, [debounceParseNgxConf]);

  const handleChangeConfig = useCallback((v: string | undefined) => {
    if (v) {
      setConfig(v);
      setSubmittedTestLocation(false);
    }
  }, []);

  const handleCheckTestUrl = useCallback(
    async (e: FormEvent<HTMLFormElement>) => {
      e.preventDefault();
      setSubmittedTestLocation(true);

      // @ts-ignore - skip, too lazy to fix
      const targetUrl = e.target['url'].value;
      // @ts-ignore - injected from WASM
      const goNgxTestLocation = window?.goNgxTestLocation;
      if (goNgxTestLocation) {
        try {
          const matchConf = await goNgxTestLocation(config, targetUrl, parseOpts);
          setParsedConf(matchConf);
        } catch (err) {
          setErrParse(String(err));
          setParsedConf(null);
        }
      }
    },
    [config, parseOpts],
  );

  useEffect(() => {
    if (config && !submittedTestLocation) {
      parseNgxConf(config);
    }
  }, [config, submittedTestLocation, parseNgxConf]);

  return (
    <div>
      <Head>
        <title>Go Nginx Config | Playground</title>
        <meta name="description" content="Nginx Config Parser and Location Tester with WASM" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <div>
        <div className="flex">
          <div className="flex-none w-1/12 p-2 bg-slate-800">
            <ParseOptions handleChangeOpts={handleChangeParseOpts} />
          </div>
          <div className="flex-none w-1/2">
            <form>
              <InputConfig value={config} onChange={handleChangeConfig} />
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
                      className="block w-full p-2 text-sm text-gray-900 border border-gray-300 bg-gray-50 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
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
                {parsedConf && <ParserResult value={parsedConf} languange="json" />}
                {errParse && <ParserResult value={errParse} languange="bash" />}
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
