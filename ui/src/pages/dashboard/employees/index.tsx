import React, { ReactElement, useEffect, useState } from "react";

import type { NextPageWithLayout } from "../../_app";
import dynamic from "next/dynamic";
import { Icon, Label, Table } from "semantic-ui-react";

import { differenceInHours } from "date-fns";

const Layout = dynamic(() => import("../../../components/layout"), { ssr: false });

export interface TimeLog {
  ID: number;
  CreatedAt: string;
  UpdatedAt: string;
  DeletedAt: string;
  UserId: number;
  Type: "entrance" | "exit";
}

export interface Employee {
  ID: number;
  CreatedAt: string;
  UpdatedAt: string;
  DeletedAt: string;
  name: string;
  email: string;
  picture: string;
  qrCodeId: string;
  qrCodeUrl: string;
  timeLogs: TimeLog[];
}

export const revalidate = 1;

const EmployeesPage: NextPageWithLayout = () => {
  const [employees, setEmployees] = useState<Employee[]>([]);

  useEffect(() => {
    const fetchData = async () => {
      let token = "";
      if (typeof window !== "undefined") {
        token = localStorage.getItem("token") || "";
      }

      const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/employees`, {
        headers: { Authorization: `Bearer ${token}` },
        next: { revalidate: 1 },
      });

      const resJson = await res.json();

      console.log("Got employees:", resJson.employees);

      const fetchedEmployees = resJson.employees as Employee[] | undefined;

      if (fetchedEmployees) setEmployees(fetchedEmployees);
    };

    fetchData();
  }, []);

  return (
    <div style={{ marginLeft: "20px", marginRight: "20px" }}>
      <Table celled striped>
        <Table.Header>
          <Table.Row>
            <Table.HeaderCell colSpan="5">Employees</Table.HeaderCell>
          </Table.Row>
        </Table.Header>

        <Table.Body>
          {employees.map((e, key) => (
            <Table.Row key={key}>
              <Table.Cell>
                <img alt="profile" src={e.picture} width={50} height={50} />
              </Table.Cell>
              <Table.Cell>{e.name}</Table.Cell>
              <Table.Cell>{e.email}</Table.Cell>
              <Table.Cell verticalAlign="middle" collapsing textAlign="right">
                <img alt="qr-code" src={e.qrCodeUrl} width={50} height={50} />
              </Table.Cell>
            </Table.Row>
          ))}
        </Table.Body>
      </Table>
    </div>
  );
};

EmployeesPage.getLayout = function getLayout(page: ReactElement) {
  return <Layout>{page}</Layout>;
};

export default EmployeesPage;
