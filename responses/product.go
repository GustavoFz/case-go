package responses

import "github.com/labstack/echo/v4"

type ProductResponse struct {
	Status   int       `json:"status"`
	Message  string    `json:"message"`
	Response *echo.Map `json:"response"`
}
