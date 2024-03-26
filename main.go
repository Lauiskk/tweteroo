package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main(){
	db, err := gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{})
	if err != nil {
        panic("failed to connect database")
    }

	db.AutoMigrate(&Tweet{})

	handler := tweet.Handler{DB: db}

	app := fiber.New()

	api := app.Group("/api")


	err := app.Listen(":8080")
	if err != nil{
		log.Fatal(err)
	}
}