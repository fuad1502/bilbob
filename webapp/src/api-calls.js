// Description: This file contains all the API calls that the client makes to the server.

/** API endpoint
 * @type string
 */
const api = process.env.REACT_APP_PROTOCOL + process.env.REACT_APP_HOSTNAME + ':' + process.env.REACT_APP_API_PORT;
/** Landing Page URL
 * @type string
 */
const landingPageUrl = process.env.REACT_APP_PROTOCOL + process.env.REACT_APP_HOSTNAME + ':' + process.env.REACT_APP_LP_PORT + '/' + process.env.REACT_APP_LP_PATH;

/**
 * @template T
 * @typedef {[payload: T, httpStatus: number]} ApiResult<T>
 */

/**
 * @template T
 * @typedef {[res: T, ok: boolean]} Result<T>
 */

/**
 * @typedef {{username: string, postText: string, postDate: string}} Post
 */

/**
 * @typedef {{username: string, name: string, animal: string}} UserInfo
 */

/**
 * @typedef {{username: string, follows: string, state: string}} FollowingInfo
 */

/** 
 * A generic GET request to the API endpoint.
 * @param {string} route API route.
 * @param {boolean} withCredentials a flag indicating whether to include Cookie or not.
 * @returns {Promise<ApiResult<any>>} A Promise for the returned payload and the HTTP status.
 * If the HTTP status code is 401 (Unauthorized), it automatically redirects to LandingPageUrl.
 */
async function genericGET(route, withCredentials) {
  const url = api + route;
  let init = {}
  if (withCredentials) {
    init['credentials'] = 'include';
  }
  const response = await fetch(url, init);
  if (response.status === 401) {
    window.location.replace(landingPageUrl);
  }
  const contentType = response.headers.get("content-type");
  if (contentType && contentType.indexOf("application/json") !== -1) {
    const payload = await response.json();
    return [payload, response.status];
  } else if (contentType && contentType.indexOf("image") !== -1) {
    const blob = await response.blob();
    return [blob, response.status];
  } else {
    return [null, response.status];
  }
}

/** 
 * A generic POST/PUT request to the API endpoint.
 * @param {string} route API route.
 * @param {boolean} withCredentials a flag indicating whether to include Cookie or not.
 * @param {any} requestBody
 * @param {String} contentType
 * @param {boolean} isPut
 * @returns {Promise<ApiResult<any>>} A Promise for the returned payload and the HTTP status.
 * If the HTTP status code is 401 (Unauthorized), it automatically redirects to LandingPageUrl.
 */
async function genericPOST(route, withCredentials, requestBody, contentType, isPut = false) {
  const url = api + route;
  let init = {
    method: 'POST',
    body: requestBody
  }
  if (isPut) {
    init['method'] = 'PUT';
  }
  if (contentType !== undefined) {
    init['headers'] = {
      'Content-Type': contentType
    };
  }
  if (withCredentials) {
    init['credentials'] = 'include';
  }
  const response = await fetch(url, init);
  if (response.status === 401) {
    window.location.replace(landingPageUrl);
  }
  const payload = await response.json();
  return [payload, response.status];
}

async function genericDELETE(route, withCredentials) {
  const url = api + route;
  let init = {
    method: 'DELETE',
  }
  if (withCredentials) {
    init['credentials'] = 'include';
  }
  const response = await fetch(url, init);
  if (response.status === 401) {
    window.location.replace(landingPageUrl);
  }
  const payload = await response.json();
  return [payload, response.status];
}

/** POST a post to the backend server
  * @param {String} username
  * @param {String} postText
  * @returns {Promise<boolean>} Returns true if post is posted.
  */
export async function postPost(username, postText) {
  const requestBody = JSON.stringify({ username: username, postText: postText });
  const [_, status] = await genericPOST('/posts', true, requestBody, 'application/json');
  if (status !== 201) {
    return false;
  }
  return true;
}

