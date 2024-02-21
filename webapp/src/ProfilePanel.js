import React, { useState } from "react";
import { getUserInfo, getFollowsUser } from "./api-calls";
import ProfileImage from "./ProfileImage";
import FollowButton from "./FollowButton";
import ProfileName from "./ProfileName";

export default function ProfilePanel({ username }) {
  const [profileInfo, setProfileInfo] = useState({});
  const [follows, setFollows] = useState({});
  const [loading, setLoading] = useState(false);
  const [loaded, setLoaded] = useState(false);

  if (loaded && username !== profileInfo.username) {
    setLoaded(false);
  }

  if (!loaded && !loading) {
    setLoading(true);
    const p1 = getUserInfo(username);
    const p2 = getFollowsUser(username);
    Promise.all([p1, p2]).then(
      (result) => {
        const [getProfile, ok1] = result[0];
        const [getFollows, ok2] = result[1];
        if (ok1) {
          setProfileInfo(getProfile);
        }
        if (ok2) {
          setFollows(getFollows);
        }
        setLoading(false);
        setLoaded(true);
      }
    );
  }

  return (
    <div id="profile-panel" className="main-panel">
      <ProfileImage />
      <ProfileName animal={profileInfo.animal} name={profileInfo.name} username={profileInfo.username} />
      <FollowButton />
    </div>
  );
}
