package models

import (
	"time"

	"github.com/JubaerHossain/gomd/gomd"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var UsersCollection gomd.MongoCollection

type User struct {
	Id        primitive.ObjectID `json:"id,omitempty"`
	Name      string             `json:"name,omitempty" validate:"required"`
	Email     string             `json:"email,omitempty" validate:"required,email"`
	Password  string             `json:"password,omitempty" validate:"required"`
	CreatedAt time.Time          `json:"created_at"`
}

func UsersSetup() {
	UsersCollection = gomd.Mongo.Collection("users")
}
