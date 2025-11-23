package area

import (
  "log"
  "context"
  "strings"
  "posso-help/internal/db"
  "posso-help/internal/utils"
  "go.mongodb.org/mongo-driver/bson"
)

type AreaParser struct {
  areas []*Area
}

func (ap *AreaParser) LoadAreasByAccount(account string) error {
  collection := db.GetCollection("areas")
  filter := bson.M{"account":account}
  cursor, err := collection.Find(context.TODO(), filter)
  if err != nil {
    log.Printf("Error reading areas for account: %v", account)
    return err
  }
  defer cursor.Close(context.TODO())

  // Iterate through the cursor and decode each document
	for cursor.Next(context.TODO()) {
    area := &Area{}
		if err := cursor.Decode(area); err != nil {
			log.Printf("Error decoding document: %v", err)
			continue 
		}
    log.Printf("LoadAreasByAccount(%s): %s  %s", 
                account, area.Name, area.Matches)
		ap.areas = append(ap.areas, area)
	}

	return cursor.Err()
}

func (ap *AreaParser) ParseAsAreaLine(line string) (string, bool) {
  line = utils.SanitizeLine(line)
  for _, area := range ap.areas {
    matches := utils.SplitAndTrim(strings.ToLower(area.Matches))
    if utils.StringContainsOneOf(line, matches) {
      log.Printf("ParseAsAreaLine: found, name=%+v\n", area.Name)
      return area.Name, true
    }
  }
  
  return "", false
}
