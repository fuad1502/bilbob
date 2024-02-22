import React from "react";
import SearchBar from "./SearchBar";

export default function SearchPanel({ onSelectUser }) {
  return (
    <div id="search-panel" className="main-panel">
      <SearchBar onSelectUser={onSelectUser} inPanel={true} />
    </div>
  )
}
