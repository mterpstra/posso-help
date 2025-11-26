package sextag

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

func TestSexTag(t *testing.T) {
  sex := New()
  tests := []TestCase{
    {"male",             true, "male",   0},
    {"m",                true, "male",   0},
    {"female",           true, "female", 0},
    {"f",                true, "female", 0},
    {"data f around",    true, "female", 0},
    {"data male around", true, "male",   0},
  }
  for index, test := range tests {
    found := sex.Parse(test.Input)
    assert.Equal(t, test.Found, found, 
                 fmt.Sprintf("test: %d", index))
    assert.Equal(t, test.Value, sex.Value(), 
                 fmt.Sprintf("test: %d", index))
    assert.Equal(t, test.ValueInt, sex.ValueAsInt(), 
                 fmt.Sprintf("test: %d", index))
  }
}
