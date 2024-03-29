import React, { useRef, useState } from "react";
import "./ProfileImage.css";
import { getProfilePicture, setProfilePicture } from "./api-calls";

export default function ProfileImage({ username, selfUsername, inPost }) {
  const [loading, setLoading] = useState(false);
  const [loaded, setLoaded] = useState(false);
  const [image, setImage] = useState(<img alt="" onClick={handleClick} onMouseEnter={handleMouseEnter} onMouseLeave={handleMouseLeave}></img>);
  const [hideOverlay, setHideOverlay] = useState(true);
  const ref = useRef(null);

  if (!loaded && !loading) {
    setLoading(true);
    getProfilePicture(username).then((result) => {
      const [url, ok] = result;
      if (ok) {
        setImage(<img src={url} alt={"An image of " + username} onClick={handleClick} onMouseEnter={handleMouseEnter} onMouseLeave={handleMouseLeave}></img>)
      }
      setLoading(false);
      setLoaded(true);
    });
  }

  function handleClick() {
    if (!inPost && username === selfUsername) {
      ref.current.click()
    }
  }

  function handleChange(e) {
    setProfilePicture(username, e.target.files[0]).then(() => setLoaded(false));
  }

  function handleMouseEnter() {
    if (!inPost) {
      console.log("Hey");
      setHideOverlay(false);
    }
  }

  function handleMouseLeave() {
    if (!inPost) {
      console.log("Ho");
      setHideOverlay(true);
    }
  }

  const overlay =
    <div hidden={hideOverlay} id="overlay">
      <span>Change profile picture</span>
    </div>;

  if (inPost || username !== selfUsername) {
    return (
      <div id="profile-image" className={inPost ? "in-post" : ""} >
        {image}
        <input ref={ref} onChange={handleChange} type="file"></input>
      </div>
    );
  } else {
    return (
      <div id="profile-image" className={inPost ? "in-post" : ""}>
        {image}
        {overlay}
        <input ref={ref} onChange={handleChange} type="file"></input>
      </div>
    );
  }
}
