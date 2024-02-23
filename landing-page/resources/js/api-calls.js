// Description: This file contains all the API calls that the client makes to the server.
import { api } from "./api-endpoint.js";

/** Checks if a username exists in the database
 * @param {String} username.
 * @return {Promise<[boolean, number]>} [exists, HTTP response status] Promise. Ignore exists if HTTP response status is not 200.
 */
export async function checkUsernameExists(username) {
  const url = api + '/users/' + username + '/exists';
  const response = await fetch(url);
  if (response.status != 200) {
    return [false, response.status];
  }
  const responseBody = await response.json();
  return [responseBody.exists, response.status];
}

/** Registers a user to the database
 * @param {{String, String, String, String}} user{username, password, name, animal} 
 * @return {Promise<number>} HTTP response status Promise. If successful, should return 201.
  */
export async function registerUser(user) {
  const requestBody = JSON.stringify(user);
  const url = api + '/users';
  const response = await fetch(url, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: requestBody
  });
  return response.status;
}

/** Checks if the username and password combination exists in the database
 * @param {{String, String}} user{username, password}
 * @return {Promise<[boolean, number]>} [verified, HTTP response status] Promise. Ignore verified if HTTP response status is not 200.
 */
export async function verifyUser(user) {
  const url = api + '/users/' + user.username + '/login?password=' + user.password;
  const response = await fetch(url,
    { credentials: 'include' }
  );
  if (response.status != 200) {
    return [false, response.status];
  }
  const responseBody = await response.json();
  return [responseBody.verified, response.status];
}
