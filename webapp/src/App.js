import React, { useState } from 'react';
import './App.css';
import NavigationPanel from './NavigationPanel';
import PostPanel from './PostPanel';
import ProfilePanel from './ProfilePanel';
import DiscoverPanel from './DiscoverPanel';
import MarketPanel from './MarketPanel';

export default function App() {
  const [view, setView] = useState("Park");

  let component;
  switch (view) {
    case "Home": component = <PostPanel />; break;
    case "Profile": component = <ProfilePanel />; break;
    case "Market": component = <MarketPanel />; break;
    default: component = <PostPanel />; break;
  }

  function handleNavigationClicked(selection) {
    setView(selection);
  }

  return (
    <div id="app">
      <NavigationPanel onSelectionClick={handleNavigationClicked} />
      {component}
      <DiscoverPanel />
    </div>
  );
}
