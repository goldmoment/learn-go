package api

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"net/http"
	"os"

	"github.com/cenkalti/dominantcolor"
	"github.com/labstack/echo"

	"github.com/goldmoment/dataloader"
	"github.com/goldmoment/model"
	"github.com/goldmoment/utils"
)

func PictureAdd(c echo.Context) error {

	file, err := c.FormFile("image")
	// Source
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Rename image
	filename := utils.GetImageName(file.Filename)

	// Destination
	dst, err := os.Create("./assets/img/" + filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	// Save pic
	src, err = file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	img, _, err := image.Decode(src)
	if err != nil {
		return err
	}

	pic := new(model.Picture)
	pic.Color = dominantcolor.Hex(dominantcolor.Find(img))
	pic.Width = img.Bounds().Max.X
	pic.Height = img.Bounds().Max.Y
	pic.Ratio = float32(pic.Width) / float32(pic.Height)
	pic.Path = "/assets/img/" + filename

	err = dbl.AddPicture(pic)
	if err != nil {
		return err
	}

	err = dbl.AddPictureTimeout(pic)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, pic)
}
