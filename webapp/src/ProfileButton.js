import React from "react";
import NavigationButton from "./NavigationButton";

export default function ProfileButton({onClick}) {
  return (
    <NavigationButton onClick={onClick} text={"My House"} logo={"🏠"}/>
  )
}
