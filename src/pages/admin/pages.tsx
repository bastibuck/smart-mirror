import React from "react";
import type { NextPageWithLayout } from "../_app";

import { getAdminLayout } from "../../layouts/AdminLayout";

const Pages: NextPageWithLayout = () => {
  return (
    <main className="flex flex-col items-center justify-center min-h-screen bg-slate-700">
      <h1 className="text-5xl md:text-[5rem] leading-normal font-extrabold text-white">
        Pages
      </h1>
    </main>
  );
};

Pages.getLayout = getAdminLayout;

export default Pages;
