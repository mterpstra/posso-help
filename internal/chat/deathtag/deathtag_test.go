package deathtag

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

func TestDeathTag(t *testing.T) {
  death := New()
  tests := []TestCase{
    {"aborto",             true, "aborto",       0},
    {"morreu",             true, "morreu",       0},
    {"nasceu morto",       true, "nasceu-morto", 0},
    {"morto",              true, "morto",        0},
    {"natimortos",         true, "natimortos",   0},
    {"natimorto",          true, "natimorto",    0},
  }
  for index, test := range tests {
    found := death.Parse(test.Input)
    assert.Equal(t, test.Found, found, 
                 fmt.Sprintf("test: %d", index))
    assert.Equal(t, test.Value, death.Value(), 
                 fmt.Sprintf("test: %d", index))
    assert.Equal(t, test.ValueInt, death.ValueAsInt(), 
                 fmt.Sprintf("test: %d", index))
  }
}
