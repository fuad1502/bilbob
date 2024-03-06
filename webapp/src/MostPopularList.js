import React, { useState } from "react";
import { getMostPopularUsers } from "./api-calls";

export default function MostPopularList({ onSelectUser }) {
  const [loading, setLoading] = useState(false);
  const [loaded, setLoaded] = useState(false);
  const [mostPopularList, setMostPopularList] = useState([]);

  if (!loaded && !loading) {
    setLoading(true);
    getMostPopularUsers().then(
      (result) => {
        const [list, ok] = result;
        if (ok) {
          setMostPopularList(list);
        }
        setLoading(false);
        setLoaded(true);
      }
    );
  }

  const element = mostPopularList.map((user) => {
    return (
      <div className="popular-list-item" key={user.username} onClick={() => onSelectUser(user.username)}>
        {user.name} (@{user.username})
      </div>
    );
  });

  return (<div id="most-popular-list">{element}</div>);
}
