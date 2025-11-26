package line

import (
  "posso-help/internal/chat/tag"
)

type LineParser struct {
  mustHave map[string]tag.Tag 
  canHave map[string]tag.Tag 
}

func NewLineParser() *LineParser {
  return &LineParser{
    mustHave:  make(map[string]tag.Tag),
    canHave: make(map[string]tag.Tag),
  }
}

func (l *LineParser) MustHave(name string, tagParser tag.Tag) *LineParser {
  l.mustHave[name] = tagParser
  return l
}

func (l *LineParser) CanHave(name string, tagParser tag.Tag) *LineParser {
  l.canHave[name] = tagParser
  return l
}

func (l *LineParser) Parse(text string) bool {

  for _, tag := range l.mustHave {
    found := tag.Parse(text)
    if !found {
      return false  
    }
  }

  for _, tag := range l.canHave {
    tag.Parse(text)
  }

  return true
}

func (l *LineParser) Value(name string) string {
	if value, ok := l.mustHave[name];ok {
    return value.Value()
	} 
  if value, ok := l.canHave[name]; ok {
    return value.Value()
	} 
  return ""
}

func (l *LineParser) ValueAsInt(name string) int {
	if value, ok := l.mustHave[name];ok {
    return value.ValueAsInt()
	} 
  if value, ok := l.canHave[name]; ok {
    return value.ValueAsInt()
	} 
  return 0
}
