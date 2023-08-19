package Exception

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"siskol_go_be/app/Providers"
	"siskol_go_be/app/Request/Validation"
)

func ErrorValidation(c *fiber.Ctx, err error) error {
	out := make([]Validation.ValidationFormatting, len(err.(validator.ValidationErrors)))
	for i, fe := range err.(validator.ValidationErrors) {
		out[i] = Validation.ValidationFormatting{Message: Validation.ValidationMessage(fe.Tag(), fe.Field(), fe.Param())}
	}

	return Providers.RESPONSEJSON(c, out, "Data Gagal Ditambahkan", fiber.StatusUnprocessableEntity)

}

//func ()  {
//
//}
