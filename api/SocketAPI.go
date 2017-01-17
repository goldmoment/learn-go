package api

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/goldmoment/model"
	"github.com/googollee/go-socket.io"
)

// JWT schema of the data it will store
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func SoServer() (*socketio.Server, error) {
	server, err := socketio.NewServer(nil)
	if err != nil {
		return nil, err
	}

	server.On("connection", func(so socketio.Socket) {

		tokenString := so.Request().Header.Get("Authorization")

		// Return a Token using the cookie
		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			// Make sure token's signature wasn't changed
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected siging method")
			}
			return []byte("secret"), nil
		})
		if err != nil {
			so.Emit("unauthorization", "404")
			return
		}

		// Grab the tokens claims and pass it into the original request
		if _, ok := token.Claims.(*Claims); ok && token.Valid {
			log.Println("on connection")

			so.Join("shop")

			so.Emit("message", "hello")

			// 		so.On("message", func(msg string) string {
			// 			log.Println(msg)
			// 			return "recieved"
			// 		})

			so.On("message", func(msg string) {
				var pro model.Product
				err := json.Unmarshal([]byte(msg), &pro)
				if err == nil {
					log.Println(pro.ID)
				}
				log.Println(msg)
			})
		} else {
			so.Emit("unauthorization", "404")
			return
		}

		so.On("disconnection", func() {
			log.Println("on disconnect")
		})
	})

	server.On("error", func(so socketio.Socket, err error) {
		log.Println("error:", err)
	})
	return server, nil
}
