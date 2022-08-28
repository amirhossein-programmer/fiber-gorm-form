package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/sixfwa/fiber-api/dist/cmd/web/database"
	"github.com/sixfwa/fiber-api/dist/cmd/web/models"
)

type User struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func createResponseUser(userModel models.User) User {
	return User{ID: userModel.Id, FirstName: userModel.FirstName, LastName: userModel.LastName, Email: userModel.Email, Password: userModel.Password}
}
func CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.Database.Db.Create(&user)
	responseUser := createResponseUser(user)
	return c.Status(200).JSON(responseUser)
}
