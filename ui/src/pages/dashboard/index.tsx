import React, { ReactElement } from "react";
import dynamic from "next/dynamic";

import type { NextPageWithLayout } from "../_app";
const Layout = dynamic(() => import("../../components/layout"), { ssr: false });

const DashboardPage: NextPageWithLayout = () => {
  return <div>Dashboard</div>;
};

DashboardPage.getLayout = function getLayout(page: ReactElement) {
  return <Layout>{page}</Layout>;
};

export default DashboardPage;
