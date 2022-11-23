import type { AppProps } from 'next/app';
import Script from 'next/script';
import { useState } from 'react';
import MainLayout from '../components/Layouts/Main';
import '../styles/globals.css';

export default function App({ Component, pageProps }: AppProps) {
  const [wasmLoaded, setWasmLoaded] = useState(false);

  const handleWasmOnLoad = () => {
    setWasmLoaded(true);
  };

  return (
    <>
      <Script src="/wasm_exec.js" onLoad={handleWasmOnLoad} />
      {wasmLoaded && <Script src="/wasm_init.js" />}
      <MainLayout>
        <Component {...pageProps} />
      </MainLayout>
    </>
  );
}
