import React from "react";
import type { NextPageWithLayout } from "../_app";

import { getAdminLayout } from "../../layouts/AdminLayout";

const Sections: NextPageWithLayout = () => {
  return (
    <main className="flex flex-col items-center justify-center min-h-screen bg-slate-700">
      <h1 className="text-5xl md:text-[5rem] leading-normal font-extrabold text-white">
        Sections
      </h1>
    </main>
  );
};

Sections.getLayout = getAdminLayout;

export default Sections;
