package utils

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestSanitizeString(t *testing.T) {
  input := " MIXED case And with WhiteSPACE\n "
  expected := "mixed case and with whitespace"
  result := SanitizeLine(input)
  assert.Equal(t, expected, result, "lines do not match")
}

func TestCapitalizeString(t *testing.T) {
  assert.Equal(t, Capitalize("MARK"), "Mark" , "capitalize failed")
  assert.Equal(t, Capitalize("mark"), "Mark" , "capitalize failed")
  assert.Equal(t, Capitalize("Mark"), "Mark" , "capitalize failed")
  assert.Equal(t, Capitalize("mARK"), "Mark" , "capitalize failed")
}
