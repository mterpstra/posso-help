package email

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestSendRegistrationEmail(t *testing.T) {
  to := "mark.terpstra.mt@gmail.com"
  code := "123456"
  err := SendRegistrationEmail(to, code)
  assert.Nil(t, err, "error was not nil")
}
