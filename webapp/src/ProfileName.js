import React from "react";
import "./ProfileName.css";

export default function ProfileName({ name, username, animal }) {

  switch (animal) {
    case "cat": animal = "🐱"; break;
    case "dog": animal = "🐶"; break;
    case "bird": animal = "🐦"; break;
    case "fish": animal = "🐠"; break;
    case "reptile": animal = "🐍"; break;
    default: animal = "🦖";
  }

  return (
    <div id="profile-name">
      <span id="animal">{animal} </span><span id="name">{name} </span><span id="username">(@{username})</span>
    </div>
  );
}

