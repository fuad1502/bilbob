import React from "react";
import './Post.css'

export default function Post ({post}) {
  return (
    <div id="post">
      <p id="username">{'@' + post.username}</p>
      <p id="postText">{post.postText}</p>
    </div>
  )
}
