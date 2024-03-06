import React from "react";
import "./FollowingsInfo.css";

export default function FollowingsInfo({ numOfFollowings, numOfFollowers }) {
  return (
    <span id="followings-info">
      <div id="followings-info-left">{numOfFollowings} Followings</div>
      <div id="followings-info-right">{numOfFollowers} Followers</div>
    </span>);
}
