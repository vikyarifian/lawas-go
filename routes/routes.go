package routes

import (
	"crypto/rand"
	"encoding/base64"
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
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func itemList(token dto.Token, page int, size int, sort string, count int64, category string, search string) ([]models.Item, int64, error) {
	var items []models.Item
	err := db.MySql.Table("items").Select("items.*", "watchlists.*", "bids.*", "categories.Name as CategoryName").
		Joins("LEFT JOIN (SELECT Count(*) count1, item_id FROM watchlists GROUP BY item_id) as watchlists ON watchlists.item_id=items.id").
		Joins("LEFT JOIN (SELECT Count(*) count2, item_id FROM bids GROUP BY item_id) as bids ON bids.item_id=items.id").
		Joins("LEFT JOIN categories ON categories.id=items.category_id").
		Where(fmt.Sprintf("CONCAT(items.name,items.description,IFNULL(categories.name,'')) LIKE '%s%s%s' ", "%", search, "%")).
		Where(fmt.Sprintf(" CASE WHEN '%s'!='' THEN IFNULL(categories.name,'') ELSE '' END LIKE '%s%s%s' ", category, "%", category, "%")).
		Preload("User").Preload("Category").Preload("Currency").Preload("Bids", func(db *gorm.DB) *gorm.DB {
		return db.Order("Bid Desc").Preload("Watchlist").Preload("User")
	}).Preload("Watchlists", func(db *gorm.DB) *gorm.DB {
		return db.Where("user_id=?", token.UserID)
	}).Order(sort).Find(&items).Count(&count).Offset((page - 1) * size).Limit(size).
		Find(&items).Error

	return items, count, err
}

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
		token, _ = auth.IsAuthenticated(c)

		err := db.MySql.Where("id=?", c.Query("id")).Preload("User").Preload("Category").Preload("Currency").Preload("Bids", func(db *gorm.DB) *gorm.DB {
			return db.Order("Bid Desc, Date").Preload("Watchlist").Preload("User")
		}).Preload("Watchlists", func(db *gorm.DB) *gorm.DB {
			return db.Where("user_id=?", token.UserID)
		}).First(&item).Error
		if err != nil {
			return utils.Render(c, pages.NotFound(auth.IsAuthenticated(c)))
		}

		return utils.Render(c, pages.Item(item, token, token.IsAuth))
	})

	app.Get("/market", func(c *fiber.Ctx) error {
		var categories []models.Category
		db.MySql.Order("name").Find(&categories)
		var token dto.Token
		token, _ = auth.IsAuthenticated(c)
		return utils.Render(c, pages.Market(categories, token, token.IsAuth))
	})

	app.Get("/items", func(c *fiber.Ctx) error {
		var items []models.Item
		var token dto.Token
		token, _ = auth.IsAuthenticated(c)

		search := c.Query("search")
		category := c.Query("category")
		// if search != "" {
		// 	time.Sleep(50 * time.Millisecond)
		// }
		sort := c.Query("sortby")
		if sort == "" {
			sort = "date"
		}
		sortBy := sort
		switch strings.ToLower(sort) {
		case "top":
			sortBy = "watchlists.count1 desc, bids.count2 desc"
		case "bid":
			sortBy = "bids.count2 desc"
		default:
			sortBy = sort
		}

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
			pageSize = "12"
		}

		size, err := strconv.Atoi(pageSize)
		if err != nil {
			size = 12
		}

		var count int64 = 0

		items, count, err = itemList(token, page, size, sortBy, count, category, search)

		if err != nil || count == 0 {
			page = 1
			items, count, err = itemList(token, page, size, sortBy, count, category, search)
		}

		return utils.Render(c, pages.ItemList(items, page, size, sort, count, category, search, token, token.IsAuth))
	})

	app.Get("/dashboard", auth.AssertAuthenticatedMiddleware, func(c *fiber.Ctx) error {
		return utils.Render(c, pages.Dashboard(auth.IsAuthenticated(c)))
	})

	app.Get("/sell", auth.AssertAuthenticatedMiddleware, func(c *fiber.Ctx) error {
		var categories []models.Category
		var currencies []models.Currency
		var token dto.Token
		db.MySql.Find(&categories)
		db.MySql.Find(&currencies)
		token, _ = auth.IsAuthenticated(c)
		return utils.Render(c, components.HtmlSell(token, categories, currencies))
	})

	app.Post("/bid", auth.AssertAuthenticatedMiddleware, func(c *fiber.Ctx) error {
		var token dto.Token
		var bid models.Bid
		var item models.Item

		token, _ = auth.IsAuthenticated(c)
		if err := c.BodyParser(&bid); err != nil {
			log.Println(err.Error())
			return utils.Render(c, components.ErrorAlert(err.Error(), "sell"), templ.WithStatus(http.StatusBadRequest))
		}
		if err := db.MySql.Where("id=?", bid.ItemID).First(&item).Error; err != nil {
			return utils.Render(c, components.ErrorAlert("Item not found!", "bid"), templ.WithStatus(http.StatusBadRequest))
		}

		maxId := db.MySql.Table("bids").Select("max(no)").Row()
		_ = maxId.Scan(&bid.No)

		idhash := utils.GetMD5Hash(strconv.Itoa(bid.No + 1))

		newBid := models.Bid{
			No:        bid.No + 1,
			ID:        idhash,
			ItemID:    bid.ItemID,
			UserID:    token.UserID,
			Bid:       bid.Bid,
			Date:      func(t time.Time) *time.Time { return &t }(time.Now()),
			CreatedBy: token.Username,
			UpdatedBy: token.Username,
		}

		if err := db.MySql.Save(&newBid).Error; err != nil {
			log.Println(err.Error())
			return utils.Render(c, components.ErrorAlert(err.Error(), "bid"), templ.WithStatus(http.StatusBadRequest))
		}

		time.Sleep(100 * time.Millisecond)
		c.Response().Header.Set("HX-Redirect", "/item?id="+item.ID)
		return utils.Render(c, components.SuccessAlert("Success!", "bid"), templ.WithStatus(http.StatusOK))

	})

	app.Get("/collection", auth.AssertAuthenticatedMiddleware, func(c *fiber.Ctx) error {
		var token dto.Token
		var bids []models.Bid
		token, _ = auth.IsAuthenticated(c)
		db.MySql.Where("user_id=?", token.UserID).Preload("Item", func(db *gorm.DB) *gorm.DB {
			return db.Preload("Category").Preload("Bids", func(db *gorm.DB) *gorm.DB {
				return db.Order("bid desc").Preload("User")
			}).Preload("User").Preload("Currency")
		}).Preload("Watchlist", func(db *gorm.DB) *gorm.DB {
			return db.Where("user_id=?", token.UserID)
		}).Preload("Cart", func(db *gorm.DB) *gorm.DB { return db.Preload("Payment") }).Find(&bids)

		return utils.Render(c, pages.Collection(bids, token, token.IsAuth))
	})

	app.Get("/offers", auth.AssertAuthenticatedMiddleware, func(c *fiber.Ctx) error {
		var token dto.Token
		var items []models.Item
		token, _ = auth.IsAuthenticated(c)
		db.MySql.Where("user_id=?", token.UserID).Preload("Category").Preload("Bids", func(db *gorm.DB) *gorm.DB {
			return db.Order("bid desc").Preload("User").Preload("Cart", func(db *gorm.DB) *gorm.DB { return db.Preload("Payment") })
		}).Preload("User").Preload("Currency").Find(&items)
		for i, item := range items {
			if len(item.Bids) == 0 {
				items[i].Bids = append(items[i].Bids, models.Bid{
					Bid: 0,
				})
			}
		}
		return utils.Render(c, pages.Offers(items, token, token.IsAuth))

	})

	app.Get("/notif-cart", auth.AssertAuthenticatedMiddleware, func(c *fiber.Ctx) error {
		var token dto.Token
		var carts []models.Cart
		token, _ = auth.IsAuthenticated(c)

		db.MySql.Table("carts").Select("carts.*", "bids.user_id").Joins("LEFT JOIN bids ON bids.id COLLATE utf8mb4_unicode_ci = carts.bid_id").
			Where("bids.user_id=? and carts.id not in (select cart_id from db_lawas.payments) ", token.UserID).Preload("Bid", func(db *gorm.DB) *gorm.DB {
			return db.Preload("Item", func(db *gorm.DB) *gorm.DB {
				return db.Preload("Currency")
			})
		}).Find(&carts)
		return utils.Render(c, components.NotifCart(carts, token))
	})

	app.Post("/checkout", auth.AssertAuthenticatedMiddleware, func(c *fiber.Ctx) error {
		var token dto.Token
		var payment models.Payment
		token, _ = auth.IsAuthenticated(c)
		if err := c.BodyParser(&payment); err != nil {
			log.Println(err.Error())
			return utils.Render(c, components.ErrorAlert(err.Error(), "checkout"), templ.WithStatus(http.StatusBadRequest))
		}

		b := make([]byte, 5)
		rand.Read(b)
		reff := base64.StdEncoding.EncodeToString(b)

		maxId := db.MySql.Table("payments").Select("max(no)").Row()
		_ = maxId.Scan(&payment.No)

		idhash := utils.GetMD5Hash(strconv.Itoa(payment.No + 1))

		newPayment := models.Payment{
			No:             payment.No + 1,
			ID:             idhash,
			Reff:           reff,
			CartID:         payment.CartID,
			ShipName:       payment.ShipName,
			ShipAddress:    payment.ShipAddress,
			ShipCity:       payment.ShipCity,
			ShipCountry:    payment.ShipCountry,
			ShipPostalCode: payment.ShipPostalCode,
			ShipPhone:      payment.ShipPhone,
			Notes:          payment.Notes,
			Status:         "O",
			CreatedBy:      token.Username,
			UpdatedBy:      token.Username,
		}

		if err := db.MySql.Save(&newPayment).Error; err != nil {
			return utils.Render(c, components.ErrorAlert(err.Error(), "checkout"), templ.WithStatus(http.StatusBadRequest))
		}

		c.Response().Header.Set("HX-Redirect", "/payment?id="+newPayment.ID)
		return utils.Render(c, components.SuccessAlert("Success!", "checkout"), templ.WithStatus(http.StatusBadRequest))

	})

	app.Get("/payment", auth.AssertAuthenticatedMiddleware, func(c *fiber.Ctx) error {
		var token dto.Token
		var payment models.Payment
		id := c.Query("id")
		token, _ = auth.IsAuthenticated(c)
		if err := db.MySql.Where("id=? and status='O'", id).Preload("Cart", func(db *gorm.DB) *gorm.DB {
			return db.Preload("Bid", func(db *gorm.DB) *gorm.DB {
				return db.Preload("Item", func(db *gorm.DB) *gorm.DB { return db.Preload("Currency") })
			})
		}).First(&payment).Error; err != nil {
			c.Response().Header.Set("HX-Redirect", "/404")
		}

		return utils.Render(c, pages.Payment(payment, token, token.IsAuth))
	})

	app.Get("checkout", auth.AssertAuthenticatedMiddleware, func(c *fiber.Ctx) error {
		cartID := c.Query("cart_id")
		var token dto.Token
		var cart models.Cart
		token, _ = auth.IsAuthenticated(c)
		if err := db.MySql.Where("id=? and id not in (select cart_id from db_lawas.payments)", cartID).Preload("Bid", func(db *gorm.DB) *gorm.DB {
			return db.Preload("Item", func(db *gorm.DB) *gorm.DB {
				return db.Preload("Currency")
			}).Preload("User", func(db *gorm.DB) *gorm.DB {
				return db.Preload("Address")
			})
		}).First(&cart).Error; err != nil {
			c.Response().Header.Set("HX-Redirect", "/404")
		}
		if token.UserID != cart.Bid.UserID {
			c.Response().Header.Set("HX-Redirect", "/401")
		}
		if cart.BidID == "" {
			c.Response().Header.Set("HX-Redirect", "/404")
		}
		if cart.Bid.User.Address.ID == "" {
			cart.Bid.User.Address = models.Address{}
		}
		return utils.Render(c, pages.Checkout(cart, cart.Bid.User.Address, token, token.IsAuth))

	})

	app.Get("/approve-bid", auth.AssertAuthenticatedMiddleware, func(c *fiber.Ctx) error {
		var token dto.Token
		var bids []models.Bid
		bidNo := c.Query("bid_no")
		bidID := c.Query("bid_id")
		userID := c.Query("user_id")
		itemID := c.Query("item_id")
		token, _ = auth.IsAuthenticated(c)
		if userID == token.UserID {
			return utils.Render(c, components.ErrorAlert("Approve failed!", "approve_bid_"+bidNo), templ.WithStatus(http.StatusBadRequest))
		}
		db.MySql.Where("item_id=?", itemID).Order("bid desc").Find(&bids)
		if bidID != bids[0].ID {
			fmt.Println(bids[0].ID)
			return utils.Render(c, components.ErrorAlert("Approve failed!", "approve_bid_"+bidNo), templ.WithStatus(http.StatusBadRequest))
		}

		no := 0
		maxId := db.MySql.Table("carts").Select("max(no)").Row()
		_ = maxId.Scan(&no)

		idhash := utils.GetMD5Hash(strconv.Itoa(no + 1))

		cart := models.Cart{
			No:        no + 1,
			ID:        idhash,
			BidID:     bidID,
			Status:    "O",
			CreatedBy: token.Username,
			UpdatedBy: token.Username,
		}
		if err := db.MySql.Save(&cart).Error; err != nil {
			fmt.Println(err.Error())
			return utils.Render(c, components.ErrorAlert("Approve failed!", "approve_bid_"+bidNo), templ.WithStatus(http.StatusBadRequest))
		}
		c.Response().Header.Set("HX-Redirect", "/offers")
		return utils.Render(c, components.SuccessAlert("Success!", "approve_bid_"+bidNo), templ.WithStatus(http.StatusOK))
	})

	app.Get("/approve-payment", auth.AssertAuthenticatedMiddleware, func(c *fiber.Ctx) error {
		var token dto.Token
		var payment models.Payment
		payNo := c.Query("pay_no")
		payID := c.Query("pay_id")
		userID := c.Query("user_id")
		// itemID := c.Query("item_id")
		token, _ = auth.IsAuthenticated(c)
		if userID != token.UserID {
			return utils.Render(c, components.ErrorAlert("Approve failed!", "approve_pay_"+payNo), templ.WithStatus(http.StatusBadRequest))
		}
		db.MySql.Where("id=?", payID).First(&payment)
		if payID != payment.ID {
			fmt.Println(payment.ID)
			return utils.Render(c, components.ErrorAlert("Approve failed!", "approve_pay_"+payNo), templ.WithStatus(http.StatusBadRequest))
		}
		payment.Status = "A"
		if err := db.MySql.Save(&payment).Error; err != nil {
			fmt.Println(err.Error())
			return utils.Render(c, components.ErrorAlert("Approve failed!", "approve_pay_"+payNo), templ.WithStatus(http.StatusBadRequest))
		}
		c.Response().Header.Set("HX-Redirect", "/offers")
		return utils.Render(c, components.SuccessAlert("Success!", "approve_pay_"+payNo), templ.WithStatus(http.StatusOK))
	})

	app.Post("/sell", auth.AssertAuthenticatedMiddleware, func(c *fiber.Ctx) error {
		var token dto.Token
		var item models.Item
		token, _ = auth.IsAuthenticated(c)
		if err := c.BodyParser(&item); err != nil {
			log.Println(err.Error())
			return utils.Render(c, components.ErrorAlert(err.Error(), "sell"), templ.WithStatus(http.StatusBadRequest))
		}
		file, err := c.FormFile("photo")
		if err != nil {
			return utils.Render(c, components.ErrorAlert(err.Error(), "sell"), templ.WithStatus(http.StatusBadRequest))
		}

		maxId := db.MySql.Table("items").Select("max(no)").Row()
		_ = maxId.Scan(&item.No)

		idhash := utils.GetMD5Hash(strconv.Itoa(item.No + 1))
		ext := strings.Split(file.Filename, ".")
		file.Filename = idhash + "." + ext[len(ext)-1]

		destination := "assets/images/products/"
		if err := c.SaveFile(file, destination+file.Filename); err != nil {
			return utils.Render(c, components.ErrorAlert(err.Error(), "sell"), templ.WithStatus(http.StatusBadRequest))
		}
		re := regexp.MustCompile("\\n")
		item.Description = re.ReplaceAllString(item.Description, "<br>")
		fmt.Println("Desc: " + item.Description)

		newItem := models.Item{
			No:          item.No + 1,
			ID:          idhash,
			UserID:      token.UserID,
			Name:        item.Name,
			Description: item.Description,
			CategoryID:  item.CategoryID,
			Brand:       item.Brand,
			Condition:   item.Condition,
			Duration:    item.Duration,
			CurrencyID:  item.CurrencyID,
			OpenBid:     item.OpenBid,
			Photo:       destination + file.Filename,
			Date:        func(t time.Time) *time.Time { return &t }(time.Now()),
			CreatedBy:   token.Username,
			UpdatedBy:   token.Username,
		}
		print(item.CategoryID + " " + item.CurrencyID)
		// db.MySql.First(&newItem.Category)
		// db.MySql.First(&newItem.Currency)

		if err := db.MySql.Save(&newItem).Error; err != nil {
			log.Println(err.Error())
			return utils.Render(c, components.ErrorAlert(err.Error(), "sell"), templ.WithStatus(http.StatusBadRequest))
		}

		time.Sleep(500 * time.Millisecond)
		c.Response().Header.Set("HX-Redirect", "/item?id="+newItem.ID)
		return utils.Render(c, components.SuccessAlert("Success!", "sell"), templ.WithStatus(http.StatusOK))
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
			return utils.Render(c, components.AddRemoveWatchlist("", token.IsAuth))
		} else {
			no := 0
			maxId := db.MySql.Table("watchlists").Select("max(no)").Row()
			_ = maxId.Scan(&no)
			idhash := utils.GetMD5Hash(strconv.Itoa(no + 1))
			watchlist.UserID = token.UserID
			watchlist.ItemID = c.Query("item_id")
			watchlist.No = no + 1
			watchlist.ID = idhash
			err := db.MySql.Save(&watchlist).Error
			if err != nil {
				fmt.Println(err.Error())
			}
			return utils.Render(c, components.AddRemoveWatchlist(watchlist.ItemID, token.IsAuth))
		}
	})

	app.Get("/my-address", auth.AssertAuthenticatedMiddleware, func(c *fiber.Ctx) error {
		var address models.Address
		var token dto.Token
		tipe := c.Query("tipe")
		token, _ = auth.IsAuthenticated(c)
		if err := db.MySql.Where("user_id=?", token.UserID).First(&address).Error; err != nil {
			address = models.Address{}
		}
		return utils.Render(c, pages.TabAddress(tipe, address))
	})

	app.Post("/save-address", auth.AssertAuthenticatedMiddleware, func(c *fiber.Ctx) error {
		var address models.Address
		var token dto.Token
		var p models.Address
		tipe := c.Query("tipe")
		token, _ = auth.IsAuthenticated(c)
		if err := db.MySql.Where("user_id=?", token.UserID).First(&address); err != nil {
			no := 0
			maxId := db.MySql.Table("addresses").Select("max(no)").Row()
			_ = maxId.Scan(&no)
			idhash := utils.GetMD5Hash(strconv.Itoa(no + 1))
			address = models.Address{}
			address.UserID = token.UserID
			address.No = no + 1
			address.ID = idhash
		}
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
		address.UpdatedAt = &time.Time{}
		address.UpdatedBy = token.Username
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

		no := 0
		maxId := db.MySql.Table("users").Select("max(no)").Row()
		_ = maxId.Scan(&no)
		idhash := utils.GetMD5Hash(strconv.Itoa(no + 1))

		t := time.Now()
		hash, _ := auth.HashPassword(password)
		newUser.Username = username
		newUser.Password = string(hash)
		newUser.Name = username
		newUser.Email = email
		// newUser.Phone = phone
		newUser.No = no + 1
		newUser.ID = idhash
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

	app.Use(func(c *fiber.Ctx) error {
		return utils.Render(c, pages.NotFound(auth.IsAuthenticated(c)))
	})

}
