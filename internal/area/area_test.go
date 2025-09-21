package area

import (
  "fmt"
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestReadAreas(t *testing.T) {
  areas, err := readAreas("16166100305")
  assert.Nil(t, err)
  for _, area := range areas {
    fmt.Printf("FOUND: %s: %s\n", area.Name, area.Matches)
  }
}

func TestParseAsAreaLine(t *testing.T) {
  area, found := ParseAsAreaLine("Jensen")
  assert.True(t, found)
  assert.Equal(t, "Jensen Beach", area)

  area, found = ParseAsAreaLine("something before Jup something after")
  assert.True(t, found)
  assert.Equal(t, "Jupiter", area)
}

func TestAddArea(t *testing.T) {
  err := AddArea("16166100305", "Miami")
  assert.Nil(t, err)
}

