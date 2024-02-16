import React from 'react';
import Post from './Post';

export default function Posts({posts}) {
  const postsJSX = posts.map((post) => {
    return <Post post={post} />
  });
  return (
    <div id="posts">
      {postsJSX}
    </div>
  );
}
