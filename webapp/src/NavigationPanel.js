import React from "react"
import HomeButton from "./HomeButton"
import ProfileButton from "./ProfileButton"
import MarketplaceButton from "./MarketplaceButton"
import './NavigationPanel.css'

export default function NavigationPanel ({onSelectionClick}) {
  return (
    <div id="navigation-panel">
      <div id="buttons-container">
        <HomeButton onClick={() => onSelectionClick("Home")}/>
        <ProfileButton onClick={() => onSelectionClick("Profile")}/>
        <MarketplaceButton onClick={() => onSelectionClick("Market")}/>
      </div>
    </div>
  )
}
