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
        const [returnedPosts, httpStatus] = result;
        if (httpStatus !== 200) {
          // TODO: Show floating error message
        }
        console.log(returnedPosts);
        setPosts(returnedPosts);
      });
    }
  }, [loading]);
  return (
    <div id="post-panel">
      <PostSubmitForm />
      <Posts posts={posts}/>
    </div>
  )
}
