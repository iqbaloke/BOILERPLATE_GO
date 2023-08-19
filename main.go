package main

import (
	"github.com/gofiber/fiber/v2"
	"siskol_go_be/config"
	"siskol_go_be/databases/migrations"
	"siskol_go_be/routes"
)

func main() {
	app := fiber.New()
	config.Connection()
	migrations.RunMigration()
	routes.Route(app)
	app.Listen(":8000")

}
