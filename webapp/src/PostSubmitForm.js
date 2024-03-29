import React from 'react';
import { useRef } from 'react';
import './PostSubmitForm.css'
import { postPost } from './api-calls';

export default function PostSubmitForm({ username, onSubmit }) {
  const ref = useRef(null);

  function handleSubmit(event) {
    event.preventDefault();
    postPost(username, ref.current.value).then((ok) => {
      if (ok) {
        onSubmit();
      } else {
        // TODO: display error floating message
      }
    });
  }

  return (
    <div id="post-submit-form">
      <p>What's on your mind?</p>
      <form onSubmit={handleSubmit}>
        <textarea ref={ref}></textarea>
        <div id="button-container">
          <div></div>
          <button type='submit'>Post</button>
        </div>
      </form>
    </div>
  );
}
