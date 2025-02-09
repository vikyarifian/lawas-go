package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"lawas-go/auth"
	"lawas-go/config"
	"lawas-go/db"
	"lawas-go/models"
	"lawas-go/routes"
	"log"
	"math/rand/v2"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	config.LoadEnv()
	db.ConnectDB()

	app := fiber.New(fiber.Config{
		Network: fiber.NetworkTCP,
	})

	app.Use(logger.New())

	app.Static("/assets", "./assets")
	app.Static("/lib", "./lib")
	app.Static("/popup", "./popup")

	app.Use(func(c *fiber.Ctx) error {
		c.Set("Cache-Control", "no-store, no-cache, must-revalidate, proxy-revalidate")
		c.Set("Pragma", "no-cache")
		c.Set("Expires", "0")
		c.Set("Surrogate-Control", "no-store")
		return c.Next()
	})

	routes.ApiRoutes(app)
	routes.WebRoutes(app)
	listen, _ := net.Listen("tcp", ":7632")
	// BulkUser()
	// BulProduct()
	log.Fatal(app.Listener(listen))

}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func BulProduct() {
	var res BulkProductResponse
	httpClient := http.Client{}
	url := "https://dummyjson.com/products?limit=194"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Content-type", "application/json")
	response, err := httpClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(responseBody, &res)

	var n1 = 0
	var n2 = 0
	var bNo = 1
	var users []models.User
	var categories []models.Category
	db.MySql.Find(&users)
	db.MySql.Find(&categories)

	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()

	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	// lastOfMonth := firstOfMonth.AddDate(0, 1, -1)

	// firstOfMonth.AddDate(0, 0, 1)
	// fmt.Println(firstOfMonth)
	// fmt.Println(lastOfMonth)
	// fmt.Println(firstOfMonth.AddDate(0, 0, 1))

	duration := []int{3, 5, 7, 10}
	for i, p := range res.Products {
		if i <= 60 {
			if p.Category != "groceries" && p.Category != "skin-care" && p.Category != "beauty" && p.Category != "womens-jewellery" {
				idhash := GetMD5Hash(strconv.Itoa(p.ID))
				cate := p.Category
				switch {
				case p.Category == "fragrances":
					cate = "a87ff679a2f3e71d9181a67b7542122c"
				case p.Category == "furniture":
					cate = "e4da3b7fbbce2345d7772b0674a318d5"
				case p.Category == "home-decoration":
					cate = "e4da3b7fbbce2345d7772b0674a318d5"
				case p.Category == "kitchen-accessories":
					cate = "e4da3b7fbbce2345d7772b0674a318d5"
				case p.Category == "laptops":
					cate = "c4ca4238a0b923820dcc509a6f75849b"
				case p.Category == "mens-shirts":
					cate = "a87ff679a2f3e71d9181a67b7542122c"
				case p.Category == "mens-shoes":
					cate = "a87ff679a2f3e71d9181a67b7542122c"
				case p.Category == "mens-watches":
					cate = "a87ff679a2f3e71d9181a67b7542122c"
				case p.Category == "mobile-accessories":
					cate = "c4ca4238a0b923820dcc509a6f75849b"
				case p.Category == "motorcycle":
					cate = "c81e728d9d4c2f636f067f89cc14862c"
				case p.Category == "smartphones":
					cate = "c4ca4238a0b923820dcc509a6f75849b"
				case p.Category == "sports-accessories":
					cate = "eccbc87e4b5ce2fe28308fd9f2a7baf3"
				case p.Category == "sunglasses":
					cate = "a87ff679a2f3e71d9181a67b7542122c"
				case p.Category == "tablets":
					cate = "c4ca4238a0b923820dcc509a6f75849b"
				case p.Category == "tops":
					cate = "a87ff679a2f3e71d9181a67b7542122c"
				case p.Category == "vehicle":
					cate = "c81e728d9d4c2f636f067f89cc14862c"
				case p.Category == "womens-bags":
					cate = "a87ff679a2f3e71d9181a67b7542122c"
				case p.Category == "womens-dresses":
					cate = "a87ff679a2f3e71d9181a67b7542122c"
				case p.Category == "womens-shoes":
					cate = "a87ff679a2f3e71d9181a67b7542122c"
				case p.Category == "womens-watches":
					cate = "a87ff679a2f3e71d9181a67b7542122c"
				}
				cond := rand.IntN(2)
				if cond == 0 {
					cond = 1
				}
				var product = models.Item{
					No:          p.ID,
					ID:          idhash,
					UserID:      users[n1].ID,
					Name:        p.Title,
					Description: p.Description,
					Brand:       p.Brand,
					OpenBid:     p.Price,
					Photo:       p.Images[0],
					CategoryID:  cate,
					CurrencyID:  "c81e728d9d4c2f636f067f89cc14862c",
					CreatedBy:   users[n1].Username,
					Condition:   cond,
					Format:      "A",
					Date:        func(t time.Time) *time.Time { return &t }(firstOfMonth.AddDate(0, 0, i+1)),
					Duration:    duration[n2],
					CreatedAt:   func(t time.Time) *time.Time { return &t }(firstOfMonth.AddDate(0, 0, i+1)),
					UpdatedBy:   users[n1].Username,
					UpdatedAt:   func(t time.Time) *time.Time { return &t }(firstOfMonth.AddDate(0, 0, i+1)),
				}
				err = db.MySql.Save(&product).Error
				if err != nil {
					fmt.Println(err)
				}

				for x := 0; x <= duration[n2]+duration[0]; x++ {
					var bidders models.User
					db.MySql.Where("no=?", rand.IntN(50)).First(&bidders)
					bidhash := GetMD5Hash(strconv.Itoa(bNo))
					var bid = models.Bid{
						No:        bNo,
						ID:        bidhash,
						UserID:    bidders.ID,
						ItemID:    product.ID,
						Bid:       p.Price + 0.2 + float64(x),
						Date:      func(t time.Time) *time.Time { return &t }(firstOfMonth.AddDate(0, 0, i+1+x)),
						CreatedBy: bidders.Username,
						UpdatedBy: bidders.Username,
						CreatedAt: func(t time.Time) *time.Time { return &t }(firstOfMonth.AddDate(0, 0, i+1+x)),
						UpdatedAt: func(t time.Time) *time.Time { return &t }(firstOfMonth.AddDate(0, 0, i+1+x)),
					}
					var watch = models.Watchlist{
						UserID:    bidders.ID,
						ItemID:    product.ID,
						Date:      func(t time.Time) *time.Time { return &t }(firstOfMonth.AddDate(0, 0, i+1+x)),
						CreatedBy: bidders.Username,
						UpdatedBy: bidders.Username,
						CreatedAt: func(t time.Time) *time.Time { return &t }(firstOfMonth.AddDate(0, 0, i+1+x)),
						UpdatedAt: func(t time.Time) *time.Time { return &t }(firstOfMonth.AddDate(0, 0, i+1+x)),
					}

					bNo++
					if bidders.No != 0 {
						db.MySql.Save(&bid)
						var count int64
						db.MySql.Model(&models.Watchlist{}).Where("user_id=? and item_id=?", bidders.ID, product.ID).Count(&count)
						if count == 0 && x == duration[n2] {
							db.MySql.Save(&watch)
						}
						if bidders.No < 49 {
							var watch2 = models.Watchlist{
								UserID:    users[bidders.No+1].ID,
								ItemID:    product.ID,
								Date:      func(t time.Time) *time.Time { return &t }(firstOfMonth.AddDate(0, 0, i+1+x)),
								CreatedBy: bidders.Username,
								UpdatedBy: bidders.Username,
								CreatedAt: func(t time.Time) *time.Time { return &t }(firstOfMonth.AddDate(0, 0, i+1+x)),
								UpdatedAt: func(t time.Time) *time.Time { return &t }(firstOfMonth.AddDate(0, 0, i+1+x)),
							}
							db.MySql.Model(&models.Watchlist{}).Where("user_id=? and item_id=?", users[bidders.No+1].ID, product.ID).Count(&count)
							if count == 0 {
								db.MySql.Save(&watch2)
							}
						}
					}
				}

				n1 = n1 + duration[n2]
				n2++
				if n2 > 3 {
					n2 = 0
				}

				if n1 > 49 {
					n1 = 0
				}

			}
		}
	}
}

