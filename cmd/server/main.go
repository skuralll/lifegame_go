package main

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com.skuralll/lifegame_go/internal/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	// Echo instance
	e := echo.New()
	// Register routes
	routes.RegisterRoutes(e)
	// Start Server
	if err := e.Start(":8080"); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("failed to start server", "error", err)
	}
}
