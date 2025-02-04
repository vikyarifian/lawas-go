package routes

import (
	"lawas-go/auth"
	"lawas-go/components"
	"lawas-go/db"
	"lawas-go/models"
	"lawas-go/pages"
	"lawas-go/utils"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func ApiRoutes(app *fiber.App) {

	// api := app.Group("/api/v1")
	app.Get("/a", func(c *fiber.Ctx) error {
		var u models.User
		db.MySql.Find(&u)
		return c.Status(fiber.StatusOK).JSON(u)
	})
}

func WebRoutes(app *fiber.App) {

	app.Get("/", func(c *fiber.Ctx) error {
		return utils.Render(c, pages.Frontpage(auth.IsAuthenticated(c)))
	})

	app.Post("/register", func(c *fiber.Ctx) error {

		var user models.User
		var count int64
		username := c.FormValue("username")
		email := c.FormValue("email")
		phone := c.FormValue("phone")
		password := c.FormValue("password")
		err := db.MySql.Model(&user).Where("username=? or email=?", username, email).Count(&count).Error
		if err != nil {
			return utils.Render(c, components.ErrorAlert(err.Error(), "register"), templ.WithStatus(http.StatusBadRequest))
		}
		if count > 0 {
			return utils.Render(c, components.ErrorAlert("Username or Email already exist!", "register"), templ.WithStatus(http.StatusBadRequest))
		}

		t := time.Now()
		hash, _ := auth.HashPassword(password)
		user.Username = username
		user.Password = string(hash)
		user.Name = username
		user.Email = email
		user.Phone = phone
		user.Level = "user"
		user.CreatedAt = &t
		user.CreatedBy = user.Username
		user.UpdatedAt = &t
		user.UpdatedBy = user.Username

		err = db.MySql.Save(&user).Error
		if err != nil {
			return utils.Render(c, components.ErrorAlert(err.Error(), "register"), templ.WithStatus(http.StatusBadRequest))
		}

		if err != nil {
			return utils.Render(c, components.ErrorAlert(err.Error(), "register"), templ.WithStatus(http.StatusBadRequest))
		}

		// c.Response().Header.Set("HX-Redirect", "/")
		// return c.SendStatus(fiber.StatusOK)
		return utils.Render(c, components.SuccessAlert("Register success!", "register"), templ.WithStatus(http.StatusOK))

	})

	app.Post("/login", func(c *fiber.Ctx) error {

		username := c.FormValue("username")
		var user models.User
		err := db.MySql.Where("username=? or email=?", username, username).First(&user).Error
		if err != nil {
			log.Println(err.Error())
			return utils.Render(c, components.ErrorAlert("Bad Credentials", "login"), templ.WithStatus(http.StatusBadRequest))
		}

		password := c.FormValue("password")
		if bcrypt.CompareHashAndPassword([]byte(strings.Trim(user.Password, " ")), []byte(password)) == nil {
			tokenString, err := auth.CreateToken(user)
			if err != nil {
				log.Println(err.Error())
				return utils.Render(c, components.ErrorAlert("Bad Credentials", "login"), templ.WithStatus(http.StatusBadRequest))
			}
			// fmt.Println(tokenString)
			c.Cookie(&fiber.Cookie{
				Name:     "session",
				Value:    tokenString,
				HTTPOnly: true,
				Secure:   true,
				SameSite: "Strict",
			})
			c.Response().Header.Set("HX-Redirect", "/")
			return c.SendStatus(fiber.StatusOK)

		} else {
			log.Println(bcrypt.CompareHashAndPassword([]byte(strings.Trim(user.Password, " ")), []byte(password)))
			return utils.Render(c, components.ErrorAlert("Invalid password", "login"), templ.WithStatus(http.StatusBadRequest))
		}

	})

	app.Get("/logout", func(c *fiber.Ctx) error {
		auth.ClearSession(c)
		c.Response().Header.Set("HX-Redirect", "/login")
		return c.Redirect("/")
	})

}
