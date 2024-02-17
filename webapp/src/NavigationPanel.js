import React from "react"
import HomeButton from "./HomeButton"
import ProfileButton from "./ProfileButton"
import MarketplaceButton from "./MarketplaceButton"
import './NavigationPanel.css'

export default function NavigationPanel () {
  return (
    <div id="navigation-panel">
      <div id="buttons-container">
        <HomeButton />
        <ProfileButton />
        <MarketplaceButton />
      </div>
    </div>
  )
}
