// Description: This file contains all the API calls that the client makes to the server.
import { api } from "./urls.js";

/** Checks if a username exists in the database
 * @param {String} username.
 * @return {Promise<[exists: boolean, ok: boolean]>}
 */
export async function checkUsernameExists(username) {
  const url = api + '/exists?username=' + username;
  const response = await fetch(url);
  if (response.status != 200) {
    return [false, false];
  }
  const responseBody = await response.json();
  return [responseBody.exists, true];
}

/** Registers a user to the database
 * @param {{String, String, String, String}} user{username, password, name, animal} 
 * @return {Promise<boolean>} Promise<ok>
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
  if (response.status == 201) {
    return true;
  } else {
    return false;
  }
}

/** Login the user if username password combination is verified
 * @param {{String, String}} user{username, password}
 * @return {Promise<[verified: boolean, ok: boolean]>}
 */
export async function login(user) {
  const url = api + '/login?username=' + user.username + '&password=' + user.password;
  const response = await fetch(url,
    { credentials: 'include' }
  );
  if (response.status != 200) {
    return [false, false];
  }
  const responseBody = await response.json();
  return [responseBody.verified, true];
}
