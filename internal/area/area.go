package area

import (
  "posso-help/internal/tag"
  "posso-help/internal/utils"
)

// Area
type Area struct {
  Name string
}

var areaTag *tag.Tag

func ParseAsAreaLine(line string) (string, bool) {

  if areaTag == nil {
    areaTag = &tag.Tag{}
    if err := areaTag.Find("16166100305", "area"); err != nil {
      return "", false
    }
  }


  line = utils.SanitizeLine(line)
  tag, found := areaTag.FindTagInLine(line)
  if found {
    return tag.Value, true
  }
  
  return "", false
}

func AddArea(account, newArea string) error {
  areaTag = &tag.Tag{
    Account:account,
    Name:"area",
  }
  return areaTag.AddValue(newArea)
}
