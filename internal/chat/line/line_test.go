package line

import (
  "fmt"
  "testing"
  "github.com/stretchr/testify/assert"

  "posso-help/internal/chat/eartag"
  "posso-help/internal/chat/sextag"
  "posso-help/internal/chat/breedtag"
)

func TestLine(t *testing.T) {
  lineParser := NewLineParser().
                MustHave("ear",  eartag.New()).
                MustHave("sex",  sextag.New()).
                CanHave("breed", breedtag.New())

  tests := []struct {
    Input             string
    ExpectedResult    bool
    ExpectedEarTag    string
    ExpectedEarTagInt int    
    ExpectedSex       string
    ExpectedBreed     string
  }{
    {"1111 f nelore",               true, "1111", 1111, "female", "nelore"},
    {"--  f 1111 nelore  --",       true, "1111", 1111, "female", "nelore"},
    {"  nalore 1111 f",             true, "1111", 1111, "female", "nelore"},
    {"male nalore 1111 f 3211 bla", true, "1111", 1111, "male",   "nelore"},
  }
  for index, test := range tests {
    found := lineParser.Parse(test.Input)
    assert.Equal(t, test.ExpectedResult, found, fmt.Sprintf("test: %d", index))
    assert.Equal(t, test.ExpectedEarTag,
                 lineParser.Value("ear"),
                 fmt.Sprintf("test: %d", index))
    assert.Equal(t, test.ExpectedEarTagInt,
                 lineParser.ValueAsInt("ear"),
                 fmt.Sprintf("test: %d", index))
    assert.Equal(t, test.ExpectedSex,
                 lineParser.Value("sex"),
                 fmt.Sprintf("test: %d", index))
    assert.Equal(t, test.ExpectedBreed,
                 lineParser.Value("breed"),
                 fmt.Sprintf("test: %d", index))
  }
}
