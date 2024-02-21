import React from "react"
import HomeButton from "./HomeButton"
import ProfileButton from "./ProfileButton"
import MarketplaceButton from "./MarketplaceButton"
import './NavigationPanel.css'

export default function NavigationPanel({ selection, onSelectionClick }) {
  return (
    <div id="navigation-panel">
      <div id="buttons-container">
        <HomeButton onClick={() => onSelectionClick("Home")} selected={selection === "Home" ? true : false} />
        <ProfileButton onClick={() => onSelectionClick("Profile")} selected={selection === "Profile" ? true : false} />
        <MarketplaceButton onClick={() => onSelectionClick("Market")} selected={selection === "Market" ? true : false} />
      </div>
    </div>
  )
}
