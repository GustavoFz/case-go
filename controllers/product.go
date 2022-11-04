package controllers

import (
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
		return c.JSON(http.StatusBadRequest, responses.ProductResponse{Status: http.StatusBadRequest, Message: "There was an error in the query"})
	}

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, responses.ProductResponse{Status: http.StatusNotFound, Message: "There are no registered products"})
	}

	return c.JSON(http.StatusOK, responses.ProductResponse{Status: http.StatusOK, Message: "success", Response: &echo.Map{"data": products}})
}

func GetProduct(c echo.Context) error {
	id := c.Param("id")
	var product model.Product

	result := db.First(&product, id)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, responses.ProductResponse{Status: http.StatusBadRequest, Message: "There was an error in the query"})
	}

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, responses.ProductResponse{Status: http.StatusNotFound, Message: "Product not found"})
	}

	return c.JSON(http.StatusOK, responses.ProductResponse{Status: http.StatusOK, Message: "success", Response: &echo.Map{"data": product}})
}

func CreateProduct(c echo.Context) error {

	name := c.FormValue("name")
	brand := c.FormValue("brand")
	price, err := strconv.ParseFloat(c.FormValue("price"), 64)

	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.ProductResponse{Status: http.StatusBadRequest, Message: "Price is not a number"})
	}

	product := model.Product{Name: name, Brand: brand, Price: price}

	result := db.Create(&product)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, responses.ProductResponse{Status: http.StatusBadRequest, Message: "An error occurred while creating the product"})
	}

	return c.JSON(http.StatusOK, responses.ProductResponse{Status: http.StatusOK, Message: "The product was created", Response: &echo.Map{"data": product}})
}

func UpdateProduct(c echo.Context) error {
	id := c.Param("id")
	name := c.FormValue("name")
	brand := c.FormValue("brand")
	price, err := strconv.ParseFloat(c.FormValue("price"), 64)

	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.ProductResponse{Status: http.StatusBadRequest, Message: "Price is not a number"})
	}

	var product model.Product

	getProduct := db.First(&product, id)

	if getProduct.Error != nil {
		return c.JSON(http.StatusNotFound, responses.ProductResponse{Status: http.StatusNotFound, Message: "Product not found"})
	}

	product = model.Product{Name: name, Brand: brand, Price: price}

	result := db.Save(&product)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, responses.ProductResponse{Status: http.StatusBadRequest, Message: "An error occurred while updating the product"})
	}

	return c.JSON(http.StatusOK, responses.ProductResponse{Status: http.StatusOK, Message: "The product has been updated", Response: &echo.Map{"data": product}})

}

func DeleteProduct(c echo.Context) error {

	id := c.Param("id")

	result := db.Delete(&model.Product{}, id)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, responses.ProductResponse{Status: http.StatusBadRequest, Message: "An error occurred while deleting the product"})
	}
	return c.JSON(http.StatusOK, responses.ProductResponse{Status: http.StatusOK, Message: "The product has been deleted"})
}
