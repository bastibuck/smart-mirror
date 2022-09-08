import React from "react";
import { getAdminLayout } from "../../layouts/AdminLayout";

import type { NextPageWithLayout } from "../_app";

const Apps: NextPageWithLayout = () => {
  return (
    <main className="flex flex-col items-center justify-center min-h-screen bg-slate-700">
      <h1 className="text-5xl md:text-[5rem] leading-normal font-extrabold text-white">
        Apps
      </h1>
    </main>
  );
};

Apps.getLayout = getAdminLayout;

export default Apps;
