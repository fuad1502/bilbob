import React from "react";
import park from "./640px-Morningside_Park_(37571327792).jpeg";
import "./ProfileBanner.css";

export default function ProfileBanner({ animal }) {
  return (
    <div id="profile-banner">
      <img src={park} alt="cozy room" />
    </div>);
}
