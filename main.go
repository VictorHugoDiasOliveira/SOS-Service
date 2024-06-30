package main

import (
	"log"
	"sosservice/src/configurations"
	"sosservice/src/configurations/logger"
	"sosservice/src/controller/routes"

	"github.com/joho/godotenv"
)

func main() {
	logger.Info("Starting application")

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	// config := &storage.Config{
	// 	Host:     os.Getenv("DB_HOST"),
	// 	User:     os.Getenv("DB_USER"),
	// 	Password: os.Getenv("DB_PASSWORD"),
	// 	DBName:   os.Getenv("DB_NAME"),
	// 	Port:     os.Getenv("DB_PORT"),
	// 	SSLMode:  os.Getenv("DB_SSLMode"),
	// }
	// storage.ConnectDatabase(config)

	// models.MigrateUsers(storage.DB)
	// models.MigrateCauses(storage.DB)
	// models.MigrateUserInfos(storage.DB)

	router := configurations.SetupRouter()

	routes.InitializeRoutes(&router.RouterGroup)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}

	// authorized := r.Group("/")
	// authorized.Use(middlewares.AuthenticationMiddleware())

	// // Get Users
	// controller.GetUsers(authorized)
	// controller.GetUsersById(authorized)

	// // Auth
	// controllers.UpdateAdminStatus(authorized)
	// r = controllers.Register(r)
	// r = controllers.Login(r)

	// // Causes
	// controllers.RegisterCause(authorized)
	// controllers.GetCauses(authorized)

	// // UserInfo
	// controllers.RegisterInfo(authorized)

}
