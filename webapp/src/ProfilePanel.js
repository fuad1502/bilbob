import React, { useState } from "react";
import { getUserInfo, getFollowsUser } from "./api-calls";

export default function ProfilePanel({username}) {
  const [profileInfo, setProfileInfo] = useState(null);
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
    <div id="profile-panel">ProfilePanel</div>
  );
}
