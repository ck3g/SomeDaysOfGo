package main

import (
	"flag"

	"github.com/ck3g/SomeDaysOfGo/hotel-reservation/api"
	"github.com/gofiber/fiber/v2"
)

func main() {
	listenAddr := flag.String("listenAddr", ":5005", "Server listen address of API")
	flag.Parse()

	app := fiber.New()
	apiv1 := app.Group("/api/v1")

	apiv1.Get("/users", api.HandleGetUsers)
	apiv1.Get("/users/:id", api.HandleGetUser)
	app.Listen(*listenAddr)
}
