import React from "react";
import "./SearchBar.css";
import { getUsers } from "./api-calls";
import { useState } from "react";

export default function SearchBar({ onSelectUser, inPanel }) {
  const [users, setUsers] = useState([]);

  function handleChange(text) {
    text = text.replaceAll(/[^ \w_]/g, '');
    if (text === "") {
      setUsers([]);
      return;
    }
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
      <div className="search-result" key={user.username} onClick={() => onSelectUser(user.username)}>
        {user.name} (@{user.username})
      </div>
    );
  });

  const hasResultClassName = searchResults.length > 0 ? "has-result" : "";

  return (
    <div id="search-bar" className={inPanel ? "in-panel" : ""}>
      <div id="search-input" className={hasResultClassName}>
        <span>🔎 </span><input onChange={(e) => handleChange(e.target.value)} type="search"></input>
      </div>
      <div id="search-results">
        {searchResults}
      </div>
    </div>
  );
}
