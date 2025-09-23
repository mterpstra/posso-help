package chat

import (
  "context"
  "log"
  "fmt"
  "strings"
  "posso-help/internal/area"
  "posso-help/internal/db"
  "posso-help/internal/date"
  "posso-help/internal/utils"
  "go.mongodb.org/mongo-driver/bson"
)

const REPLY_BIRTHS = `Posso Help has detected birth data.  We added %d births to area %s. To claim your data and see a report sign up at https://possohelp.com` 

type Birth struct {
  EntryId   string `json:"entry_id"`
  MessageId string `json:"message_id"`
  Phone     string `json:"phone"`
  Name      string `json:"phone"`
  Date      string `json:"date"`
  Tag       int64  `json:"tag" bson:"tag, omitempty"`
  Pure      bool   `json:"pure_bread"`
  Sex       string `json:"sex"`
  Breed     string `json:"breed"`
  Area      string `json:"area"`
}

type BirthEntry struct {
  PureBred bool   `json:"pure_bred"`
  Id       int    `json:"tag"`
  Sex      string `json:"sex"`
  Breed    string `json:"breed"`
}

type BirthMessage struct {
  Date string
  Entries []*BirthEntry
  Area *area.Area
  AreaParser *area.AreaParser
  NewAreaFound bool
  Total int
}

func (b *BirthMessage) Parse(message string) bool {
  found := false
  lines := strings.Split(message, "\n")
  parsedLines := map[int]bool{}
  for index, line := range lines {
    if date, found := date.ParseAsDateLine(line); found {
      b.Date = date
      parsedLines[index] = true
    }
    if entry := b.parseAsBirthLine(line); entry != nil {
      b.Entries = append(b.Entries, entry)
      b.Total++
      found = true
      parsedLines[index] = true
    }
    if areaName, found := b.AreaParser.ParseAsAreaLine(line); found {
      b.Area = &area.Area{Name:areaName}
      parsedLines[index] = true
    }
  }


  // If we found at least one birth, and there was no "known area"
  // and the last line was not parsed, maybe the last line is a 
  // new area.
  if found && b.Area == nil && !parsedLines[len(lines)-1] {
    newArea := utils.SanitizeLine(lines[len(lines)-1])
    b.Area = &area.Area{Name:newArea}
    log.Printf("New Area Found \"%s\"", newArea)
    b.NewAreaFound = true
  }

  if found && b.Area == nil {
    b.Area = &area.Area{Name: "unknown"}
  }

  return found 
}

func (b *BirthMessage) parseAsBirthLine(line string) (*BirthEntry) {
  var num int
  var pure, sex, breed string
  line = utils.SanitizeLine(line)

  // Standard Birth Line
  n, err := fmt.Sscanf(line, "%d %s %s", &num, &sex, &breed)
  if err == nil && n == 3 && num > 0 &&
    (utils.StringIsOneOf(sex, SEXES)) && (utils.StringIsOneOf(breed, BREEDS)) {
    return &BirthEntry{false, num, sex, breed}
  }

  // Pure Breed Birth Line
  n, err = fmt.Sscanf(line, "%s %d %s %s", &pure, &num, &sex, &breed)
  if err == nil && n == 4 && num > 0 &&
    (utils.StringIsOneOf(sex, SEXES)) && 
    (utils.StringIsOneOf(breed, BREEDS) &&
    (pure == PURE_BREED)) {
    return &BirthEntry{true, num, sex, breed}
  }

  return nil
}

func (b *BirthMessage) Text() string {
  return fmt.Sprintf(REPLY_BIRTHS, b.Total, b.Area.Name)
}

func (b *BirthMessage) Insert(bmv *BaseMessageValues) error {
  births := db.GetCollection("births")

  for _, birth := range b.Entries {
    document := bmv.ToMap()
    document = append(document, bson.E{Key: "tag", Value: birth.Id})
    document = append(document, bson.E{Key: "pure_bread", Value: birth.PureBred})
    document = append(document, bson.E{Key: "sex", Value: birth.Sex})
    document = append(document, bson.E{Key: "breed", Value: birth.Breed})
    document = append(document, bson.E{Key: "area", Value: b.Area.Name})

    // If the message text has a date, use it over the message date
    if b.Date != "" {
      document = append(document, bson.E{Key: "date", Value: b.Date})
    }

    _, err := births.InsertOne(context.TODO(), document)
    if err != nil {
      return err
    }
  }

  if b.NewAreaFound {
    err := area.AddArea(bmv.Account, b.Area.Name, b.Area.Name) 
    if err != nil {
      fmt.Printf("Could not add new area %v", err)
    }
  }
  
  return nil
}
