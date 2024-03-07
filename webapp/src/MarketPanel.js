import React from "react";
import gopher from "./1024px-Go_gopher_pencil_wrench.jpg";
import "./MarketPanel.css";

export default function MarketPanel() {
  return (
    <div id="market-panel" className="main-panel">
      <h2>Under construction 🚧</h2>
      <img src={gopher} alt="Golang's gopher constructing something good" />
    </div>
  );
}
