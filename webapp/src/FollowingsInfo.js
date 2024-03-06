import React, { useState } from "react";
import { getFollowers, getFollowings } from "./api-calls";
import "./FollowingsInfo.css";

export default function FollowingsInfo({ username }) {
  const [loading, setLoading] = useState(false);
  const [loaded, setLoaded] = useState(false);
  const [numOfFollowings, setNumOfFollowings] = useState(0);
  const [numOfFollowers, setNumOfFollowers] = useState(0);

  if (!loaded && !loading) {
    setLoading(true);
    const p1 = getFollowings(username);
    const p2 = getFollowers(username);
    Promise.all([p1, p2]).then(
      (result) => {
        const [followings, ok1] = result[0];
        const [followers, ok2] = result[1];
        if (ok1) {
          setNumOfFollowings(followings.length);
        }
        if (ok2) {
          setNumOfFollowers(followers.length);
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
