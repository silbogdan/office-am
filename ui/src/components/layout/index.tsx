import React, { useEffect, useState } from "react";
import { useRouter } from "next/router";

import Navbar from "../navbar";

export default function Layout({ children }) {
  let token: string | null = null;
  const router = useRouter();

  if (typeof window !== "undefined") {
    console.log("Getting token");
    token = localStorage.getItem("token");
  }

  return token ? (
    <div>
      <Navbar />
      {children}
    </div>
  ) : (
    <div></div>
  );
}
