package api

import (
	// "fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"github.com/goldmoment/learn-go/dataloader"
	"github.com/goldmoment/learn-go/model"
)

// func ProductAdd(c echo.Context) error {
// 	product := new(model.Product)

// 	categoryid := c.FormValue("categoryid")
// 	product.Name = c.FormValue("name")
// 	product.Price, _ = strconv.Atoi(c.FormValue("price"))
// 	product.Quantity, _ = strconv.Atoi(c.FormValue("quantity"))

// 	// Multipart form
// 	form, err := c.MultipartForm()
// 	if err != nil {
// 		return err
// 	}

// 	descriptions := form.Value["descriptions[]"]
// 	pics := form.Value["pics[]"]

// 	files := form.File["files"]
// 	for index, file := range files {
// 		// Source
// 		src, err := file.Open()
// 		if err != nil {
// 			return err
// 		}
// 		defer src.Close()

// 		// Rename image
// 		filename := utils.GetImageName(file.Filename)

// 		// Destination
// 		dst, err := os.Create("./assets/img/" + filename)
// 		if err != nil {
// 			return err
// 		}
// 		defer dst.Close()

// 		// Copy
// 		if _, err = io.Copy(dst, src); err != nil {
// 			return err
// 		}

// 		// Save pic
// 		src, err = file.Open()
// 		if err != nil {
// 			return err
// 		}
// 		defer src.Close()

// 		img, _, err := image.Decode(src)
// 		if err != nil {
// 			return err
// 		}

// 		pic := new(model.Picture)
// 		pic.Color = dominantcolor.Hex(dominantcolor.Find(img))
// 		pic.Width = img.Bounds().Max.X
// 		pic.Height = img.Bounds().Max.Y
// 		pic.Ratio = float32(pic.Width) / float32(pic.Height)
// 		pic.Path = "/assets/img/" + filename
// 		pic.Description = descriptions[index]

// 		err = dbl.AddPicture(pic)
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	dbl.AddProduct(categoryid, product)
// 	return c.JSON(http.StatusOK, product)
// }

func AddProduct(c echo.Context) error {
	product := model.Product{}

	categoryid := c.FormValue("categoryid")
	product.Name = c.FormValue("name")
	product.Price, _ = strconv.Atoi(c.FormValue("price"))
	product.Quantity, _ = strconv.Atoi(c.FormValue("quantity"))

	descriptions := c.Request().Form["descriptions[]"]
	pics := c.Request().Form["pics[]"]

	for i, pic := range pics {
		err := dbl.UpdatePicture(pic, descriptions[i])
		if err != nil {
			return err
		}
	}

	err := dbl.AddProduct(categoryid, &product)
	if err != nil {
		return err
	}

	for i, pic := range pics {
		err = dbl.AddProductPicture(product.ID, pic, i)
		if err != nil {
			return err
		}

		err = dbl.RemovePictureTimeoutByPicID(pic)
		if err != nil {
			return err
		}
	}

	return c.JSON(http.StatusOK, product)
}

func Product(c echo.Context) error {
	categoryid := c.FormValue("categoryid")

	products := dbl.GetProducts(categoryid)
	return c.JSON(http.StatusOK, products)
}

func ProductHotest(c echo.Context) error {
	products := dbl.GetProductsHotest()
	return c.JSON(http.StatusOK, products)
}
