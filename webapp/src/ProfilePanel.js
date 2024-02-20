import React, { useState } from "react";
import { getUserInfo, getFollowsUser, redirectWrap } from "./api-calls";

export default function ProfilePanel({username}) {
  const [profileInfo, setProfileInfo] = useState(null);
  const [isFollowing, setIsFollowing] = useState("NA");
  const [loaded, setLoaded] = useState(false);

  if (!loaded) {
    setLoaded(true);
    redirectWrap(() => getUserInfo(username)).then(
      (result) => {
        const [getProfile, status] = result;
        if (status === 200) {
          setProfileInfo(getProfile);
        }
      }
    );
    redirectWrap(() => getFollowsUser(username)).then(
      (result) => {
        const [getIsFollowing, status] = result;
        if (status === 200) {
          setIsFollowing(getIsFollowing);
        }
      }
    );
  }

  return (
    <div id="profile-panel">ProfilePanel</div>
  );
}
