package tag

import (
  "fmt"
  "log"
  "regexp"
  "strconv"
)

type Number struct {
  value  string  
  asint  int
  minLen int
  maxLen int
}

func NewNumber(minLen, maxLen int) *Number {
  return &Number {
    minLen: minLen,
    maxLen: maxLen,
  }
}

func (n *Number) Parse(text string) string {
  format := fmt.Sprintf("\\b\\d{%d,%d}\\b", n.minLen, n.maxLen)
  pattern := regexp.MustCompile(format)
  matches := pattern.FindAllString(text, 1)
  if len(matches) == 0 {
    return ""
  }

  n.value = matches[0]
  var err error
  if n.asint, err = strconv.Atoi(n.value); err != nil {
		log.Fatalf("Error converting string to int: %v", err)
    return ""
	}

  return n.value
}

func (n *Number) AsNumber() int {
  return n.asint
}
