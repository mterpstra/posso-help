package puretag

import (
  "fmt"
  "testing"
  "github.com/stretchr/testify/assert"
)

type TestCase struct {
  Input string
  Found bool 
  Value string
  ValueInt int
}

func TestPureTag(t *testing.T) {
  pure := New()
  tests := []TestCase{
    {"fft",  true, "fft", 0},
    {"pure", true, "fft", 0},
    {" data pure stuff", true, "fft", 0},
  }
  for index, test := range tests {
    found := pure.Parse(test.Input)
    assert.Equal(t, test.Found, found, 
                 fmt.Sprintf("test: %d", index))
    assert.Equal(t, test.Value, pure.Value(), 
                 fmt.Sprintf("test: %d", index))
    assert.Equal(t, test.ValueInt, pure.ValueAsInt(), 
                 fmt.Sprintf("test: %d", index))
  }
}
