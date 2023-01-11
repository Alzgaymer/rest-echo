package electronic

import (
	"fmt"
	"server/config"
	service "server/src/service"

	"github.com/ilyakaznacheev/cleanenv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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

	service := service.New()
	service.SetURI(cfg.MongoURI)
}

func Start() {

	e.Validator = &ProductValidator{validator: v}

	e.Use(ServerMessage)
	e.Pre(AnotherServerMessage, middleware.RemoveTrailingSlash()) //always works before Use method

	e.DELETE("/products/:id", DeleteByID)
	e.PUT("/products/:name", PutByName)
	e.POST("/products", PostAdd)
	e.GET("/", GetInit)
	e.GET("/products/:id", GetByID)
	e.GET("/products", GetAll)

	log.Printf("Listening on port :%s...", cfg.Port)
	log.Fatal(e.Start(fmt.Sprintf(":%s", cfg.Port)))
}
