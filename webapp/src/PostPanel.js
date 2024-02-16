import React from "react"
import PostSubmitForm from "./PostSubmitForm"
import Posts from "./Posts"

export default function PostPanel () {
  return (
    <div id="post-panel">
      <PostSubmitForm />
      <Posts />
    </div>
  )
}
