package tag

import (
  "fmt"
  "log"
  "regexp"
  "strings"
)

type String struct {
  value string
  variants []string
}

func NewString(value string, variants []string) *String{
  return &String {
    value: value,
    variants: variants,
  }
}

func (s *String) Parse(text string) string {
  pattern := fmt.Sprintf("\\b(%s)\\b", 
    strings.Join(s.variants, "|"))

  matched, err := regexp.MatchString(pattern, text)
  if err != nil {
    log.Printf("Could not compile regexp for tag: %s", s.variants)
    return ""
  }

  if (!matched) {
    return ""
  }

  return s.value
}
