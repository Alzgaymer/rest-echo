package electronic

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"server/config"
	builder "server/src/builder"
	service "server/src/service"
	"time"

	"github.com/ilyakaznacheev/cleanenv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"gopkg.in/go-playground/validator.v9"
)

var (
	e = echo.New()
	v = validator.New()

	cfg = config.GetConfig()
)

func init() {
	err := cleanenv.ReadEnv(cfg)
	e.Logger.Printf("%+v", cfg)
	if err != nil {
		e.Logger.Fatal("Unable to load configurations")
	}

	s := service.New()
	MongoBuilder := builder.MongoBuilder{
		Service: s,
		Config:  cfg,
		Log:     e.Logger,
	}
	MongoBuilder.Build()

}

func Start() {
	e.Validator = &ProductValidator{validator: v}
	e.Logger.SetLevel(log.INFO)
	e.Use(ServerMessage)
	e.Pre(AnotherServerMessage, middleware.RemoveTrailingSlash()) //always works before Use method

	e.DELETE("/product/:id", DeleteByID)
	e.PUT("/product/:name", PutByName)
	e.POST("/product", PostAdd)
	e.GET("/", GetInit)
	e.GET("/product/:name", GetByID)
	e.GET("/products", GetAll)

	go func() {
		e.Logger.Printf("Listening on port :%s...", cfg.Port)
		if err := e.Start(fmt.Sprintf(":%s", cfg.Port)); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("Shutting down the server", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
