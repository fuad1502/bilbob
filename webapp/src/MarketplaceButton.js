import React from "react";
import NavigationButton from "./NavigationButton";

export default function MarketplaceButton({ onClick, selected }) {
  return (
    <NavigationButton onClick={onClick} selected={selected} text={"Marketplace"} logo={"🛍️"} />
  )
}
