package breedtag

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

func TestBreedTag(t *testing.T) {
  breed := New()
  tests := []TestCase{
    {"nelore",             true, "nelore", 0},
    {"nalore",             true, "nelore", 0},
    {"angus",              true, "angus",  0},
    {"data nelore around", true, "nelore", 0},
    {"data nalore around", true, "nelore", 0},
    {"data angus around",  true, "angus", 0},
  }
  for index, test := range tests {
    found := breed.Parse(test.Input)
    assert.Equal(t, test.Found, found, 
                 fmt.Sprintf("test: %d", index))
    assert.Equal(t, test.Value, breed.Value(), 
                 fmt.Sprintf("test: %d", index))
    assert.Equal(t, test.ValueInt, breed.ValueAsInt(), 
                 fmt.Sprintf("test: %d", index))
  }
}
