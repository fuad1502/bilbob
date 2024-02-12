const floatingMessage = document.getElementById('floating-message');

function clearFloatingMessage() {
  floatingMessage.innerHTML = '';
  floatingMessage.className = '';
}

function setFloatingMessage(message, type, timeout) {
  floatingMessage.innerHTML = message;
  floatingMessage.className = type;
  setTimeout(clearFloatingMessage, timeout);
}

export function setFloatingMessageSuccess(message, timeout = 3000) {
  setFloatingMessage(message, 'success', timeout);
}

export function setFloatingMessageError(message, timeout = 3000) {
  setFloatingMessage(message, 'error', timeout);
}
