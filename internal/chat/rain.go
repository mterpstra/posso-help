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

const REPLY_RAINFALL = `Posso Help has detected rainfall data.  We added %d mm of rain. To claim your data and see a report sign up at https://possohelp.com` 

type Rain struct {
  EntryId   string `json:"entry_id"`
  MessageId string `json:"message_id"`
  Phone     string `json:"phone"`
  Name      string `json:"phone"`
  Date      string `json:"date"`
  Amount    int    `json:"amount"`
}

type RainEntry struct {
  Date string
  Amount int // in mm
}

type RainMessage struct {
  Entries []*RainEntry
  Area *area.Area
  Total int
}

// RainMessage returns an array of rain entries
func (r *RainMessage) Parse(message string) bool {
  found := false
  lines := strings.Split(message, "\n")
  for _, line := range lines {
    if entry := r.parseRainLine(line); entry != nil {
      r.Entries = append(r.Entries, entry)
      r.Total += entry.Amount
      found = true
    }
  }
  return found 
}

func (r *RainMessage) parseRainLine(line string) (*RainEntry) {
  var day, month int
  var rainfall int // in millimeters
  line = utils.SanitizeLine(line)

  // Support both 15mm and 15 mm (with space)
  line = strings.Replace(line, "mm", " mm", 1)
  n, err := fmt.Sscanf(line, "%d/%d %d mm", &day, &month, &rainfall)
  if err == nil && n == 3 {
    return &RainEntry{
      Date: date.MonthDayToUTC(month, day),
      Amount: rainfall,
    }
  }
  return nil
}

func (r *RainMessage) Text() string {
  return fmt.Sprintf(REPLY_RAINFALL, r.Total)
}

func (b *RainMessage) Insert(bmv *BaseMessageValues) error {
  rains := db.GetCollection("rain")
  for _, rain := range b.Entries {
    document := bmv.ToMap()
    document = append(document, bson.E{Key: "amount", Value: rain.Amount})
    document = append(document, bson.E{Key: "date", Value: rain.Date})
    _, err := rains.InsertOne(context.TODO(), document)
    if err != nil {
      return err
    }
  }
  return nil
}
