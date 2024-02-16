import React from "react";

export default function Post ({post}) {
  return (
    <div id="post">
      <p>{post.username}</p>
      <p>{post.postText}</p>
    </div>
  )
}
