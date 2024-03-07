import React, { useState } from "react"
import PostSubmitForm from "./PostSubmitForm"
import Posts from "./Posts"
import { getPosts } from "./api-calls";

export default function PostPanel({ username, onSelectUser }) {
  const [posts, setPosts] = useState([]);
  const [loaded, setLoaded] = useState(false);
  const [loading, setLoading] = useState(false);
  const [bottomed, setBottomed] = useState(false);

  if (!bottomed && !loaded && !loading) {
    setLoading(true);
    let fromTimestamp = "";
    if (posts.length > 0) {
      fromTimestamp = posts[posts.length - 1].postDate;
    }
    getPosts(username, true, fromTimestamp).then((result) => {
      const [returnedPosts, ok] = result;
      if (ok) {
        if (returnedPosts.length > 0) {
          setPosts(posts.concat(returnedPosts));
        } else {
          setBottomed(true);
        }
      }
      setLoading(false);
      setLoaded(true);
    });
  }

  function handleSubmit() {
    setPosts([]);
    setBottomed(false);
    setLoaded(false);
  }

  if (!bottomed) {
    window.onscroll = function(_e) {
      if ((window.innerHeight * 1.2 + Math.round(window.scrollY)) >= document.body.offsetHeight) {
        setLoaded(false);
      }
    };
  }

  return (
    <div id="post-panel" className="main-panel">
      <PostSubmitForm username={username} onSubmit={handleSubmit} />
      <Posts posts={posts} onSelectUser={onSelectUser} />
    </div>
  )
}
