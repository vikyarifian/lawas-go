package auth

import (
	"fmt"
	"lawas-go/dto"
	"lawas-go/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	secretKey = []byte("53cr3tk3Y")
)

func CreateToken(user models.User) (string, error) {

	claims := jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"name":     cases.Title(language.English, cases.Compact).String(user.Name), //user.Employee.FullName,
		"email":    user.Email,
		"level":    user.Level,
		"expired":  time.Now().Add(time.Minute * 30).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil

}

func GetToken(c *fiber.Ctx) (dto.Token, error) {

	var token dto.Token
	claims, err := VerifyToken(c.Cookies("session"))
	if err != nil {
		return token, err
	}

	// var items []models.Item
	// var bids []models.Bid
	// var watchs []models.Watchlist
	// // var user models.User
	var notif []dto.Notification
	// // db.MySql.Find(&user, "username=? AND status='A'", fmt.Sprintf("%s", claims["username"]))
	// db.MySql.Find(&items, "user_id=? AND (date + interval duration day) > now()", fmt.Sprintf("%s", claims["user_id"]))
	// db.MySql.Find(&bids, "user_id=? AND item_id in (SELECT id FROM items WHERE (date + interval duration day) > now())", fmt.Sprintf("%s", claims["user_id"]))
	// db.MySql.Find(&watchs, "user_id=? AND item_id in (SELECT id FROM items WHERE (date + interval duration day) > now())", fmt.Sprintf("%s", claims["user_id"]))
	// notif = append(notif, dto.Notification{
	// 	Code:  "SELL",
	// 	Name:  "Sells",
	// 	Count: len(items),
	// })
	// notif = append(notif, dto.Notification{
	// 	Code:  "BID",
	// 	Name:  "Bids",
	// 	Count: len(bids),
	// })
	// notif = append(notif, dto.Notification{
	// 	Code:  "WATCH",
	// 	Name:  "Watchlist",
	// 	Count: len(watchs),
	// })

	return dto.Token{
		UserID:        fmt.Sprintf("%s", claims["user_id"]),
		Username:      fmt.Sprintf("%s", claims["username"]),
		Name:          cases.Title(language.English, cases.Compact).String(fmt.Sprintf("%s", claims["name"])),
		Email:         fmt.Sprintf("%s", claims["email"]),
		Level:         fmt.Sprintf("%s", claims["level"]),
		Token:         c.Get("Authorization"),
		Notifications: notif,
	}, nil
}

func VerifyToken(tokenString string) (jwt.MapClaims, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.ErrUnauthorized
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fiber.ErrUnauthorized
	}

}
