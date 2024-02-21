import React, { useEffect, useState } from "react"
import PostSubmitForm from "./PostSubmitForm"
import Posts from "./Posts"
import { getPosts } from "./api-calls";

export default function PostPanel () {
  const [posts, setPosts] = useState([]);
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    if (!loading) {
      setLoading(true);
      getPosts().then((result) => {
        const [returnedPosts, _] = result;
        setPosts(returnedPosts);
      });
    }
  }, [loading]);

  function handleSubmit() {
    setLoading(false);
  }

  return (
    <div id="post-panel" className="main-panel">
      <PostSubmitForm onSubmit={handleSubmit} />
      <Posts posts={posts}/>
    </div>
  )
}
