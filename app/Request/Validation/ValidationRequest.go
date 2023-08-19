package Validation

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ValidationRequest(ctx *fiber.Ctx, model interface{}) error {

	validate := validator.New()
	errValidate := validate.Struct(model)
	if err := ctx.BodyParser(model); err != nil {
		return err
	}
	if err := errValidate; err != nil {
		return err
	}
	return nil
}
