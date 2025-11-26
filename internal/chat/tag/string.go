package tag

import (
  "fmt"
  "log"
  "regexp"
  "strings"
)

type String struct {
  found bool
  value string
  variants []string
}

func NewString(value string, variants []string) *String{
  return &String {
    value: value,
    variants: variants,
  }
}

func (s *String) Parse(text string) bool {
  pattern := fmt.Sprintf("\\b(%s)\\b", 
    strings.Join(s.variants, "|"))

  var err error
  s.found, err = regexp.MatchString(pattern, text)
  if err != nil {
    log.Printf("Could not compile regexp for tag: %s", s.variants)
    return false
  }

  return s.found
}

func (s *String) Value() string {
  if s.found {
    return s.value
  }
  return ""
}

func (s *String) ValueAsInt() int {
  return 0
}
