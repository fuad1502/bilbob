import React from "react";
import NavigationButton from "./NavigationButton";

export default function ProfileButton({ onClick, selected }) {
  return (
    <NavigationButton onClick={onClick} selected={selected} text={"My House"} logo={"🏠"} />
  )
}
