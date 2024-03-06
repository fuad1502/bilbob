import React, { useState } from "react";
import "./ProfilePanel.css";
import { getUserInfo, getFollowState, requestFollow, unfollow, unrequest, getPosts } from "./api-calls";
import ProfileImage from "./ProfileImage";
import FollowButton from "./FollowButton";
import ProfileName from "./ProfileName";
import Posts from "./Posts";
import FollowingsInfo from "./FollowingsInfo";

export default function ProfilePanel({ username, selfUsername }) {
  const [profileInfo, setProfileInfo] = useState({});
  const [followState, setFollowState] = useState({});
  const [loading, setLoading] = useState(false);
  const [loaded, setLoaded] = useState(false);
  const [posts, setPosts] = useState([]);

  if (loaded && username !== profileInfo.username) {
    setLoaded(false);
  }

  if (!loaded && !loading) {
    setLoading(true);
    const p1 = getUserInfo(username);
    const p2 = getFollowState(selfUsername, username);
    const p3 = getPosts(username, false);
    Promise.all([p1, p2, p3]).then(
      (result) => {
        const [getProfile, ok1] = result[0];
        const [getFollows, ok2] = result[1];
        const [returnedPosts, ok3] = result[2];
        if (ok1) {
          setProfileInfo(getProfile);
        }
        if (ok2) {
          setFollowState(getFollows);
        }
        if (ok3) {
          setPosts(returnedPosts);
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
      <ProfileImage key={username} username={username} selfUsername={selfUsername} />
      <ProfileName animal={profileInfo.animal} name={profileInfo.name} username={profileInfo.username} />
      <div id="followings-info-container">
        <FollowingsInfo username={username} />
      </div>
      <FollowButton state={followState} onClick={handleClick} />
      <Posts posts={posts} />
    </div>
  );
}
