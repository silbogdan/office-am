import React, { ReactElement } from "react";

import type { NextPageWithLayout } from "../../_app";
import dynamic from "next/dynamic";

const Layout = dynamic(() => import("../../../components/layout"), { ssr: false });

const EmployeesPage: NextPageWithLayout = () => {
  return <div>Employees</div>;
};

EmployeesPage.getLayout = function getLayout(page: ReactElement) {
  return <Layout>{page}</Layout>;
};

export default EmployeesPage;
