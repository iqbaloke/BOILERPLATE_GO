package Providers

import "github.com/gofiber/fiber/v2"

func SUCCESS(ctx *fiber.Ctx, data interface{}, message string, response int) error {
	return ctx.Status(response).JSON(fiber.Map{
		"message":       message,
		"response_code": response,
		"data":          data,
	})
}

func VALIDATION(ctx *fiber.Ctx, data interface{}, message string, statusCode int) error {
	return ctx.Status(statusCode).JSON(fiber.Map{
		"message":     message,
		"status_code": statusCode,
		"data":        data,
	})
}
