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
	"strconv"
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

	app.Get("/contact", func(c *fiber.Ctx) error {
		return utils.Render(c, pages.Contact(auth.IsAuthenticated(c)))
	})

	app.Get("/about", func(c *fiber.Ctx) error {
		return utils.Render(c, pages.About(auth.IsAuthenticated(c)))
	})

	app.Get("/faq", func(c *fiber.Ctx) error {
		return utils.Render(c, pages.Faq(auth.IsAuthenticated(c)))
	})

	app.Get("/item", func(c *fiber.Ctx) error {
		var item models.Item
		var token dto.Token
		token, _ = auth.GetToken(c)

		db.MySql.Where("id=?", c.Query("id")).First(&item)

		return utils.Render(c, pages.Item(item, token, token.IsAuth))
	})

	app.Get("/dashboard", auth.AssertAuthenticatedMiddleware, func(c *fiber.Ctx) error {
		return utils.Render(c, pages.Dashboard(auth.IsAuthenticated(c)))
	})

	app.Get("/add-remove-watchlist", auth.AssertAuthenticatedMiddleware, func(c *fiber.Ctx) error {
		var watchlist models.Watchlist
		var token dto.Token
		var count int64
		token, _ = auth.IsAuthenticated(c)
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

	app.Get("/my-address", auth.AssertAuthenticatedMiddleware, func(c *fiber.Ctx) error {
		var address models.Address
		var token dto.Token
		tipe := c.Query("tipe")
		token, _ = auth.IsAuthenticated(c)
		db.MySql.Where("user_id=?", token.UserID).First(&address)
		return utils.Render(c, pages.TabAddress(tipe, address))
	})

	app.Post("/save-address", auth.AssertAuthenticatedMiddleware, func(c *fiber.Ctx) error {
		var address models.Address
		var token dto.Token
		var p models.Address
		tipe := c.Query("tipe")
		token, _ = auth.IsAuthenticated(c)
		db.MySql.Where("user_id=?", token.UserID).First(&address)
		// return utils.Render(c, components.ErrorAlert("errr", tipe), templ.WithStatus(http.StatusBadRequest))
		if err := c.BodyParser(&p); err != nil {
			log.Println(err.Error())
			return utils.Render(c, components.ErrorAlert(err.Error(), tipe), templ.WithStatus(http.StatusBadRequest))
		}

		if tipe == "billing" {
			address.BillName = p.BillName
			address.BillPhone = p.BillPhone
			address.BillAddress = p.BillAddress
			address.BillCity = p.BillCity
			address.BillPostalCode = p.BillPostalCode
			address.BillCountry = p.BillCountry
		} else {
			address.ShipAddress = p.ShipName
			address.ShipPhone = p.ShipPhone
			address.ShipAddress = p.ShipAddress
			address.ShipCity = p.ShipCity
			address.ShipPostalCode = p.ShipPostalCode
			address.ShipCountry = p.ShipCountry
		}
		err := db.MySql.Save(&address).Error
		if err != nil {
			log.Println(err.Error())
			return utils.Render(c, components.ErrorAlert(err.Error(), tipe), templ.WithStatus(http.StatusBadRequest))
		}
		return utils.Render(c, components.SuccessAlert("Update success!", tipe), templ.WithStatus(http.StatusOK))
	})

	app.Get("/my-account", auth.AssertAuthenticatedMiddleware, func(c *fiber.Ctx) error {
		var user models.User
		var token dto.Token
		token, _ = auth.IsAuthenticated(c)
		err := db.MySql.Where("id=?", token.UserID).First(&user).Error
		if err != nil {
			log.Println(err.Error())
			return utils.Render(c, components.ErrorAlert(err.Error(), "account"), templ.WithStatus(http.StatusBadRequest))
		}
		return utils.Render(c, pages.TabAccount(user))
	})

	app.Post("/save-account", auth.AssertAuthenticatedMiddleware, func(c *fiber.Ctx) error {
		var user models.User
		type Alias models.User
		var token dto.Token
		token, _ = auth.IsAuthenticated(c)
		var userForm = struct {
			Alias
			CurrenPassword  string `form:"current_password"`
			NewPassword     string `form:"new_password"`
			ConfirmPassword string `form:"confirm_password"`
		}{}
		err := db.MySql.Where("id=?", token.UserID).First(&user).Error
		if err != nil {
			log.Println(err.Error())
			return utils.Render(c, components.ErrorAlert(err.Error(), "account"), templ.WithStatus(http.StatusBadRequest))
		}
		if err := c.BodyParser(&userForm); err != nil {
			log.Println(err.Error())
			return utils.Render(c, components.ErrorAlert(err.Error(), "account"), templ.WithStatus(http.StatusBadRequest))
		}

		if bcrypt.CompareHashAndPassword([]byte(strings.Trim(user.Password, " ")), []byte(userForm.CurrenPassword)) != nil {
			return utils.Render(c, components.ErrorAlert("Invalid password", "account"), templ.WithStatus(http.StatusBadRequest))
		}

		if userForm.NewPassword != userForm.ConfirmPassword {
			return utils.Render(c, components.ErrorAlert("New password not matched", "account"), templ.WithStatus(http.StatusBadRequest))
		}

		if strings.Trim(userForm.NewPassword, " ") != "" {
			hash, _ := auth.HashPassword(userForm.NewPassword)
			user.Password = string(hash)
		}

		user.Name = userForm.Name
		user.Email = userForm.Email

		err = db.MySql.Save(&user).Error
		if err != nil {
			log.Println(err.Error())
			return utils.Render(c, components.ErrorAlert(err.Error(), "account"), templ.WithStatus(http.StatusBadRequest))
		}
		return utils.Render(c, components.SuccessAlert("Update success!", "account"), templ.WithStatus(http.StatusOK))
	})

	app.Get("/my-sells", auth.AssertAuthenticatedMiddleware, func(c *fiber.Ctx) error {
		var sells []models.Item
		var token dto.Token
		var count int64 = 0

		pageStr := c.Query("page")
		if pageStr == "" {
			pageStr = "1"
		}

		page, err := strconv.Atoi(pageStr)
		if err != nil {
			page = 1
		}

		pageSize := c.Query("size")
		if pageSize == "" {
			pageSize = "5"
		}

		size, err := strconv.Atoi(pageSize)
		if err != nil {
			size = 5
		}
		token, _ = auth.IsAuthenticated(c)

		db.MySql.Where("user_id=?", token.UserID).Preload("User").Preload("Bids").Preload("Category").Preload("Currency").Find(&sells).
			Count(&count).Offset((page - 1) * size).Limit(size).Find(&sells)
		// for _, sell := range sells {
		// 	fmt.Println(sell.User.Username)
		// }
		return utils.Render(c, pages.TabSell(page, size, count, sells))
	})

	app.Get("/my-watchs", auth.AssertAuthenticatedMiddleware, func(c *fiber.Ctx) error {
		var watchlist []models.Watchlist
		var token dto.Token
		var count int64 = 0

		pageStr := c.Query("page")
		if pageStr == "" {
			pageStr = "1"
		}

		page, err := strconv.Atoi(pageStr)
		if err != nil {
			page = 1
		}

		pageSize := c.Query("size")
		if pageSize == "" {
			pageSize = "5"
		}

		size, err := strconv.Atoi(pageSize)
		if err != nil {
			size = 5
		}
		token, _ = auth.IsAuthenticated(c)

		db.MySql.Where("user_id=?", token.UserID).Preload("Item", func(db *gorm.DB) *gorm.DB {
			return db.Preload("Category").Preload("Bids").Preload("User").Preload("Currency")
		}).Find(&watchlist).Count(&count).Offset((page - 1) * size).Limit(size).Find(&watchlist)
		// for _, sell := range sells {
		// 	fmt.Println(sell.User.Username)
		// }
		return utils.Render(c, pages.TabWatch(page, size, count, watchlist))
	})

	app.Get("/my-bids", auth.AssertAuthenticatedMiddleware, func(c *fiber.Ctx) error {
		var bids []models.Bid
		var token dto.Token
		var count int64 = 0

		pageStr := c.Query("page")
		if pageStr == "" {
			pageStr = "1"
		}

		page, err := strconv.Atoi(pageStr)
		if err != nil {
			page = 1
		}

		pageSize := c.Query("size")
		if pageSize == "" {
			pageSize = "5"
		}

		size, err := strconv.Atoi(pageSize)
		if err != nil {
			size = 5
		}

		token, _ = auth.IsAuthenticated(c)

		db.MySql.Where("user_id=?", token.UserID).Preload("Item", func(db *gorm.DB) *gorm.DB {
			return db.Preload("Category").Preload("Bids").Preload("User").Preload("Currency")
		}).Preload("Watchlist", func(db *gorm.DB) *gorm.DB {
			return db.Where("user_id=?", token.UserID)
		}).Find(&bids).Count(&count).Offset((page - 1) * size).Limit(size).Find(&bids)
		// for _, sell := range bids {
		// 	fmt.Println(sell.User.Username)
		// }
		return utils.Render(c, pages.TabBid(page, size, count, bids))
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
