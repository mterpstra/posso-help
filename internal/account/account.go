package account 

import (
  "context"
  "posso-help/internal/db"
  "go.mongodb.org/mongo-driver/bson"
)

type Team struct {
  Account      string `bson:"account"`
  PhoneNumber  string `bson:"phone_number"`
  Name         string `bson:"name"`
}

func FindAccountByPhoneNumber(phoneNumber string) (string, error) {
  teams := db.GetCollection("teams")
  filter := bson.M{"phone_number":phoneNumber}
  team := &Team{}
  err := teams.FindOne(context.TODO(), filter).Decode(team)
  return team.Account, err
}
