package electronic

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
)

type Product struct {
	Product_name string `json:"product_name" validate:"required,min=4"`
}

type ProductValidator struct {
	validator *validator.Validate
}

func (p *ProductValidator) Validate(i interface{}) error {
	return p.validator.Struct(i)
}

var market = []map[int]string{{1: "tvs"}, {2: "smartphones"}, {3: "iphones"}}

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
	var (
		product map[int]string
		temp    Product
	)
	pid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	for _, p := range market {
		for index := range p {
			if index == pid {
				product = p
			}
		}
	}
	if err := c.Bind(&temp); err != nil {
		return err
	}
	if err := c.Validate(&temp); err != nil {
		return err
	}
	product[pid] = temp.Product_name
	return c.JSON(http.StatusOK, market)
}

func PostAdd(c echo.Context) error {
	var temp Product

	if err := c.Bind(&temp); err != nil {
		return err
	}
	if err := c.Validate(temp); err != nil {
		return err
	}

	newProduct := map[int]string{
		len(market) + 1: temp.Product_name,
	}
	market = append(market, newProduct)

	return c.JSON(http.StatusOK, newProduct)
}

func GetInit(c echo.Context) error {
	return c.String(http.StatusOK, "ECHO INIT!")
}

func GetByID(c echo.Context) error {
	for _, product := range market {
		for index := range product {
			pID, err := strconv.Atoi(c.Param("id"))
			if err != nil {
				return err
			}
			if pID == index {
				return c.JSON(http.StatusOK, market[index-1])
			}
		}
	}
	return c.JSON(http.StatusNotFound, "Product Not Found")
}

func GetAll(c echo.Context) error {
	return c.JSON(http.StatusNotFound, market)
}
