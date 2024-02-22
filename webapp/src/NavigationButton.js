import React from "react";
import './NavigationButton.css'

export default function NavigationButton({ text, logo, selected, onClick }) {
  return (
    <button id="navigation-button" className={selected ? "selected" : ""} onClick={onClick}>
      <div>
        <span className="logo">{logo}</span><span className="text">{text}</span>
      </div>
    </button>
  )
}
