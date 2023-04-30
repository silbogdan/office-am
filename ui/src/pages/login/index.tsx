import React, { useState } from "react";
import Link from "next/link";
import { useRouter } from "next/router";

import axios from "axios";

import { Button, Header, Input } from "semantic-ui-react";

import styles from "@/styles/Login.module.css";

export default function LoginPage() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const router = useRouter();

  const login = () => {
    axios.post<{ token: string }>(`${process.env.NEXT_PUBLIC_API_URL}/auth/login`, { email, password }).then((res) => {
      localStorage.setItem("token", res.data.token);
      router.push("/dashboard");
    });
  };

  return (
    <div className={styles.container}>
      <div className={styles["button-group"]}>
        <Header as="h1">
          <p>Log in to Admin Portal</p>
        </Header>
        <div className={styles["hor-container"]}>
          <Input placeholder="Email" value={email} onChange={(e) => setEmail(e.target.value)} />
          <Input
            type="password"
            placeholder="Password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />
          <div className={styles["buttons-container"]}>
            <Button color="green" onClick={login}>
              Log In
            </Button>
            <Link href="/">
              <Button color="red">Cancel</Button>
            </Link>
          </div>
        </div>
      </div>
    </div>
  );
}
