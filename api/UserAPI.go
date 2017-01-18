package api

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"

	"github.com/goldmoment/learn-go/dataloader"
)

func Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	userid := dbl.ValidateUser(username, password)
	if userid != "nil" {
		// Create token
		token := jwt.New(jwt.SigningMethodHS256)

		// Set claims
		expired := time.Now().Add(time.Hour * 72)
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = username
		claims["admin"] = true
		claims["exp"] = expired.Unix()

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}

		// Save login status in database
		dbl.UpdateUser(userid, t, expired, "user")

		return c.JSON(http.StatusOK, map[string]string{
			"token":  t,
			"userid": userid,
		})
	}

	return echo.ErrUnauthorized
}

func Register(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	res := dbl.RegisterUser(username, password)
	if res == false {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Can't register user " + username,
		})
	} else {
		return c.JSON(http.StatusCreated, map[string]string{
			"username": username,
		})
	}
}
