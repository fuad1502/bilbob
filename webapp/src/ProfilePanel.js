import React, { useState } from "react";
import "./ProfilePanel.css";
import { getUserInfo, getFollowState, requestFollow, unfollow, unrequest, getPosts, getFollowings, getFollowers } from "./api-calls";
import ProfileImage from "./ProfileImage";
import FollowButton from "./FollowButton";
import ProfileName from "./ProfileName";
import Posts from "./Posts";
import FollowingsInfo from "./FollowingsInfo";
import ProfileBanner from "./ProfileBanner";

export default function ProfilePanel({ username, selfUsername }) {
  const [profileInfo, setProfileInfo] = useState({});
  const [followState, setFollowState] = useState({});
  const [loading, setLoading] = useState(false);
  const [loaded, setLoaded] = useState(false);
  const [posts, setPosts] = useState([]);
  const [numOfFollowings, setNumOfFollowings] = useState(0);
  const [numOfFollowers, setNumOfFollowers] = useState(0);

  if (loaded && username !== profileInfo.username) {
    setLoaded(false);
  }

  if (!loaded && !loading) {
    setLoading(true);
    const p1 = getUserInfo(username);
    const p2 = getFollowState(selfUsername, username);
    const p3 = getPosts(username, false);
    const p4 = getFollowings(username);
    const p5 = getFollowers(username);
    Promise.all([p1, p2, p3, p4, p5]).then(
      (result) => {
        const [getProfile, ok1] = result[0];
        const [getFollows, ok2] = result[1];
        const [returnedPosts, ok3] = result[2];
        const [returnedFollowings, ok4] = result[3];
        const [returnedFollowers, ok5] = result[4];
        if (ok1) {
          setProfileInfo(getProfile);
        }
        if (ok2) {
          setFollowState(getFollows);
        }
        if (ok3) {
          setPosts(returnedPosts);
        }
        if (ok4) {
          setNumOfFollowings(returnedFollowings.length);
        }
        if (ok5) {
          setNumOfFollowers(returnedFollowers.length);
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
      <ProfileBanner animal={profileInfo.animal} />
      <ProfileImage key={username} username={username} selfUsername={selfUsername} />
      <div id="profile-name-container">
        <ProfileName animal={profileInfo.animal} name={profileInfo.name} username={profileInfo.username} />
      </div>
      <FollowButton state={followState} onClick={handleClick} />
      <div id="followings-info-container">
        <FollowingsInfo numOfFollowings={numOfFollowings} numOfFollowers={numOfFollowers} />
      </div>
      <Posts posts={posts} />
    </div>
  );
}
