import React from "react";
import './NavigationButton.css'

export default function NavigationButton({ text, logo, selected, onClick }) {
  return (
    <div id="navigation-button" className={selected ? "selected" : ""} onClick={onClick}>
      <div>
        <span className="logo">{logo}</span><span>{text}</span>
      </div>
    </div>
  )
}
