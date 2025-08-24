package date

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestParseAsDateLine(t *testing.T) {
  date, found := ParseAsDateLine("20/01")
  assert.Equal(t, found, true)
  assert.Equal(t, date, "2025-01-20T00:00:00Z")

  date, found = ParseAsDateLine("01/11")
  assert.Equal(t, found, true)
  assert.Equal(t, date, "2025-11-01T00:00:00Z")

  date, found = ParseAsDateLine("2/8")
  assert.Equal(t, found, true)
  assert.Equal(t, date, "2025-08-02T00:00:00Z")

  date, found = ParseAsDateLine("  2/8  ")
  assert.Equal(t, found, true)
  assert.Equal(t, date, "2025-08-02T00:00:00Z")

  date, found = ParseAsDateLine("text before  2/8  text after")
  assert.Equal(t, found, true)
  assert.Equal(t, date, "2025-08-02T00:00:00Z")
}
