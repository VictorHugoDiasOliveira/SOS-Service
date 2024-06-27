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
	config.DB.AutoMigrate(&models.User{}, &models.Cause{})

	authorized := r.Group("/")
	authorized.Use(middlewares.AuthenticationMiddleware())

	// Get Users
	controllers.GetUsers(authorized)
	controllers.GetUsersById(authorized)

	// Auth
	controllers.UpdateAdminStatus(authorized)
	r = controllers.Register(r)
	r = controllers.Login(r)

	// Causes
	controllers.RegisterCause(authorized)
	controllers.GetCauses(authorized)

	r.Run(":8080")
}
