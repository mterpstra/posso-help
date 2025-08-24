package tag 

import (
  "strings"
)

type TagValue struct {
  Value   string   `json:"value"`
  Display string   `json:"display"`
  Matches []string `json:"matches"`
}

type Tag struct {
  Account string     `json:"account"`
  Name    string     `json:"name"`
  Type    string     `json:"type"`
  Values  []TagValue `json:"values"`
}

func (t *Tag) Find(account, name string) error {
  db := NewDb("tags")
  if err := db.Connect(); err != nil {
    return err
  }
  defer db.Close()

  filter := map[string]string{
    "account":account,
    "name":name,
  }

  if err := db.Read(filter, t); err != nil {
    return err
  }

  return nil
}

func (t *Tag) FindTagInLine(line string) (*TagValue, bool) {
  for _, tagValue := range t.Values {
    for _, match := range tagValue.Matches {
      if strings.Contains(line, match) {
        return &tagValue, true
      }
    }
  }
  return nil, false
}

func (t *Tag) AddValue(newvalue string) error {
  db := NewDb("tags")
  if err := db.Connect(); err != nil {
    return err
  }
  defer db.Close()

  // @todo: Generate display name and any other matches
  return db.AddValue(t.Account, t.Name, newvalue)

}

