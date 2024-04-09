package main

import (
	user "Tweteroo/api"
	"Tweteroo/model"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Tweet{})

	handler := user.Handler{DB: db}

	app := fiber.New()

	app.Use(cors.New())

	api := app.Group("/api")

	api.Get("/", handler.GetAllUsers)

	api.Get("/users/:id", handler.GetUserByID)
	api.Get("/tweets", handler.GetAllTweets)
	api.Get("/users/:id/tweets", handler.GetTweetsByUser)

	api.Post("/users", handler.CreateUser)
	api.Post("/users/:id/tweet", handler.CreateTweet)

	api.Put("/users/:id", handler.UpdateUser)

	api.Delete("/users/:id", handler.DeleteUser)
	api.Delete("/tweets/:id", handler.DeleteTweet)

	err = app.Listen(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
