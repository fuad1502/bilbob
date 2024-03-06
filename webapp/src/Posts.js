import React from 'react';
import Post from './Post';

export default function Posts({ posts, onSelectUser }) {
  const postsJSX = posts.map((post) => {
    return <Post post={post} onSelectUser={onSelectUser} />
  });
  return (
    <div id="posts">
      {postsJSX}
    </div>
  );
}
