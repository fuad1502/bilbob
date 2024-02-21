import React from "react"
import "./DiscoverPanel.css"
import SearchBar from "./SearchBar"
import SearchResult from "./SearchResult"

export default function DiscoverPanel() {
  return (
    <div id="discover-panel">
      <SearchBar />
      <SearchResult />
    </div>
  )
}
