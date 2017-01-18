package router

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"../api"
)

func InitServer() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/assets", "assets")

	// Login route
	e.POST("/login", api.Login)
	// Register router
	e.POST("/register", api.Register)

	// Unauthenticated route
	e.GET("/producthotest", api.ProductHotest)
	e.GET("/category", api.PublicCategory)
	e.GET("/product", api.Product)
	e.POST("/product", api.AddProduct)
	e.POST("/checkout", api.Checkout)
	e.POST("/picture", api.PictureAdd)

	// Restricted group
	r := e.Group("/pri")
	r.Use(middleware.JWT([]byte("secret")))
	r.GET("/category", api.Category)
	r.GET("/product", api.Product)

	// WebSocket
	server, _ := api.SoServer()
	e.GET("/socket.io/", echo.WrapHandler(server))

	e.Start(":8080")
}
