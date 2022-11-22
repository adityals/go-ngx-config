import Head from 'next/head';
import { ChangeEvent, useState } from 'react';
import InputConfig from '../components/InputConfig';
import styles from '../styles/Home.module.css';
import { debounce } from '../utilities/common/debounce';

export default function Home() {
  const [parsedConf, setParsedConf] = useState(null);
  const [errParse, setErrParse] = useState('');

  const handleChangeConfig = debounce(async (e: ChangeEvent<HTMLTextAreaElement>) => {
    const configVal = e.target.value;

    setErrParse('');

    // @ts-expect-error - for now
    const goNgxParseConfig = window?.goNgxParseConfig;

    if (goNgxParseConfig) {
      try {
        const parsed = await goNgxParseConfig(configVal);
        setParsedConf(parsed);
      } catch (err) {
        setErrParse(String(err));
        setParsedConf(null);
      }
    }
  }, 500);

  return (
    <div className={styles.container}>
      <Head>
        <title>Go Nginx Config | Playground</title>
        <meta name="description" content="Nginx Config Parser and Tester with WASM" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <main>
        <div className="grid grid-cols-2">
          <form>
            <InputConfig name="nginx_config" onChange={handleChangeConfig} />
          </form>
          <div className="h-screen">
            {parsedConf && <pre className="py-2 pl-3 pr-3 h-full">{parsedConf}</pre>}
            {errParse && <div className="py-2 pl-3 pr-3 h-full text-red-500">{errParse}</div>}
          </div>
        </div>
      </main>
    </div>
  );
}
