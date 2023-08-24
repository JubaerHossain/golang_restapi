package routes

import (
	attendanceRoute "github.com/JubaerHossain/golang_restapi/services/attendance/routes"
	userRoute "github.com/JubaerHossain/golang_restapi/services/users/routes"
	"github.com/JubaerHossain/gomd/config"
	. "github.com/JubaerHossain/gomd/gomd"
	"github.com/gin-gonic/gin"
)

func Register() {
	Router.GET("/", func(c *gin.Context) {
		data := map[string]interface{}{
			"app": config.Config.GetString("App.Name"),
		}
		Res.Code(200).
			Message("success").
			Data(data).Json(c)
	})
	attendanceRoute.AttendanceSetup()
	userRoute.UsersSetup()
}
