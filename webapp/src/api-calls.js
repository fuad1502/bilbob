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

/**
 * @typedef {{username: string, name: string, animal: string}} UserInfo
 */

/**
 * @typedef {{username: string, state: string}} FollowsInfo
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
  * @param {String} postText the text to post
  * @returns {Promise<ApiResult<null>>}
  * a Promise for a null object wrapped with an HTTP status code. 
  * 201 (Created) status code signifies a sucessful POST.
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
 * a Promise for an array of Post objects wrapped with an HTTP status code. 
 * Array is set to empty array if HTTP status code is not 200 (OK).
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

/** GET username of logged in session.
  * @returns {Promise<ApiResult<string>>}
  * a Promise for a username string wrapped with an HTTP status code. 
  * username is set to empty string if HTTP status code is not 200 (OK).
  */
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

/** GET information about a user.
  * @returns {Promise<ApiResult<UserInfo>} 
  * a Promise for single UserInfo object wrapped with an HTTP status code. 
  * Object is set to null if HTTP status code is not 200 (OK).
  */
export async function getUserInfo(username) {
  const url = api + '/users/' + username;
  const response = await fetch(url, {
    credentials: 'include'
  });
  if (response.status !== 200) {
    return [null, response.status];
  }
  const payload = await response.json();
  return [payload, response.status];
}

/** GET the followings information of the logged in user filtered to a certain 
  * user.
  * @returns {Promise<ApiResult<FollowsInfo>>} 
  * a Promise for single FollowsInfo object wrapped with an HTTP status code. 
  * Object is set to null if HTTP status code is not 200 (OK).
  */
export async function getFollowsUser(username) {
  const url = api + '/users/' + 'follows?username=' + username;
  const response = await fetch(url, {
    credentials: 'include'
  });
  if (response.status !== 200) {
    return [null, response.status];
  }
  const payload = await response.json();
  return [payload, response.status];
}
