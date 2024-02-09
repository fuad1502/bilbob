export default class FormHandler {
  constructor(sectionName) {
    this.form = document.querySelector(`#${sectionName} form`);
    this.sectionName = sectionName;
    this.inputs = [];
    this.errors = [];
    this.valueMissingMessages = [];
    this.patternMismatchMessages = [];
    return this;
  }

  // Add an input to the form handler.
  addFormInput(inputName, valueMissingMessage, patternMismatchMessage) {
    let input = document.querySelector(`#${this.sectionName} form input[name="${inputName}"]`);
    let error = document.querySelector(`#${this.sectionName} form input[name="${inputName}"] + span.error`);
    this.inputs.push(input);
    this.errors.push(error);
    this.valueMissingMessages.push(valueMissingMessage);
    this.patternMismatchMessages.push(patternMismatchMessage);
    let formHandler = this;
    let index = this.inputs.length - 1;
    input.addEventListener('invalid', function(event) {formHandler.invalidHandler(index)});
    input.addEventListener('input', function(event) {formHandler.inputHandler(index)});
  }

  // Clear the error message for the input at the given index once the input is valid.
  inputHandler(inputIndex) {
    let input = this.inputs[inputIndex];
    let error = this.errors[inputIndex];
    if (input.validity.valid) {
      error.textContent = '';
      error.className = 'error';
    }
  }

  // Show the error message for the input at the given index.
  invalidHandler(inputIndex) {
    // If there is an error in the previous input, don't show the error in the current input.
    if (!this.noErrorTillIndex(inputIndex - 1)) {
      return;
    }
    // Clear all errors from the current input to the end.
    this.clearErrorFromIndex(inputIndex + 1);
    // Show the error in the current input.
    let input = this.inputs[inputIndex];
    let error = this.errors[inputIndex];
    if (input.validity.valueMissing) {
      error.textContent = this.valueMissingMessages[inputIndex];
      error.className = 'error active';
    } else if (input.validity.patternMismatch) {
      error.textContent = this.patternMismatchMessages[inputIndex];
      error.className = 'error active';
    }
  }

  // Clear all errors from the given index to the end.
  clearErrorFromIndex(index) {
    for (let i = index; i < this.errors.length; i++) {
      this.errors[i].textContent = '';
      this.errors[i].className = 'error';
    }
  }

  // Check if there is an error in any of the inputs from the start till the given index.
  noErrorTillIndex(index) {
    for (let i = 0; i <= index; i++) {
      if (!this.inputs[i].validity.valid) {
        return false;
      }
    }
    return true;
  }
}

