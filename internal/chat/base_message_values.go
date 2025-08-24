package chat

import "go.mongodb.org/mongo-driver/bson"

type BaseMessageValues struct {
  EntryID     string `json:"entry_id"`
  MessageID   string `json:"message_id"`
  PhoneNumber string `json:"phone"`
  Name        string `json:"name"`
  Date        string `json:"date"`
}

func (bmv *BaseMessageValues) ToMap() bson.D {
  return bson.D {
    {"entry_id", bmv.EntryID},
    {"message_id", bmv.MessageID},
    {"phone", bmv.PhoneNumber},
    {"name", bmv.Name},
    {"date", bmv.Date},
  }
}
