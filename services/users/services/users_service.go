package services

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/JubaerHossain/golang_restapi/services/users/models"
	"github.com/JubaerHossain/golang_restapi/services/users/validation"
	pagination "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func AllUsers(requestFilter map[string]interface{}) ([]models.User, pagination.PaginationData, error) {
	var users []models.User

	// Initialize the filter
	filter := bson.M{}

	// Parse page and limit from requestFilter
	pageStr, pageOk := requestFilter["page"].(string)
	limitStr, limitOk := requestFilter["limit"].(string)
	if !pageOk || !limitOk {
		return nil, pagination.PaginationData{}, errors.New("invalid or missing page or limit")
	}

	page, err := strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		return nil, pagination.PaginationData{}, err
	}

	limit, err := strconv.ParseInt(limitStr, 10, 64)
	if err != nil {
		return nil, pagination.PaginationData{}, err
	}

	// Initialize the pagination data
	paginatedData, err := pagination.New(models.UsersCollection.Collection).
		Page(page).
		Limit(limit).
		Decode(&users).
		Filter(filter).
		Find()

	if err != nil {
		return nil, pagination.PaginationData{}, err
	}

	return users, paginatedData.Pagination, nil
}


func CreateAUsers(createUsers validation.CreateUsersRequest) models.User {
	user := models.User{
		Id:        primitive.NewObjectID(),
		Name:      createUsers.Name,
	}

	result, err := models.UsersCollection.InsertOne(user)
	if err != nil || result == nil {
		panic(err)
	}

	return user
}

func UpdateAUsers(userId string, updateUsers validation.UpdateUsersRequest) (models.User, error) {

	objId, _ := primitive.ObjectIDFromHex(userId)

	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}

	result := models.UsersCollection.FindOneAndUpdate(
		bson.M{"_id": objId},
		bson.D{
			{"$set", bson.M{
				"name":       updateUsers.Name,
				"updated_at": time.Now(),
			}},
		},
		&opt,
	)

	if result.Err() != nil {
		log.Println("Err ", result.Err())
		return models.User{}, result.Err()
	}

	var user models.User
	if err := result.Decode(&user); err != nil {
		return models.User{}, err
	}

	return user, nil
}

func AUsers(userId string) models.User {
	var user models.User

	objId, _ := primitive.ObjectIDFromHex(userId)

	err := models.UsersCollection.FindOne(bson.M{"_id": objId}).Decode(&user)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	return user
}

func DeleteAUsers(userId string) bool {
	objId, _ := primitive.ObjectIDFromHex(userId)

	result := models.UsersCollection.FindOneAndDelete(bson.D{{"_id", objId}})

	if result.Err() != nil {
		return false
	}

	return true
}
