// Description: This file contains all the API calls that the client makes to the server.

/** API endpoint
 */
const api = 'http://localhost:8081';

/** POST a post to the backend server
  * @param {String} postText 
  * @returns {Promise<number>} HTTP response status
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
  return response.status; 
}

/** GET posts from the backend server.
 * @returns {Promise<[[{String, String}], number]>} posts[[{username, postText}], httpStatus]
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
