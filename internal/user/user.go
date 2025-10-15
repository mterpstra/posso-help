package user

import (
  "context"
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
  PhoneNumber string             `bson:"phone_number" json:"phone_number"`
  Name        string             `json:"name"`
  CreatedAt   time.Time          `json:"created_at"`
  UpdatedAt   time.Time          `json:"updated_at"`
  IsActive    bool               `bson:"is_active" json:"is_active"`
  Language    string             `bson:"lang" json:"lang"`
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
  return user, nil
}

func (u *User) GetDisplayName() string {

  if (u.Name != "") {
    return u.Name
  }

  if (u.Username != "") {
    return u.Username
  }

  if (u.Email != "") {
    return u.Email
  }

  return u.ID.Hex()
}
