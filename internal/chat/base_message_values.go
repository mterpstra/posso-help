package chat

import "go.mongodb.org/mongo-driver/bson"

type BaseMessageValues struct {
  Account     string `json:"account"`
  PhoneNumber string `json:"phone"`
  Name        string `json:"name"`
  Date        string `json:"date"`
}

func (bmv *BaseMessageValues) ToMap() bson.D {
  return bson.D {
    {"account", bmv.Account},
    {"phone", bmv.PhoneNumber},
    {"created_by", bmv.Name},
    {"date", bmv.Date},
  }
}
