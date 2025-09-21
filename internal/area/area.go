package area

import (
  "log"
  "context"
  "strings"
  "posso-help/internal/db"
  "posso-help/internal/utils"
  "go.mongodb.org/mongo-driver/bson"
)

// Area
type Area struct {
  Name string     `bson:"area_name"`
  Matches string  `bson:"matches"`
}

func readAreas(phoneNumber string) ([]Area, error) {
  areas := []Area{}
  collection := db.GetCollection("areas")
  filter := bson.D{{"phone", phoneNumber}} 

  cursor, err := collection.Find(context.Background(), filter)
  if err != nil {
    return areas, err
  }
  defer cursor.Close(context.Background()) 

  for cursor.Next(context.Background()) {
    var doc Area 
    if err := cursor.Decode(&doc); err != nil {
      return areas, err
    }
    areas = append(areas, doc)
  }

  if err := cursor.Err(); err != nil { 
    return areas, err
  }
  return areas, nil
}

func ParseAsAreaLine(line string) (string, bool) {

  phoneNumber := "16166100305"
  areas, err := readAreas(phoneNumber)
  if err != nil {
    log.Printf("Error reading areas for %s: %v\n", phoneNumber, err)
    return "", false
  }
  line = utils.SanitizeLine(line)
  for _, area := range areas {
    matches := utils.SplitAndTrim(strings.ToLower(area.Matches))
    if utils.StringContainsOneOf(line, matches) {
      return area.Name, true
    }
  }
  
  return "", false
}

func AddArea(phone, newArea string) error {
  data := make(map[string]interface{})
  collection := db.GetCollection("areas")

  data["area_name"]  = newArea
  data["matches"]    = newArea
  data["entry_id"]   = "Pull this from message"
  data["message_id"] = "Pull this from message"
  data["phone"]      = phone
  data["name"]       = "Pull this from message"

  _, err := collection.InsertOne(context.TODO(), data)
  if err != nil {
    println("ERROR")
    return err
  }

  println("SUCCESS")
  return nil
}
