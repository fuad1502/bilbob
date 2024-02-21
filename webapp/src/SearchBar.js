import React from "react";
import "./SearchBar.css";
import { getUsers } from "./api-calls";
import { useState } from "react";

export default function SearchBar() {
  const [users, setUsers] = useState([]);

  function handleChange(text) {
    getUsers(text).then(
      (result) => {
        const [users, ok] = result;
        if (ok) {
          setUsers(users);
        }
      }
    );
  }

  const searchResults = users.map((user) => {
    return (
      <div className="search-result" key={user.username}>
        {user.name} (@{user.username})
      </div>
    );
  });

  const hasResultClassName = searchResults.length > 0 ? "has-result" : "";

  return (
    <div id="search-bar">
      <div id="search-input" className={hasResultClassName}>
        <span>🔎 </span><input onChange={(e) => handleChange(e.target.value)} type="search"></input>
      </div>
      <div id="search-results">
        {searchResults}
      </div>
    </div>
  );
}
