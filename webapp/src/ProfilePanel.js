import React, { useState } from "react";
import { getUserInfo, getFollowState, requestFollow, unfollow, unrequest } from "./api-calls";
import ProfileImage from "./ProfileImage";
import FollowButton from "./FollowButton";
import ProfileName from "./ProfileName";

export default function ProfilePanel({ username, selfUsername }) {
  const [profileInfo, setProfileInfo] = useState({});
  const [followState, setFollowState] = useState({});
  const [loading, setLoading] = useState(false);
  const [loaded, setLoaded] = useState(false);

  if (loaded && username !== profileInfo.username) {
    setLoaded(false);
  }

  if (!loaded && !loading) {
    setLoading(true);
    const p1 = getUserInfo(username);
    const p2 = getFollowState(selfUsername, username);
    Promise.all([p1, p2]).then(
      (result) => {
        const [getProfile, ok1] = result[0];
        const [getFollows, ok2] = result[1];
        if (ok1) {
          setProfileInfo(getProfile);
        }
        if (ok2) {
          setFollowState(getFollows);
        }
        setLoading(false);
        setLoaded(true);
      }
    );
  }

  function handleClick() {
    if (followState === "no") {
      requestFollow(selfUsername, username).then(() => setLoaded(false));
    } else if (followState === "follows") {
      unfollow(selfUsername, username).then(() => setLoaded(false));
    } else if (followState === "requested") {
      unrequest(selfUsername, username).then(() => setLoaded(false));
    }
  }

  return (
    <div id="profile-panel" className="main-panel">
      <ProfileImage />
      <ProfileName animal={profileInfo.animal} name={profileInfo.name} username={profileInfo.username} />
      <FollowButton state={followState} onClick={handleClick} />
    </div>
  );
}
