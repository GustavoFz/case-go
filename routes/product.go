package routes

import (
	"github.com/gustavofz/case-eulabs/controllers"
	"github.com/labstack/echo/v4"
)

func ProductRoute(e *echo.Echo) {
	// Get all products
	e.GET("/products", controllers.GetAllProducts)

	// Get a product
	e.GET("/products/:id", controllers.GetProduct)

	// Create a new product
	e.POST("/products", controllers.CreateProduct)

	// Update a product
	e.PUT("/products/:id", controllers.UpdateProduct)

	// Delete a product
	e.DELETE("products/:id", controllers.DeleteProduct)
}
