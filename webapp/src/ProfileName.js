import React from "react";
import "./ProfileName.css";

export default function ProfileName({ name, username, animal }) {

  switch (animal) {
    case "cat": animal = "🐱"; break;
    default:
  }

  return (
    <div id="profile-name">
      <div>
        <span id="animal">{animal} </span><span id="name">{name} </span><span id="username">(@{username})</span>
      </div>
    </div>
  );
}

