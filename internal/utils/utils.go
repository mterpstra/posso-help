package utils

import (
  "log"
  "strings"
)

func StringIsOneOf(in string, oneOf []string) bool {
  for i:=0; i < len(oneOf); i++ {
    if (in == oneOf[i]) {
      return true
    }
  }
  return false
}

func StringContainsOneOf(in string, oneOf []string) bool {
  for i:=0; i < len(oneOf); i++ {
    if (strings.Contains(in, oneOf[i])) {
      return true
    }
  }
  return false
}

func SanitizeLine(str string) string {
  sanitized := strings.ToLower(strings.TrimSpace(str))
  sanitized = strings.Replace(sanitized, "sta. zelia", "sta.zelia", -1)
  return sanitized
}

func Capitalize(str string) string {
  first := strings.ToUpper(str[:1])
  last := strings.ToLower(str[1:])
  return first + last 
}

func SplitAndTrim(str string) []string {
  // Need to split with simi-colon as comma's mess up CSV upload
  log.Printf("SplitAndTrim(%s)", str)
  parts := []string{}
  for _, part := range strings.Split(str, ";") {
    part = strings.TrimSpace(part)
    if len(part) > 0 {
      parts = append(parts, part)
    }
  }
  log.Printf("SplitAndTrim(%s): [%+v] len: %d", str, parts, len(parts))
  return parts
}
