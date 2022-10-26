import React, { useState } from "react";
import { FiEdit, FiPlus } from "react-icons/fi";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";

import Modal from "../../components/modal/Modal";
import { getAdminLayout } from "../../layouts/AdminLayout";
import { AppType, UpsertAppSchema } from "../../types/shared";
import { trpc } from "../../utils/trpc";

import type { NextPageWithLayout } from "../_app";
import { z } from "zod";

const Apps: NextPageWithLayout = () => {
  const [upsertAppModalOpen, setUpsertAppModalOpen] = useState(false);
  const utils = trpc.useContext();
  const { data } = trpc.useQuery(["apps.getAll"]);
  const upsertAppMutation = trpc.useMutation("apps.upsertApp", {
    onSuccess: () => {
      utils.invalidateQueries(["apps.getAll"]);
      setUpsertAppModalOpen(false);
    },
  });

  const { register, handleSubmit, reset } = useForm<
    z.infer<typeof UpsertAppSchema>
  >({
    resolver: zodResolver(UpsertAppSchema),
  });

  const onSubmit = handleSubmit((data) => {
    upsertAppMutation.mutate(data);
    reset();
  });

  return (
    <main className="flex flex-col items-center min-h-screen bg-slate-700">
      <h1 className="text-5xl md:text-[5rem] leading-normal font-extrabold text-white mb-12">
        Apps
      </h1>

      <div className="overflow-x-auto w-4/5">
        <div className="flex justify-end w-full mb-4">
          <Modal
            isOpen={upsertAppModalOpen}
            onOpen={() => {
              setUpsertAppModalOpen(true);
            }}
            onClose={() => setUpsertAppModalOpen(false)}
            buttonClass={"flex items-center gap-2"}
            title={"New app"}
            openButon={
              <>
                <span>New app</span>
                <FiPlus className="text-xl" />
              </>
            }
          >
            <form onSubmit={onSubmit}>
              <div className="form-control w-full mb-4">
                <label className="label">
                  <span className="label-text">Name</span>
                </label>
                <input
                  {...register("name")}
                  type="text"
                  placeholder="Type here"
                  className="input input-bordered w-full"
                />
              </div>

              <div className="form-control w-full mb-8">
                <label className="label">
                  <span className="label-text">Type</span>
                </label>
                <select
                  className="select select-bordered"
                  defaultValue={"-"}
                  {...register("type")}
                >
                  <option disabled value={"-"}>
                    Select app type
                  </option>

                  <option value={AppType.SIMPLE_TEXT}>Simple text</option>
                </select>
              </div>

              <div className="flex justify-end">
                <button type="submit" className="btn btn-primary">
                  Submit
                </button>
              </div>
            </form>
          </Modal>
        </div>

        <table className="table table-zebra w-full">
          <thead>
            <tr>
              <th>Name</th>
              <th>Type</th>
              <th />
            </tr>
          </thead>

          <tbody>
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

            {data?.length === 0 ? (
              <tr>
                <td
                  colSpan={3}
                  className="text-center italic text-neutral-content"
                >
                  No apps configured
                </td>
              </tr>
            ) : null}
          </tbody>
        </table>
      </div>
    </main>
  );
};

Apps.getLayout = getAdminLayout;

export default Apps;
