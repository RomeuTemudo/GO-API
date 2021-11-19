package controllers

import (
	"auth-go/database"
	"auth-go/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func GetRoles(c *fiber.Ctx) error {

	var roles []models.UsersRoles

	database.DB.Find(&roles)

	return c.JSON(roles)

}

func GetUsers(c *fiber.Ctx) error {

	//use model without password
	var userList []models.UserOutput
	//just to specify table
	user := new(models.User)

	//database.DB.Model(user).Find(&userList)

	database.DB.Model(user).Select("users.user_id,users.user_email,users.role_id ,users_roles.role_description").Joins("inner join users_roles on users.role_id = users_roles.role_id").Scan(&userList)

	return c.JSON(userList)

}

func AddUser(c *fiber.Ctx) error {

	user := new(models.User)

	userFull := new(models.UserOutput)

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)

	user_hashed := models.User{

		UserEmail: user.UserEmail,
		Password:  string(passwordHash),
		RoleID:    user.RoleID,
	}

	database.DB.Create(&user_hashed)

	database.DB.Where("user_id = ?", &user_hashed.UserID).Model(user).Select("users.user_id,users.user_email,users.role_id ,users_roles.role_description").Joins("inner join users_roles on users.role_id = users_roles.role_id").Find(&userFull)

	return c.JSON(&userFull)
}

func DeleteUser(c *fiber.Ctx) error {

	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return err
	}

	//database.DB.Model(&sensor).Where("sensor_id=?", sensor.SensorID).Delete(&sensor)

	//database.DB.Delete(models.Sensor{SensorID: sensor.SensorID}, sensor.SensorID)

	database.DB.Delete(&user, user.UserID)

	return c.JSON(fiber.Map{
		"Status":  "Success",
		"Message": "Boa burro , apagaste um sensor!",
	})

}
