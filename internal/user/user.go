package user

import (
  "time"
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
}
