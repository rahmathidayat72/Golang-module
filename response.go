package golangmodule

import "github.com/labstack/echo/v4"

type Response struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
}

func BuildResponse(data interface{}, status int, message string, c echo.Context) error {
	success := true
	if status >= 400 {
		success = false
	}

	result := Response{
		Message: message,
		Code:    status,
		Success: success,
		Data:    data,
	}

	return c.JSON(status, result)
}
