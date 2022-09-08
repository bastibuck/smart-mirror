import Head from "next/head";
import type { NextPageWithLayout } from "../_app";
import { getAdminLayout } from "../../layouts/AdminLayout";

const Admin: NextPageWithLayout = () => {
  return (
    <>
      <Head>
        <title>Admin | Smart mirror</title>
        <meta
          name="description"
          content="Administration of Smart mirror application"
        />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main className="flex flex-col items-center justify-center min-h-screen bg-slate-700">
        <h1 className="text-5xl md:text-[5rem] leading-normal font-extrabold text-white">
          ADMIN Smart mirror
        </h1>
      </main>
    </>
  );
};

Admin.getLayout = getAdminLayout;

export default Admin;
