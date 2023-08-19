package Controllers

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"siskol_go_be/app/Providers"
	"siskol_go_be/config"
	"siskol_go_be/databases/migrations"
)

func UsersIndex(ctx *fiber.Ctx) error {
	var users []migrations.User
	result := config.DB.Find(&users)

	if result.Error != nil {
		log.Println(result.Error)
	}
	return Providers.SUCCESS(ctx, users, "Data Collection Berhasil Diambil", 200)
}
