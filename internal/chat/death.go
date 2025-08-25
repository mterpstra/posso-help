package chat

import (
  "log"
  "fmt"
  "strings"
  "context"
  "posso-help/internal/area"
  "posso-help/internal/db"
  "posso-help/internal/date"
  "posso-help/internal/utils"
  "go.mongodb.org/mongo-driver/bson"
)

const REPLY_DEATHS = `Posso Help has detected death data.  We added %d deaths to area %s. To claim your data and see a report sign up at https://possohelp.com` 

type DeathEntry struct {
  Id       int    `json:"tag"`
  Sex      string `json:"sex"`
  Cause    string `json:"cause"`
}

type DeathMessage struct {
  Date string
  Entries []*DeathEntry
  Area *area.Area
  Total int
}

func (d *DeathMessage) Parse(message string) bool {
  found := false
  lines := strings.Split(message, "\n")
  for _, line := range lines {
    if date, found := date.ParseAsDateLine(line); found {
      d.Date = date
    }
    if entry := d.parseAsDeathLine(line); entry != nil {
      d.Entries = append(d.Entries, entry)
      d.Total++
      found = true
    }
    if areaName, found := area.ParseAsAreaLine(line); found {
      d.Area = &area.Area{Name: areaName}
    }
  }

  if found && d.Area == nil {
    d.Area = &area.Area{Name: "unknown"}
  }
  return found 
}

func (d *DeathMessage) parseAsDeathLine(line string) (*DeathEntry) {
  var num int
  var sex, cause string
  line = utils.SanitizeLine(line)

  // Death Line with Sex
  n, err := fmt.Sscanf(line, "%d %s %s", &num, &sex, &cause)
  if err == nil && n == 3 && num > 0 &&
    (utils.StringIsOneOf(sex, SEXES)) && (utils.StringIsOneOf(cause, DEATHS)) {
      return &DeathEntry{Id:num, Sex:sex, Cause:cause}
  }

  // Death line with NO Sex 
  n, err = fmt.Sscanf(line, "%d %s", &num, &cause)
  if err == nil && n == 2 && num > 0 &&
    (utils.StringIsOneOf(cause, DEATHS)) {
      return &DeathEntry{Id:num, Cause:cause}
  }

  return nil
}

func (d *DeathMessage) Text() string {
  return fmt.Sprintf(REPLY_DEATHS, d.Total, d.Area.Name)
}

func (d *DeathMessage) Insert(bmv *BaseMessageValues) error {
  deaths := db.GetCollection("deaths")
  log.Printf("inserting death message to collection: %v\n", deaths)
  for _, death := range d.Entries {
    document := bmv.ToMap()
    document = append(document, bson.E{Key: "tag", Value: death.Id})
    document = append(document, bson.E{Key: "sex", Value: death.Sex})
    document = append(document, bson.E{Key: "cause", Value: death.Cause})
    document = append(document, bson.E{Key: "area", Value: d.Area.Name})

    // If the message text has a date, use it over the message date
    if d.Date != "" {
      document = append(document, bson.E{Key: "date", Value: d.Date})
    }

    result, err := deaths.InsertOne(context.TODO(), document)
    if err != nil {
      log.Printf("error inserting death: %v\n", err)
      return err
    }

    log.Printf("insert result: %v\n", result)
  }

  log.Printf("death inserted successfully")
  return nil
}
