package routes

import (
	"github.com/JubaerHossain/gomd/gomd"
	c "github.com/JubaerHossain/golang_restapi/services/users/controllers"
)

func UsersSetup() {
    v1 := gomd.Router.Group("api/v1")
	v1.GET("users", c.UsersIndex())
	v1.POST("users", c.UsersCreate())
	v1.GET("users/:userId", c.UsersShow())
	v1.PUT("users/:userId", c.UsersUpdate())
	v1.DELETE("users/:userId", c.UsersDelete())
}
