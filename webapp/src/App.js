import React from 'react';
import './App.css';
import NavigationPanel from './NavigationPanel';
import PostPanel from './PostPanel';
import DiscoverPanel from './DiscoverPanel';

export default function App() {
  return (
    <div id="app">
      <NavigationPanel />
      <PostPanel />
      <DiscoverPanel />
    </div>
  );
}
