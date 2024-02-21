import React, { useState } from 'react';
import './App.css';
import NavigationPanel from './NavigationPanel';
import PostPanel from './PostPanel';
import ProfilePanel from './ProfilePanel';
import DiscoverPanel from './DiscoverPanel';
import MarketPanel from './MarketPanel';

export default function App({ username }) {
  const [view, setView] = useState("Home");
  const [otherProfile, setOtherProfile] = useState("");

  let component;
  switch (view) {
    case "Home": component = <PostPanel />; break;
    case "Profile": component = <ProfilePanel username={username} />; break;
    case "Market": component = <MarketPanel />; break;
    case "OtherProfile": component = <ProfilePanel username={otherProfile} />; break;
    default: component = <PostPanel />; break;
  }

  function handleNavigationClicked(selection) {
    setView(selection);
  }

  function handleSelectUser(username) {
    setView("OtherProfile");
    setOtherProfile(username);
  }

  return (
    <div id="app">
      <NavigationPanel onSelectionClick={handleNavigationClicked} selection={view} />
      {component}
      <DiscoverPanel onSelectUser={handleSelectUser} />
    </div>
  );
}
