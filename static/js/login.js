document.addEventListener('DOMContentLoaded', function() {
  const login_button    = document.getElementById('login-button');

  login_button.addEventListener('click', function(e) {
    LoginButtonClicked(e);
  });
});

function LoginButtonClicked(e) {
  console.log("LoginButtonClicked", e);
}
