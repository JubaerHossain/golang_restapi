package main

import (
	"github.com/JubaerHossain/golang_restapi/config"
	"github.com/JubaerHossain/golang_restapi/routes"
	"github.com/JubaerHossain/gomd/gomd"
)

func main() {

	gomd.New()
	config.Register()
	routes.Register()

	gomd.NoSqlConnection()
	// gomd.DatabaseConnection()
	config.Boot()

	gomd.Run()
}
