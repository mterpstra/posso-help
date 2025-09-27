package chat

import (
  "log"
  "strconv"
  "strings"
  "time"
  "posso-help/internal/area"
  "posso-help/internal/account"
  "posso-help/internal/textmsg"
)

type ChatMessage struct {
  Object  string   `json:"object"`
  Entries []Entry  `json:"entry"`
}

type Entry struct {
	Changes []Changes `json:"changes"`
	ID      string    `json:"id"`
}

type Changes struct {
	Field string `json:"field"`
	Value Value  `json:"value"`
}

type Value struct {
	Contacts         []Contacts `json:"contacts"`
	Messages         []Messages `json:"messages"`
	MessagingProduct string     `json:"messaging_product"`
	Metadata         Metadata   `json:"metadata"`
}

type Contacts struct {
	Profile Profile `json:"profile"`
	WaID    string  `json:"wa_id"`
}

type Messages struct {
	ID        string `json:"id"`
	Text      Text   `json:"text"`
	Timestamp string `json:"timestamp"`
	Type      string `json:"type"`
	From      string `json:"from"`
}

type Metadata struct {
	DisplayPhoneNumber string `json:"display_phone_number"`
	PhoneNumberID      string `json:"phone_number_id"`
}

type Profile struct {
	Name string `json:"name"`
}

type Text struct {
	Body string `json:"body"`
}

type Parser interface {
  Parse(string) bool 
  Text() string
  Insert(*BaseMessageValues) error
}

func ProcessEntries(entries []Entry) error {
  for _, entry := range entries {
    entry.Process()
  }
  return nil
}

func (e Entry) Process() error {

  birthMessageParser := &BirthMessage{
    AreaParser: &area.AreaParser{},
  }
  deathMessageParser := &DeathMessage{
    AreaParser: &area.AreaParser{},
  }

  parsers := make(map[string]Parser)
  parsers["birth"] = birthMessageParser
  parsers["death"] = deathMessageParser
  parsers["rain"] = &RainMessage{}
  parsers["temperature"] = &TemperatureMessage{}
  parsers["weather"] = &WeatherMessage{}

  for _, change := range e.Changes {

    name := "unknown"
    for _, contact := range change.Value.Contacts {
      name = contact.Profile.Name
    }

    for _, message := range change.Value.Messages {

      timestamp, err := strconv.ParseInt(message.Timestamp, 10, 64)
      if err != nil {
        now := time.Now()
        timestamp = now.Unix()
      }

      unixTimestamp := int64(timestamp)
      t := time.Unix(unixTimestamp, 0)

      account, err := account.FindAccountByPhoneNumber(message.From)
      if (err != nil) {
        log.Printf("WARNING: Could not find account for %s\n", message.From)
      }

      areaParser := &area.AreaParser{}
      err = areaParser.LoadAreasByAccount(account)
      if err != nil {
        log.Printf("WARNING: Could not load areas from account: %v\n", account)
      } 
      birthMessageParser.AreaParser = areaParser
      deathMessageParser.AreaParser = areaParser

      baseMessageValues := &BaseMessageValues {
        Account      : account,
        PhoneNumber  : message.From,
        Name         : name,
        Date         : t.Format(time.RFC3339),
      }

      for name, parser := range parsers {
        msg := strings.TrimSpace(message.Text.Body)
        if found := parser.Parse(msg); found {
          log.Printf("message parsed with parser: %v\n", name)
          if err := parser.Insert(baseMessageValues); err != nil {
            log.Printf("Error insert record into DB: %v\n", err)
          }
          text := textmsg.NewMessageSender(message.From, parser.Text())
          if err := text.Send(); err != nil {
            log.Printf("Error during text reply: %v\n", err)
          }
        }
      }
    }
  }

  return nil
}
