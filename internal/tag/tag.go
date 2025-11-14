package tag

import (
  "fmt"
  "log"
  "regexp"
  "strings"
)

type Tag struct {
  Value string
  Variants []string
}

type TagList struct {
  Tags []Tag
}

// IsInText checks to see if any of the tag
// variants exist in the passed text.
func (t *Tag) IsInText(text string) bool {
  pattern := t.toRegexPattern()
  matched, err := regexp.MatchString(pattern, text)
  if err != nil {
    log.Printf("Could not compile regexp for tag: %s", t.Value)
    return false
  }
  return matched
}


// ToRegexPattern returns a string that can be compiled
// into a regex and then used to match strings
func (t *Tag) toRegexPattern() string {
  return fmt.Sprintf("\\b(%s)\\b", strings.Join(t.Variants, "|"))
}


func (t *TagList) IsInText(text string) (bool, *Tag) {
  for _, tag := range t.Tags {
    if tag.IsInText(text) {
      return true, &tag
    }
  }
  return false, nil
}

