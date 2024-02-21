import React, { useState } from "react";
import { getUserInfo, getFollowsUser } from "./api-calls";
import ProfileImage from "./ProfileImage";
import FollowButton from "./FollowButton";
import ProfileName from "./ProfileName";

export default function ProfilePanel({ username }) {
  const [profileInfo, setProfileInfo] = useState({});
  const [isFollowing, setIsFollowing] = useState("NA");
  const [loaded, setLoaded] = useState(false);

  if (!loaded) {
    setLoaded(true);
    getUserInfo(username).then(
      (result) => {
        const [getProfile, ok] = result;
        if (ok) {
          setProfileInfo(getProfile);
        }
      }
    );
    getFollowsUser(username).then(
      (result) => {
        const [getIsFollowing, ok] = result;
        if (ok) {
          setIsFollowing(getIsFollowing);
        }
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
