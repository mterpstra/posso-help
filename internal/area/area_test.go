package area

import (
  "fmt"
  "testing"
  "github.com/stretchr/testify/assert"
)

/*
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
  err := AddArea("68ceeb2883e5fb7567af0423", "Miami", "Miami")
  assert.Nil(t, err)
}
*/

func TestLoadAreasByAccount(t *testing.T) {
  areaParser := &AreaParser{}
  err := areaParser.LoadAreasByAccount("68ceeb2883e5fb7567af0423")
  assert.Nil(t, err)
  for _, area := range areaParser.areas {
    fmt.Printf("Area: %v\n", area)
  }
}

