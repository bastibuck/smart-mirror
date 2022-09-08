import type { NextPage } from "next";
import Head from "next/head";

const Home: NextPage = () => {
  return (
    <>
      <Head>
        <title>Smart mirror</title>
        <meta name="description" content="Smart mirror application" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main className="flex flex-col items-center justify-center min-h-screen bg-black">
        <h1 className="text-5xl md:text-[5rem] leading-normal font-extrabold text-white">
          Smart mirror
        </h1>
      </main>
    </>
  );
};

export default Home;
