package Controllers

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"siskol_go_be/app/Exception"
	"siskol_go_be/app/Models"
	"siskol_go_be/app/Providers"
	"siskol_go_be/app/Request/Validation"
	"siskol_go_be/config"
	"siskol_go_be/databases/migrations"
)

func UsersIndex(ctx *fiber.Ctx) error {
	var users []Models.UserModel
	result := config.DB.Find(&users)

	if result.Error != nil {
		log.Println(result.Error)
	}
	return Providers.RESPONSEJSON(ctx, users, "Data Collection Berhasil Diambil", fiber.StatusOK)
}

func UserStore(ctx *fiber.Ctx) error {
	user := new(Models.UserModel)
	err := ctx.BodyParser(user)

	if err := Validation.ValidationRequest(ctx, user); err != nil {
		return Exception.ErrorValidation(ctx, err)
	}

	if err != nil {
		return err
	}

	newUser := migrations.UserMigration{
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
	}

	errCreateDb := config.DB.Create(&newUser).Error
	if errCreateDb != nil {
		return Providers.RESPONSEJSON(ctx, errCreateDb.Error(), "Data Gagal Ditambahkan", fiber.StatusOK)
	}

	return Providers.RESPONSEJSON(ctx, newUser, "Data Berhasil Ditambahkan", fiber.StatusOK)
}
