// Description: This file contains all the API calls that the client makes to the server.

/** API endpoint
 */
const api = 'http://localhost:8081';

/** POST a post to the backend server
  * @param {String} postText 
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
