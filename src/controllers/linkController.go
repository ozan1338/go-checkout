package controllers

import (
	"encoding/json"
	"fmt"
	"go-checkout/src/database"
	"go-checkout/src/models"
	"go-checkout/src/services"

	"github.com/gofiber/fiber/v2"
)


func GetLink(c *fiber.Ctx) error {
	code := c.Params("code")

	var link models.Link

	database.DB.Preload("Products").Where("code = ?", code).First(&link)

	if link.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Invalid Link",
		})
	}

	resp,err := services.UserService.Get(fmt.Sprintf("users/%d", link.UserId), "")

	if err != nil {
		panic(err)
	}

	var user models.User

	json.NewDecoder(resp.Body).Decode(&user)

	link.User = user

	return c.JSON(link)
}
