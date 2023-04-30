import React from "react";

import { Menu } from "semantic-ui-react";

export default function Navbar() {
  return (
    <Menu>
      <Menu.Item name="employees" onClick={() => {}}>
        Employees
      </Menu.Item>

      <Menu.Item name="add-employee" onClick={() => {}}>
        Add New Employee
      </Menu.Item>

      <Menu.Item name="logs" onClick={() => {}}>
        Logs
      </Menu.Item>
    </Menu>
  );
}
