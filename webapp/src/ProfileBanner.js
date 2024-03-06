import React from "react";
import catHouse from "./pexels-cats-coming-674580.jpg";
import dogPark from "./pexels-leo-bao-2663202.jpg";
import reef from "./pexels-francesco-ungaro-3854025.jpg";
import sky from "./pexels-pixabay-531767.jpg";
import rock from "./mossy-rock-560385_1280.jpg";
import "./ProfileBanner.css";

export default function ProfileBanner({ animal }) {
  let banner;
  switch (animal) {
    case 'cat': banner = catHouse; break;
    case 'dog': banner = dogPark; break;
    case 'fish': banner = reef; break;
    case 'bird': banner = sky; break;
    case 'reptile': banner = rock; break;
    default: banner = catHouse;
  }
  return (
    <div id="profile-banner">
      <img src={banner} alt="cozy room" />
    </div>);
}
