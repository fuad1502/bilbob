import React from 'react';
import { useRef } from 'react';
import './PostSubmitForm.css'
import { postPost } from './api-calls';

export default function PostSubmitForm() {
  const ref = useRef(null);
  function onsubmit (event) {
    event.preventDefault();
    postPost(ref.current.value);
    // TODO: display floating message of result
  }
  return (
    <div id="post-submit-form">
      <p>What's on your mind?</p>
      <form onSubmit={onsubmit}>
        <textarea ref={ref}></textarea>
        <div id="button-container">
          <div></div>
          <button type='submit'>Post</button>
        </div>
      </form>
    </div>
  );
}
