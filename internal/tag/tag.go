package tag

type Tag interface {
  Parse(string) string
}

/*
type TagList struct {
  Tags []Tag
}

// IsInText returns the matching tag if one is found
func (t *TagList) IsInText(text string) (bool, *Tag) {
  for _, tag := range t.Tags {
    if tag.IsInText(text) {
      return true, &tag
    }
  }
  return false, nil
}
*/
