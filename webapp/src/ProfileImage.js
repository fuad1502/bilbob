import React, { useRef, useState } from "react";
import "./ProfileImage.css";
import { getProfilePicture, setProfilePicture } from "./api-calls";

export default function ProfileImage({ username, selfUsername }) {
  const [loading, setLoading] = useState(false);
  const [loaded, setLoaded] = useState(false);
  const [image, setImage] = useState(<img alt=""></img>);
  const ref = useRef(null);

  if (!loaded && !loading) {
    setLoading(true);
    getProfilePicture(username).then((result) => {
      const [url, ok] = result;
      if (ok) {
        setImage(<img src={url} alt={"An image of " + username}></img>)
      }
      setLoading(false);
      setLoaded(true);
    });
  }

  function handleClick() {
    if (username === selfUsername) {
      ref.current.click()
    }
  }

  function handleChange(e) {
    setProfilePicture(username, e.target.files[0]).then(() => setLoaded(false));
  }

  return (
    <div id="profile-image" onClick={handleClick}>
      {image}
      <input ref={ref} onChange={handleChange} type="file"></input>
    </div>
  );
}
