document.addEventListener('DOMContentLoaded', function () {
    const form = document.querySelector('form');
    const usernameInput = document.querySelector('input[type="text"][placeholder="Enter your username"]');
    const emailInput = document.querySelector('input[type="text"][placeholder="Enter your email"]');
    const passwordInput = document.querySelector('input[type="password"][placeholder="Create password"]');
    const confirmPasswordInput = document.querySelector('input[type="password"][placeholder="Confirm password"]');
    const termsCheckbox = document.querySelector('input[type="checkbox"]');
    
    form.addEventListener('submit', function (event) {
      event.preventDefault();
  
      // Validate username, email, password, and terms acceptance
      if (!validateUsername(usernameInput.value)) {
        alert('Please enter a valid username.');
        return;
      }
      if (!validateEmail(emailInput.value)) {
        alert('Please enter a valid email address.');
        return;
      }
      if (!validatePassword(passwordInput.value)) {
        alert('Password must be at least 6 characters long.');
        return;
      }
      if (passwordInput.value !== confirmPasswordInput.value) {
        alert('Passwords do not match.');
        return;
      }
      if (!termsCheckbox.checked) {
        alert('Please accept the terms & conditions.');
        return;
      }
  
      // If all validations pass, submit the form
      form.submit();
    });
  
    function validateUsername(username) {
      // You can add custom validation logic for the username here
      return username.length > 0;
    }
  
    function validateEmail(email) {
      // You can add custom validation logic for the email here
      const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
      return emailRegex.test(email);
    }
  
    function validatePassword(password) {
      // You can add custom validation logic for the password here
      return password.length >= 6;
    }
  });
  