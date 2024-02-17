import React from "react";
import './NavigationButton.css'

export default function NavigationButton({text, logo}) {
  return (
    <div id="navigation-button">
      <div>
        <span class="logo">{logo}</span><span>{text}</span>
      </div>
    </div>
  )
}
