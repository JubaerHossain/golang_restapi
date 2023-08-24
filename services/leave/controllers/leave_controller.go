package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/JubaerHossain/gomd/gomd"
	"github.com/JubaerHossain/golang_restapi/services/leave/validation"
	"github.com/JubaerHossain/golang_restapi/services/leave/services"
	"net/http"
)
func LeaveIndex() gin.HandlerFunc {
	return func(c *gin.Context) {
        page := c.DefaultQuery("page", "1")
        limit := c.DefaultQuery("limit", "10")
        status := c.DefaultQuery("status", "")

        var filter map[string]interface{} = make(map[string]interface{})
        filter["page"] = page
        filter["limit"] = limit
        filter["status"] = status

        leaves, paginate := services.AllLeave(filter)

        gomd.Res.Code(200).Data(leaves).Raw(map[string]interface{}{
            "meta": paginate,
        }).Json(c)
	}
}


func LeaveCreate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var createLeave validation.CreateLeaveRequest

		defer func() {
			if err := recover(); err != nil {
				gomd.Res.Code(http.StatusUnprocessableEntity).Message("error").Data(err).Json(c)
			}
		}()

		if err := c.ShouldBind(&createLeave); err != nil {
			gomd.Res.Code(http.StatusBadRequest).Message("Bad Request").Data(err.Error()).AbortWithStatusJSON(c)
			return
		}

		leave := services.CreateALeave(createLeave)

		gomd.Res.Code(http.StatusCreated).Message("success").Data(leave).Json(c)
	}
}


func LeaveShow() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				gomd.Res.Code(http.StatusNotFound).Message(http.StatusText(http.StatusNotFound)).Json(c)
			}
		}()

		leaveId := c.Param("leaveId")

		leave := services.ALeave(leaveId)

		gomd.Res.Code(http.StatusOK).Message("success").Data(leave).Json(c)
	}
}


func LeaveUpdate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var updateLeave validation.UpdateLeaveRequest

		defer func() {
			if err := recover(); err != nil {
				gomd.Res.Code(http.StatusUnprocessableEntity).Message(http.StatusText(http.StatusUnprocessableEntity)).Data(err).Json(c)
			}
		}()

		leaveId := c.Param("leaveId")

		if err := c.ShouldBind(&updateLeave); err != nil {
			gomd.Res.Code(http.StatusBadRequest).Message(http.StatusText(http.StatusBadRequest)).Data(err.Error()).AbortWithStatusJSON(c)
			return
		}

		leave, err := services.UpdateALeave(leaveId, updateLeave)

		if err != nil {
			gomd.Res.Code(http.StatusInternalServerError).Message(http.StatusText(http.StatusInternalServerError)).Json(c)
			return
		}

		gomd.Res.Code(http.StatusOK).Message("Successfully Updated !!!").Data(leave).Json(c)
	}
}


func LeaveDelete() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				gomd.Res.Code(http.StatusUnprocessableEntity).Message("error").Data(err).Json(c)
			}
		}()

		leaveId := c.Param("leaveId")
		err := services.DeleteALeave(leaveId)

		if !err {
			gomd.Res.Code(http.StatusInternalServerError).Message("something wrong").Json(c)
			return
		}

		gomd.Res.Code(http.StatusOK).Message("Successfully Delete !!!").Json(c)
	}
}