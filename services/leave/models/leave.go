package models

import (
	"github.com/JubaerHossain/gomd/gomd"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var LeaveCollection gomd.MongoCollection

type Leave struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id"`
    Task      string             `json:"task" bson:"task"`
    Status    string             `json:"status" bson:"status"`
    CreatedAt time.Time          `json:"created_at,omitempty" bson:"created_at"`
    UpdatedAt time.Time          `json:"updated_at,omitempty" bson:"updated_at"`
}

func LeaveSetup() {
	LeaveCollection = gomd.Mongo.Collection("leaves")
}