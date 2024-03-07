import React, { useEffect, useState } from "react"
import PostSubmitForm from "./PostSubmitForm"
import Posts from "./Posts"
import { getPosts } from "./api-calls";

export default function PostPanel({ username, onSelectUser }) {
  const [posts, setPosts] = useState([]);
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    if (!loading) {
      setLoading(true);
      let fromTimestamp = "";
      if (posts.length > 0) {
        fromTimestamp = posts[posts.length - 1].postDate;
      }
      getPosts(username, true, fromTimestamp).then((result) => {
        const [returnedPosts, _] = result;
        setPosts(posts.concat(returnedPosts));
      });
    }
  }, [loading]);

  function handleSubmit() {
    setPosts([]);
    setLoading(false);
  }

  window.onscroll = function(_e) {
    if ((window.innerHeight + Math.round(window.scrollY)) >= document.body.offsetHeight) {
      setLoading(false);
    }
  };

  return (
    <div id="post-panel" className="main-panel">
      <PostSubmitForm username={username} onSubmit={handleSubmit} />
      <Posts posts={posts} onSelectUser={onSelectUser} />
    </div>
  )
}
