package controllers

import (
	"net/http"

	"github.com/JubaerHossain/golang_restapi/services/users/services"
	"github.com/JubaerHossain/golang_restapi/services/users/validation"
	"github.com/JubaerHossain/gomd/gomd"
	"github.com/gin-gonic/gin"
)

func UsersIndex() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get query parameters
		page := c.DefaultQuery("page", "1")
		limit := c.DefaultQuery("limit", "10")
		status := c.DefaultQuery("status", "")

		// Create a filter map
		filter := map[string]interface{}{
			"page":   page,
			"limit":  limit,
			"status": status,
		}

		// Fetch users with the provided filter
		users, paginate, err := services.AllUsers(filter)
		if err != nil {
			gomd.Res.Code(http.StatusInternalServerError).Message("error").Data(err).Json(c)
			return
		}

		// Respond with the user data and pagination information
		responseData := map[string]interface{}{
			"data": users,
			"pagination": paginate,
		}
		gomd.Res.Code(http.StatusOK).Data(responseData).Json(c)
	}
}


func UsersCreate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var createUsers validation.CreateUsersRequest

		defer func() {
			if err := recover(); err != nil {
				gomd.Res.Code(http.StatusUnprocessableEntity).Message("error").Data(err).Json(c)
			}
		}()

		if err := c.ShouldBind(&createUsers); err != nil {
			gomd.Res.Code(http.StatusBadRequest).Message("Bad Request").Data(err.Error()).AbortWithStatusJSON(c)
			return
		}

		user := services.CreateAUsers(createUsers)

		gomd.Res.Code(http.StatusCreated).Message("success").Data(user).Json(c)
	}
}

func UsersShow() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				gomd.Res.Code(http.StatusNotFound).Message(http.StatusText(http.StatusNotFound)).Json(c)
			}
		}()

		userId := c.Param("userId")

		user := services.AUsers(userId)

		gomd.Res.Code(http.StatusOK).Message("success").Data(user).Json(c)
	}
}

func UsersUpdate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var updateUsers validation.UpdateUsersRequest

		defer func() {
			if err := recover(); err != nil {
				gomd.Res.Code(http.StatusUnprocessableEntity).Message(http.StatusText(http.StatusUnprocessableEntity)).Data(err).Json(c)
			}
		}()

		userId := c.Param("userId")

		if err := c.ShouldBind(&updateUsers); err != nil {
			gomd.Res.Code(http.StatusBadRequest).Message(http.StatusText(http.StatusBadRequest)).Data(err.Error()).AbortWithStatusJSON(c)
			return
		}

		user, err := services.UpdateAUsers(userId, updateUsers)

		if err != nil {
			gomd.Res.Code(http.StatusInternalServerError).Message(http.StatusText(http.StatusInternalServerError)).Json(c)
			return
		}

		gomd.Res.Code(http.StatusOK).Message("Successfully Updated !!!").Data(user).Json(c)
	}
}

func UsersDelete() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				gomd.Res.Code(http.StatusUnprocessableEntity).Message("error").Data(err).Json(c)
			}
		}()

		userId := c.Param("userId")
		err := services.DeleteAUsers(userId)

		if !err {
			gomd.Res.Code(http.StatusInternalServerError).Message("something wrong").Json(c)
			return
		}

		gomd.Res.Code(http.StatusOK).Message("Successfully Delete !!!").Json(c)
	}
}
