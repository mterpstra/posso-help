package chat

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestRainMessage(t *testing.T) {
  input := "02/04 40mm\n25/12 5MM\n15/10 22 mm\nanystring"
  rm := &RainMessage{}
  assert.Equal(t, rm.Parse(input), true, "Could not parse rain message")
  assert.Equal(t, len(rm.Entries), 3, "Wrong number of rain entries")
  assert.Equal(t, rm.Total, 67, "Wrong rain total")
}
