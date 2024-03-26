package user

import (
	"Tweteroo/model"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Handler struct{
	DB *gorm.DB
}

func (h *Handler) GetUserByID(c *fiber.Ctx) error{
	var user model.User
	id := c.Params("id")
	result := h.DB.Where("id = ?", id).First(&user)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error" : "Erro ao buscar tarefa",
		})
	}
	return c.JSON(user)
}

func (h *Handler) GetAllUsers(c *fiber.Ctx) error{
	var users []model.User
	result := h.DB.Find(&users)
	if result.Error != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error" : "Erro ao buscar todas as tarefas",
		})
	}
	return c.JSON(users)
}

func (h *Handler) CreateUser(c *fiber.Ctx) error{
	user := new(model.User)
	if err := c.BodyParser(&user); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error" : "Dados de usuário invalido",
		})
	}

	if user.Avatar == "" && user.Username == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error" : "O seu nome de usuario e avatar nao podem estar vazios",
		})
	}

	h.DB.Create(&user)
	return c.Status(fiber.StatusCreated).JSON(&user)
}

func (h *Handler) UpdateUser(c *fiber.Ctx) error{
	id := c.Params("id")
	user := new(model.User)
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	var existingUser model.User
	result := h.DB.First(&existingUser, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).SendString("Sem users com esse ID" + id)
	}
	existingUser.Username = user.Username
	existingUser.Avatar = user.Avatar
	h.DB.Save(&existingUser)
	return c.JSON(&existingUser)
}

func (h *Handler) DeleteUser(c *fiber.Ctx) error{
	id := c.Params("id")
	var user model.User
	result := h.DB.First(&user, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).SendString("Sem users com esse ID" + id)
	}
	h.DB.Delete(&user)
	return c.SendString("Usuário deletado com sucesso")
}




