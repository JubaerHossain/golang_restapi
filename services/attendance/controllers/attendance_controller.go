package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/JubaerHossain/gomd/gomd"
	"github.com/JubaerHossain/gomd/services/attendance/validation"
	"github.com/JubaerHossain/gomd/services/attendance/services"
	"net/http"
)
func AttendanceIndex() gin.HandlerFunc {
	return func(c *gin.Context) {
        page := c.DefaultQuery("page", "1")
        limit := c.DefaultQuery("limit", "10")
        status := c.DefaultQuery("status", "")

        var filter map[string]interface{} = make(map[string]interface{})
        filter["page"] = page
        filter["limit"] = limit
        filter["status"] = status

        attendances, paginate := services.AllAttendance(filter)

        gomd.Res.Code(200).Data(attendances).Raw(map[string]interface{}{
            "meta": paginate,
        }).Json(c)
	}
}


func AttendanceCreate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var createAttendance validation.CreateAttendanceRequest

		defer func() {
			if err := recover(); err != nil {
				gomd.Res.Code(http.StatusUnprocessableEntity).Message("error").Data(err).Json(c)
			}
		}()

		if err := c.ShouldBind(&createAttendance); err != nil {
			gomd.Res.Code(http.StatusBadRequest).Message("Bad Request").Data(err.Error()).AbortWithStatusJSON(c)
			return
		}

		attendance := services.CreateAAttendance(createAttendance)

		gomd.Res.Code(http.StatusCreated).Message("success").Data(attendance).Json(c)
	}
}


func AttendanceShow() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				gomd.Res.Code(http.StatusNotFound).Message(http.StatusText(http.StatusNotFound)).Json(c)
			}
		}()

		attendanceId := c.Param("attendanceId")

		attendance := services.AAttendance(attendanceId)

		gomd.Res.Code(http.StatusOK).Message("success").Data(attendance).Json(c)
	}
}


func AttendanceUpdate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var updateAttendance validation.UpdateAttendanceRequest

		defer func() {
			if err := recover(); err != nil {
				gomd.Res.Code(http.StatusUnprocessableEntity).Message(http.StatusText(http.StatusUnprocessableEntity)).Data(err).Json(c)
			}
		}()

		attendanceId := c.Param("attendanceId")

		if err := c.ShouldBind(&updateAttendance); err != nil {
			gomd.Res.Code(http.StatusBadRequest).Message(http.StatusText(http.StatusBadRequest)).Data(err.Error()).AbortWithStatusJSON(c)
			return
		}

		attendance, err := services.UpdateAAttendance(attendanceId, updateAttendance)

		if err != nil {
			gomd.Res.Code(http.StatusInternalServerError).Message(http.StatusText(http.StatusInternalServerError)).Json(c)
			return
		}

		gomd.Res.Code(http.StatusOK).Message("Successfully Updated !!!").Data(attendance).Json(c)
	}
}


func AttendanceDelete() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				gomd.Res.Code(http.StatusUnprocessableEntity).Message("error").Data(err).Json(c)
			}
		}()

		attendanceId := c.Param("attendanceId")
		err := services.DeleteAAttendance(attendanceId)

		if !err {
			gomd.Res.Code(http.StatusInternalServerError).Message("something wrong").Json(c)
			return
		}

		gomd.Res.Code(http.StatusOK).Message("Successfully Delete !!!").Json(c)
	}
}