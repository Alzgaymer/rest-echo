package electronic

import (
	"fmt"
	"net/http"
	"server/src/model"
	"strconv"

	"github.com/labstack/echo"
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

	pid, err := strconv.Atoi(c.Param("id"))
	pid--
	if err != nil {
		return err
	}
	market = append(market[:pid], market[pid+1:]...)
	return c.JSON(http.StatusOK, market)
}

func PutByID(c echo.Context) error {
	return c.JSON(http.StatusOK, "")
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
