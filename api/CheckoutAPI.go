package api

import (
	"fmt"
	"net/http"

	"../parameter"

	"github.com/labstack/echo"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
)

func Checkout(c echo.Context) error {
	stripe.Key = "sk_test_0lN7kTlPqSc0qVcRyQpqOKld"

	checkout := &parameter.Checkout{}
	if err := c.Bind(checkout); err != nil {
		return err
	}

	for _, product := range checkout.Products {
		fmt.Println(product.ID)
	}

	chargeParams := &stripe.ChargeParams{
		Amount:   12300,
		Currency: "usd",
		Customer: "cus_9WTDQpR4TB9yCY",
		Desc:     "Charge for isabella.garcia@example.com",
	}
	chargeParams.SetSource("card_19GjuTLikvBi66EuVPonK1X4@")

	ch, err := charge.New(chargeParams)

	if err != nil {
		return c.JSON(402, "@@")
	}

	return c.JSON(http.StatusOK, ch)
}
