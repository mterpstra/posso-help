package tag

import (
  "fmt"
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestNumber(t *testing.T) {
  tag := NewNumber(4,8)
  tests := []struct {
    Input string
    Result string
    IntResult int
  }{
    {"11 m angus", "", 0},
    {"nalore 123456789 m", "", 0},
    {"nalore m something else", "", 0},
    {"1111 m angus", "1111", 1111},
    {"nalore 12345678 m", "12345678", 12345678},
    {"nalore m 99999999 something else", "99999999", 99999999},
  }
  for index, test := range tests {
    str := tag.Parse(test.Input)
    assert.Equal(t, test.Result, str, 
                 fmt.Sprintf("test: %d", index))
    assert.Equal(t, test.IntResult, tag.AsNumber(), 
                 fmt.Sprintf("test: %d", index))
  }
}

func TestString(t *testing.T) {
  tag := NewString("male", []string{"male", "m", "strange-variant"})
  tests := []struct {
    Input string
    Result string
  }{
    {"1111 angus", ""},
    {"1111 female angus", ""},
    {"1111 m angus", "male"},
    {"m 1111 angus", "male"},
    {"biggy smalls  female 1111 angus male  ", "male"},
    {"  male, 1111 angus", "male"},
    {"  1111 angus strange-variant blue moon", "male"},
  }
  for index, test := range tests {
    str := tag.Parse(test.Input)
    assert.Equal(t, test.Result, str,
                 fmt.Sprintf("test: %d", index))
  }
}

func TestTagSet(t *testing.T) {
  m := NewString("male", []string{"male", "m"})
  f := NewString("female", []string{"female", "f"})
  tagSet := NewStringSet(m,f)
  tests := []struct {
    Input string
    Result string
  }{
    {"female", "female"},
    {"f", "female"},
    {"male", "male"},
    {"m", "male"},
    {"1111 m nalore", "male"},
    {"1111 female nalore", "female"},
    {"11118888 male nalore f female", "male"},
    {"11118888 m f female", "male"},
  }
  for index, test := range tests {
    str := tagSet.Parse(test.Input)
    assert.Equal(t, test.Result, str,
                 fmt.Sprintf("test: %d", index))
  }
}

func TestLine(t *testing.T) {
  sex := NewStringSet(
    NewString("male", []string{"male", "m"}),
    NewString("female", []string{"female", "f"}),
  )
  breeds := NewStringSet(
    NewString("angus", []string{"angus"}),
    NewString("nelore", []string{"nelore", "nalore"}),
  )
  tag := NewNumber(4,8)
  tests := []struct {
    Input        string
    ResultString string
    ResultInt    int    
    ResultSex    string
    ResultBreed  string
  }{
    {"1111 f nelore", "1111", 1111, "female", "nelore"},
    {"--  f 1111 nelore  --", "1111", 1111, "female", "nelore"},
    {"  nalore 1111 f", "1111", 1111, "female", "nelore"},
    {"male nalore 1111 f 3211 bla", "1111", 1111, "male", "nelore"},
  }
  for index, test := range tests {
    assert.Equal(t, test.ResultSex,
                 sex.Parse(test.Input),
                 fmt.Sprintf("test: %d", index))
    assert.Equal(t, test.ResultBreed, 
                 breeds.Parse(test.Input),
                 fmt.Sprintf("test: %d", index))
    eartag := tag.Parse(test.Input)
    assert.Equal(t, test.ResultString, eartag, 
                 fmt.Sprintf("test: %d", index))
    assert.Equal(t, test.ResultInt, tag.AsNumber(),
                 fmt.Sprintf("test: %d", index))
  }
}
