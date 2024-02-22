import React from "react";
import NavigationButton from "./NavigationButton";
import "./SearchButton.css";

export default function SearchButton({ onClick, selected }) {
  return (
    <div id="search-button-container">
      <NavigationButton onClick={onClick} selected={selected} text={""} logo={"🔎"} />
    </div>
  )
}
