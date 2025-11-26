package tag

type StringSet struct {
  value string
  strings []*String
}

func NewStringSet(strings ...*String) *StringSet {
  return &StringSet{
    strings: strings,
  }
}

func (ss *StringSet) Parse(text string) bool {
  for _, stringTag := range ss.strings {
    found := stringTag.Parse(text)
    if found {
      ss.value = stringTag.Value()
      return true
    }
  }
  return false
}

func (ss *StringSet) Value() string {
  return ss.value
}

func (ss *StringSet) ValueAsInt() int {
  return 0
}
