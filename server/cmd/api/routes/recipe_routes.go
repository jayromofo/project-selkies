package routes

import (
	"github.com/jayromofo/project-selkies/server/cmd/api/handlers"
	"github.com/labstack/echo/v4"
)

func GetRecipeRoutes(e *echo.Echo) {
	e.PUT("/api/recipe/", handlers.AddRecipe)
	e.GET("/api/recipe/", handlers.ViewAllRecipes)
	e.GET("/api/recipe/:id", handlers.ViewRecipe)
	e.DELETE("/api/recipe/:id", handlers.DeleteRecipe)
	e.PATCH("/api/recipe/:id", handlers.EditRecipe)
}
