package services

import (
	"fmt"
	pagination "github.com/gobeam/mongo-go-pagination"
	"github.com/JubaerHossain/gomd/services/attendance/validation"
	"github.com/JubaerHossain/gomd/services/attendance/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"strconv"
    "time"
)

func AllAttendance(requestFilter map[string]interface{}) ([]models.Attendance, pagination.PaginationData) {
	var attendances []models.Attendance

	filter := bson.M{}

	if requestFilter["status"] != "" {
		filter["status"] = requestFilter["status"]
	}

	page, _ := strconv.ParseInt(requestFilter["page"].(string), 10, 64)
	limit, _ := strconv.ParseInt(requestFilter["limit"].(string), 10, 64)

	paginatedData, err := pagination.New(models.AttendanceCollection.Collection).
		Page(page).
		Limit(limit).
		Sort("created_at", -1).
		Decode(&attendances).
		Filter(filter).
		Find()

	if err != nil {
		panic(err)
	}
	return attendances, paginatedData.Pagination
}

func CreateAAttendance(createAttendance validation.CreateAttendanceRequest) models.Attendance {
	attendance := models.Attendance{
		Id:                 primitive.NewObjectID(),
		Task:               createAttendance.Task,
		Status:             createAttendance.Status,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}

	result, err := models.AttendanceCollection.InsertOne(attendance)
	if err != nil || result == nil {
		panic(err)
	}

	return attendance
}

func UpdateAAttendance(attendanceId string, updateAttendance validation.UpdateAttendanceRequest) (models.Attendance, error) {

	objId, _ := primitive.ObjectIDFromHex(attendanceId)

	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}

	result := models.AttendanceCollection.FindOneAndUpdate(
		bson.M{"_id": objId},
		bson.D{
			{"$set", bson.M{
				"task":                updateAttendance.Task,
				"status":              updateAttendance.Status,
				"updated_at":          time.Now(),
			}},
		},
		&opt,
	)

	if result.Err() != nil {
		log.Println("Err ", result.Err())
		return models.Attendance{}, result.Err()
	}

	var attendance models.Attendance
	if err := result.Decode(&attendance); err != nil {
		return models.Attendance{}, err
	}

	return attendance, nil
}

func AAttendance(attendanceId string) models.Attendance {
	var attendance models.Attendance

	objId, _ := primitive.ObjectIDFromHex(attendanceId)

	err := models.AttendanceCollection.FindOne(bson.M{"_id": objId}).Decode(&attendance)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	return attendance
}

func DeleteAAttendance(attendanceId string) bool {
	objId, _ := primitive.ObjectIDFromHex(attendanceId)

	result := models.AttendanceCollection.FindOneAndDelete(bson.D{{"_id", objId}})

	if result.Err() != nil {
		return false
	}

	return true
}