// Description: This file contains all the API calls that the client makes to the server.

/** API endpoint
 * @type string
 */
const api = 'http://localhost:8081';
/** Landing Page URL
 * @type string
 */
export const LandingPageUrl = 'http://localhost:8080';

/**
 * @template T
 * @typedef {[payload: T, httpStatus: number]} ApiResult<T>
 */

/**
 * @template T
 * @typedef {[res: T, ok: boolean]} Result<T>
 */

/**
 * @typedef {{username: string, postText: string}} Post
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
    window.location.replace(LandingPageUrl);
  }
  const contentType = response.headers.get("content-type");
  if (contentType && contentType.indexOf("application/json") !== -1) {
    const payload = await response.json();
    return [payload, response.status];
  } else {
    return [null, response.status];
  }
}

/** 
 * A generic POST request to the API endpoint.
 * @param {string} route API route.
 * @param {boolean} withCredentials a flag indicating whether to include Cookie or not.
 * @returns {Promise<ApiResult<any>>} A Promise for the returned payload and the HTTP status.
 * If the HTTP status code is 401 (Unauthorized), it automatically redirects to LandingPageUrl.
 */
async function genericPOST(route, withCredentials, requestBody) {
  const url = api + route;
  let init = {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: requestBody
  }
  if (withCredentials) {
    init['credentials'] = 'include';
  }
  const response = await fetch(url, init);
  if (response.status === 401) {
    window.location.replace(LandingPageUrl);
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
    window.location.replace(LandingPageUrl);
  }
  const payload = await response.json();
  return [payload, response.status];
}

/** POST a post to the backend server
  * @param {String} postText the text to post
  * @returns {Promise<ApiResult<null>>}
  * a Promise for a null object wrapped with an HTTP status code. 
  * 201 (Created) status code signifies a sucessful POST.
  */
export async function postPost(postText) {
  const requestBody = JSON.stringify({ "postText": postText });
  const [payload, status] = await genericPOST('/posts', true, requestBody);
  if (status !== 201) {
    return [null, false];
  }
  return [payload, true];
}

/** GET posts from the backend server.
 * @returns {Promise<Result<[Post]>>}
 * a Promise for an array of Post objects inside a Result struct.
 */
export async function getPosts() {
  const [payload, status] = await genericGET('/posts', true);
  if (status !== 200) {
    return [[], false];
  }
  return [payload, true];
}

/** GET username of logged in session.
  * @returns {Promise<Result<string>>}
  * a Promise for a username string inside a Result struct. 
  */
export async function getUsername() {
  const [payload, status] = await genericGET('/authorize', true);
  if (status !== 200) {
    return ["", false];
  }
  return [payload.username, true];
}

/** GET information about a user.
  * @returns {Promise<Result<UserInfo>} 
  * a Promise for single UserInfo object inside a Result struct. 
  */
export async function getUserInfo(username) {
  const [payload, status] = await genericGET('/users/' + username, true);
  if (status !== 200) {
    return [null, false];
  }
  return [payload, true];
}

/** GET the the following information if 'username' follows 'follows' 
  * @returns {Promise<Result<string>>} 
  * a Promise for single FollowsInfo object in a Result struct. 
  */
export async function getFollowState(username, follows) {
  if (username === follows) {
    return [null, true];
  }
  const [payload, status] = await genericGET('/followings/' + username + '?follows=' + follows, true);
  if (status === 404) {
    return ["no", true];
  }
  if (status !== 200) {
    return [null, false];
  }
  return [payload.state, true];
}

export async function requestFollow(username, follows) {
  const requestBody = JSON.stringify({ "username": username, "follows": follows, "state": "requested" });
  const [_, status] = await genericPOST("/followings/" + username + "?follows=" + follows, true, requestBody);
  if (status !== 201) {
    return false;
  }
  return true;
}

export async function unfollow(username, follows) {
  const [_, status] = await genericDELETE("/followings/" + username + "?follows=" + follows, true);
  if (status !== 200) {
    return false;
  }
  return true;
}

export async function unrequest(username, follows) {
  return await unfollow(username, follows);
}

export async function logout() {
  const [_, status] = await genericGET("/logout", true);
  if (status !== 200) {
    return false;
  }
  window.location.replace(LandingPageUrl);
  return true;
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
