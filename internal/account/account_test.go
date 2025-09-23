package account

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestFindAccountByPhoneNumber(t *testing.T) {

  account, err := FindAccountByPhoneNumber("16168956304")
  println(account)
  assert.Nil(t, err)
  assert.Greater(t, len(account), 5)
}
