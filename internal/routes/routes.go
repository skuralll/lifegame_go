package routes

import (
	"github.com.skuralll/lifegame_go/internal/handler"
	"github.com/labstack/echo/v4"
)

// エンドポイント登録
func RegisterRoutes(e *echo.Echo) {
	e.GET("/hello", handler.SayHello)
}
