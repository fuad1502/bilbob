import React from "react";
import './Post.css'
import ProfileImage from "./ProfileImage";

export default function Post({ post, onSelectUser }) {
  const postDateString = displayPostDate(post.postDate);
  return (
    <div id="post">
      <div id="post-header">
        <ProfileImage key={post.username} username={post.username} inPost={true} />
        <span onClick={() => onSelectUser(post.username)} id="username">{'@' + post.username} · {postDateString}</span>
      </div>
      <p id="postText">{post.postText}</p>
    </div>
  )
}

function displayPostDate(postDate) {
  const timestamp = Date.parse(postDate);
  const secondsAgo = (Date.now() - timestamp) / 1000;
  const minutesAgo = secondsAgo / 60;
  const hoursAgo = minutesAgo / 60;
  const daysAgo = hoursAgo / 24;
  if (secondsAgo < 60) {
    return 'now';
  } else if (minutesAgo < 60) {
    return (Math.round(minutesAgo)).toString() + 'm';
  } else if (hoursAgo < 24) {
    return (Math.round(hoursAgo)).toString() + 'h';
  } else if (daysAgo < 7) {
    return (Math.round(daysAgo)).toString() + 'd';
  } else {
    return (new Date(timestamp)).toDateString();
  }
}
