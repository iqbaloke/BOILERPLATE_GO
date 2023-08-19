package Providers

import "github.com/gofiber/fiber/v2"

func RESPONSEJSON(ctx *fiber.Ctx, data interface{}, message string, response int) error {
	return ctx.Status(response).JSON(fiber.Map{
		"message":       message,
		"response_code": response,
		"data":          data,
	})
}

func RESPONSENOTFOUND(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"status_code": fiber.StatusNotFound,
		"message":     "Data Not Found",
	})
}
