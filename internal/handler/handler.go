package handler

import (
	"net/http"

	"github.com.skuralll/lifegame_go/internal/game"
	"github.com/labstack/echo/v4"
)

// サンプルハンドラ
func SayHello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

// 次世代の盤面情報を取得するハンドラ
func GetNextBoard(c echo.Context) error {
	var board game.Board
	if err := c.Bind(&board); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid board data"})
	}

	return c.JSON(http.StatusOK, board.GetNext())
}