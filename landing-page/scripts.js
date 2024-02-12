import FormHandler from './form-handler.js';

const api = 'http://localhost:8081';

// Show the signup form and hide the login form.
function showSignup() {
  document.getElementById('signup').style.display = 'flex';
  document.getElementById('login').style.display = 'none';
}
document.querySelector('#login a').addEventListener('click', showSignup);

// Show the login form and hide the signup form.
function showLogin() {
  document.getElementById('signup').style.display = 'none';
  document.getElementById('login').style.display = 'flex';
}
document.querySelector('#signup a').addEventListener('click', showLogin);

// Create a form handler for the signup form.
const signupFormHandler = new FormHandler('signup');
signupFormHandler.addFormInput('name', 'Please input your name.', 'Name can only contain letters and spaces.');
signupFormHandler.addFormInput('username', 'Please input a username.', 'Username can only contain letters, numbers, and underscores.');
signupFormHandler.addFormInput('password', 'Please input a password.', 'Password must be at least 8 characters long and contain at least one letter, one number, and one special character.');

// Handle signup form submission.
const password = document.querySelector('#signup form input[name="password"]');
const confirmPassword = document.querySelector('#signup form input[name="confirm-password"]');
const confirmPasswordError = document.querySelector('#signup form input[name="confirm-password"] + span.error');
const username = document.querySelector('#signup form input[name="username"]');
const name = document.querySelector('#signup form input[name="name"]');
const animal = document.querySelector('#signup form select[name="animal"]');
document.querySelector('#signup form').addEventListener('submit', async function(event) {
  event.preventDefault();

  // Check if the passwords match.
  if (password.value !== confirmPassword.value) {
    confirmPasswordError.textContent = 'Passwords do not match.';
    return;
  }
  confirmPasswordError.textContent = '';

  // Check if the username is already taken.
  let url = api + '/users/' + username.value;
  let response = await fetch(url);
  if (response.status != 200) {
    // TODO: Display internal server error
    return;
  }
  let responseBody = await response.json();
  if (responseBody.exists) {
    signupFormHandler.setErrorMessage('username', 'Sorry, username is not available');
    return;
  }

  // Register the user through the API.
  url = api + '/users';
  let requestBody = JSON.stringify({username: username.value, password: password.value, name: name.value, animal: animal.value});
  response = await fetch(url, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: requestBody
  });

  // Check API call response
  if (response.status === 201) {
    // Redirect to the login page.
    showLogin();
    // TODO: Display a success message.
    console.log('User account created successfully.');
  } else {
    // TODO: Display an error message.
    console.log('Failed to create user account.');
  }
});

// Create a form handler for the login form.
const loginFormHandler = new FormHandler('login');
loginFormHandler.addFormInput('username', 'Please input a username.', '');
loginFormHandler.addFormInput('password', 'Please input a password.', '');

// Handle login form submission.
const login_username = document.querySelector('#login form input[name="username"]');
const login_password = document.querySelector('#login form input[name="password"]');
document.querySelector('#login form').addEventListener('submit', function(event) {
  // TODO: Check if the username and password are correct.
});
