import React from "react";
import './NavigationButton.css'

export default function NavigationButton({text, logo, onClick}) {
  return (
    <div id="navigation-button" onClick={onClick}>
      <div>
        <span class="logo">{logo}</span><span>{text}</span>
      </div>
    </div>
  )
}
