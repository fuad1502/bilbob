import React from 'react';
import './PostSubmitForm.css'
import { postPost } from './api-calls';

export default function PostSubmitForm() {
  function onsubmit (event) {
    event.preventDefault();
    postPost();
  }
  return (
    <div id="post-submit-form">
      What's on your mind?
      <form onSubmit={onsubmit}>
        <textarea></textarea>
        <div id="button-container">
          <div></div>
          <button type='submit'>Post</button>
        </div>
      </form>
    </div>
  );
}
