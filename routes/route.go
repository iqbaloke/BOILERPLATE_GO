package routes

import (
	"github.com/gofiber/fiber/v2"
	"siskol_go_be/app/Controllers"
)

func Route(route *fiber.App) {

	route.Get("/", func(route *fiber.Ctx) error {
		return route.JSON(fiber.Map{
			"message": "golang",
			"version": "1.20.7",
		})
	})

	route.Get("/users", Controllers.UsersIndex)
	route.Post("/users", Controllers.UserStore)

}
