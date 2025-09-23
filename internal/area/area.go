package area

import (
  "log"
  "context"
  "posso-help/internal/db"
)

// Area
type Area struct {
  Name string     `bson:"area_name"`
  Matches string  `bson:"matches"`
}


func AddArea(account, name, matches string) error {
  data := make(map[string]interface{})
  collection := db.GetCollection("areas")

  data["account"] = account
  data["name"]    = name
  data["matches"] = matches

  _, err := collection.InsertOne(context.TODO(), data)
  if err != nil {
    log.Printf("Error inserting new area: %v", err)
    return err
  }

  log.Printf("Successfully inserted new Area")
  return nil
}
