package Providers

import "github.com/gofiber/fiber/v2"

func RESPONSEJSON(ctx *fiber.Ctx, data interface{}, message string, response int) error {
	return ctx.Status(response).JSON(fiber.Map{
		"message":       message,
		"response_code": response,
		"data":          data,
	})
}
