document.addEventListener('DOMContentLoaded', function() {
  const register_button = document.getElementById('register-button');

  register_button.addEventListener('click', function(e) {
    RegisterButtonClicked(e);
  });
});

function RegisterButtonClicked(e) {
  let form_valid = true;
  const email_input = document.getElementById('email-input');
  const password_input = document.getElementById('password-input');
  const confirm_input = document.getElementById('confirm-input');
  const error_message = document.getElementById('error-message');
  const register_form = document.getElementById('register-form');
  error_message.innerHTML = "";

  if (!isValidEmail(email_input.value)) {
    error_message.innerHTML = "Email is invalid";
    form_valid = false;
  }

  if (form_valid) {
    const password_valid = isValidPassword(password_input.value);
    if (!password_valid.isValid) {
      error_message.innerHTML = password_valid.message;
      form_valid = false;
    }
  }

  if (form_valid) {
    if (password_input.value != confirm_input.value) {
      error_message.innerHTML = "Passwords do not match";
      form_valid = false;
    }
  }
  
  if (!form_valid) {
    error_message.style.display = "block";
    e.stopPropagation();
    e.preventDefault();
    return;
  }

  register_form.requestSubmit();
}

function isValidEmail(email) {
  // Regular expression for basic email validation
  // This regex checks for:
  // - One or more characters that are not whitespace or '@' before '@'
  // - An '@' symbol
  // - One or more characters that are not whitespace or '@' after '@'
  // - A period '.'
  // - One or more characters that are not whitespace or '@' after the period
  const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;

  // Test the email string against the regex
  return emailRegex.test(email);
}

function isValidPassword(password) {
  // Define criteria for a strong password
  const minLength = 8;
  const hasUppercase = /[A-Z]/;
  const hasLowercase = /[a-z]/;
  const hasDigit = /\d/;
  const hasSpecialChar = /[!@#$%^&*()_+\-=[\]{};':"\\|,.<>/?]/; // Common special characters

  // Check minimum length
  if (password.length < minLength) {
    return {
      isValid: false,
      message: `Password must be at least ${minLength} characters long.`
    };
  }

  // Check for uppercase letter
  if (!hasUppercase.test(password)) {
    return {
      isValid: false,
      message: 'Password must contain at least one uppercase letter.'
    };
  }

  // Check for lowercase letter
  if (!hasLowercase.test(password)) {
    return {
      isValid: false,
      message: 'Password must contain at least one lowercase letter.'
    };
  }

  // Check for digit
  if (!hasDigit.test(password)) {
    return {
      isValid: false,
      message: 'Password must contain at least one digit.'
    };
  }

  // Check for special character
  if (!hasSpecialChar.test(password)) {
    return {
      isValid: false,
      message: 'Password must contain at least one special character.'
    };
  }

  // If all criteria are met
  return {
    isValid: true,
    message: 'Password is strong.'
  };
}
