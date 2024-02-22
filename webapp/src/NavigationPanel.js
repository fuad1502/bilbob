import React from "react"
import HomeButton from "./HomeButton"
import ProfileButton from "./ProfileButton"
import MarketplaceButton from "./MarketplaceButton"
import SearchButton from "./SearchButton"
import LogoutButton from "./LogoutButton"
import './NavigationPanel.css'

export default function NavigationPanel({ selection, onSelectionClick }) {
  return (
    <div id="navigation-panel">
      <div id="buttons-container">
        <HomeButton onClick={() => onSelectionClick("Home")} selected={selection === "Home" ? true : false} />
        <ProfileButton onClick={() => onSelectionClick("Profile")} selected={selection === "Profile" ? true : false} />
        <MarketplaceButton onClick={() => onSelectionClick("Market")} selected={selection === "Market" ? true : false} />
        <SearchButton onClick={() => onSelectionClick("Search")} selected={selection === "Search" ? true : false} />
        <LogoutButton onClick={() => onSelectionClick("Logout")} selected={selection === "Logout" ? true : false} />
      </div>
    </div>
  )
}
