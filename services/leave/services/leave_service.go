package services

import (
	"fmt"
	pagination "github.com/gobeam/mongo-go-pagination"
	"github.com/JubaerHossain/golang_restapi/services/leave/validation"
	"github.com/JubaerHossain/golang_restapi/services/leave/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"strconv"
    "time"
)

func AllLeave(requestFilter map[string]interface{}) ([]models.Leave, pagination.PaginationData) {
	var leaves []models.Leave

	filter := bson.M{}

	if requestFilter["status"] != "" {
		filter["status"] = requestFilter["status"]
	}

	page, _ := strconv.ParseInt(requestFilter["page"].(string), 10, 64)
	limit, _ := strconv.ParseInt(requestFilter["limit"].(string), 10, 64)

	paginatedData, err := pagination.New(models.LeaveCollection.Collection).
		Page(page).
		Limit(limit).
		Sort("created_at", -1).
		Decode(&leaves).
		Filter(filter).
		Find()

	if err != nil {
		panic(err)
	}
	return leaves, paginatedData.Pagination
}

func CreateALeave(createLeave validation.CreateLeaveRequest) models.Leave {
	leave := models.Leave{
		Id:                 primitive.NewObjectID(),
		Task:               createLeave.Task,
		Status:             createLeave.Status,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}

	result, err := models.LeaveCollection.InsertOne(leave)
	if err != nil || result == nil {
		panic(err)
	}

	return leave
}

func UpdateALeave(leaveId string, updateLeave validation.UpdateLeaveRequest) (models.Leave, error) {

	objId, _ := primitive.ObjectIDFromHex(leaveId)

	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}

	result := models.LeaveCollection.FindOneAndUpdate(
		bson.M{"_id": objId},
		bson.D{
			{"$set", bson.M{
				"task":                updateLeave.Task,
				"status":              updateLeave.Status,
				"updated_at":          time.Now(),
			}},
		},
		&opt,
	)

	if result.Err() != nil {
		log.Println("Err ", result.Err())
		return models.Leave{}, result.Err()
	}

	var leave models.Leave
	if err := result.Decode(&leave); err != nil {
		return models.Leave{}, err
	}

	return leave, nil
}

func ALeave(leaveId string) models.Leave {
	var leave models.Leave

	objId, _ := primitive.ObjectIDFromHex(leaveId)

	err := models.LeaveCollection.FindOne(bson.M{"_id": objId}).Decode(&leave)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	return leave
}

func DeleteALeave(leaveId string) bool {
	objId, _ := primitive.ObjectIDFromHex(leaveId)

	result := models.LeaveCollection.FindOneAndDelete(bson.D{{"_id", objId}})

	if result.Err() != nil {
		return false
	}

	return true
}