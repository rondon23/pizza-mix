package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rondon23/pizza-mix/database"
	"github.com/rondon23/pizza-mix/models"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data["senha"] != data["senha_confirmacao"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "a senha nao confere",
		})
	}

	funcionario := models.Funcionario{}

	funcionario.DefinirSenha(data["senha"])

	database.DB.Create(&funcionario)

	return c.JSON(funcionario)
}
