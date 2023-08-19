package migrations

import (
	"fmt"
	"log"
	"siskol_go_be/config"
)

func RunMigration() {
	//checkUser := database.DB.Migrator().HasTable(&User{})
	//if checkUser {
	//	fmt.Println("ready")
	//} else {
	err := config.DB.AutoMigrate(&User{})

	if err != nil {
		log.Println(err)
		fmt.Println("Migration faild")
	}

	fmt.Println("Migration Succeess")
	//}

}
