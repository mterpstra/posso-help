package password

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestGetSalted(t *testing.T) {
  input := "somePassword01%"
  expected := "5c8fb4e2b681c52bb02c86dd2bea0c5d"
  actual, err := GetSalted(input)
  assert.Nil(t, err, "error was not nil")
  assert.Equal(t, expected, actual, "salted password doesn't match")
}
