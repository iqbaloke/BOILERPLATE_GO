package FormattingMain

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

type messageRequest struct {
	Message string `json:"message"`
}

func messageFormatting(tag string, msg string, parms string) string {
	switch tag {
	case "required":
		return msg + " Wajib Diisi"
	case "max":
		return msg + " Maksimal " + parms + " Karakter"
	case "email":
		return "Invalid email"
	}
	return ""
}

func ValidationFormatting(model interface{}) []messageRequest {

	validate := validator.New()
	errValidate := validate.Struct(model)

	if errValidate != nil {
		out := make([]messageRequest, len(errValidate.(validator.ValidationErrors)))
		for i, fe := range errValidate.(validator.ValidationErrors) {
			out[i] = messageRequest{messageFormatting(fe.Tag(), fe.Field(), fe.Param())}
		}
		return out
	}
	return nil
}