/** GET posts from the backend server.
 * @param {String} username
 * @param {boolean} includeFollowings
 * @returns {Promise<Result<[Post]>>}
 * a Promise for an array of Post objects inside a Result struct.
 */
export async function getPosts(username, includeFollowings, fromTimestamp) {
  const [payload, status] = await genericGET('/posts?username=' + username + '&includeFollowings=' + includeFollowings.toString() + '&fromTimestamp=' + fromTimestamp, true);
  if (status !== 200) {
    return [[], false];
  }
  return [payload, true];
}

/** GET username of logged in session.
  * @returns {Promise<Result<string>>}
  * a Promise for the logged in username string inside a Result struct. 
  */
export async function getUsername() {
  const [payload, status] = await genericGET('/authorize', true);
  if (status !== 200) {
    return ["", false];
  }
  return [payload.username, true];
}

/** GET information about a user.
  * @param {String} username
  * @returns {Promise<Result<UserInfo>} 
  * a Promise for a single UserInfo object inside a Result struct. 
  */
export async function getUserInfo(username) {
  const [payload, status] = await genericGET('/users/' + username, true);
  if (status !== 200) {
    return [null, false];
  }
  return [payload, true];
}

/** GET the the following state of 'username' in relation to 'follows'.
  * @param {String} username
  * @param {String} follows
  * @returns {Promise<Result<string>>} 
  * a Promise for the following state inside a Result struct. 
  */
export async function getFollowState(username, follows) {
  if (username === follows) {
    return [null, true];
  }
  const [payload, status] = await genericGET('/users/' + username + '/followings?username=' + follows, true);
  if (status === 404) {
    return ["no", true];
  }
  if (status !== 200) {
    return [null, false];
  }
  return [payload.state, true];
}

export async function getFollowings(username) {
  const [payload, status] = await genericGET('/users/' + username + '/followings', true);
  if (status !== 200) {
    return [[], false];
  }
  return [payload, true];
}

export async function requestFollow(username, follows) {
  const requestBody = JSON.stringify({ "username": follows, "state": "requested" });
  const [_, status] = await genericPOST("/users/" + username + "/followings", true, requestBody, 'application/json');
  if (status !== 201) {
    return false;
  }
  return true;
}

export async function unfollow(username, follows) {
  const [_, status] = await genericDELETE("/users/" + username + "/followings/" + follows, true);
  if (status !== 200) {
    return false;
  }
  return true;
}

export async function getFollowers(username) {
  const [payload, status] = await genericGET('/users/' + username + '/followers', true);
  if (status !== 200) {
    return [[], false];
  }
  return [payload, true];
}

export async function unrequest(username, follows) {
  return await unfollow(username, follows);
}

export async function logout() {
  const [_, status] = await genericGET("/logout", true);
  if (status !== 204) {
    return false;
  }
  window.location.replace(landingPageUrl);
  return true;
}

export async function getProfilePicture(username) {
  const [blob, status] = await genericGET("/users/" + username + "/profilePicture", true);
  if (status !== 200) {
    return ["", false];
  }
  const objURL = URL.createObjectURL(blob);
  return [objURL, true];
}

export async function setProfilePicture(username, file) {
  let formData = new FormData();
  formData.append("profile-picture", file);
  const [_, status] = await genericPOST("/users/" + username + "/profilePicture", true, formData, undefined, true);
  if (status !== 201) {
    return false;
  }
  return true;
}

export async function getMostPopularUsers() {
  const [payload, status] = await genericGET('/users?sortBy=popularity', true);
  if (status !== 200) {
    return [[], false];
  }
  return [payload, true];
}

/** GET all users that has a name or username matching the 'like' filter. 
  * @param {string} like name/username filter
  * @returns {Promise<Result<UserInfo>>} 
  * a Promise for an array of UserInfo object in a Result struct. 
  */
export async function getUsers(like) {
  const [payload, status] = await genericGET('/users?like=' + like, true);
  if (status !== 200) {
    return [[], false];
  }
  return [payload, true];
}
