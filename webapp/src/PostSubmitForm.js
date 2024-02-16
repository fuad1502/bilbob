import React from 'react';
import './PostSubmitForm.css'

export default function PostSubmitForm() {
  return (
    <div id="post-submit-form">
      What's on your mind?
      <form>
        <textarea></textarea>
        <div id="button-container">
          <div></div>
          <button type='submit'>Post</button>
        </div>
      </form>
    </div>
  );
}
