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
 * @typedef {()=>Promise<ApiResult<T>>} ApiClosure<T>
 */

/**
 * @typedef {{username: string, postText: string}} Post
 */

/** Wrap an API call with a redirect if the API call returns 401 (Unauthorized)
 * @template T
 * @param {ApiClosure<T>} apiCall
 * @param {string} redirectUrl
 * @returns {Promise<ApiResult<T>>}
 */
export async function redirectWrap(apiCall, redirectUrl) {
  const [body, status] = await apiCall();
  if (status === 401) {
    window.location.replace(redirectUrl);
  }
  return [body, status]
}

/** POST a post to the backend server
  * @param {String} postText 
  * @returns {Promise<ApiResult<null>>}
  */
export async function postPost(postText) {
  const requestBody = JSON.stringify({"postText": postText});
  const url = api + '/posts'; 
  const response = await fetch(url, {
    method: 'POST',
    credentials: 'include',
    headers: {
      'Content-Type': 'application/json'
    },
    body: requestBody
  });
  return [null, response.status]; 
}

/** GET posts from the backend server.
 * @returns {Promise<ApiResult<[Post]>>}
 */
export async function getPosts() {
  const url = api + '/posts';
  const response = await fetch(url, {
    credentials: 'include'
  });
  if (response.status !== 200) {
    return [[], response.status];
  }
  const payload = await response.json();
  return [payload, response.status];
}

export async function getUsername() {
  const url = api + '/authorize';
  const response = await fetch(url, {
    credentials: 'include'
  });
  if (response.status !== 200) {
    return ["", response.status];
  }
  const payload = await response.json();
  return [payload.username, response.status];
}

export async function getUserInfo(username) {
  const url = api + '/users/' + username + '/info';
  const response = await fetch(url, {
    credentials: 'include'
  });
  if (response.status !== 200) {
    return [null, response.status];
  }
  const payload = await response.json();
  return [payload, response.status];
}

export async function getUserFollows(username) {
  const url = api + '/users/' + username + '/follows';
  const response = await fetch(url, {
    credentials: 'include'
  });
  if (response.status !== 200) {
    return [null, response.status];
  }
  const payload = await response.json();
  return [payload.followsState, response.status];
}
