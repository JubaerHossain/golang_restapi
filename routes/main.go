package routes

import (
	attendanceRoute "github.com/JubaerHossain/golang_restapi/services/attendance/routes"
	"github.com/JubaerHossain/gomd/config"
	. "github.com/JubaerHossain/gomd/gomd"
	"github.com/gin-gonic/gin"
)

func Register() {
	BaseRoute()

	attendanceRoute.AttendanceSetup()
}

func BaseRoute() {
	Router.GET("/", func(c *gin.Context) {
		data := map[string]interface{}{
			"app": config.Config.GetString("App.Name"),
		}
		Res.Code(200).
			Message("success").
			Data(data).Json(c)
	})
}
