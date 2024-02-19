import React, { useEffect, useState } from "react"
import PostSubmitForm from "./PostSubmitForm"
import Posts from "./Posts"
import { getPosts, redirectWrap, LandingPageUrl } from "./api-calls";

export default function PostPanel () {
  const [posts, setPosts] = useState([]);
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    if (!loading) {
      setLoading(true);
      redirectWrap(getPosts, LandingPageUrl).then((result) => {
        const [returnedPosts, httpStatus] = result;
        if (httpStatus !== 200) {
          // TODO: Show floating error message
        }
        setPosts(returnedPosts);
      });
    }
  }, [loading]);

  function handleSubmit() {
    setLoading(false);
  }

  return (
    <div id="post-panel">
      <PostSubmitForm onSubmit={handleSubmit} />
      <Posts posts={posts}/>
    </div>
  )
}
