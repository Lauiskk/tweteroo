package user

import (
	"Tweteroo/model"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func (h *Handler) GetUserByID(c *fiber.Ctx) error {
	var user model.User
	id := c.Params("id")
	result := h.DB.Where("id = ?", id).First(&user)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Erro ao buscar usuario",
		})
	}

	if user.Tweets == nil {
		var tweets []model.Tweet
		h.DB.Model(&user).Association("Tweets").Find(&tweets)
		user.Tweets = tweets
	}

	return c.JSON(user)
}

func (h *Handler) GetAllUsers(c *fiber.Ctx) error {
	var users []model.User
	result := h.DB.Find(&users)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Erro ao buscar todos os usuarios",
		})
	}
	return c.JSON(users)
}

func (h *Handler) CreateUser(c *fiber.Ctx) error {
	user := new(model.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Dados de usuário invalido",
		})
	}

	if user.Avatar == "" && user.Username == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "O seu nome de usuario e avatar nao podem estar vazios",
		})
	}

	var existingUser model.User

	result := h.DB.Where("Username = ?", user.Username).First(&existingUser)
	if result.RowsAffected != 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Este usuário já existe",
		})
	}

	h.DB.Create(&user)
	return c.Status(fiber.StatusCreated).JSON(user)
}

func (h *Handler) UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user model.User
	result := h.DB.First(&user, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error ": "Sem users com esse ID",
		})
	}

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "error",
		})
	}

	h.DB.Save(&user)
	return c.JSON(&user)
}

func (h *Handler) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user model.User
	result := h.DB.First(&user, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).SendString("Sem users com esse ID" + id)
	}
	h.DB.Delete(&user)
	return c.SendString("Usuário deletado com sucesso")
}

func (h *Handler) CreateTweet(c *fiber.Ctx) error {
	id := c.Params("id")
	tweet := new(model.Tweet)
	var user model.User
	result := h.DB.First(&user, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).SendString("Sem usuários com esse ID " + id)
	}

	if err := c.BodyParser(tweet); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Dados inválidos",
		})
	}

	if tweet.Tweet == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Não é possível criar um tweet vazio",
		})
	}

	tweet.UserID = user.ID

	h.DB.Create(&tweet)

	h.DB.Model(tweet).Association("User").Find(&tweet.User)

	return c.Status(fiber.StatusCreated).JSON(&tweet)
}

func (h *Handler) GetAllTweets(c *fiber.Ctx) error {
	var tweets []model.Tweet

	result := h.DB.Preload("User").Find(&tweets)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Erro ao buscar todos os tweets",
		})
	}

	return c.JSON(tweets)
}

func (h *Handler) DeleteTweet(c *fiber.Ctx) error {
	id := c.Params("id")
	var tweet model.Tweet
	result := h.DB.First(&tweet, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).SendString("Sem tweet com esse ID" + id)
	}
	h.DB.Delete(&tweet)

	return c.SendString("Tweet deletado com sucesso")
}

func (h *Handler) GetTweetsByUser(c *fiber.Ctx) error {
	id := c.Params("id")

	var tweets []model.Tweet

	result := h.DB.Preload("User").Where("user_id = ?", id).Find(&tweets)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Erro ao buscar os tweets do usuário",
		})
	}

	return c.JSON(tweets)
}
