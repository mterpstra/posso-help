package user

import (
  "context"
  "log"
  "time"
  "posso-help/internal/db"
  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
  ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
  Username    string             `json:"username"`
  Email       string             `json:"email"`
  Password    string             `json:"-"` // Excluded from JSON responses
  PhoneNumber string             `json:"phone_number"`
  Name        string             `json:"name"`
  CreatedAt   time.Time          `json:"created_at"`
  UpdatedAt   time.Time          `json:"updated_at"`
  IsActive    bool               `bson:"is_active" json:"is_active"`
  PhoneNumbers []string          `bson:"phone_numbers" json:"phone_numbers"`
}

func Read(Id string) (*User, error) {
  users := db.GetCollection("users")
  objectID, err := primitive.ObjectIDFromHex(Id)
  if err != nil {
    return nil, err
  }
  filter := bson.M{"_id": objectID}
  user := &User{}
  err = users.FindOne(context.TODO(), filter).Decode(user)
  if err != nil {
    return nil, err
  }
  log.Printf("User Found: %+v\n", user)
  return user, nil
}
