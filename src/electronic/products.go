package electronic

import (
	"fmt"
	"net/http"
	"server/src/model"
	service "server/src/service"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/go-playground/validator.v9"
)

type ProductValidator struct {
	validator *validator.Validate
}

func (p *ProductValidator) Validate(i interface{}) error {
	return p.validator.Struct(i)
}

var market []*model.Product

func DeleteByID(c echo.Context) error {
	return c.JSON(http.StatusOK, "")
}

func PutByName(c echo.Context) error {

	var (
		product model.Product
	)

	product.Product_name = c.Param("name")
	product.ID = primitive.NewObjectID()
	product.CreationTime = product.ID.Timestamp()

	if err := c.Bind(&product); err != nil {
		log.Fatal(err)
	}

	if err := c.Validate(&product); err != nil {
		log.Fatal(err)
	}

	service := service.New()

	service.Collection = append(service.Collection, &product)
	res := service.InsertOne(c.Request().Context(), product)

	log.Print(res)

	return c.JSON(http.StatusOK, service.Collection)
}

func PostAdd(c echo.Context) error {
	return c.JSON(http.StatusOK, "")
}

func GetInit(c echo.Context) error {
	return c.String(http.StatusOK, "ECHO INIT!")
}

func GetByID(c echo.Context) error {

	return c.JSON(http.StatusNotFound, "Product Not Found")
}

func GetAll(c echo.Context) error {
	return c.JSON(http.StatusNotFound, market)
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
