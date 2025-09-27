package account 

import (
  "log"
  "fmt"
  "context"
  "posso-help/internal/db"
  "go.mongodb.org/mongo-driver/bson"
)

type Team struct {
  Account      string `bson:"account"`
  PhoneNumber  string `bson:"phone_number"`
  Name         string `bson:"name"`
}

func getAllPhoneNumberVariants(phoneNumber string) ([]string) {
  variants := []string{}
  variants = append(variants, phoneNumber);

  // 16166100305 -> 616-610-0305
  if (len(phoneNumber)==11) {
    tmp := fmt.Sprintf("%s-%s-%s", phoneNumber[1:4], phoneNumber[4:7], phoneNumber[7:11])
    variants = append(variants, tmp);
  }

  // 12123451234 -> 12-12345-1234
  if (len(phoneNumber)==11) {
    tmp := fmt.Sprintf("%s-%s-%s", phoneNumber[0:2], phoneNumber[2:7], phoneNumber[7:11])
    variants = append(variants, tmp);
  }

  return variants;
}

func FindAccountByPhoneNumber(phoneNumber string) (string, error) {
  teams := db.GetCollection("teams")
  variants := getAllPhoneNumberVariants(phoneNumber)
  log.Printf("PhoneNumber: %s,  Variants: %v", phoneNumber, variants)
  filter := bson.M{"phone_number": bson.M{"$in": variants}}
  team := &Team{}
  err := teams.FindOne(context.TODO(), filter).Decode(team)
  log.Printf("FindAccountByPhoneNumber(%s):  returned %s:  %v", 
             phoneNumber, team.Account, err)
  return team.Account, err
}
