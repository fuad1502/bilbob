import React from "react"
import "./DiscoverPanel.css"
import SearchBar from "./SearchBar"
import MostPopularList from "./MostPopularList"

export default function DiscoverPanel({ onSelectUser }) {
  return (
    <div id="discover-panel">
      <div id="discover-items-container">
        <SearchBar onSelectUser={onSelectUser} />
        <MostPopularList onSelectUser={onSelectUser} />
      </div>
    </div>
  )
}
