package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// サンプルハンドラ
func SayHello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
