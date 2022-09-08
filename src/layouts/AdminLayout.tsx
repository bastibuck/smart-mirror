import React, { PropsWithChildren, ReactElement } from "react";
import FloatingMenu from "../components/actions/FloatingMenu";

const AdminLayout: React.FC<PropsWithChildren> = ({ children }) => {
  return (
    <>
      {children}

      <FloatingMenu />
    </>
  );
};

export const getAdminLayout = (page: ReactElement) => {
  return <AdminLayout>{page}</AdminLayout>;
};
