package controllers

import (
	"auth-go/database"
	"auth-go/models"

	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

const SecretKey = "secret"

func Login(c *fiber.Ctx) error {

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {

		return err
	}

	var user models.User

	database.DB.Where("user_email = ?", data["email"]).First(&user)

	if user.UserID == 0 {

		c.Status(fiber.StatusNotFound)

		return c.JSON(fiber.Map{
			"message": "User not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"])); err != nil {

		c.Status(fiber.StatusBadRequest)

		return c.JSON(fiber.Map{
			"message": "Wrong password",
		})

	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{

		Issuer:    strconv.Itoa(int(user.UserID)),
		ExpiresAt: time.Now().Add(time.Hour * 12).Unix(),
	})

	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {

		c.Status(fiber.StatusInternalServerError)

		return c.JSON(fiber.Map{
			"message": "Could not login",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true, //nao pode ser chamado pelo frontend,

	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "Success",
	})

}

//valida se o utilizador tem login(jwt+cookie)
func LoggedIn(c *fiber.Ctx) error {

	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {

		return []byte(SecretKey), nil

	})

	if err != nil {

		c.Status(fiber.StatusUnauthorized)

		return c.JSON(fiber.Map{
			"message": "Unauthenticated",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User

	database.DB.Where("user_id = ?", claims.Issuer).First(&user)

	return c.JSON(user)

}

func Logout(c *fiber.Ctx) error {

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true, //nao pode ser chamado pelo frontend,

	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{

		"message": "you successfully logged out",
	})

}
