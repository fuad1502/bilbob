function showSignup() {
  document.getElementById('signup').style.display = 'flex';
  document.getElementById('login').style.display = 'none';
}

function showLogin() {
  document.getElementById('signup').style.display = 'none';
  document.getElementById('login').style.display = 'flex';
}

// Handle name validation
const signup_name = document.querySelector('#signup form input[name="name"]');
const signup_name_error = document.querySelector('#signup form input[name="name"] + span.error');

signup_name.addEventListener('input', function(event) {
  if (signup_name.validity.valid) {
    signup_name_error.textContent = '';
    signup_name_error.className = 'error';
  }
});

signup_name.addEventListener('invalid', function(event) {
  if (signup_name.validity.valueMissing) {
    signup_name_error.textContent = 'You need to enter a name.';
    signup_name_error.className = 'error active';
  } else if (signup_name.validity.patternMismatch) {
    signup_name_error.textContent = 'Name can only contain letters and spaces.';
    signup_name_error.className = 'error active';
  } else {
    signup_name_error.textContent = '';
    signup_name_error.className = 'error';
  }
});

// Handle username validation
const signup_username = document.querySelector('#signup form input[name="username"]');
const signup_username_error = document.querySelector('#signup form input[name="username"] + span.error');

signup_username.addEventListener('input', function(event) {
  if (signup_username.validity.valid) {
    signup_username_error.textContent = '';
    signup_username_error.className = 'error';
  }
});

signup_username.addEventListener('invalid', function(event) {
  if (signup_username.validity.valueMissing) {
    signup_username_error.textContent = 'You need to enter a username.';
    signup_username_error.className = 'error active';
  } else if (signup_username.validity.patternMismatch) {
    signup_username_error.textContent = 'Username can only contain letters and numbers.';
    signup_username_error.className = 'error active';
  } else {
    signup_username_error.textContent = '';
    signup_username_error.className = 'error';
  }
});

// Handle password validation
const signup_password = document.querySelector('#signup form input[name="password"]');
const signup_password_error = document.querySelector('#signup form input[name="password"] + span.error');

signup_password.addEventListener('input', function(event) {
  if (signup_password.validity.valid) {
    signup_password_error.textContent = '';
    signup_password_error.className = 'error';
  }
});

signup_password.addEventListener('invalid', function(event) {
  if (signup_password.validity.valueMissing) {
    signup_password_error.textContent = 'You need to enter a password.';
    signup_password_error.className = 'error active';
  } else if (signup_password.validity.patternMismatch) {
    signup_password_error.textContent = 'Password must be at least 8 characters long and contains a lowercase letter, an uppercase letter, a number, and a special character.';
    signup_password_error.className = 'error active';
  } else {
    signup_password_error.textContent = '';
    signup_password_error.className = 'error';
  }
});

const signup_form = document.querySelector('#signup form');
const confirm_password = document.querySelector('#signup form input[name="confirm_password"]');
const confirm_password_error = document.querySelector('#signup form input[name="confirm_password"] + span.error');

// Handle password confirmation before submitting the form
signup_form.addEventListener('submit', function(event) {
  if (signup_password.value !== confirm_password.value) {
    confirm_password_error.textContent = 'Passwords do not match.';
    confirm_password_error.className = 'error active';
    event.preventDefault();
  } else {
    confirm_password_error.textContent = '';
    confirm_password_error.className = 'error';
  }
});
