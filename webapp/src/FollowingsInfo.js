import React, { useState } from "react";
import { getFollowings } from "./api-calls";
import "./FollowingsInfo.css";

export default function FollowingsInfo({ username }) {
  const [loading, setLoading] = useState(false);
  const [loaded, setLoaded] = useState(false);
  const [numOfFollowings, setNumOfFollowings] = useState(0);
  const [numOfFollowers, setNumOfFollowers] = useState(0);

  if (!loaded && !loading) {
    setLoading(true);
    getFollowings(username).then(
      (result) => {
        const [followings, ok] = result;
        if (ok) {
          setNumOfFollowings(followings.length);
        }
        setLoading(false);
        setLoaded(true);
      }
    );
  }

  return (
    <span id="followings-info">
      <div id="followings-info-left">{numOfFollowings} Followings</div>
      <div id="followings-info-right">{numOfFollowers} Followers</div>
    </span>);
}
