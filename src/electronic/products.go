package electronic

import (
	"context"
	"fmt"
	"net/http"
	"server/src/model"
	"server/src/service"
	"sync"
	"time"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/go-playground/validator.v9"
)

type ProductValidator struct {
	validator *validator.Validate
}

func (p *ProductValidator) Validate(i interface{}) error {
	return p.validator.Struct(i)
}

var (
	lock sync.Mutex
)

func DeleteByID(c echo.Context) error {
	return c.JSON(http.StatusOK, "")
}

func PutByName(c echo.Context) error {

	return c.JSON(http.StatusOK, "")
}

func PostAdd(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	var (
		product model.Product
		filter  bson.M
		log     echo.Logger = c.Logger()
	)

	if err := c.Bind(&product); err != nil {
		return err
	}
	if err := c.Validate(&product); err != nil {
		return err
	}
	product.ID = primitive.NewObjectID()
	product.CreationTime = product.ID.Timestamp().Add(2 * time.Hour)

	filter = bson.M{"product_name": product.Product_name}

	log.Print("Finding in mongo...")

	res := service.New().Col().FindOne(context.Background(), filter)
	log.Print(res.Err())
	log.Print(model.Decode(res))
	if res.Err() != nil { //if true adding to mongo
		res, err := service.New().Col().InsertOne(
			context.Background(),
			model.MgoToBson(model.ModelToMgo(&product)),
		)
		log.Print(res)
		if err != nil {
			log.Print(err)
		}
	} else {
		str := fmt.Sprintf("Product: %v, already exist", product)
		log.Printf(str)
		return c.JSON(http.StatusOK, str)
	}

	return c.JSON(http.StatusOK, product)
}

func GetInit(c echo.Context) error {
	return c.String(http.StatusOK, "ECHO INIT!")
}

func GetByID(c echo.Context) error {
	return c.JSON(http.StatusOK, "")
}

func GetAll(c echo.Context) error {
	return c.JSON(http.StatusNotFound, "")
}

func AnotherServerMessage(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		fmt.Println("[1]: inside middleware")
		return next(ctx)
	}
}
func ServerMessage(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		fmt.Println("[2]: inside middleware")
		return next(ctx)
	}
}
