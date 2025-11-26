package tag

type StringSet struct {
  strings []*String
}

func NewStringSet(strings ...*String) *StringSet {
  return &StringSet{
    strings: strings,
  }
}

func (ss *StringSet) Parse(text string) string {
  for _, stringTag := range ss.strings {
    value := stringTag.Parse(text)
    if value != "" {
      return value
    }
  }
  return ""
}
