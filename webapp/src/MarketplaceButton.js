import React from "react";
import NavigationButton from "./NavigationButton";

export default function MarketplaceButton({onClick}) {
  return (
    <NavigationButton onClick={onClick} text={"Marketplace"} logo={"🛍️"}/>
  )
}
