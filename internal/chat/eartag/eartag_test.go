package eartag

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

func TestEarTag(t *testing.T) {
  ear := New()
  tests := []TestCase{
    {"4444",               true,  "4444",      4444},
    {"88888888",           true,  "88888888",  88888888},
    {" 666666 ",           true,  "666666",    666666},
    {"before 55555 after", true,  "55555",     55555},
    {"12",                 false, "",          0},
    {"123456789",          false, "",          0},
  }
  for index, test := range tests {
    found := ear.Parse(test.Input)
    assert.Equal(t, test.Found, found, 
                 fmt.Sprintf("test: %d", index))
    assert.Equal(t, test.Value, ear.Value(), 
                 fmt.Sprintf("test: %d", index))
    assert.Equal(t, test.ValueInt, ear.ValueAsInt(), 
                 fmt.Sprintf("test: %d", index))
  }
}
