package routes

import (
	"fmt"
	"lawas-go/auth"
	"lawas-go/components"
	"lawas-go/db"
	"lawas-go/dto"
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
	"gorm.io/gorm"
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

	app.Get("/401", func(c *fiber.Ctx) error {
		return utils.Render(c, pages.Forbidden(auth.IsAuthenticated(c)))
	})

	app.Get("/dashboard", auth.AssertAuthenticatedMiddleware, func(c *fiber.Ctx) error {
		return utils.Render(c, pages.Dashboard(auth.IsAuthenticated(c)))
	})

	app.Get("/add-remove-watchlist", auth.AssertAuthenticatedMiddleware, func(c *fiber.Ctx) error {
		var watchlist models.Watchlist
		var token dto.Token
		var count int64
		token, _ = auth.IsAuthenticated(c)
		fmt.Println(c.Query("item_id"))
		db.MySql.Where("user_id=? AND item_id=?", token.UserID, c.Query("item_id")).First(&watchlist).Count(&count)
		fmt.Println(count)
		if count > 0 {
			db.MySql.Where("user_id=? AND item_id=?", token.UserID, c.Query("item_id")).First(&watchlist)
			err := db.MySql.Delete(&watchlist).Error
			if err != nil {
				fmt.Println(err.Error())
			}
			return utils.Render(c, components.AddRemoveWatchlist(""))
		} else {
			watchlist.UserID = token.UserID
			watchlist.ItemID = c.Query("item_id")
			err := db.MySql.Save(&watchlist).Error
			if err != nil {
				fmt.Println(err.Error())
			}
			return utils.Render(c, components.AddRemoveWatchlist(watchlist.ItemID))
		}
	})

	app.Get("/my-sells", auth.AssertAuthenticatedMiddleware, func(c *fiber.Ctx) error {
		var sells []models.Item
		var token dto.Token
		var user models.User
		token, _ = auth.IsAuthenticated(c)
		db.MySql.First(&user, "username=?", token.Username)
		db.MySql.Where("user_id=?", user.ID).Preload("User").Preload("Bids").Preload("Category").Preload("Currency").Find(&sells)
		// for _, sell := range sells {
		// 	fmt.Println(sell.User.Username)
		// }
		return utils.Render(c, pages.TabSell(sells))
	})

	app.Get("/my-bids", auth.AssertAuthenticatedMiddleware, func(c *fiber.Ctx) error {
		var bids []models.Bid
		var token dto.Token
		var user models.User
		token, _ = auth.IsAuthenticated(c)
		db.MySql.First(&user, "username=?", token.Username)
		db.MySql.Where("user_id=?", user.ID).Preload("Item", func(db *gorm.DB) *gorm.DB {
			return db.Preload("Category").Preload("Bids").Preload("User").Preload("Currency")
		}).Preload("Watchlist", func(db *gorm.DB) *gorm.DB {
			return db.Where("user_id=?", user.ID)
		}).Find(&bids)
		// for _, sell := range bids {
		// 	fmt.Println(sell.User.Username)
		// }
		return utils.Render(c, pages.TabBid(bids))
	})

	app.Post("/register", func(c *fiber.Ctx) error {

		var users []models.User
		var newUser models.User
		var count int64
		username := c.FormValue("username")
		email := c.FormValue("email")
		phone := c.FormValue("phone")
		password := c.FormValue("password")
		err := db.MySql.Where("username=? or email=? or phone=?", username, email, phone).Find(&users).Count(&count).Error
		if err != nil {
			return utils.Render(c, components.ErrorAlert(err.Error(), "register"), templ.WithStatus(http.StatusBadRequest))
		}
		if count > 0 {
			msg := ""
			for _, user := range users {
				if strings.Trim(user.Username, " ") == strings.Trim(username, " ") {
					if msg != "" {
						msg = msg + ", "
					}
					msg = "Username"
				}
				if strings.Trim(user.Email, " ") == strings.Trim(email, " ") {
					if msg != "" {
						msg = msg + ", "
					}
					msg = msg + "Email"
				}
				// if strings.Trim(user.Phone, " ") == strings.Trim(phone, " ") {
				// 	if msg != "" {
				// 		msg = msg + ", "
				// 	}
				// 	msg = msg + "Phone"
				// }
			}

			return utils.Render(c, components.ErrorAlert(msg+" already exist!", "register"), templ.WithStatus(http.StatusBadRequest))
		}

		if len(strings.Trim(password, " ")) < 6 {
			return utils.Render(c, components.ErrorAlert("Password must be at least 6 characters!", "register"), templ.WithStatus(http.StatusBadRequest))
		}
		t := time.Now()
		hash, _ := auth.HashPassword(password)
		newUser.Username = username
		newUser.Password = string(hash)
		newUser.Name = username
		newUser.Email = email
		// newUser.Phone = phone
		newUser.Level = "user"
		newUser.CreatedAt = &t
		newUser.CreatedBy = newUser.Username
		newUser.UpdatedAt = &t
		newUser.UpdatedBy = newUser.Username

		err = db.MySql.Save(&newUser).Error
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
