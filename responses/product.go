package responses

import "github.com/labstack/echo/v4"

type ProductResponse struct {
	Message string    `json:"message"`
	Data    *echo.Map `json:"data"`
}
