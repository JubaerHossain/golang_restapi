package routes

import (
	"github.com/JubaerHossain/gomd/gomd"
	c "github.com/JubaerHossain/golang_restapi/services/attendance/controllers"
)

func AttendanceSetup() {
    v1 := gomd.Router.Group("api/v1")
	v1.GET("attendances", c.AttendanceIndex())
	v1.POST("attendances", c.AttendanceCreate())
	v1.GET("attendances/:attendanceId", c.AttendanceShow())
	v1.PUT("attendances/:attendanceId", c.AttendanceUpdate())
	v1.DELETE("attendances/:attendanceId", c.AttendanceDelete())
}
