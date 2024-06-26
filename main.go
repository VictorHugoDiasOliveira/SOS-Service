package main

import (
	"authservice/config"
	"authservice/controllers"
	"authservice/middlewares"
	"authservice/models"
)

func main() {
	r := config.SetupRouter()

	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.User{})

	authorized := r.Group("/")
	authorized.Use(middlewares.AuthMiddleware())

	controllers.GetUsers(authorized)
	controllers.GetUsersById(authorized)
	controllers.UpdateAdminStatus(authorized)

	r = controllers.Register(r)
	r = controllers.Login(r)

	r.Run(":8080")
}
