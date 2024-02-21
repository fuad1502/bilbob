import React from "react";
import NavigationButton from "./NavigationButton";

export default function HomeButton({ onClick, selected }) {
  return (
    <NavigationButton onClick={onClick} selected={selected} text={"Park"} logo={"🌳"} />
  )
}
