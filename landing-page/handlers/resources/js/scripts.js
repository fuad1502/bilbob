import FormHandler from './form-handler.js';
import { checkUsernameExists, registerUser, login } from './api-calls.js';
import { setFloatingMessageError, setFloatingMessageSuccess } from './floating-message.js';
import { webappUrl } from './urls.js';

// TODO: Check if stored session_id is still authorized, if it is, redirect to webapp

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

// Create a form handler for the login form.
const loginFormHandler = new FormHandler('login');
loginFormHandler.addFormInput('username', 'Please input a username.', '');
loginFormHandler.addFormInput('password', 'Please input a password.', '');

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
    confirmPasswordError.className = 'error active';
    return;
  }
  confirmPasswordError.textContent = '';
  confirmPasswordError.className = '';

  // Check if the username is already taken.
  let [exists, ok] = await checkUsernameExists(username.value);
  if (!ok) {
    setFloatingMessageError('Internal server error');
    return;
  }
  if (exists) {
    signupFormHandler.setErrorMessage('username', 'Sorry, username is not available');
    return;
  }

  // Register the user through the API.
  ok = await registerUser({ username: username.value, password: password.value, name: name.value, animal: animal.value });
  if (ok) {
    showLogin();
    setFloatingMessageSuccess('User account created successfully!');
  } else {
    setFloatingMessageError('Internal server error');
  }
});

// Handle login form submission.
const login_username = document.querySelector('#login form input[name="username"]');
const login_password = document.querySelector('#login form input[name="password"]');
document.querySelector('#login form').addEventListener('submit', async function(event) {
  event.preventDefault();
  const [verified, ok] = await login({ username: login_username.value, password: login_password.value });
  if (!ok) {
    setFloatingMessageError('Internal server error');
    return;
  }
  if (!verified) {
    setFloatingMessageError('Invalid username or password');
    return;
  }
  setFloatingMessageSuccess('Login successful');
  window.location.replace(webappUrl);
});
