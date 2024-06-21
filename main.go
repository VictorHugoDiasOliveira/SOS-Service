package main

import (
	"authservice/config"
	"authservice/handlers"
)

func main() {
	r := config.SetupRouter()

	r = handlers.GetUsers(r)
	r = handlers.Register(r)
	r = handlers.Login(r)

	r.Run(":8080")
}
