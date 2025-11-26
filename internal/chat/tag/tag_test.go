package tag

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

func TestNumber(t *testing.T) {
  tag := NewNumber(4,8)
  tests := []TestCase{
    {"11 m angus", false, "", 0},
    {"nalore 123456789 m", false, "", 0},
    {"nalore m something else", false, "", 0},
    {"1111 m angus", true, "1111", 1111},
    {"nalore 12345678 m", true, "12345678", 12345678},
    {"nalore m 99999999 something else", true, "99999999", 99999999},
  }
  for index, test := range tests {
    found := tag.Parse(test.Input)
    assert.Equal(t, test.Found, found, 
                 fmt.Sprintf("test: %d", index))
    assert.Equal(t, test.Value, tag.Value(), 
                 fmt.Sprintf("test: %d", index))
    assert.Equal(t, test.ValueInt, tag.ValueAsInt(), 
                 fmt.Sprintf("test: %d", index))
  }
}

func TestString(t *testing.T) {
  tag := NewString("male", []string{"male", "m", "strange-variant"})
  tests := []TestCase{
    {"1111 angus",        false, "",     0},
    {"1111 female angus", false, "",     0},
    {"1111 m angus",      true,  "male", 0},
    {"m 1111 angus",      true,  "male", 0},

    {"biggy smalls  female 1111 angus male  ", true, "male", 0},
    {"  male, 1111 angus", true, "male", 0},
    {"  1111 angus strange-variant blue moon", true, "male", 0},
  }
  for index, test := range tests {
    found := tag.Parse(test.Input)
    assert.Equal(t, test.Found, found, 
                 fmt.Sprintf("test: %d", index))
    assert.Equal(t, test.Value, tag.Value(), 
                 fmt.Sprintf("test: %d", index))
    assert.Equal(t, test.ValueInt, tag.ValueAsInt(), 
                 fmt.Sprintf("test: %d", index))
  }
}

func TestTagSet(t *testing.T) {
  m := NewString("male", []string{"male", "m"})
  f := NewString("female", []string{"female", "f"})
  tagSet := NewStringSet(m,f)
  tests := []TestCase{
    {"female", true, "female", 0},
    {"f", true, "female", 0},
    {"male", true, "male", 0},
    {"m", true, "male", 0},
    {"1111 m nalore", true, "male", 0},
    {"1111 female nalore", true, "female", 0},
    {"11118888 male nalore f female", true, "male", 0},
    {"11118888 m f female", true, "male", 0},
  }
  for index, test := range tests {
    found := tagSet.Parse(test.Input)
    assert.Equal(t, test.Found, found, 
                 fmt.Sprintf("test: %d", index))
    assert.Equal(t, test.Value, tagSet.Value(), 
                 fmt.Sprintf("test: %d", index))
    assert.Equal(t, test.ValueInt, tagSet.ValueAsInt(), 
                 fmt.Sprintf("test: %d", index))
  }
}

func TestDate(t *testing.T) {
  date := NewDate()

  tests := []TestCase{
    {"2025/12/02",  true, "2025-12-02", 0},
    {"2025/12/2",   true, "2025-12-02", 0},
    {"2025-7-05",   true, "2025-07-05", 0},
    {"something before 2025/7/1 and after", true, "2025-07-01", 0},

    //DD/MM
    {"25/12", true, "2025-12-25", 0},
    {"5/8",   true, "2025-08-05", 0},

    // Invalids
    {"2025/13/02", false, "", 0},
    {"32/02",      false, "", 0},
  }
  for index, test := range tests {
    found := date.Parse(test.Input)
    assert.Equal(t, test.Found, found, 
                 fmt.Sprintf("test: %d", index))
    assert.Equal(t, test.Value, date.Value(), 
                 fmt.Sprintf("test: %d", index))
    assert.Equal(t, test.ValueInt, date.ValueAsInt(), 
                 fmt.Sprintf("test: %d", index))
  }
}
