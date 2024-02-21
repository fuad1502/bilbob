import React from "react";
import "./SearchBar.css";
import { getUsers } from "./api-calls";

export default function SearchBar() {
  function handleInput(text) {
    getUsers(text)
  }

  return (
    <div id="search-bar">
      <span>🔎 </span><input onInput={(e) => handleInput(e.target.value)} type="search"></input>
    </div>
  );
}
