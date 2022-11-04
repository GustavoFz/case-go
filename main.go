package main

import (
	"github.com/gustavofz/case-eulabs/routes"
	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	// Routes products
	routes.ProductRoute(e)

	e.Logger.Fatal(e.Start(":3000"))
}
