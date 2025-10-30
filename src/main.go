package main

import (
	"Weapons-API/src/env"
	"Weapons-API/src/web/routers"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
 	server := echo.New()

	// setting up me routers
	routers.InitHealthRouter().Bind(server)
	routers.InitWeaponsRouter().Bind(server)
	routers.InitUsersRouter().Bind(server)

	defaultPort := 80
	port := env.GetInt("PORT", &defaultPort)
	log.Fatal(server.Start(fmt.Sprintf(":%d", port)))
}