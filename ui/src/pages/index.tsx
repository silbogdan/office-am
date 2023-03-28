import styles from "@/styles/Home.module.css";
import bg from "../../public/background.jpeg";

import Link from "next/link";
import { useState } from "react";

import { Button, Header, Modal, Image } from "semantic-ui-react";
import axios from "axios";

const checkedInModalHeader = "Welcome to the office";
const checkedInHeader = "You have checked in";
const checkedInDescription = () =>
  "The time of your check in is: " +
  new Date().toLocaleTimeString() +
  " " +
  new Date().toLocaleDateString() +
  ".\nGood luck with your work!";

const checkedOutModalHeader = "See you tomorrow";
const checkedOutHeader = "You have checked out";
const checkedOutDescription = () =>
  "The time of your check out is: " +
  new Date().toLocaleTimeString() +
  " " +
  new Date().toLocaleDateString() +
  ".\nEnjoy the rest of your day!";

export default function Home() {
  const [open, setOpen] = useState(false);
  const [name, setName] = useState("");
  const [picture, setPicture] = useState("");
  const [checkedIn, setCheckedIn] = useState(false);
  const [checkingIn, setCheckingIn] = useState(false);
  const [checkingOut, setCheckingOut] = useState(false);

  const checkIn = () => {
    setCheckingIn(true);

    axios
      .post(`${process.env.NEXT_PUBLIC_API_URL}/timelogs/scan`, { type: "entrance" })
      .then((res) => {
        const { data } = res;

        if (data?.qrCodeData?.Type) {
          setName(data.qrCodeData.Name);
          setPicture(data.qrCodeData.Picture);
          setCheckedIn(true);
          setOpen(true);
          setCheckingIn(false);
        } else {
          setCheckingIn(false);
          // TODO: Show error message
        }
      })
      .catch((err) => {
        setCheckingIn(false);
        // TODO: Show error message
      });
  };

  const checkOut = () => {
    setCheckingOut(true);

    axios
      .post(`${process.env.NEXT_PUBLIC_API_URL}/timelogs/scan`, { type: "exit" })
      .then((res) => {
        const { data } = res;

        if (data?.qrCodeData?.Type) {
          setName(data.qrCodeData.Name);
          setPicture(data.qrCodeData.Picture);
          setCheckedIn(false);
          setOpen(true);
          setCheckingOut(false);
        } else {
          setCheckingOut(false);
          // TODO: Show error message
        }
      })
      .catch((err) => {
        setCheckingOut(false);
        // TODO: Show error message
      });
  };

  return (
    <div className={styles.container} style={{ backgroundImage: `url(${bg.src})`, backgroundSize: "cover" }}>
      <div className={styles["button-group"]}>
        <Header as="h1">
          <p style={{ color: "white" }}>Welcome to the Office</p>
        </Header>
        <Button loading={checkingIn} onClick={checkIn} color="green">
          Check In
        </Button>
        <Button loading={checkingOut} onClick={checkOut} primary>
          Check Out
        </Button>
        <Link href="/login">
          <Button color="red">Log In</Button>
        </Link>
      </div>

      <Modal onClose={() => setOpen(false)} onOpen={() => setOpen(true)} open={open}>
        <Modal.Header>{checkedIn ? checkedInModalHeader : checkedOutModalHeader}</Modal.Header>
        <Modal.Content image>
          <Image size="medium" src={picture} wrapped alt="profile-pic" />
          <Modal.Description>
            <Header as="h2">Hello, {name}</Header>
            <Header>{checkedIn ? checkedInHeader : checkedOutHeader}</Header>
            <p>{checkedIn ? checkedInDescription() : checkedOutDescription()}</p>
          </Modal.Description>
        </Modal.Content>
        <Modal.Actions>
          <Button content="Okay" labelPosition="right" icon="checkmark" onClick={() => setOpen(false)} positive />
        </Modal.Actions>
      </Modal>
    </div>
  );
}
