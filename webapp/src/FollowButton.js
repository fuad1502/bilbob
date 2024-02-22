import React from "react";
import "./FollowButton.css";

export default function FollowButton({ state, onClick }) {
  let text, hoverText, hoverClass, defaultClass;
  let display = true;
  switch (state) {
    case "requested":
      text = "Requested";
      hoverText = "Unrequest";
      defaultClass = "blue"
      hoverClass = "red";
      break;
    case "follows":
      text = "Following";
      hoverText = "Unfollow";
      defaultClass = "green";
      hoverClass = "red";
      break;
    case "no":
      text = "Not following";
      hoverText = "Follow";
      defaultClass = "gray";
      hoverClass = "green";
      break;
    default:
      display = false;
      break;
  }

  function handleMouseLeave(e) {
    e.target.innerText = text;
    e.target.className = defaultClass;
  }

  function handleMouseEnter(e) {
    e.target.innerText = hoverText;
    e.target.className = hoverClass;
  }

  if (display) {
    return (
      <div id="follow-button">
        <button className={defaultClass} onMouseLeave={handleMouseLeave} onMouseEnter={handleMouseEnter} onClick={onClick}>
          {text}
        </button>
      </div>
    );
  }
}
