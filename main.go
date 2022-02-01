package main

import (
	"log"
	"start-up/user"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/start-up?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	userInput := user.RegisterUserInput{}
	userInput.Name = "Test User Input"
	userInput.Email = "din@testnih.com"
	userInput.Occupation = "Nganggur"
	userInput.Password = "inipass"

	userService.RegisterUser(userInput)
}
