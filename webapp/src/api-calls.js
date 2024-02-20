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
 * @typedef {{username: string, state: string}} FollowsInfo
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
  const payload = await response.json();
  return [payload, response.status];
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

/** POST a post to the backend server
  * @param {String} postText the text to post
  * @returns {Promise<ApiResult<null>>}
  * a Promise for a null object wrapped with an HTTP status code. 
  * 201 (Created) status code signifies a sucessful POST.
  */
export async function postPost(postText) {
  const requestBody = JSON.stringify({"postText": postText});
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

/** GET the followings information of the logged in user filtered to a certain 
  * user.
  * @returns {Promise<Result<FollowsInfo>>} 
  * a Promise for single FollowsInfo object in a Result struct. 
  */
export async function getFollowsUser(username) {
  const [payload, status] = await genericGET('/follows?username=' + username, true);
  if (status !== 200) {
    return [null, false];
  }
  return [payload, true];
}
