import React from "react";
import SearchBar from "./SearchBar";
import MostPopularList from "./MostPopularList";
import "./DiscoverPanel.css";

export default function SearchPanel({ onSelectUser }) {
  return (
    <div id="discover-items-container" className="in-panel">
      <SearchBar onSelectUser={onSelectUser} inPanel={true} />
      <MostPopularList onSelectUser={onSelectUser} inPanel={true} />
    </div>
  )
}
