package api

import (
	"net/http"

	"github.com/labstack/echo"

	"../dataloader"
)

func Category(c echo.Context) error {
	userid := c.FormValue("userid")

	categories := dbl.GetCategories(userid)
	return c.JSON(http.StatusOK, categories)
}

func PublicCategory(c echo.Context) error {
	categories := dbl.GetPublicCategories()
	return c.JSON(http.StatusOK, categories)
}
