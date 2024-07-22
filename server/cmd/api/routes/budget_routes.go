package routes

import (
	"github.com/jayromofo/project-selkies/server/cmd/api/handlers"
	"github.com/labstack/echo/v4"
)

func GetBudgetRoutes(e *echo.Echo) {
	e.POST("/api/budget/create/", handlers.AddBudget)
	e.GET("/api/budget/create/", handlers.GetBudget)
	e.DELETE("/api/budget/:id", handlers.DeleteBudget)
	e.PATCH("/api/budget/:id", handlers.EditBudget)
}
