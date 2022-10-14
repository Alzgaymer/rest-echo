package electronic

import (
	"fmt"
	"server/config"

	"github.com/ilyakaznacheev/cleanenv"

	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
)

var e = echo.New()
var v = validator.New()
var log = e.Logger
var cfg = config.GetConfig()

func init() {
	err := cleanenv.ReadEnv(cfg)
	e.Logger.Printf("%+v", cfg)
	if err != nil {
		log.Fatal("Unable to load configurations")
	}
}

func Start() {

	e.Validator = &ProductValidator{validator: v}

	e.DELETE("/products/:id", DeleteByID)
	e.PUT("/products/:id", PutByID)
	e.POST("/products", PostAdd)
	e.GET("/", GetInit)
	e.GET("/products/:id", GetByID)
	e.GET("/products", GetAll)

	log.Printf("Listening on port :%s...", cfg.Port)
	log.Fatal(e.Start(fmt.Sprintf(":%s", cfg.Port)))
}
