package routes

import (
	"github.com/JubaerHossain/gomd/gomd"
	c "github.com/JubaerHossain/golang_restapi/services/leave/controllers"
)

func LeaveSetup() {
    v1 := gomd.Router.Group("api/v1")
	v1.GET("leaves", c.LeaveIndex())
	v1.POST("leaves", c.LeaveCreate())
	v1.GET("leaves/:leaveId", c.LeaveShow())
	v1.PUT("leaves/:leaveId", c.LeaveUpdate())
	v1.DELETE("leaves/:leaveId", c.LeaveDelete())
}
