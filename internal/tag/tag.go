package tag

import (
  "fmt"
  "strings"
)

type Tag struct {
  Value string
  Variants []string
}

// IsVariantFound returns true if the passed text value
// is a variant of the given tag
func (t *Tag) IsVariantFound(text string) bool {
  for _, variant := range t.Variants {
    if variant == text {
      return true
    }
  }
  return false
}

// ToRegexPattern returns a string that can be compiled
// into a regex and then used to match strings
func (t *Tag) ToRegexPattern() string {
  return fmt.Sprintf("(%s)", strings.Join(t.Variants, "|"))
}
