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

func TestSplitAndTrim(t *testing.T) {
  parts := SplitAndTrim("  zero , one , two ")
  assert.Equal(t, "zero", parts[0], "part 0 is wrong")
  assert.Equal(t, "one",  parts[1], "part 1 is wrong")
  assert.Equal(t, "two",  parts[2], "part 2 is wrong")
}
