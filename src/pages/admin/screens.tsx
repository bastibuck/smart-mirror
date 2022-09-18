import React from "react";
import type { NextPageWithLayout } from "../_app";

import { getAdminLayout } from "../../layouts/AdminLayout";
import { trpc } from "../../utils/trpc";
import { useRouter } from "next/router";
import Link from "next/link";

const Screens: NextPageWithLayout = () => {
  const { data } = trpc.useQuery(["screen.getAll"]);

  const router = useRouter();

  const isActive = (id: string) => {
    return router.asPath.includes(id);
  };

  return (
    <main className="flex flex-col items-center justify-center min-h-screen bg-slate-700">
      <h1 className="text-5xl md:text-[5rem] leading-normal font-extrabold text-white">
        Screens
      </h1>

      <div className="carousel w-4/5 gap-10">
        {data?.map((screen) => {
          return (
            <div
              id={screen.id}
              key={screen.id}
              className="carousel-item w-full bg-black h-64 flex justify-center items-center"
            >
              <h3 className="text-xl">{screen.id}</h3>
            </div>
          );
        })}
      </div>

      <div className="flex justify-center w-full py-2 gap-2">
        {data?.map((screen, idx) => (
          <Link key={screen.id} href={`#${screen.id}`}>
            <a
              className={`btn btn-square ${
                isActive(screen.id) ? "btn-primary" : ""
              }`}
            >
              {idx + 1}
            </a>
          </Link>
        ))}
      </div>
    </main>
  );
};

Screens.getLayout = getAdminLayout;

export default Screens;
