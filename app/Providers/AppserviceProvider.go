package Providers

import "github.com/gofiber/fiber/v2"

func SUCCESS(ctx *fiber.Ctx, data interface{}, message string, response int) error {
	return ctx.Status(response).JSON(fiber.Map{
		"message":       message,
		"response_code": response,
		"data":          data,
	})
}

func VALIDATION(ctx *fiber.Ctx, data interface{}, message string, status_code int) error {
	return ctx.Status(status_code).JSON(fiber.Map{
		"message":     message,
		"status_code": status_code,
		"data":        data,
	})
}
