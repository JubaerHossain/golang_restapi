package models

import (
	"github.com/JubaerHossain/gomd/gomd"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var AttendanceCollection gomd.MongoCollection

type Attendance struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id"`
    Task      string             `json:"task" bson:"task"`
    Status    string             `json:"status" bson:"status"`
    CreatedAt time.Time          `json:"created_at,omitempty" bson:"created_at"`
    UpdatedAt time.Time          `json:"updated_at,omitempty" bson:"updated_at"`
}

func AttendanceSetup() {
	AttendanceCollection = gomd.Mongo.Collection("attendances")
}