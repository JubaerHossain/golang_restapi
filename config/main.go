package config

import (
	attendance "github.com/JubaerHossain/golang_restapi/services/attendance/models"
	user "github.com/JubaerHossain/golang_restapi/services/users/models"
	"github.com/JubaerHossain/gomd/config"
)

func Register() {
	config.Config.AddConfig("App", new(AppConfig))
	config.Config.AddConfig("NoSql", new(MongoConfig))
	config.Config.Load()
}

func Boot() {
	attendance.AttendanceSetup()
	user.UsersSetup()
}
