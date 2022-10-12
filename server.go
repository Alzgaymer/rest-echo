package main

import (
	"fmt"
	"net/http"
	"os"
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

func main() {
	market := []map[int]string{{1: "tvs"}, {2: "smartphones"}, {3: "iphones"}}
	e := echo.New()
	v := validator.New()
	e.Validator = &ProductValidator{validator: v}
	port := os.Getenv("MY_APP_PORT")
	if port == "" {
		port = "8080"
	}
	e.DELETE("/products/:id", func(c echo.Context) error {
		var newMarket []map[int]string
		pid, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return err
		}
		for _, p := range market {
			for index, _ := range p {
				if index == pid {
					continue
				}
				newMarket = append(newMarket, p)
			}
		}
		if len(newMarket) == len(market) {
			return c.JSON(http.StatusNotFound, "No such element")
		}
		market = newMarket
		return c.JSON(http.StatusOK, market)

	})
	e.PUT("/products/:id", func(c echo.Context) error {
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

	})
	e.POST("/products", func(c echo.Context) error {
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
	})
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "ECHO INIT!")
	})
	e.GET("/querry/:q", func(c echo.Context) error {
		return c.JSON(http.StatusOK, c.QueryParam("querryParam"))
	})
	e.GET("/products/:id", func(c echo.Context) error {
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
	})
	e.GET("/products", func(c echo.Context) error {
		return c.JSON(http.StatusNotFound, market)
	})
	log := e.Logger
	log.Printf("Listening on port :%s...", port)
	log.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
