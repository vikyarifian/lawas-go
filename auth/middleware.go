package auth

import (
	"lawas-go/dto"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func IsAuthenticated(c *fiber.Ctx) (dto.Token, bool) {

	token := c.Cookies("session")
	_, err := VerifyToken(token)
	var jwt dto.Token
	if err == nil {
		jwt, _ := GetToken(c)
		return jwt, true
	}
	return jwt, false

}

func AssertAuthenticatedMiddleware(c *fiber.Ctx) error {
	if _, r := IsAuthenticated(c); !r {
		c.Set("HX-Redirect", "/login")
		return c.Redirect("/login")
	}
	return c.Next()
}

func GetUserSessionId(c *fiber.Ctx) string {
	return c.Cookies("session")
}

func SetSession(c *fiber.Ctx) string {
	newSessionId := c.Get("Authorization")
	c.Cookie(&fiber.Cookie{
		Name:     "session",
		Value:    newSessionId,
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Strict",
	})
	return newSessionId
}

func ClearSession(c *fiber.Ctx) {
	c.Cookie(&fiber.Cookie{
		Name:     "session",
		Value:    "",
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Strict",
	})
}

func UserMiddleware(c *fiber.Ctx) error {

	token := c.Get("Authorization")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(dto.Response1{
			Data:    nil,
			Message: "Unauthorized",
			Code:    401,
			Status:  "Error",
		})
	}

	claims, err := VerifyToken(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(dto.Response1{
			Data:    nil,
			Message: "Unauthorized",
			Code:    401,
			Status:  "Error",
		})
	}

	c.Locals("username", claims["username"])

	return c.Next()
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
