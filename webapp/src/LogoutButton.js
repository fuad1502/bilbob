import React from "react";
import NavigationButton from "./NavigationButton";

export default function LogoutButton({ onClick, selected }) {
  return (
    <NavigationButton onClick={onClick} selected={selected} text={"Logout"} logo={"🚪"} />
  )
}
