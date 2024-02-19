import React from 'react';
import ReactDOM from 'react-dom';
import App from './App';
import './index.css';
import { getUsername, redirectWrap } from './api-calls';

redirectWrap(getUsername).then(
  (result) => {
    const [username, _] = result;
    console.log(username)
    ReactDOM.render(
      <App username={username}/>,
      document.getElementById('root')
    );
  }
)
