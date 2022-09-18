import React from "react";
import { FiEdit, FiPlus } from "react-icons/fi";
import { getAdminLayout } from "../../layouts/AdminLayout";
import { trpc } from "../../utils/trpc";

import type { NextPageWithLayout } from "../_app";

const Apps: NextPageWithLayout = () => {
  const { data } = trpc.useQuery(["apps.getAll"]);

  return (
    <main className="flex flex-col items-center min-h-screen bg-slate-700">
      <h1 className="text-5xl md:text-[5rem] leading-normal font-extrabold text-white mb-12">
        Apps
      </h1>

      <div className="overflow-x-auto w-4/5">
        <table className="table table-zebra w-full">
          <thead>
            <tr>
              <th>Name</th>
              <th>Type</th>
              <th />
            </tr>
          </thead>

          <tbody>
            <tr>
              <td
                colSpan={3}
                className="bg-primary-content hover:bg-primary-focus hover:cursor-pointer"
              >
                <div
                  className=" flex justify-end items-center w-full gap-2 select-none"
                  onClick={() => {
                    console.log("open add modal");
                  }}
                >
                  <span>New app</span>
                  <FiPlus className="text-xl" />
                </div>
              </td>
            </tr>
            {data?.map((app) => (
              <tr className="hover" key={app.id}>
                <td>{app.name}</td>
                <td>{app.type}</td>
                <td className="text-right">
                  <FiEdit
                    className="inline hover:text-primary hover:cursor-pointer"
                    onClick={() => {
                      console.log("open edit modal");
                    }}
                  />
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </main>
  );
};

Apps.getLayout = getAdminLayout;

export default Apps;
