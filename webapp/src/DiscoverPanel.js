import React from "react"
import "./DiscoverPanel.css"
import SearchBar from "./SearchBar"

export default function DiscoverPanel({ onSelectUser }) {
  return (
    <div id="discover-panel">
      <SearchBar onSelectUser={onSelectUser} />
    </div>
  )
}
