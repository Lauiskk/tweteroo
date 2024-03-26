package main

import (
	"Tweteroo/model"
	"Tweteroo/user"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main(){
	db, err := gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{})
	if err != nil {
        panic("failed to connect database")
    }

	db.AutoMigrate(&model.User{})

	handler := user.Handler{DB: db}

	app := fiber.New()

	api := app.Group("/api")

	api.Get("/", handler.GetAllUsers)
	
	api.Get("/user/:id", handler.GetUserByID)

	api.Post("/user", handler.CreateUser)

	api.Put("/user/:id", handler.UpdateUser)

	api.Delete("/user/:id", handler.DeleteUser)

	app.Listen(":8080")

	//err = app.Listen(":8080")
	//if err != nil{
	//	log.Fatal(err)
	//}
}