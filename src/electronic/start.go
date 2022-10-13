package electronic

import (
	"fmt"
	"os"

	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
)

var e = echo.New()
var v = validator.New()

func Start() {

	e.Validator = &ProductValidator{validator: v}
	port := os.Getenv("MY_APP_PORT")
	if port == "" {
		port = "8080"
	}

	e.DELETE("/products/:id", DeleteByID)
	e.PUT("/products/:id", PutByID)
	e.POST("/products", PostAdd)
	e.GET("/", GetInit)
	e.GET("/products/:id", GetByID)
	e.GET("/products", GetAll)

	log := e.Logger
	log.Printf("Listening on port :%s...", port)
	log.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
