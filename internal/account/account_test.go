package account

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestFindAccountByPhoneNumber(t *testing.T) {
  phone := "16166100305"
  account, err := FindAccountByPhoneNumber(phone)
  println(phone, account)
  assert.Nil(t, err)
  assert.Greater(t, len(account), 5)

  phone = "12123451234"
  account, err = FindAccountByPhoneNumber(phone)
  println(phone, account)
  assert.Nil(t, err)
  assert.Greater(t, len(account), 5)
}

func TestVariants(t *testing.T) {
  variants := getAllPhoneNumberVariants("12123451234")
  for _, variant := range variants {
    println(variant)
  }
}
