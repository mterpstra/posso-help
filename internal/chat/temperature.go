package chat

import (
  "fmt"
  "strings"
  "context"
  "posso-help/internal/area"
  "posso-help/internal/db"
  "posso-help/internal/date"
  "posso-help/internal/utils"
  "go.mongodb.org/mongo-driver/bson"
)

// Data formats for Temperature data
// "dd/m 35C"
// "dd/m 35 C"

const REPLY_TEMPERATURE = `Zap Manejo has detected temperature data.  We added %d days of temperature data.  To claim your data and see a report sign up at https://dashboard.zapmanejo.com/`

type Temperature struct {
  EntryId     string `json:"entry_id"`
  MessageId   string `json:"message_id"`
  Phone       string `json:"phone"`
  Name        string `json:"name"`
  Date        string `json:"date"`
  Temperature int    `json:"temperature"`
}

type TemperatureEntry struct {
  Date string
  Temperature int // in Celcius
}

type TemperatureMessage struct {
  Entries []*TemperatureEntry
  Area *area.Area
}

// ParseTemperatureMessage returns a parsed Temperaturemessage
func (t *TemperatureMessage) Parse(message string) bool {
  found := false
  lines := strings.Split(message, "\n")
  for _, line := range lines {
    if entry := t.parseTemperatureLine(line); entry != nil {
      t.Entries = append(t.Entries, entry)
      found = true
    }
  }
  return found 
}

func (t *TemperatureMessage) parseTemperatureLine(line string) (*TemperatureEntry) {
  var day, month int
  var temperature int // in celcius 
  var celcius rune
  line = utils.SanitizeLine(line)

  n, err := fmt.Sscanf(line, "%d/%d %d%c\n", &day, &month, &temperature, &celcius)
  if err == nil && n == 4 && celcius == 'c' {
    return &TemperatureEntry{
      Date: date.MonthDayToUTC(month, day),
      Temperature: temperature,
    }
  }

  n, err = fmt.Sscanf(line, "%d/%d %d %c\n", &day, &month, &temperature, &celcius)
  if err == nil && n == 4 && celcius == 'c' {
    return &TemperatureEntry{
      Date: date.MonthDayToUTC(month, day),
      Temperature: temperature,
    }
  }

  return nil
}

func (r *TemperatureMessage) Text() string {
  return fmt.Sprintf(REPLY_TEMPERATURE, len(r.Entries))
}

func (b *TemperatureMessage) Insert(bmv *BaseMessageValues) error {
  temps := db.GetCollection("temperature")
  for _, temp := range b.Entries {
    document := bmv.ToMap()
    document = append(document, bson.E{Key: "temperature", Value: temp.Temperature})
    document = append(document, bson.E{Key: "date", Value: temp.Date})
    _, err := temps.InsertOne(context.TODO(), document)
    if err != nil {
      return err
    }
  }
  return nil
}