func BulkUser() {
	var res BulkUserResponse
	httpClient := http.Client{}
	url := "https://dummyjson.com/users?limit=50"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Content-type", "application/json")
	response, err := httpClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(string(responseBody))
	json.Unmarshal(responseBody, &res)
	for _, user := range res.Users {
		if user.ID >= 3 && user.ID <= 50 {
			hash, _ := auth.HashPassword(user.Password)
			idhash := GetMD5Hash(strconv.Itoa(user.ID))
			var u = models.User{
				No:        user.ID,
				ID:        idhash,
				Username:  user.Username,
				Email:     user.Email,
				Password:  hash,
				Name:      user.FirstName + " " + user.LastName,
				Level:     "user",
				CreatedBy: "admin",
				UpdatedBy: "admin",
				Address: models.Address{
					No:             user.ID,
					ID:             idhash,
					UserID:         idhash,
					Phone:          user.Phone,
					BillAddress:    user.Address.Address,
					BillCity:       user.Address.City,
					BillPostalCode: user.Address.PostalCode,
					BillCountry:    user.Address.Country,
					ShipAddress:    user.Address.Address,
					ShipCity:       user.Address.City,
					ShipPostalCode: user.Address.PostalCode,
					ShipCountry:    user.Address.Country,
					CreatedAt:      func(t time.Time) *time.Time { return &t }(time.Now()),
					CreatedBy:      "admin",
					UpdatedBy:      "admin",
				},
			}
			err = db.MySql.Save(&u).Error
			db.MySql.First(&u)
			u.Address.UserID = u.ID
			err = db.MySql.Save(&u.Address).Error
			if err != nil {
				fmt.Println(err)
			}
		}
	}

}

type BulkProductResponse struct {
	Products []struct {
		ID          int      `json:"id"`
		Title       string   `json:"title"`
		Description string   `json:"description"`
		Brand       string   `json:"brand"`
		Price       float64  `json:"price"`
		Images      []string `json:"images"`
		Category    string   `json:"category"`
	}
}

type BulkUserResponse struct {
	Users []struct {
		ID        int    `json:"id"`
		Username  string `json:"username"`
		Email     string `json:"email"`
		Password  string `json:"password"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Phone     string `json:"phone"`
		Address   struct {
			Address    string `json:"address"`
			City       string `json:"city"`
			PostalCode string `json:"postalCode"`
			Country    string `json:"country"`
		} `json:"address"`
	} `json:"users"`
}
