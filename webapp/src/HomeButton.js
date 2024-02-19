import React from "react";
import NavigationButton from "./NavigationButton";

export default function HomeButton({onClick}) {
  return (
    <NavigationButton onClick={onClick} text={"Park"} logo={"🌳"} />
  )
}
