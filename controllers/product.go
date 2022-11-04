package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gustavofz/case-eulabs/config"
	"github.com/gustavofz/case-eulabs/model"
	"github.com/gustavofz/case-eulabs/responses"
	"github.com/labstack/echo/v4"
)

var db = config.InitConnection()

func GetAllProducts(c echo.Context) error {
	products := []*model.Product{}
	result := db.Find(&products)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, responses.ProductResponse{Message: "error", Data: &echo.Map{"data": result.Error}})

	}

	return c.JSON(http.StatusOK, responses.ProductResponse{Message: "error", Data: &echo.Map{"data": products}})
}

func GetProduct(c echo.Context) error {
	id := c.Param("id")
	var product model.Product

	result := db.First(&product, id)

	if result.Error != nil {
		fmt.Print(result.Error)
		return c.JSON(http.StatusBadRequest, responses.ProductResponse{Message: "error", Data: &echo.Map{"data": result.Error}})
	}
	return c.JSON(http.StatusOK, product)
}

func CreateProduct(c echo.Context) error {

	name := c.FormValue("name")
	brand := c.FormValue("brand")
	price, err := strconv.ParseFloat(c.FormValue("price"), 64)

	if err != nil {
		fmt.Print(err)
	}

	product := model.Product{Name: name, Brand: brand, Price: price}

	result := db.Create(&product)

	if result.Error != nil {
		fmt.Print(result.Error)
		return c.JSON(http.StatusOK, result.Error)
	}

	return c.JSON(http.StatusOK, product)
}

func UpdateProduct(c echo.Context) error {
	id := c.Param("id")
	name := c.FormValue("name")
	brand := c.FormValue("brand")
	price, err := strconv.ParseFloat(c.FormValue("price"), 64)

	if err != nil {
		log.Fatal(err)
	}

	var product model.Product

	getProduct := db.First(&product, id)

	if getProduct != nil {
		fmt.Print(getProduct.Error)
	}

	product = model.Product{Name: name, Brand: brand, Price: price}

	result := db.Save(&product)

	if result.Error != nil {
		fmt.Print(result.Error)
	}

	return c.JSON(http.StatusOK, product)

}

func DeleteProduct(c echo.Context) error {

	id := c.Param("id")

	result := db.Delete(&model.Product{}, id)

	if result.Error != nil {
		fmt.Print(result.Error)
	}
	return c.JSON(http.StatusOK, result)
}